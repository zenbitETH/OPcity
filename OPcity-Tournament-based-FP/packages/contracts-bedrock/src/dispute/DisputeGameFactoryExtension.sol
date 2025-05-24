
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { IDisputeGame } from "../../interfaces/dispute/IDisputeGame.sol";
import { ITournamentGame } from "../../interfaces/dispute/ITournamentGame.sol";
import { TournamentFactory } from "./TournamentFactory.sol";
import { GameTypes } from "./lib/GameTypes.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { Semver } from "../universal/Semver.sol";

/**
 * @title DisputeGameFactoryExtension
 * @notice Factory contract for creating and managing dispute games with support for DAVE's tournament-based games
 * @dev This contract extends the OP Stack's dispute game factory pattern to support tournament-based dispute games
 */
contract DisputeGameFactoryExtension is Ownable, Semver {
    /**
     * @notice Mapping of game type to implementation address
     * @dev Maps game type identifier to the address of the factory contract for that game type
     */
    mapping(uint8 => address) public gameImpls;

    /**
     * @notice Mapping of all created games
     * @dev Maps game UUID to the address of the created game
     */
    mapping(bytes32 => address) public games;

    /**
     * @notice Array of all created game addresses
     */
    address[] public allGames;

    /**
     * @notice Tournament factory contract
     * @dev Used to create tournament-based dispute games
     */
    TournamentFactory public immutable TOURNAMENT_FACTORY;

    /**
     * @notice Emitted when a new dispute game is created
     * @param disputeProxy Address of the created dispute game
     * @param gameType Type of the dispute game
     * @param rootClaim Root claim of the dispute game
     */
    event DisputeGameCreated(
        address indexed disputeProxy,
        uint8 indexed gameType,
        bytes32 indexed rootClaim
    );

    /**
     * @notice Emitted when a game implementation is set
     * @param gameType Type of the dispute game
     * @param impl Address of the implementation contract
     */
    event ImplementationSet(uint8 indexed gameType, address indexed impl);

    /**
     * @notice Error thrown when attempting to create a game with an unregistered game type
     */
    error UnknownGameType(uint8 gameType);

    /**
     * @notice Error thrown when attempting to create a game that already exists
     */
    error GameAlreadyExists(bytes32 uuid);

    /**
     * @notice Constructor for the DisputeGameFactoryExtension contract
     * @param _owner Address that will own the contract
     * @param _tournamentFactory Address of the TournamentFactory contract
     */
    constructor(
        address _owner,
        address _tournamentFactory
    ) Semver(1, 0, 0) {
        if (_owner == address(0)) {
            revert("Owner cannot be address(0)");
        }
        if (_tournamentFactory == address(0)) {
            revert("TournamentFactory cannot be address(0)");
        }

        _transferOwnership(_owner);
        TOURNAMENT_FACTORY = TournamentFactory(_tournamentFactory);
    }

    /**
     * @notice Sets the implementation address for a game type
     * @param _gameType Type of the dispute game
     * @param _impl Address of the implementation contract
     * @dev Can only be called by the owner
     */
    function setImplementation(uint8 _gameType, address _impl) external onlyOwner {
        gameImpls[_gameType] = _impl;
        emit ImplementationSet(_gameType, _impl);
    }

    /**
     * @notice Creates a new dispute game
     * @param _gameType Type of the dispute game
     * @param _rootClaim Root claim of the dispute game
     * @param _extraData Extra data supplied to the dispute game
     * @return The address of the created dispute game
     */
    function create(
        uint8 _gameType,
        bytes32 _rootClaim,
        bytes calldata _extraData
    ) external payable returns (address) {
        // Check if the game type is registered
        address impl = gameImpls[_gameType];
        if (impl == address(0)) {
            revert UnknownGameType(_gameType);
        }

        // Generate a unique identifier for the game
        bytes32 uuid = keccak256(abi.encode(_gameType, _rootClaim, _extraData));

        // Check if the game already exists
        if (games[uuid] != address(0)) {
            revert GameAlreadyExists(uuid);
        }

        // Create the game based on its type
        address proxy;
        if (_gameType == GameTypes.TOURNAMENT) {
            // For tournament games, use the TournamentFactory
            proxy = _createTournamentGame(msg.sender, _rootClaim, _extraData);
        } else {
            // For other game types, delegate to the appropriate implementation
            (bool success, bytes memory returndata) = impl.delegatecall(
                abi.encodeWithSignature(
                    "create(address,bytes32,bytes)",
                    msg.sender,
                    _rootClaim,
                    _extraData
                )
            );

            if (!success) {
                // Revert with the error message if the call failed
                assembly {
                    revert(add(returndata, 32), mload(returndata))
                }
            }

            // Extract the proxy address from the return data
            proxy = abi.decode(returndata, (address));
        }

        // Register the game
        games[uuid] = proxy;
        allGames.push(proxy);

        emit DisputeGameCreated(proxy, _gameType, _rootClaim);

        return proxy;
    }

    /**
     * @notice Registers an existing game with the factory
     * @param _game Address of the game to register
     * @dev This is used by the TournamentFactory to register created tournaments
     */
    function registerGame(address _game) external {
        // Only the TournamentFactory can register games
        require(
            msg.sender == address(TOURNAMENT_FACTORY),
            "Only TournamentFactory can register games"
        );

        // Get the game data
        IDisputeGame game = IDisputeGame(_game);
        (uint8 gameType, bytes32 rootClaim, bytes memory extraData) = game.gameData();

        // Generate the UUID
        bytes32 uuid = keccak256(abi.encode(gameType, rootClaim, extraData));

        // Register the game if it doesn't already exist
        if (games[uuid] == address(0)) {
            games[uuid] = _game;
            allGames.push(_game);

            emit DisputeGameCreated(_game, gameType, rootClaim);
        }
    }

    /**
     * @notice Gets a dispute game by its identifier
     * @param _gameType Type of the dispute game
     * @param _rootClaim Root claim of the dispute game
     * @param _extraData Extra data supplied to the dispute game
     * @return The address of the dispute game
     */
    function getGame(
        uint8 _gameType,
        bytes32 _rootClaim,
        bytes calldata _extraData
    ) external view returns (address) {
        return games[keccak256(abi.encode(_gameType, _rootClaim, _extraData))];
    }

    /**
     * @notice Gets the total number of games created
     * @return The total number of games
     */
    function gameCount() external view returns (uint256) {
        return allGames.length;
    }

    /**
     * @notice Creates a new tournament game
     * @param _gameCreator Address that created the game
     * @param _rootClaim Root claim of the game
     * @param _extraData Extra data supplied to the game
     * @return The address of the created tournament game
     */
    function _createTournamentGame(
        address _gameCreator,
        bytes32 _rootClaim,
        bytes calldata _extraData
    ) internal returns (address) {
        // Get the current L1 head hash
        bytes32 l1Head = blockhash(block.number - 1);

        // Create a new tournament using the TournamentFactory
        return TOURNAMENT_FACTORY.createTournament(
            _gameCreator,
            _rootClaim,
            l1Head,
            _extraData
        );
    }
}
