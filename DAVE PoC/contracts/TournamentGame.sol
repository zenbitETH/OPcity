// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title ITournamentGame Interface
 * @notice Interface for tournament-based dispute games in the Cartesi DAVE fraud proofs system
 * @dev Extends the IDisputeGame interface with tournament-specific functionality
 */
interface IDisputeGame {
    /**
     * @notice Enum representing the status of the dispute game
     */
    enum GameStatus {
        IN_PROGRESS,
        CHALLENGER_WINS,
        DEFENDER_WINS
    }

    /**
     * @notice Emitted when the dispute game is resolved
     * @param status The status of the game after resolution
     */
    event Resolved(GameStatus indexed status);

    /**
     * @notice Returns the timestamp when the dispute game was created
     * @return The timestamp when the dispute game was created
     */
    function createdAt() external view returns (uint256);

    /**
     * @notice Returns the timestamp when the dispute game was resolved
     * @return The timestamp when the dispute game was resolved
     */
    function resolvedAt() external view returns (uint256);

    /**
     * @notice Returns the current status of the dispute game
     * @return The current status of the dispute game
     */
    function status() external view returns (GameStatus);

    /**
     * @notice Returns the type of the dispute game
     * @return The type of the dispute game
     */
    function gameType() external pure returns (uint8);

    /**
     * @notice Returns the address that created the dispute game
     * @return The address that created the dispute game
     */
    function gameCreator() external view returns (address);

    /**
     * @notice Returns the root claim of the dispute game
     * @return The root claim of the dispute game
     */
    function rootClaim() external view returns (bytes32);

    /**
     * @notice Returns the L1 head hash at the time the dispute game was created
     * @return The L1 head hash at the time the dispute game was created
     */
    function l1Head() external view returns (bytes32);

    /**
     * @notice Returns extra data supplied to the dispute game
     * @return Extra data supplied to the dispute game
     */
    function extraData() external view returns (bytes memory);

    /**
     * @notice Returns the game type, root claim, and extra data
     * @return The game type, root claim, and extra data
     */
    function gameData() external view returns (uint8, bytes32, bytes memory);

    /**
     * @notice Resolves the dispute game
     * @return The status of the game after resolution
     */
    function resolve() external returns (GameStatus);
}

/**
 * @title ITournamentGame Interface
 * @notice Interface for tournament-based dispute games in the Cartesi DAVE fraud proofs system
 * @dev Extends the IDisputeGame interface with tournament-specific functionality
 */
interface ITournamentGame is IDisputeGame {
    /**
     * @notice Struct representing a node in the tournament tree
     * @param participant Address of the participant
     * @param claim The claim made by the participant
     * @param hasJoined Whether the participant has officially joined
     * @param bondAmount Amount of ETH bonded by the participant
     */
    struct Node {
        address participant;
        bytes32 claim;
        bool hasJoined;
        uint256 bondAmount;
    }

    /**
     * @notice Struct representing a match between two participants
     * @param nodeA Index of the first participant node
     * @param nodeB Index of the second participant node
     * @param winner Index of the winning node
     * @param startTime Timestamp when the match started
     * @param resolved Whether the match has been resolved
     */
    struct Match {
        uint256 nodeA;
        uint256 nodeB;
        uint256 winner;
        uint256 startTime;
        bool resolved;
    }

    /**
     * @notice Emitted when the tournament is started
     * @param rootClaim The root claim being disputed
     * @param gameCreator The address that created the tournament
     */
    event TournamentStarted(bytes32 indexed rootClaim, address indexed gameCreator);

    /**
     * @notice Emitted when a participant joins the tournament
     * @param participant Address of the participant
     * @param nodeIndex Index of the node in the tournament tree
     * @param claim The counter-claim made by the participant
     */
    event ParticipantJoined(address indexed participant, uint256 indexed nodeIndex, bytes32 claim);

    /**
     * @notice Emitted when a match is created
     * @param matchIndex Index of the match
     * @param nodeA Index of the first participant node
     * @param nodeB Index of the second participant node
     */
    event MatchCreated(uint256 indexed matchIndex, uint256 nodeA, uint256 nodeB);

