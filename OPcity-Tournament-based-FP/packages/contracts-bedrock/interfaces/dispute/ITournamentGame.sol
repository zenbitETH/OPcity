// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { IDisputeGame } from "./IDisputeGame.sol";

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
