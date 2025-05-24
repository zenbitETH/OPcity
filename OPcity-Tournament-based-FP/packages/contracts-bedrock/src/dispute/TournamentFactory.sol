// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { ITournamentGame } from "../../interfaces/dispute/ITournamentGame.sol";
import { TournamentGame } from "./TournamentGame.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { Address } from "@openzeppelin/contracts/utils/Address.sol";
import { Create2 } from "@openzeppelin/contracts/utils/Create2.sol";
import { Semver } from "../universal/Semver.sol";
import { Clone } from "../libraries/Clone.sol";

/**
 * @title TournamentFactory
 * @notice Factory contract for creating and managing tournament-based dispute games
 * @dev This contract follows the factory pattern used in the OP Stack for creating dispute games
 *      It creates new tournament instances and registers them with the DisputeGameFactory
 */
contract TournamentFactory is Ownable, Semver {
    /**
     * @notice Implementation of the TournamentGame contract
     * @dev This is the base contract that will be cloned for each new tournament
     */
    TournamentGame public immutable TOURNAMENT_IMPLEMENTATION;

    /**
     * @notice Address of the DisputeGameFactory contract
     * @dev This is the contract that will register the created tournaments
     */
    address public immutable DISPUTE_GAME_FACTORY;

    /**
     * @notice Default match duration for tournaments
     * @dev This is the minimum time that must pass before a match can be resolved
     */
    uint256 public defaultMatchDuration;

    /**
     * @notice Default minimum bond amount for tournaments
     * @dev This is the minimum amount of ETH required to join a tournament
     */
    uint256 public defaultMinBondAmount;

    /**
     * @notice Mapping of tournament addresses to their creation data
     * @dev Maps tournament address to a tuple of (gameCreator, rootClaim, l1Head, extraData)
     */
    mapping(address => bytes) public tournamentData;

    /**
     * @notice Array of all created tournament addresses
     */
    address[] public tournaments;

    /**
     * @notice Emitted when a new tournament is created
     * @param tournament Address of the created tournament
     * @param gameCreator Address that created the tournament
     * @param rootClaim Root claim of the tournament
     * @param l1Head L1 head hash at the time the tournament was created
     */
    event TournamentCreated(
        address indexed tournament,
        address indexed gameCreator,
        bytes32 indexed rootClaim,
        bytes32 l1Head
    );

    /**
     * @notice Emitted when the default match duration is updated
     * @param oldDuration Previous match duration
     * @param newDuration New match duration
     */
    event DefaultMatchDurationUpdated(uint256 oldDuration, uint256 newDuration);

    /**
     * @notice Emitted when the default minimum bond amount is updated
     * @param oldAmount Previous minimum bond amount
     * @param newAmount New minimum bond amount
     */
    event DefaultMinBondAmountUpdated(uint256 oldAmount, uint256 newAmount);

    /**
     * @notice Error thrown when attempting to create a tournament with an invalid root claim
     */
    error InvalidRootClaim();

    /**
     * @notice Error thrown when attempting to create a tournament with an invalid match duration
     */
    error InvalidMatchDuration();

    /**
     * @notice Error thrown when attempting to create a tournament with an invalid minimum bond amount
     */
    error InvalidMinBondAmount();

    /**
     * @notice Constructor for the TournamentFactory contract
     * @param _disputeGameFactory Address of the DisputeGameFactory contract
     * @param _matchDuration Default match duration for tournaments
     * @param _minBondAmount Default minimum bond amount for tournaments
     */
    constructor(
        address _disputeGameFactory,
        uint256 _matchDuration,
        uint256 _minBondAmount
    ) Semver(1, 0, 0) {
        if (_disputeGameFactory == address(0)) {
            revert Address.AddressZero();
        }
        if (_matchDuration == 0) {
            revert InvalidMatchDuration();
        }
        if (_minBondAmount == 0) {
            revert InvalidMinBondAmount();
        }

        DISPUTE_GAME_FACTORY = _disputeGameFactory;
        defaultMatchDuration = _matchDuration;
        defaultMinBondAmount = _minBondAmount;

        // Deploy the implementation contract
        TOURNAMENT_IMPLEMENTATION = new TournamentGame(_matchDuration, _minBondAmount);
    }

    /**
     * @notice Creates a new tournament with the default parameters
     * @param _gameCreator Address that created the tournament
     * @param _rootClaim Root claim of the tournament
     * @param _l1Head L1 head hash at the time the tournament was created
     * @param _extraData Extra data supplied to the tournament
     * @return The address of the created tournament
     */
    function createTournament(
        address _gameCreator,
        bytes32 _rootClaim,
        bytes32 _l1Head,
        bytes calldata _extraData
    ) external returns (address) {
        return _createTournament(
            _gameCreator,
            _rootClaim,
            _l1Head,
            _extraData,
            defaultMatchDuration,
            defaultMinBondAmount
        );
    }

    /**
     * @notice Creates a new tournament with custom parameters
     * @param _gameCreator Address that created the tournament
     * @param _rootClaim Root claim of the tournament
     * @param _l1Head L1 head hash at the time the tournament was created
     * @param _extraData Extra data supplied to the tournament
     * @param _matchDuration Custom match duration for the tournament
     * @param _minBondAmount Custom minimum bond amount for the tournament
     * @return The address of the created tournament
     */
    function createTournamentWithCustomParams(
        address _gameCreator,
        bytes32 _rootClaim,
        bytes32 _l1Head,
        bytes calldata _extraData,
        uint256 _matchDuration,
        uint256 _minBondAmount
    ) external returns (address) {
        if (_matchDuration == 0) {
            revert InvalidMatchDuration();
        }
        if (_minBondAmount == 0) {
            revert InvalidMinBondAmount();
        }

        return _createTournament(
            _gameCreator,
            _rootClaim,
            _l1Head,
            _extraData,
            _matchDuration,
            _minBondAmount
        );
    }

    /**
     * @notice Updates the default match duration
     * @param _matchDuration New default match duration
     * @dev Can only be called by the owner
     */
    function setDefaultMatchDuration(uint256 _matchDuration) external onlyOwner {
        if (_matchDuration == 0) {
            revert InvalidMatchDuration();
        }

        uint256 oldDuration = defaultMatchDuration;
        defaultMatchDuration = _matchDuration;

        emit DefaultMatchDurationUpdated(oldDuration, _matchDuration);
    }

    /**
     * @notice Updates the default minimum bond amount
     * @param _minBondAmount New default minimum bond amount
     * @dev Can only be called by the owner
     */
    function setDefaultMinBondAmount(uint256 _minBondAmount) external onlyOwner {
        if (_minBondAmount == 0) {
            revert InvalidMinBondAmount();
        }

        uint256 oldAmount = defaultMinBondAmount;
        defaultMinBondAmount = _minBondAmount;

        emit DefaultMinBondAmountUpdated(oldAmount, _minBondAmount);
    }

    /**
     * @notice Gets the total number of tournaments created
     * @return The total number of tournaments
     */
    function getTournamentCount() external view returns (uint256) {
        return tournaments.length;
    }

    /**
     * @notice Gets a tournament address by index
     * @param _index The index of the tournament
     * @return The address of the tournament
     */
    function getTournamentAtIndex(uint256 _index) external view returns (address) {
        require(_index < tournaments.length, "Index out of bounds");
        return tournaments[_index];
    }

    /**
     * @notice Gets all tournament addresses
     * @return An array of all tournament addresses
     */
    function getAllTournaments() external view returns (address[] memory) {
        return tournaments;
    }

    /**
     * @notice Checks if an address is a tournament created by this factory
     * @param _tournament The address to check
     * @return True if the address is a tournament created by this factory, false otherwise
     */
    function isTournament(address _tournament) external view returns (bool) {
        return tournamentData[_tournament].length > 0;
    }

    /**
     * @notice Internal function to create a new tournament
     * @param _gameCreator Address that created the tournament
     * @param _rootClaim Root claim of the tournament
     * @param _l1Head L1 head hash at the time the tournament was created
     * @param _extraData Extra data supplied to the tournament
     * @param _matchDuration Match duration for the tournament
     * @param _minBondAmount Minimum bond amount for the tournament
     * @return The address of the created tournament
     */
    function _createTournament(
        address _gameCreator,
        bytes32 _rootClaim,
        bytes32 _l1Head,
        bytes calldata _extraData,
        uint256 _matchDuration,
        uint256 _minBondAmount
    ) internal returns (address) {
        if (_rootClaim == bytes32(0)) {
            revert InvalidRootClaim();
        }

        // Generate a unique salt for the tournament
        bytes32 salt = keccak256(
            abi.encode(
                _gameCreator,
                _rootClaim,
                _l1Head,
                _extraData,
                _matchDuration,
                _minBondAmount,
                block.timestamp
            )
        );

        // Deploy a new tournament using Create2 for deterministic addresses
        TournamentGame tournament;
        if (_matchDuration == defaultMatchDuration && _minBondAmount == defaultMinBondAmount) {
            // Use the implementation contract directly if using default parameters
            tournament = TournamentGame(
                address(
                    Clone.clone(
                        address(TOURNAMENT_IMPLEMENTATION),
                        salt
                    )
                )
            );
        } else {
            // Deploy a new implementation with custom parameters
            tournament = new TournamentGame{salt: salt}(_matchDuration, _minBondAmount);
        }

        // Initialize the tournament
        tournament.initialize(_gameCreator, _rootClaim, _l1Head, _extraData);

        // Store the tournament data
        tournamentData[address(tournament)] = abi.encode(_gameCreator, _rootClaim, _l1Head, _extraData);
        tournaments.push(address(tournament));

        // Register the tournament with the DisputeGameFactory
        _registerWithDisputeGameFactory(address(tournament));

        emit TournamentCreated(address(tournament), _gameCreator, _rootClaim, _l1Head);

        return address(tournament);
    }

    /**
     * @notice Registers a tournament with the DisputeGameFactory
     * @param _tournament Address of the tournament to register
     */
    function _registerWithDisputeGameFactory(address _tournament) internal {
        // Call the register function on the DisputeGameFactory
        // The exact interface may vary based on the DisputeGameFactory implementation
        // This is a placeholder implementation
        (bool success, ) = DISPUTE_GAME_FACTORY.call(
            abi.encodeWithSignature("registerGame(address)", _tournament)
        );
        require(success, "Failed to register with DisputeGameFactory");
    }
}
