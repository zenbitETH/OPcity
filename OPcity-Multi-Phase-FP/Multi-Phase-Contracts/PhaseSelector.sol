// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { IPhaseSelector } from "interfaces/dispute/IPhaseSelector.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ISemver } from "interfaces/universal/ISemver.sol";

/**
 * @title PhaseSelector
 * @notice Implementation of the Phase Selector that determines which phase to use
 */
contract PhaseSelector is IPhaseSelector, Ownable, Semver {
    // Constants
    uint8 private constant PHASE_1 = 1;
    uint8 private constant PHASE_2 = 2;

    // State variables
    uint256 public phaseTransitionThreshold;
    mapping(uint256 => mapping(bytes32 => bool)) private verifiedStateRoots;

    /**
     * @notice Constructor for the PhaseSelector
     * @param _initialThreshold The initial threshold for phase transitions
     */
    constructor(uint256 _initialThreshold) Semver(1, 0, 0) {
        require(_initialThreshold > 0, "Threshold must be positive");
        phaseTransitionThreshold = _initialThreshold;
    }

    /**
     * @notice Determines the appropriate phase for a given dispute
     * @param disputeId The ID of the dispute
     * @param complexity The complexity measure of the dispute (e.g., segment size)
     * @return The recommended phase for dispute resolution
     */
    function determinePhase(
        uint256 disputeId,
        uint256 complexity
    ) external view override returns (uint8) {
        // If complexity is below threshold, use Phase 2, otherwise use Phase 1
        if (complexity <= phaseTransitionThreshold) {
            return PHASE_2;
        } else {
            return PHASE_1;
        }
    }

    /**
     * @notice Verifies a segment state root
     * @param disputeId The ID of the dispute
     * @param startIndex The start index of the segment
     * @param endIndex The end index of the segment
     * @param stateRoot The state root to verify
     * @return True if the state root is valid, false otherwise
     */
    function verifySegmentStateRoot(
        uint256 disputeId,
        uint256 startIndex,
        uint256 endIndex,
        bytes32 stateRoot
    ) external view override returns (bool) {
        // For now, trust registered state roots
        // In a production implementation, this would verify against proven state
        bytes32 segmentKey = keccak256(abi.encodePacked(
            disputeId,
            startIndex,
            endIndex
        ));

        return verifiedStateRoots[disputeId][segmentKey];
    }

    /**
     * @notice Registers a verified state root for a segment
     * @param disputeId The ID of the dispute
     * @param startIndex The start index of the segment
     * @param endIndex The end index of the segment
     * @param stateRoot The verified state root
     */
    function registerVerifiedStateRoot(
        uint256 disputeId,
        uint256 startIndex,
        uint256 endIndex,
        bytes32 stateRoot
    ) external onlyOwner {
        bytes32 segmentKey = keccak256(abi.encodePacked(
            disputeId,
            startIndex,
            endIndex
        ));

        verifiedStateRoots[disputeId][segmentKey] = true;
    }

    /**
     * @notice Sets the phase transition threshold
     * @param newThreshold The new threshold for transitioning from Phase 1 to Phase 2
     */
    function setPhaseTransitionThreshold(
        uint256 newThreshold
    ) external override onlyOwner {
        require(newThreshold > 0, "Threshold must be positive");

        uint256 oldThreshold = phaseTransitionThreshold;
        phaseTransitionThreshold = newThreshold;

        emit PhaseTransitionThresholdUpdated(oldThreshold, newThreshold);
    }

    /**
     * @notice Gets the current phase transition threshold
     * @return The current threshold for phase transitions
     */
    function getPhaseTransitionThreshold()
        external view override returns (uint256) {
        return phaseTransitionThreshold;
    }
}