    /**
     * @notice Emitted when a match is resolved
     * @param matchIndex Index of the match
     * @param winner Index of the winning node
     */
    event MatchResolved(uint256 indexed matchIndex, uint256 indexed winner);

    /**
     * @notice Emitted when a bond is claimed by the tournament winner
     * @param winner Address of the winner
     * @param amount Amount of ETH claimed
     */
    event BondClaimed(address indexed winner, uint256 amount);

    /**
     * @notice Emitted when the tournament advances to a new round
     * @param newRound The new round number
     */
    event RoundAdvanced(uint256 indexed newRound);

    /**
     * @notice Error thrown when attempting to perform an action on an already resolved tournament
     */
    error TournamentAlreadyResolved();

    /**
     * @notice Error thrown when an invalid participant attempts to join
     */
    error InvalidParticipant();

    /**
     * @notice Error thrown when the bond amount is insufficient
     */
    error InsufficientBond();

    /**
     * @notice Error thrown when a participant attempts to join multiple times
     */
    error AlreadyParticipated();

    /**
     * @notice Error thrown when attempting to resolve a match that hasn't started
     */
    error MatchNotStarted();

    /**
     * @notice Error thrown when attempting to resolve an already resolved match
     */
    error MatchAlreadyResolved();

    /**
     * @notice Error thrown when attempting to resolve a match that is still in progress
     */
    error MatchInProgress();

    /**
     * @notice Error thrown when a non-participant attempts to resolve a match
     */
    error NotMatchParticipant();

    /**
     * @notice Error thrown when attempting to claim a bond before the tournament is resolved
     */
    error TournamentNotResolved();

    /**
     * @notice Error thrown when a non-winner attempts to claim the bond
     */
    error NotTournamentWinner();

    /**
     * @notice Join the tournament with a counter-claim
     * @param claim The counter-claim being made
     * @dev Must send sufficient ETH to cover the bond amount
     */
    function joinTournament(bytes32 claim) external payable;

    /**
     * @notice Resolve a match in the tournament
     * @param matchIndex The index of the match to resolve
     * @param winnerIndex The index of the winning node
     * @dev Can only be called by a participant in the match after the match duration has passed
     */
    function resolveMatch(uint256 matchIndex, uint256 winnerIndex) external;

    /**
     * @notice Claim the bond as the tournament winner
     * @dev Can only be called by the tournament winner after the tournament is resolved
     */
    function claimBond() external;

    /**
     * @notice Get information about a node in the tournament tree
     * @param index The index of the node
     * @return The node information (participant, claim, hasJoined, bondAmount)
     */
    function nodes(uint256 index) external view returns (Node memory);

    /**
     * @notice Get information about a match in the tournament
     * @param index The index of the match
     * @return The match information (nodeA, nodeB, winner, startTime, resolved)
     */
    function matches(uint256 index) external view returns (Match memory);

    /**
     * @notice Get the current round of the tournament
     * @return The current round number
     */
    function currentRound() external view returns (uint256);

    /**
     * @notice Get the total number of participants in the tournament
     * @return The total number of participants
     */
    function totalParticipants() external view returns (uint256);

    /**
     * @notice Check if an address has participated in the tournament
     * @param participant The address to check
     * @return True if the address has participated, false otherwise
     */
    function hasParticipated(address participant) external view returns (bool);

    /**
     * @notice Get the start time of a specific round
     * @param round The round number
     * @return The timestamp when the round started
     */
    function roundStartTime(uint256 round) external view returns (uint256);
}

/**
 * @title TournamentGame
 * @notice Implementation of a tournament-based dispute game for the Cartesi DAVE fraud proofs system
 * @dev This contract implements the ITournamentGame interface and manages a tournament tree structure
 *      for dispute resolution. It allows participants to join with counter-claims, handles match
 *      resolution, and determines the final result of the dispute.
 */
