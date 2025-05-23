// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { IMultiPhaseResolver } from "interfaces/dispute/IMultiPhaseResolver.sol";
import { IPhaseSelector } from "interfaces/dispute/IPhaseSelector.sol";
import { ICannonFPVM } from "interfaces/dispute/ICannonFPVM.sol";
import { IFaultDisputeGame } from "interfaces/dispute/IFaultDisputeGame.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ReentrancyGuard } from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";

/**
 * @title MultiPhaseResolver
 * @notice Implementation of the Multi-Phase Resolver that coordinates dispute resolution
 */
contract MultiPhaseResolver is IMultiPhaseResolver, Ownable, ReentrancyGuard, ISemver {
    // Constants
    uint256 private constant PHASE_TRANSITION_THRESHOLD = 1000;
    uint256 private constant MAX_BISECTION_DEPTH = 32;

    // State variables
    mapping(uint256 => Dispute) private disputes;
    uint256 private nextDisputeId;
    IPhaseSelector public phaseSelector;
    ICannonFPVM public cannonFPVM;
    IFaultDisputeGame public faultDisputeGame;

    // Data structures
    struct Dispute {
        address proposer;
        address challenger;
        DisputePhase currentPhase;
        DisputeStatus status;
        bytes32 claimedStateRoot;
        bytes32 challengedStateRoot;
        uint256 timestamp;

        // Phase 1 specific data
        Phase1Data phase1Data;

        // Phase 2 specific data
        Phase2Data phase2Data;
    }

    struct Phase1Data {
        uint256 startIndex;
        uint256 endIndex;
        bytes32 startStateRoot;
        bytes32 endStateRoot;
        uint256 bisectionCount;
        mapping(uint256 => BisectionStep) bisectionSteps;
    }

    struct Phase2Data {
        uint256 segmentStart;
        uint256 segmentEnd;
        bytes32 segmentStateRoot;
        uint256 instructionIndex;
        bytes32 preStateRoot;
        bytes32 postStateRoot;
        bytes32 witnessHash;
    }

    struct BisectionStep {
        uint256 bisectionPoint;
        bytes32 leftStateRoot;
        bytes32 rightStateRoot;
        uint256 timestamp;
    }

    // Modifiers
    modifier onlyDisputeParticipant(uint256 _disputeId) {
        require(
            msg.sender == disputes[_disputeId].proposer ||
            msg.sender == disputes[_disputeId].challenger,
            "Not a dispute participant"
        );
        _;
    }

    modifier disputeExists(uint256 _disputeId) {
        require(_disputeId < nextDisputeId, "Dispute does not exist");
        _;
    }

    modifier inPhase(uint256 _disputeId, DisputePhase _phase) {
        require(
            disputes[_disputeId].currentPhase == _phase,
            "Invalid phase"
        );
        _;
    }

    /**
     * @notice Constructor for the MultiPhaseResolver
     * @param _phaseSelector The address of the Phase Selector contract
     * @param _cannonFPVM The address of the Cannon FPVM contract
     * @param _faultDisputeGame The address of the Fault Dispute Game contract
     */
    constructor(
        address _phaseSelector,
        address _cannonFPVM,
        address _faultDisputeGame
    ) Semver(1, 0, 0) {
        require(_phaseSelector != address(0), "Invalid phase selector address");
        require(_cannonFPVM != address(0), "Invalid cannon FPVM address");
        require(_faultDisputeGame != address(0), "Invalid dispute game address");

        phaseSelector = IPhaseSelector(_phaseSelector);
        cannonFPVM = ICannonFPVM(_cannonFPVM);
        faultDisputeGame = IFaultDisputeGame(_faultDisputeGame);
        nextDisputeId = 1;
    }

    /**
     * @notice Creates a new dispute
     * @param claimedStateRoot The state root claimed by the proposer
     * @param challengedStateRoot The state root claimed by the challenger
     * @return disputeId The ID of the created dispute
     */
    function createDispute(
        bytes32 claimedStateRoot,
        bytes32 challengedStateRoot
    ) external override nonReentrant returns (uint256 disputeId) {
        require(claimedStateRoot != challengedStateRoot, "State roots must differ");

        disputeId = nextDisputeId++;
        Dispute storage dispute = disputes[disputeId];

        dispute.proposer = msg.sender;
        dispute.challenger = address(0); // Will be set when someone challenges
        dispute.currentPhase = DisputePhase.PHASE_1;
        dispute.status = DisputeStatus.ACTIVE;
        dispute.claimedStateRoot = claimedStateRoot;
        dispute.challengedStateRoot = challengedStateRoot;
        dispute.timestamp = block.timestamp;

        // Initialize Phase 1 data
        dispute.phase1Data.startIndex = 0;
        dispute.phase1Data.endIndex = type(uint256).max; // Full range initially
        dispute.phase1Data.startStateRoot = claimedStateRoot;
        dispute.phase1Data.endStateRoot = challengedStateRoot;
        dispute.phase1Data.bisectionCount = 0;

        emit DisputeCreated(
            disputeId,
            msg.sender,
            address(0),
            claimedStateRoot,
            challengedStateRoot
        );

        return disputeId;
    }

    /**
     * @notice Gets the current phase of a dispute
     * @param disputeId The ID of the dispute
     * @return The current phase of the dispute
     */
    function getCurrentPhase(
        uint256 disputeId
    ) external view override disputeExists(disputeId) returns (DisputePhase) {
        return disputes[disputeId].currentPhase;
    }

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
    ) external override
      disputeExists(disputeId)
      inPhase(disputeId, DisputePhase.PHASE_1)
      onlyDisputeParticipant(disputeId)
      nonReentrant {
        Dispute storage dispute = disputes[disputeId];

        require(segmentEnd > segmentStart, "Invalid segment range");
        require(segmentEnd - segmentStart <= PHASE_TRANSITION_THRESHOLD, "Segment too large");
        require(
            segmentStart >= dispute.phase1Data.startIndex &&
            segmentEnd <= dispute.phase1Data.endIndex,
            "Segment out of bounds"
        );

        // Verify segment state root via Phase Selector
        require(
            phaseSelector.verifySegmentStateRoot(
                disputeId,
                segmentStart,
                segmentEnd,
                segmentStateRoot
            ),
            "Invalid segment state root"
        );

        // Transition to Phase 2
        DisputePhase oldPhase = dispute.currentPhase;
        dispute.currentPhase = DisputePhase.PHASE_2;

        // Initialize Phase 2 data
        dispute.phase2Data.segmentStart = segmentStart;
        dispute.phase2Data.segmentEnd = segmentEnd;
        dispute.phase2Data.segmentStateRoot = segmentStateRoot;
        dispute.phase2Data.instructionIndex = segmentStart; // Start at beginning of segment

        emit PhaseTransition(disputeId, oldPhase, DisputePhase.PHASE_2);
    }

    /**
     * @notice Gets the status of a dispute
     * @param disputeId The ID of the dispute
     * @return The status of the dispute
     */
    function getDisputeStatus(
        uint256 disputeId
    ) external view override disputeExists(disputeId) returns (DisputeStatus) {
        return disputes[disputeId].status;
    }

    /**
     * @notice Resolves a dispute and determines the winner
     * @param disputeId The ID of the dispute
     * @param winner The address of the winning party
     */
    function resolveDispute(
        uint256 disputeId,
        address winner
    ) external override
      disputeExists(disputeId)
      nonReentrant {
        Dispute storage dispute = disputes[disputeId];

        // Ensure dispute is active
        require(dispute.status == DisputeStatus.ACTIVE, "Dispute not active");

        // Ensure winner is either proposer or challenger
        require(
            winner == dispute.proposer || winner == dispute.challenger,
            "Invalid winner"
        );

        // Set dispute as resolved
        dispute.status = DisputeStatus.RESOLVED;
        dispute.currentPhase = DisputePhase.RESOLVED;

        // Determine winning state root
        bytes32 winningStateRoot = winner == dispute.proposer
            ? dispute.claimedStateRoot
            : dispute.challengedStateRoot;

        // Notify the FaultDisputeGame contract of the resolution
        faultDisputeGame.resolveGame(disputeId, winner);

        emit DisputeResolved(disputeId, winner, winningStateRoot);
    }

    /**
     * @notice Submits a claim in Phase 1
     * @param disputeId The ID of the dispute
     * @param startIndex The start index of the claim segment
     * @param endIndex The end index of the claim segment
     * @param stateRoot The state root at the specified index range
     */
    function submitPhase1Claim(
        uint256 disputeId,
        uint256 startIndex,
        uint256 endIndex,
        bytes32 stateRoot
    ) external override
      disputeExists(disputeId)
      inPhase(disputeId, DisputePhase.PHASE_1)
      onlyDisputeParticipant(disputeId)
      nonReentrant {
        Dispute storage dispute = disputes[disputeId];

        // Verify claim segment
        require(endIndex > startIndex, "Invalid segment range");
        require(
            startIndex >= dispute.phase1Data.startIndex &&
            endIndex <= dispute.phase1Data.endIndex,
            "Segment out of bounds"
        );

        // Set challenger if not set
        if (dispute.challenger == address(0) && msg.sender != dispute.proposer) {
            dispute.challenger = msg.sender;
        }

        // Update Phase 1 data
        dispute.phase1Data.startIndex = startIndex;
        dispute.phase1Data.endIndex = endIndex;
        dispute.phase1Data.startStateRoot = stateRoot;

        // Check if segment is small enough to transition to Phase 2
        if ((endIndex - startIndex) <= PHASE_TRANSITION_THRESHOLD) {
            // Auto-transition to phase 2 once we've narrowed down the segment
            _autoTransitionToPhase2(disputeId);
        }
    }

    /**
     * @notice Challenges a claim in Phase 1
     * @param disputeId The ID of the dispute
     * @param challengeIndex The index at which the challenge is made
     * @param expectedStateRoot The expected state root at the challenge index
     */
    function challengePhase1Claim(
        uint256 disputeId,
        uint256 challengeIndex,
        bytes32 expectedStateRoot
    ) external override
      disputeExists(disputeId)
      inPhase(disputeId, DisputePhase.PHASE_1)
      onlyDisputeParticipant(disputeId)
      nonReentrant {
        Dispute storage dispute = disputes[disputeId];

        // Verify challenge index
        require(
            challengeIndex > dispute.phase1Data.startIndex &&
            challengeIndex < dispute.phase1Data.endIndex,
            "Challenge index out of bounds"
        );

        // Set challenger if not set
        if (dispute.challenger == address(0) && msg.sender != dispute.proposer) {
            dispute.challenger = msg.sender;
        }

        // Update dispute data based on which side of the segment is challenged
        if (challengeIndex == dispute.phase1Data.startIndex) {
            dispute.phase1Data.startStateRoot = expectedStateRoot;
        } else if (challengeIndex == dispute.phase1Data.endIndex) {
            dispute.phase1Data.endStateRoot = expectedStateRoot;
        } else {
            // For intermediate points, require bisection
            revert("Use bisection for intermediate points");
        }
    }

    /**
     * @notice Performs bisection in Phase 1
     * @param disputeId The ID of the dispute
     * @param bisectionPoint The point at which to bisect the segment
     * @param leftStateRoot The state root of the left segment
     * @param rightStateRoot The state root of the right segment
     */
    function bisectPhase1Claim(
        uint256 disputeId,
        uint256 bisectionPoint,
        bytes32 leftStateRoot,
        bytes32 rightStateRoot
    ) external override
      disputeExists(disputeId)
      inPhase(disputeId, DisputePhase.PHASE_1)
      onlyDisputeParticipant(disputeId)
      nonReentrant {
        Dispute storage dispute = disputes[disputeId];

        // Ensure the bisection point is within bounds
        require(
            bisectionPoint > dispute.phase1Data.startIndex &&
            bisectionPoint < dispute.phase1Data.endIndex,
            "Invalid bisection point"
        );

        // Ensure we haven't exceeded max bisection depth
        require(
            dispute.phase1Data.bisectionCount < MAX_BISECTION_DEPTH,
            "Max bisection depth exceeded"
        );

        // Store the bisection step
        uint256 bisectionIndex = dispute.phase1Data.bisectionCount++;
        dispute.phase1Data.bisectionSteps[bisectionIndex] = BisectionStep({
            bisectionPoint: bisectionPoint,
            leftStateRoot: leftStateRoot,
            rightStateRoot: rightStateRoot,
            timestamp: block.timestamp
        });

        // Check if segment is small enough to transition to Phase 2
        uint256 leftSize = bisectionPoint - dispute.phase1Data.startIndex;
        uint256 rightSize = dispute.phase1Data.endIndex - bisectionPoint;

        if (leftSize <= PHASE_TRANSITION_THRESHOLD || rightSize <= PHASE_TRANSITION_THRESHOLD) {
            // Auto-transition to phase 2 once we've narrowed down the segment
            _autoTransitionToPhase2(disputeId);
        }
    }

    /**
     * @notice Verifies a state root in Phase 1
     * @param disputeId The ID of the dispute
     * @param startIndex The start index of the segment
     * @param endIndex The end index of the segment
     * @param stateRoot The state root to verify
     * @return Whether the state root is valid
     */
    function verifyPhase1StateRoot(
        uint256 disputeId,
        uint256 startIndex,
        uint256 endIndex,
        bytes32 stateRoot
    ) external view override
      disputeExists(disputeId)
      returns (bool) {
        // Delegate to the Phase Selector for verification
        return phaseSelector.verifySegmentStateRoot(
            disputeId,
            startIndex,
            endIndex,
            stateRoot
        );
    }

    /**
     * @notice Submits a claim in Phase 2
     * @param disputeId The ID of the dispute
     * @param instructionIndex The index of the instruction
     * @param preStateRoot The state root before the instruction
     * @param postStateRoot The state root after the instruction
     * @param witnessHash The hash of the witness data
     */
    function submitPhase2Claim(
        uint256 disputeId,
        uint256 instructionIndex,
        bytes32 preStateRoot,
        bytes32 postStateRoot,
        bytes32 witnessHash
    ) external override
      disputeExists(disputeId)
      inPhase(disputeId, DisputePhase.PHASE_2)
      onlyDisputeParticipant(disputeId)
      nonReentrant {
        Dispute storage dispute = disputes[disputeId];

        // Verify instruction index is within the segment
        require(
            instructionIndex >= dispute.phase2Data.segmentStart &&
            instructionIndex < dispute.phase2Data.segmentEnd,
            "Instruction index out of bounds"
        );

        // Update Phase 2 data
        dispute.phase2Data.instructionIndex = instructionIndex;
        dispute.phase2Data.preStateRoot = preStateRoot;
        dispute.phase2Data.postStateRoot = postStateRoot;
        dispute.phase2Data.witnessHash = witnessHash;
    }

    /**
     * @notice Challenges a claim in Phase 2
     * @param disputeId The ID of the dispute
     * @param instructionIndex The index of the instruction
     * @param expectedPostStateRoot The expected post-state root
     */
    function challengePhase2Claim(
        uint256 disputeId,
        uint256 instructionIndex,
        bytes32 expectedPostStateRoot
    ) external override
      disputeExists(disputeId)
      inPhase(disputeId, DisputePhase.PHASE_2)
      onlyDisputeParticipant(disputeId)
      nonReentrant {
        Dispute storage dispute = disputes[disputeId];

        // Verify instruction index matches the current claim
        require(
            instructionIndex == dispute.phase2Data.instructionIndex,
            "Instruction index mismatch"
        );

        // Verify the expected post-state root differs from the claimed one
        require(
            expectedPostStateRoot != dispute.phase2Data.postStateRoot,
            "Post-state root must differ"
        );

        // Update the challenged post-state root
        dispute.phase2Data.postStateRoot = expectedPostStateRoot;
    }

    /**
     * @notice Executes a disputed instruction
     * @param disputeId The ID of the dispute
     * @param instructionIndex The index of the instruction
     * @param witnessData The witness data for the instruction
     * @return success Whether the execution was successful
     * @return postStateRoot The post-state root after execution
     */
    function executeDisputedInstruction(
        uint256 disputeId,
        uint256 instructionIndex,
        bytes calldata witnessData
    ) external override
      disputeExists(disputeId)
      inPhase(disputeId, DisputePhase.PHASE_2)
      nonReentrant
      returns (bool success, bytes32 postStateRoot) {
        Dispute storage dispute = disputes[disputeId];

        // Verify instruction index
        require(
            instructionIndex == dispute.phase2Data.instructionIndex,
            "Instruction index mismatch"
        );

        // Verify witness hash
        require(
            keccak256(witnessData) == dispute.phase2Data.witnessHash,
            "Invalid witness data"
        );

        // Execute the instruction using Cannon FPVM
        (success, postStateRoot) = cannonFPVM.executeInstruction(
            dispute.phase2Data.preStateRoot,
            instructionIndex,
            witnessData
        );

        // If execution is successful and we can determine the winner, resolve the dispute
        if (success) {
            address winner;
            if (postStateRoot == dispute.phase2Data.postStateRoot) {
                // The proposer's claim was correct
                winner = dispute.proposer;
            } else {
                // The challenger's claim was correct
                winner = dispute.challenger;
            }

            // Resolve the dispute
            _resolveDispute(disputeId, winner);
        }

        return (success, postStateRoot);
    }

    /**
     * @notice Verifies a witness in Phase 2
     * @param disputeId The ID of the dispute
     * @param instructionIndex The index of the instruction
     * @param witnessData The witness data for the instruction
     * @return Whether the witness is valid
     */
    function verifyPhase2Witness(
        uint256 disputeId,
        uint256 instructionIndex,
        bytes calldata witnessData
    ) external view override
      disputeExists(disputeId)
      returns (bool) {
        Dispute storage dispute = disputes[disputeId];

        // Verify the instruction index is within bounds
        require(
            instructionIndex >= dispute.phase2Data.segmentStart &&
            instructionIndex < dispute.phase2Data.segmentEnd,
            "Instruction index out of bounds"
        );

        // Verify the witness using Cannon FPVM
        return cannonFPVM.verifyWitness(
            dispute.phase2Data.preStateRoot,
            instructionIndex,
            witnessData,
            dispute.phase2Data.postStateRoot
        );
    }

    /**
     * @notice Gets the details of a dispute
     * @param disputeId The ID of the dispute
     * @return Dispute memory representing the dispute details
     * Note: This does not return mappings which are part of Phase1Data
     */
    function getDisputeDetails(
        uint256 disputeId
    ) external view disputeExists(disputeId) returns (
        address proposer,
        address challenger,
        DisputePhase currentPhase,
        DisputeStatus status,
        bytes32 claimedStateRoot,
        bytes32 challengedStateRoot,
        uint256 timestamp,
        uint256 p1StartIndex,
        uint256 p1EndIndex,
        bytes32 p1StartStateRoot,
        bytes32 p1EndStateRoot,
        uint256 p1BisectionCount,
        uint256 p2SegmentStart,
        uint256 p2SegmentEnd,
        bytes32 p2SegmentStateRoot,
        uint256 p2InstructionIndex,
        bytes32 p2PreStateRoot,
        bytes32 p2PostStateRoot,
        bytes32 p2WitnessHash
    ) {
        Dispute storage dispute = disputes[disputeId];

        return (
            dispute.proposer,
            dispute.challenger,
            dispute.currentPhase,
            dispute.status,
            dispute.claimedStateRoot,
            dispute.challengedStateRoot,
            dispute.timestamp,
            dispute.phase1Data.startIndex,
            dispute.phase1Data.endIndex,
            dispute.phase1Data.startStateRoot,
            dispute.phase1Data.endStateRoot,
            dispute.phase1Data.bisectionCount,
            dispute.phase2Data.segmentStart,
            dispute.phase2Data.segmentEnd,
            dispute.phase2Data.segmentStateRoot,
            dispute.phase2Data.instructionIndex,
            dispute.phase2Data.preStateRoot,
            dispute.phase2Data.postStateRoot,
            dispute.phase2Data.witnessHash
        );
    }

    /**
     * @notice Gets a specific bisection step from Phase 1
     * @param disputeId The ID of the dispute
     * @param bisectionIndex The index of the bisection step
     * @return The bisection step details
     */
    function getBisectionStep(
        uint256 disputeId,
        uint256 bisectionIndex
    ) external view
      disputeExists(disputeId)
      returns (
        uint256 bisectionPoint,
        bytes32 leftStateRoot,
        bytes32 rightStateRoot,
        uint256 timestamp
      ) {
        require(
            bisectionIndex < disputes[disputeId].phase1Data.bisectionCount,
            "Bisection index out of bounds"
        );

        BisectionStep storage step = disputes[disputeId].phase1Data.bisectionSteps[bisectionIndex];
        return (
            step.bisectionPoint,
            step.leftStateRoot,
            step.rightStateRoot,
            step.timestamp
        );
    }

    /**
     * @notice Internal function to auto-transition to Phase 2
     * @param disputeId The ID of the dispute
     */
    function _autoTransitionToPhase2(uint256 disputeId) internal {
        Dispute storage dispute = disputes[disputeId];

        DisputePhase oldPhase = dispute.currentPhase;
        dispute.currentPhase = DisputePhase.PHASE_2;

        // Initialize Phase 2 data
        dispute.phase2Data.segmentStart = dispute.phase1Data.startIndex;
        dispute.phase2Data.segmentEnd = dispute.phase1Data.endIndex;
        dispute.phase2Data.segmentStateRoot = dispute.phase1Data.startStateRoot;
        dispute.phase2Data.instructionIndex = dispute.phase1Data.startIndex;

        emit PhaseTransition(disputeId, oldPhase, DisputePhase.PHASE_2);
    }

    /**
     * @notice Internal function to resolve a dispute
     * @param disputeId The ID of the dispute
     * @param winner The address of the winning party
     */
    function _resolveDispute(uint256 disputeId, address winner) internal {
        Dispute storage dispute = disputes[disputeId];

        // Set dispute as resolved
        dispute.status = DisputeStatus.RESOLVED;
        dispute.currentPhase = DisputePhase.RESOLVED;

        // Determine winning state root
        bytes32 winningStateRoot = winner == dispute.proposer
            ? dispute.claimedStateRoot
            : dispute.challengedStateRoot;

        emit DisputeResolved(disputeId, winner, winningStateRoot);
    }
}