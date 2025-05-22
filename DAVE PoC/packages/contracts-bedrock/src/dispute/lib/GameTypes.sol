
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

/**
 * @title GameTypes
 * @notice Library defining constants for different types of dispute games in the OP Stack
 * @dev This library provides a centralized location for game type identifiers
 */
library GameTypes {
    /**
     * @notice Type identifier for fault dispute games
     */
    uint8 public constant FAULT = 0;

    /**
     * @notice Type identifier for validity dispute games
     */
    uint8 public constant VALIDITY = 1;

    /**
     * @notice Type identifier for attestation dispute games
     */
    uint8 public constant ATTESTATION = 2;

    /**
     * @notice Type identifier for tournament-based dispute games (DAVE)
     * @dev Used for Cartesi DAVE's tournament-based dispute resolution system
     */
    uint8 public constant TOURNAMENT = 3;
}
