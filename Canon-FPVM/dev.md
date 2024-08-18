# Cannon Developer Documentation

## Table of Contents

1. [Introduction](#introduction)
2. [System Architecture](#system-architecture)
3. [Core Components](#core-components)
   3.1. [MIPS.sol](#mipssol)
   3.2. [mipsevm](#mipsevm)
4. [Key Concepts](#key-concepts)
   4.1. [State](#state)
   4.2. [Witness Data](#witness-data)
   4.3. [Pre-image Oracle](#pre-image-oracle)
5. [Main Workflow](#main-workflow)
6. [Important Functions](#important-functions)
7. [Testing and Debugging](#testing-and-debugging)
8. [Performance Considerations](#performance-considerations)
9. [Future Improvements](#future-improvements)

## 1. Introduction

Cannon is an onchain MIPS instruction emulator designed to support EVM-equivalent fault proofs. It enables Geth to run onchain, one instruction at a time, as part of an interactive dispute game. The system consists of two main components: an onchain MIPS implementation (`MIPS.sol`) and an offchain Go implementation (`mipsevm`) to produce proofs for MIPS instructions.

## 2. System Architecture

Cannon's architecture is divided into two main parts:

1. **Onchain Component**: Implemented in Solidity (`MIPS.sol`), this part verifies the execution of individual MIPS instructions within the EVM.

2. **Offchain Component**: Implemented in Go (`mipsevm`), this part emulates MIPS instructions, generates proofs, and interacts with the onchain component.

The system also includes a CLI for loading programs, transitioning states, and generating proofs.

## 3. Core Components

### 3.1. MIPS.sol

`MIPS.sol` is the onchain implementation of big-endian 32-bit MIPS instruction execution. It covers MIPS III, R3000 instructions and implements a minimal subset of the Linux kernel syscalls.

Key features:
- Instruction execution verification
- Memory operations
- Syscall handling
- State management

### 3.2. mipsevm

`mipsevm` is the Go package that implements the offchain MIPS emulator. It's responsible for:

- Loading and executing MIPS programs
- Generating proofs for instruction execution
- Managing VM state
- Handling pre-image requests
- Providing interfaces for different VM implementations (e.g., single-threaded, multi-threaded)

Key sub-packages:
- `memory`: Implements the VM's memory model
- `program`: Handles program loading and patching
- `exec`: Contains core execution logic for MIPS instructions

## 4. Key Concepts

### 4.1. State

The VM state consists of:
- Memory content
- Register values
- Program counter (PC)
- Other CPU-specific values (e.g., HI, LO registers)
- VM metadata (e.g., step count, exit status)

State is encoded and decoded for proof generation and verification.

### 4.2. Witness Data

Witness data is crucial for onchain verification. It includes:
- Packed State: Compressed representation of the VM state
- Memory Proofs: Merkle proofs for memory access
- Pre-image Data: Data retrieved from the pre-image oracle

### 4.3. Pre-image Oracle

The pre-image oracle provides external data to the VM during execution. It's implemented using a specific ABI for communication through syscalls.

## 5. Main Workflow

1. Load MIPS program into initial state
2. Execute instructions step-by-step
3. Generate witness data for each step
4. Verify steps onchain using `MIPS.sol`
5. Continue execution or resolve disputes based on verification results

## 6. Important Functions

### mipsevm

#### State Management
- `CreateInitialState`: Initialize the VM state
- `GetState`: Retrieve current VM state
- `EncodeWitness`: Encode the current state as witness data

#### Execution
- `Step`: Execute a single MIPS instruction
- `ExecMipsCoreStepLogic`: Core logic for executing MIPS instructions
- `handleSyscall`: Handle system calls during execution

#### Memory Operations
- `SetMemory`: Write to VM memory
- `GetMemory`: Read from VM memory
- `MerkleProof`: Generate Merkle proof for memory access

#### Pre-image Oracle
- `GetPreimage`: Retrieve pre-image data
- `Hint`: Provide hints to the pre-image oracle

### MIPS.sol

- `step`: Verify a single step of MIPS execution
- `executeInstruction`: Execute a MIPS instruction onchain
- `handleMemoryOp`: Handle memory operations
- `handlePreimageOp`: Handle pre-image oracle operations

## 7. Testing and Debugging

Cannon includes extensive testing capabilities:

- Unit tests for individual components
- Integration tests for end-to-end workflows
- Fuzzing tests to identify edge cases and vulnerabilities
- Debugging tools for tracing execution and analyzing state

Key testing functions:
- `TestEVM`: Test EVM execution of MIPS instructions
- `FuzzStateSyscall*`: Fuzz testing for various syscalls
- `RunVMTest_*`: Run specific VM tests (e.g. Claim)

## 8. Performance Considerations

- Memory management: Efficient allocation and deallocation of pages
- Proof generation: Optimize Merkle proof generation for large memory spaces
- Onchain verification: Minimize gas costs for instruction verification
- Pre-image oracle: Efficient handling of external data requests

## 9. Future Improvements

Potential areas for enhancement:
- Support for additional MIPS instructions
- Optimizations for faster offchain execution
- Enhanced debugging and tracing capabilities
- Expanded test coverage and fuzzing scenarios
