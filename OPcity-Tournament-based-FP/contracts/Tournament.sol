
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

/**
 * @title Tournament Contract for Cartesi DAVE Fraud Proofs
 * @notice This contract implements a tournament tree structure for dispute resolution
 * @dev Implements the IDisputeGame interface for compatibility with OP Stack
 */

/**
 * @title IDisputeGame Interface
 * @notice Interface for dispute games in the OP Stack
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
 * @title Tournament Contract for Cartesi DAVE Fraud Proofs
 * @notice This contract implements a tournament tree structure for dispute resolution
 * @dev Implements the IDisputeGame interface for compatibility with OP Stack
 */
contract Tournament is IDisputeGame {
    // Constants
    uint8 private constant GAME_TYPE = 3; // DAVE Tournament type
    uint256 private constant MATCH_DURATION = 1 days; // Duration for each match
    uint256 private constant BOND_AMOUNT = 0.1 ether; // Bond amount required to join

    // Tournament tree structure
    struct Node {
        address participant;
        bytes32 claim;
        bool hasJoined;
        uint256 bondAmount;
    }

    struct Match {
        uint256 nodeA;
        uint256 nodeB;
        uint256 winner;
        uint256 startTime;
        bool resolved;
    }

    // Tournament state variables
    bytes32 private immutable _rootClaim;
    bytes32 private immutable _l1Head;
    address private immutable _gameCreator;
    bytes private _extraData;
    uint256 private immutable _createdAt;
    uint256 private _resolvedAt;
    GameStatus private _status;

    // Tournament specific variables
    Node[] public nodes;
    Match[] public matches;
    uint256 public currentRound;
    uint256 public totalParticipants;
    mapping(address => bool) public hasParticipated;
    mapping(uint256 => uint256) public roundStartTime;

    // Events
    event TournamentStarted(bytes32 rootClaim, address gameCreator);
    event ParticipantJoined(address participant, uint256 nodeIndex, bytes32 claim);
    event MatchCreated(uint256 matchIndex, uint256 nodeA, uint256 nodeB);
    event MatchResolved(uint256 matchIndex, uint256 winner);
    event BondClaimed(address winner, uint256 amount);
    event RoundAdvanced(uint256 newRound);

    // Errors
    error TournamentAlreadyResolved();
    error InvalidParticipant();
    error InsufficientBond();
    error AlreadyParticipated();
    error MatchNotStarted();
    error MatchAlreadyResolved();
    error MatchInProgress();
    error NotMatchParticipant();
    error TournamentNotResolved();
    error NotTournamentWinner();

    /**
     * @notice Constructor for the Tournament contract
     * @param initialRootClaim The root claim being disputed
     * @param initialL1Head The L1 head hash at creation time
     * @param initialExtraData Additional data for the tournament
     */
    constructor(bytes32 initialRootClaim, bytes32 initialL1Head, bytes memory initialExtraData) {
        _rootClaim = initialRootClaim;
        _l1Head = initialL1Head;
        _gameCreator = msg.sender;
        _extraData = initialExtraData;
        _createdAt = block.timestamp;
        _status = GameStatus.IN_PROGRESS;

        // Initialize the tournament with the root node (defender)
        nodes.push(Node({
            participant: msg.sender,
            claim: initialRootClaim,
            hasJoined: true,
            bondAmount: 0 // Creator doesn't need to bond
        }));
        totalParticipants = 1;

        // Initialize the first round
        roundStartTime[currentRound] = block.timestamp;

        emit TournamentStarted(initialRootClaim, msg.sender);
    }

    /**
     * @notice Join the tournament with a counter-claim
     * @param claim The counter-claim being made
     */
    function joinTournament(bytes32 claim) external payable {
        if (_status != GameStatus.IN_PROGRESS) revert TournamentAlreadyResolved();
        if (msg.sender == address(0)) revert InvalidParticipant();
        if (msg.value < BOND_AMOUNT) revert InsufficientBond();
        if (hasParticipated[msg.sender]) revert AlreadyParticipated();

        // Add the participant to the tournament
        uint256 nodeIndex = nodes.length;
        nodes.push(Node({
            participant: msg.sender,
            claim: claim,
            hasJoined: true,
            bondAmount: msg.value
        }));
        hasParticipated[msg.sender] = true;
        totalParticipants++;

        emit ParticipantJoined(msg.sender, nodeIndex, claim);

        // If we have a power of 2 number of participants, create matches for the new round
        if (isPowerOfTwo(totalParticipants)) {
            createMatchesForNewRound();
        }
    }

    /**
     * @notice Resolve a match in the tournament
     * @param matchIndex The index of the match to resolve
     * @param winnerIndex The index of the winning node
     */
    function resolveMatch(uint256 matchIndex, uint256 winnerIndex) external {
        if (_status != GameStatus.IN_PROGRESS) revert TournamentAlreadyResolved();
        if (matchIndex >= matches.length) revert MatchNotStarted();
        
        Match storage currentMatch = matches[matchIndex];
        if (currentMatch.resolved) revert MatchAlreadyResolved();
        if (block.timestamp < currentMatch.startTime + MATCH_DURATION) revert MatchInProgress();
        
        // Verify the caller is a participant in the match
        if (msg.sender != nodes[currentMatch.nodeA].participant && 
            msg.sender != nodes[currentMatch.nodeB].participant) {
            revert NotMatchParticipant();
        }
        
        // Verify the winner is part of the match
        if (winnerIndex != currentMatch.nodeA && winnerIndex != currentMatch.nodeB) {
            revert NotMatchParticipant();
        }
        
        // Resolve the match
        currentMatch.winner = winnerIndex;
        currentMatch.resolved = true;
        
        emit MatchResolved(matchIndex, winnerIndex);
        
        // Check if all matches in the current round are resolved
        bool allMatchesResolved = true;
        for (uint256 i = 0; i < matches.length; i++) {
            if (!matches[i].resolved) {
                allMatchesResolved = false;
                break;
            }
        }
        
        // If all matches are resolved, advance to the next round or resolve the tournament
        if (allMatchesResolved) {
            if (matches.length == 1) {
                // Final match resolved, tournament is over
                _resolvedAt = block.timestamp;
                _status = determineWinner();
                emit Resolved(_status);
            } else {
                // Create matches for the next round
                currentRound++;
                roundStartTime[currentRound] = block.timestamp;
                createMatchesForNextRound();
                emit RoundAdvanced(currentRound);
            }
        }
    }

    /**
     * @notice Claim the bond as the tournament winner
     */
    function claimBond() external {
        if (_status == GameStatus.IN_PROGRESS) revert TournamentNotResolved();
        
        // Determine the winner address based on the final status
        address winner;
        if (_status == GameStatus.DEFENDER_WINS) {
            winner = nodes[0].participant; // Root claim defender
        } else {
            // Find the challenger who won
            winner = findWinningChallenger();
        }
        
        if (msg.sender != winner) revert NotTournamentWinner();
        
        // Calculate total bond amount
        uint256 totalBond = 0;
        for (uint256 i = 1; i < nodes.length; i++) { // Skip the defender (index 0)
            totalBond += nodes[i].bondAmount;
        }
        
        // Transfer the bond to the winner
        payable(winner).transfer(totalBond);
        
        emit BondClaimed(winner, totalBond);
    }

    /**
     * @notice Find the winning challenger
     * @return The address of the winning challenger
     */
    function findWinningChallenger() internal view returns (address) {
        // In a tournament, the winner is the participant of the final match winner
        if (matches.length > 0) {
            Match storage finalMatch = matches[matches.length - 1];
            if (finalMatch.resolved) {
                return nodes[finalMatch.winner].participant;
            }
        }
        
        // Fallback to the first challenger if no matches were played
        return nodes.length > 1 ? nodes[1].participant : address(0);
    }

    /**
     * @notice Create matches for a new round when the number of participants is a power of 2
     */
    function createMatchesForNewRound() internal {
        // Clear previous matches
        delete matches;
        
        // Create matches for the current round
        for (uint256 i = 0; i < totalParticipants; i += 2) {
            uint256 matchIndex = matches.length;
            matches.push(Match({
                nodeA: i,
                nodeB: i + 1,
                winner: 0,
                startTime: block.timestamp,
                resolved: false
            }));
            
            emit MatchCreated(matchIndex, i, i + 1);
        }
    }

    /**
     * @notice Create matches for the next round based on winners of the current round
     */
    function createMatchesForNextRound() internal {
        uint256[] memory winners = new uint256[](matches.length);
        
        // Collect winners from the current round
        for (uint256 i = 0; i < matches.length; i++) {
            winners[i] = matches[i].winner;
        }
        
        // Clear previous matches
        delete matches;
        
        // Create new matches with winners
        for (uint256 i = 0; i < winners.length; i += 2) {
            if (i + 1 < winners.length) {
                uint256 matchIndex = matches.length;
                matches.push(Match({
                    nodeA: winners[i],
                    nodeB: winners[i + 1],
                    winner: 0,
                    startTime: block.timestamp,
                    resolved: false
                }));
                
                emit MatchCreated(matchIndex, winners[i], winners[i + 1]);
            }
        }
    }

    /**
     * @notice Determine the final winner of the tournament
     * @return The final status of the game
     */
    function determineWinner() internal view returns (GameStatus) {
        if (matches.length == 0) {
            // No matches played, defender wins by default
            return GameStatus.DEFENDER_WINS;
        }
        
        // Get the final match
        Match storage finalMatch = matches[matches.length - 1];
        
        // If the winner is the root node (defender), defender wins
        if (finalMatch.winner == 0) {
            return GameStatus.DEFENDER_WINS;
        }
        
        // Otherwise, challenger wins
        return GameStatus.CHALLENGER_WINS;
    }

    /**
     * @notice Check if a number is a power of 2
     * @param x The number to check
     * @return True if the number is a power of 2, false otherwise
     */
    function isPowerOfTwo(uint256 x) internal pure returns (bool) {
        return x > 0 && (x & (x - 1)) == 0;
    }

    /**
     * @notice Returns the timestamp when the dispute game was created
     * @return The timestamp when the dispute game was created
     */
    function createdAt() external view override returns (uint256) {
        return _createdAt;
    }

    /**
     * @notice Returns the timestamp when the dispute game was resolved
     * @return The timestamp when the dispute game was resolved
     */
    function resolvedAt() external view override returns (uint256) {
        return _resolvedAt;
    }

    /**
     * @notice Returns the current status of the dispute game
     * @return The current status of the dispute game
     */
    function status() external view override returns (GameStatus) {
        return _status;
    }

    /**
     * @notice Returns the type of the dispute game
     * @return The type of the dispute game (3 for DAVE Tournament)
     */
    function gameType() external pure override returns (uint8) {
        return GAME_TYPE;
    }

    /**
     * @notice Returns the address that created the dispute game
     * @return The address that created the dispute game
     */
    function gameCreator() external view override returns (address) {
        return _gameCreator;
    }

    /**
     * @notice Returns the root claim of the dispute game
     * @return The root claim of the dispute game
     */
    function rootClaim() external view override returns (bytes32) {
        return _rootClaim;
    }

    /**
     * @notice Returns the L1 head hash at the time the dispute game was created
     * @return The L1 head hash at the time the dispute game was created
     */
    function l1Head() external view override returns (bytes32) {
        return _l1Head;
    }

    /**
     * @notice Returns extra data supplied to the dispute game
     * @return Extra data supplied to the dispute game
     */
    function extraData() external view override returns (bytes memory) {
        return _extraData;
    }

    /**
     * @notice Returns the game type, root claim, and extra data
     * @return The game type, root claim, and extra data
     */
    function gameData() external view override returns (uint8, bytes32, bytes memory) {
        return (GAME_TYPE, _rootClaim, _extraData);
    }

    /**
     * @notice Resolves the dispute game
     * @return The status of the game after resolution
     */
    function resolve() external override returns (GameStatus) {
        if (_status != GameStatus.IN_PROGRESS) {
            return _status;
        }
        
        // If all matches are resolved, determine the winner
        bool allMatchesResolved = true;
        for (uint256 i = 0; i < matches.length; i++) {
            if (!matches[i].resolved) {
                allMatchesResolved = false;
                break;
            }
        }
        
        if (allMatchesResolved) {
            _resolvedAt = block.timestamp;
            _status = determineWinner();
            emit Resolved(_status);
        }
        
        return _status;
    }
}
