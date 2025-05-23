// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

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
