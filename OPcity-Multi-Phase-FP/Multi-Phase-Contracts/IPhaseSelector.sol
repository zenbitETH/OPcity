// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

/**
 * @title IPhaseSelector
 * @notice Interface for the Phase Selector that determines which phase to use for dispute resolution
 */
interface IPhaseSelector {
    /**
     * @notice Event emitted when the phase transition threshold is updated
     * @param oldThreshold The previous threshold
     * @param newThreshold The new threshold
     */
    event PhaseTransitionThresholdUpdated(
        uint256 oldThreshold,
        uint256 newThreshold
    );

    /**
     * @notice Determines the appropriate phase for a given dispute
     * @param disputeId The ID of the dispute
     * @param complexity The complexity measure of the dispute (e.g., segment size)
     * @return The recommended phase for dispute resolution
     */
    function determinePhase(
        uint256 disputeId,
        uint256 complexity
    ) external view returns (uint8);

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
    ) external view returns (bool);

    /**
     * @notice Sets the phase transition threshold
     * @param newThreshold The new threshold for transitioning from Phase 1 to Phase 2
     */
    function setPhaseTransitionThreshold(uint256 newThreshold) external;

    /**
     * @notice Gets the current phase transition threshold
     * @return The current threshold for phase transitions
     */
    function getPhaseTransitionThreshold() external view returns (uint256);
}