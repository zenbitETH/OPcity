// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

/**
 * @title IMultiPhaseResolver
 * @notice Interface for the Multi-Phase Resolver that coordinates dispute resolution across phases
 */
interface IMultiPhaseResolver {
    /**
     * @notice Enum representing the possible phases of a dispute
     */
    enum DisputePhase {
        INVALID,
        PHASE_1,
        PHASE_2,
        RESOLVED
    }

    /**
     * @notice Enum representing the status of a dispute
     */
    enum DisputeStatus {
        INVALID,
        ACTIVE,
        RESOLVING,
        RESOLVED
    }

    /**
     * @notice Event emitted when a new dispute is created
     * @param disputeId The ID of the created dispute
     * @param proposer The address that proposed the claim
     * @param challenger The address that challenged the claim
     * @param claimedStateRoot The state root claimed by the proposer
     * @param challengedStateRoot The state root claimed by the challenger
     */
    event DisputeCreated(
        uint256 indexed disputeId,
        address indexed proposer,
        address indexed challenger,
        bytes32 claimedStateRoot,
        bytes32 challengedStateRoot
    );

    /**
     * @notice Event emitted when a dispute phase changes
     * @param disputeId The ID of the dispute
     * @param oldPhase The previous phase
     * @param newPhase The new phase
     */
    event PhaseTransition(
        uint256 indexed disputeId,
        DisputePhase oldPhase,
        DisputePhase newPhase
    );

    /**
     * @notice Event emitted when a dispute is resolved
     * @param disputeId The ID of the resolved dispute
     * @param winner The address of the winning party
     * @param winningStateRoot The state root that was determined to be correct
     */
    event DisputeResolved(
        uint256 indexed disputeId,
        address winner,
        bytes32 winningStateRoot
    );

    /**
     * @notice Creates a new dispute
     * @param claimedStateRoot The state root claimed by the proposer
     * @param challengedStateRoot The state root claimed by the challenger
     * @return disputeId The ID of the created dispute
     */
    function createDispute(
        bytes32 claimedStateRoot,
        bytes32 challengedStateRoot
    ) external returns (uint256 disputeId);

    /**
     * @notice Gets the current phase of a dispute
     * @param disputeId The ID of the dispute
     * @return The current phase of the dispute
     */
    function getCurrentPhase(
        uint256 disputeId
    ) external view returns (DisputePhase);

    /**
     * @notice Transitions a dispute from Phase 1 to Phase 2
     * @param disputeId The ID of the dispute
     * @param segmentStart The start index of the disputed segment
     * @param segmentEnd The end index of the disputed segment
     * @param segmentStateRoot The state root at the segment boundary
     */
    function transitionToPhase2(
        uint256 disputeId,
        uint256 segmentStart,
        uint256 segmentEnd,
        bytes32 segmentStateRoot
    ) external;

    /**
     * @notice Resolves a dispute
     * @param disputeId The ID of the dispute
     * @param winner The address of the winning party
     */
    function resolveDispute(
        uint256 disputeId,
        address winner
    ) external;

    /**
     * @notice Gets the status of a dispute
     * @param disputeId The ID of the dispute
     * @return The status of the dispute
     */
    function getDisputeStatus(
        uint256 disputeId
    ) external view returns (DisputeStatus);

    /**
     * @notice Phase 1 verification functions
     */
    function submitPhase1Claim(
        uint256 disputeId,
        uint256 startIndex,
        uint256 endIndex,
        bytes32 stateRoot
    ) external;

    function challengePhase1Claim(
        uint256 disputeId,
        uint256 challengeIndex,
        bytes32 expectedStateRoot
    ) external;

    function bisectPhase1Claim(
        uint256 disputeId,
        uint256 bisectionPoint,
        bytes32 leftStateRoot,
        bytes32 rightStateRoot
    ) external;

    function verifyPhase1StateRoot(
        uint256 disputeId,
        uint256 startIndex,
        uint256 endIndex,
        bytes32 stateRoot
    ) external view returns (bool);

    /**
     * @notice Phase 2 verification functions
     */
    function submitPhase2Claim(
        uint256 disputeId,
        uint256 instructionIndex,
        bytes32 preStateRoot,
        bytes32 postStateRoot,
        bytes32 witnessHash
    ) external;

    function challengePhase2Claim(
        uint256 disputeId,
        uint256 instructionIndex,
        bytes32 expectedPostStateRoot
    ) external;

    function executeDisputedInstruction(
        uint256 disputeId,
        uint256 instructionIndex,
        bytes calldata witnessData
    ) external returns (bool success, bytes32 postStateRoot);

    function verifyPhase2Witness(
        uint256 disputeId,
        uint256 instructionIndex,
        bytes calldata witnessData
    ) external view returns (bool);
}