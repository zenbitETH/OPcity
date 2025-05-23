// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

/**
 * @title ICannonFPVM
 * @notice Interface for the Cannon FPVM that executes and verifies instructions
 */
interface ICannonFPVM {
    /**
     * @notice Executes an instruction with the given witness data
     * @param preStateRoot The state root before execution
     * @param instructionIndex The index of the instruction to execute
     * @param witnessData The witness data for the instruction
     * @return success Whether the execution was successful
     * @return postStateRoot The state root after execution
     */
    function executeInstruction(
        bytes32 preStateRoot,
        uint256 instructionIndex,
        bytes calldata witnessData
    ) external returns (bool success, bytes32 postStateRoot);

    /**
     * @notice Verifies a witness for an instruction
     * @param preStateRoot The state root before execution
     * @param instructionIndex The index of the instruction
     * @param witnessData The witness data for the instruction
     * @param expectedPostStateRoot The expected state root after execution
     * @return True if the witness is valid, false otherwise
     */
    function verifyWitness(
        bytes32 preStateRoot,
        uint256 instructionIndex,
        bytes calldata witnessData,
        bytes32 expectedPostStateRoot
    ) external view returns (bool);

    /**
     * @notice Gets the details of an instruction
     * @param stateRoot The state root at which to fetch the instruction
     * @param instructionIndex The index of the instruction
     * @return The instruction details
     */
    function getInstructionDetails(
        bytes32 stateRoot,
        uint256 instructionIndex
    ) external view returns (bytes memory);
}