contract TournamentGame is ITournamentGame, Initializable, ReentrancyGuard {
    /**
     * @notice Duration of a match in seconds
     * @dev This is the minimum time that must pass before a match can be resolved
     */
    uint256 public immutable MATCH_DURATION;

    /**
     * @notice Minimum bond amount required to join the tournament
     */
    uint256 public immutable MIN_BOND_AMOUNT;

    /**
     * @notice Timestamp when the tournament was created
     */
    uint256 private _createdAt;

    /**
     * @notice Timestamp when the tournament was resolved
     */
    uint256 private _resolvedAt;

    /**
     * @notice Current status of the dispute game
     */
    GameStatus private _status;

    /**
     * @notice Address that created the dispute game
     */
    address private _gameCreator;

    /**
     * @notice Root claim of the dispute game
     */
    bytes32 private _rootClaim;

    /**
     * @notice L1 head hash at the time the dispute game was created
     */
    bytes32 private _l1Head;

    /**
     * @notice Extra data supplied to the dispute game
     */
    bytes private _extraData;

    /**
     * @notice Current round of the tournament
     */
    uint256 private _currentRound;

    /**
     * @notice Total number of participants in the tournament
     */
    uint256 private _totalParticipants;

    /**
     * @notice Array of nodes in the tournament tree
     */
    Node[] private _nodes;

    /**
     * @notice Array of matches in the tournament
     */
    Match[] private _matches;

    /**
     * @notice Mapping of round start times
     * @dev Maps round number to the timestamp when the round started
     */
    mapping(uint256 => uint256) private _roundStartTimes;

    /**
     * @notice Mapping to track if an address has participated in the tournament
     * @dev Maps participant address to a boolean indicating participation
     */
    mapping(address => bool) private _hasParticipated;

    /**
     * @notice Mapping to track the node index for each participant
     * @dev Maps participant address to their node index in the tournament tree
     */
    mapping(address => uint256) private _participantNodeIndex;

    /**
     * @notice Modifier to ensure the tournament is not resolved
     */
    modifier tournamentNotResolved() {
        if (_resolvedAt != 0) {
            revert TournamentAlreadyResolved();
        }
        _;
    }

    /**
     * @notice Modifier to ensure the tournament is resolved
     */
    modifier tournamentResolved() {
        if (_resolvedAt == 0) {
            revert TournamentNotResolved();
        }
        _;
    }

    /**
     * @notice Constructor for the TournamentGame contract
     * @param matchDuration Duration of a match in seconds
     * @param minBondAmount Minimum bond amount required to join the tournament
     */
    constructor(uint256 matchDuration, uint256 minBondAmount) {
        MATCH_DURATION = matchDuration;
        MIN_BOND_AMOUNT = minBondAmount;
        _disableInitializers();
    }

    /**
     * @notice Initializes the tournament game
     * @param gameCreator Address that created the dispute game
     * @param rootClaim Root claim of the dispute game
     * @param l1Head L1 head hash at the time the dispute game was created
     * @param extraData Extra data supplied to the dispute game
     */
    function initialize(
        address gameCreator,
        bytes32 rootClaim,
        bytes32 l1Head,
        bytes calldata extraData
    ) external initializer {
        _gameCreator = gameCreator;
        _rootClaim = rootClaim;
        _l1Head = l1Head;
        _extraData = extraData;
        _createdAt = block.timestamp;
        _status = GameStatus.IN_PROGRESS;
        _currentRound = 0;
        _roundStartTimes[0] = block.timestamp;

        // Create the root node with the root claim
        _nodes.push(Node({
            participant: gameCreator,
            claim: rootClaim,
            hasJoined: true,
            bondAmount: 0
        }));
        _totalParticipants = 1;
        _hasParticipated[gameCreator] = true;
        _participantNodeIndex[gameCreator] = 0;

        emit TournamentStarted(rootClaim, gameCreator);
        emit RoundAdvanced(0);
    }

    /**
     * @notice Join the tournament with a counter-claim
     * @param claim The counter-claim being made
     * @dev Must send sufficient ETH to cover the bond amount
     */
    function joinTournament(bytes32 claim) external payable tournamentNotResolved nonReentrant {
        // Ensure the sender is not the game creator
        if (msg.sender == _gameCreator) {
            revert InvalidParticipant();
        }

        // Ensure the sender has not already participated
        if (_hasParticipated[msg.sender]) {
            revert AlreadyParticipated();
        }

        // Ensure the bond amount is sufficient
        if (msg.value < MIN_BOND_AMOUNT) {
            revert InsufficientBond();
        }

        // Create a new node for the participant
        uint256 nodeIndex = _nodes.length;
        _nodes.push(Node({
            participant: msg.sender,
            claim: claim,
            hasJoined: true,
            bondAmount: msg.value
        }));

        // Update participant tracking
        _totalParticipants++;
        _hasParticipated[msg.sender] = true;
        _participantNodeIndex[msg.sender] = nodeIndex;

        // Create a match if there's an available opponent
        _createMatchIfPossible(nodeIndex);

        emit ParticipantJoined(msg.sender, nodeIndex, claim);
    }

    /**
     * @notice Resolve a match in the tournament
     * @param matchIndex The index of the match to resolve
     * @param winnerIndex The index of the winning node
     * @dev Can only be called by a participant in the match after the match duration has passed
     */
    function resolveMatch(uint256 matchIndex, uint256 winnerIndex) external tournamentNotResolved nonReentrant {
        // Ensure the match exists
        if (matchIndex >= _matches.length) {
            revert MatchNotStarted();
        }

        Match storage currentMatch = _matches[matchIndex];

        // Ensure the match is not already resolved
        if (currentMatch.resolved) {
            revert MatchAlreadyResolved();
        }

        // Ensure the match duration has passed
        if (block.timestamp < currentMatch.startTime + MATCH_DURATION) {
            revert MatchInProgress();
        }

        // Ensure the caller is a participant in the match
        if (msg.sender != _nodes[currentMatch.nodeA].participant && msg.sender != _nodes[currentMatch.nodeB].participant) {
            revert NotMatchParticipant();
        }

        // Ensure the winner is a participant in the match
        if (winnerIndex != currentMatch.nodeA && winnerIndex != currentMatch.nodeB) {
            revert NotMatchParticipant();
        }

        // Resolve the match
        currentMatch.winner = winnerIndex;
        currentMatch.resolved = true;

        // Update the tournament tree
        _updateParent(matchIndex, winnerIndex);

        emit MatchResolved(matchIndex, winnerIndex);

        // Check if the tournament is resolved
        _checkTournamentResolution();
    }

    /**
     * @notice Claim the bond as the tournament winner
     * @dev Can only be called by the tournament winner after the tournament is resolved
     */
    function claimBond() external tournamentResolved nonReentrant {
        // Ensure the tournament has a winner
        if (_status != GameStatus.CHALLENGER_WINS && _status != GameStatus.DEFENDER_WINS) {
            revert TournamentNotResolved();
        }

        // Determine the winner address
        address winner;
        if (_status == GameStatus.DEFENDER_WINS) {
            winner = _gameCreator;
        } else {
            // Find the challenger who won
            for (uint256 i = 1; i < _nodes.length; i++) {
                if (_nodes[i].participant != address(0) && _nodes[i].hasJoined) {
                    bool isWinner = true;
                    for (uint256 j = 0; j < _matches.length; j++) {
                        Match storage currentMatch = _matches[j];
                        if (currentMatch.resolved && (currentMatch.nodeA == i || currentMatch.nodeB == i) && currentMatch.winner != i) {
                            isWinner = false;
                            break;
                        }
                    }
                    if (isWinner) {
                        winner = _nodes[i].participant;
                        break;
                    }
                }
            }
        }

        // Ensure the caller is the winner
        if (msg.sender != winner) {
            revert NotTournamentWinner();
        }

        // Calculate the total bond amount
        uint256 totalBond = 0;
        for (uint256 i = 1; i < _nodes.length; i++) {
            totalBond += _nodes[i].bondAmount;
        }

        // Reset bond amounts to prevent re-entrancy
        for (uint256 i = 1; i < _nodes.length; i++) {
            _nodes[i].bondAmount = 0;
        }

        // Transfer the bond to the winner
        if (totalBond > 0) {
            Address.sendValue(payable(winner), totalBond);
            emit BondClaimed(winner, totalBond);
        }
    }

    /**
     * @notice Resolves the dispute game
     * @return The status of the game after resolution
     */
    function resolve() external tournamentNotResolved nonReentrant returns (GameStatus) {
        // Check if all matches are resolved
        for (uint256 i = 0; i < _matches.length; i++) {
            if (!_matches[i].resolved) {
                // If any match is not resolved, check if it has timed out
                if (block.timestamp >= _matches[i].startTime + MATCH_DURATION * 2) {
                    // If the match has timed out, resolve it with a default winner
                    // The default winner is the defender (node A)
                    _matches[i].winner = _matches[i].nodeA;
                    _matches[i].resolved = true;
                    _updateParent(i, _matches[i].nodeA);
                    emit MatchResolved(i, _matches[i].nodeA);
                } else {
                    // If the match has not timed out, the tournament cannot be resolved yet
                    return _status;
                }
            }
        }

        // Determine the winner
        bool defenderWins = true;
        for (uint256 i = 0; i < _matches.length; i++) {
            Match storage currentMatch = _matches[i];
            if (currentMatch.resolved) {
                Node storage winnerNode = _nodes[currentMatch.winner];
                if (winnerNode.participant != _gameCreator) {
                    defenderWins = false;
                    break;
                }
            }
        }

        // Set the status and resolved timestamp
        _status = defenderWins ? GameStatus.DEFENDER_WINS : GameStatus.CHALLENGER_WINS;
        _resolvedAt = block.timestamp;

        emit Resolved(_status);
        return _status;
    }

    /**
     * @notice Get information about a node in the tournament tree
     * @param index The index of the node
     * @return The node information (participant, claim, hasJoined, bondAmount)
     */
    function nodes(uint256 index) external view returns (Node memory) {
        require(index < _nodes.length, "Invalid node index");
        return _nodes[index];
    }

    /**
     * @notice Get information about a match in the tournament
     * @param index The index of the match
     * @return The match information (nodeA, nodeB, winner, startTime, resolved)
     */
    function matches(uint256 index) external view returns (Match memory) {
        require(index < _matches.length, "Invalid match index");
        return _matches[index];
    }

    /**
     * @notice Get the current round of the tournament
     * @return The current round number
     */
    function currentRound() external view returns (uint256) {
        return _currentRound;
    }

    /**
     * @notice Get the total number of participants in the tournament
     * @return The total number of participants
     */
    function totalParticipants() external view returns (uint256) {
        return _totalParticipants;
    }

    /**
     * @notice Check if an address has participated in the tournament
     * @param participant The address to check
     * @return True if the address has participated, false otherwise
     */
    function hasParticipated(address participant) external view returns (bool) {
        return _hasParticipated[participant];
    }

    /**
     * @notice Get the start time of a specific round
     * @param round The round number
     * @return The timestamp when the round started
     */
    function roundStartTime(uint256 round) external view returns (uint256) {
        return _roundStartTimes[round];
    }

    /**
     * @notice Returns the timestamp when the dispute game was created
     * @return The timestamp when the dispute game was created
     */
    function createdAt() external view returns (uint256) {
        return _createdAt;
    }

    /**
     * @notice Returns the timestamp when the dispute game was resolved
     * @return The timestamp when the dispute game was resolved
     */
    function resolvedAt() external view returns (uint256) {
        return _resolvedAt;
    }

    /**
     * @notice Returns the current status of the dispute game
     * @return The current status of the dispute game
     */
    function status() external view returns (GameStatus) {
        return _status;
    }

    /**
     * @notice Returns the type of the dispute game
     * @return The type of the dispute game
     */
    function gameType() external pure returns (uint8) {
        return 3; // Assuming 3 is the type for TournamentGame
    }

    /**
     * @notice Returns the address that created the dispute game
     * @return The address that created the dispute game
     */
    function gameCreator() external view returns (address) {
        return _gameCreator;
    }

    /**
     * @notice Returns the root claim of the dispute game
     * @return The root claim of the dispute game
     */
    function rootClaim() external view returns (bytes32) {
        return _rootClaim;
    }

    /**
     * @notice Returns the L1 head hash at the time the dispute game was created
     * @return The L1 head hash at the time the dispute game was created
     */
    function l1Head() external view returns (bytes32) {
        return _l1Head;
    }

    /**
     * @notice Returns extra data supplied to the dispute game
     * @return Extra data supplied to the dispute game
     */
    function extraData() external view returns (bytes memory) {
        return _extraData;
    }

    /**
     * @notice Returns the game type, root claim, and extra data
     * @return The game type, root claim, and extra data
     */
    function gameData() external view returns (uint8, bytes32, bytes memory) {
        return (gameType(), _rootClaim, _extraData);
    }

    /**
     * @notice Creates a match if there's an available opponent
     * @param nodeIndex The index of the node to create a match for
     */
    function _createMatchIfPossible(uint256 nodeIndex) internal {
        // Find an opponent for the new participant
        for (uint256 i = 0; i < _nodes.length; i++) {
            if (i != nodeIndex && _nodes[i].hasJoined) {
                bool alreadyMatched = false;
                for (uint256 j = 0; j < _matches.length; j++) {
                    Match storage currentMatch = _matches[j];
                    if (!currentMatch.resolved && (
                        (currentMatch.nodeA == i && currentMatch.nodeB == nodeIndex) ||
                        (currentMatch.nodeA == nodeIndex && currentMatch.nodeB == i)
                    )) {
                        alreadyMatched = true;
                        break;
                    }
                }

                if (!alreadyMatched) {
                    // Create a new match
                    uint256 matchIndex = _matches.length;
                    _matches.push(Match({
                        nodeA: i,
                        nodeB: nodeIndex,
                        winner: 0,
                        startTime: block.timestamp,
                        resolved: false
                    }));

                    emit MatchCreated(matchIndex, i, nodeIndex);

                    // Check if we need to advance to a new round
                    if (_currentRound == 0 || _matches.length >= (1 << _currentRound)) {
                        _currentRound++;
                        _roundStartTimes[_currentRound] = block.timestamp;
                        emit RoundAdvanced(_currentRound);
                    }

                    break;
                }
            }
        }
    }

    /**
     * @notice Updates the parent node in the tournament tree
     * @param matchIndex The index of the resolved match
     * @param winnerIndex The index of the winning node
     */
    function _updateParent(uint256 matchIndex, uint256 winnerIndex) internal {
        // In a tournament tree, the winner advances to the next round
        // This is a simplified implementation that doesn't create a full binary tree
        // Instead, it creates matches as participants join and tracks the winners
    }

    /**
     * @notice Checks if the tournament is resolved
     * @dev If all matches are resolved and there's a clear winner, sets the status and resolved timestamp
     */
    function _checkTournamentResolution() internal {
        // Check if all matches are resolved
        bool allResolved = true;
        for (uint256 i = 0; i < _matches.length; i++) {
            if (!_matches[i].resolved) {
                allResolved = false;
                break;
            }
        }

        if (allResolved) {
            // Determine the winner
            bool defenderWins = true;
            for (uint256 i = 0; i < _matches.length; i++) {
                Match storage currentMatch = _matches[i];
                Node storage winnerNode = _nodes[currentMatch.winner];
                if (winnerNode.participant != _gameCreator) {
                    defenderWins = false;
                    break;
                }
            }

            // Set the status and resolved timestamp
            _status = defenderWins ? GameStatus.DEFENDER_WINS : GameStatus.CHALLENGER_WINS;
            _resolvedAt = block.timestamp;

            emit Resolved(_status);
        }
    }

    /**
     * @notice Checks if this contract supports a given interface
     * @param interfaceId The interface identifier to check
     * @return True if the interface is supported, false otherwise
     */
    function supportsInterface(bytes4 interfaceId) external pure returns (bool) {
        return
            interfaceId == type(ITournamentGame).interfaceId ||
            interfaceId == 0x01ffc9a7; // ERC165 Interface ID
    }
}
