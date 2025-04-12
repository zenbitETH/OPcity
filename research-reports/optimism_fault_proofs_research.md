# Optimism's Fault Proofs & Canon FPVM Research

## Architecture of Optimism's Fault Proof System

Optimism's fault proof system is designed to ensure the validity of Layer 2 transactions by providing a mechanism to challenge incorrect state transitions. The system consists of several key components that work together to enable secure and verifiable dispute resolution.

### Overview

The fault proof system is built around a dispute game mechanism where participants can challenge state transitions they believe are incorrect. The system uses a bisection protocol to narrow down disagreements to a single instruction, which can then be verified on-chain.

Key components include:
- **Dispute Game Protocol**: Manages the overall dispute resolution process
- **Fault Proof Virtual Machine (FPVM)**: Executes and verifies state transitions
- **OP-Challenger**: Handles initiating challenges and interacting with dispute games
- **OP-Program**: Handles derivation of L2 output from L1 inputs

## Canon FPVM Implementation Details

Canon is Optimism's default Fault Proof Virtual Machine (FPVM) and consists of two main components:

1. **Onchain MIPS.sol**: EVM implementation to verify execution of a single MIPS instruction
2. **Offchain Cannon**: Go implementation to produce a proof for any MIPS instruction to verify onchain

### Onchain vs. Offchain Components

The Canon FPVM has both onchain and offchain components that work together:

- **Onchain (MIPS.sol)**: Runs a single MIPS instruction to definitively prove the correct post-state
- **Offchain (Cannon)**: Runs many MIPS instructions and generates state witness hashes and proofs

### Control Flow

The fault proof process follows a specific flow:

1. During an active dispute game, participants disagree on L2 block state transitions
2. The bisection game narrows down to a single L2 block state transition in dispute
3. OP-Challenger runs Cannon to process MIPS instructions within the FPVM
4. Cannon generates state witness hashes as commitments to computation results
5. The bisection continues until a single MIPS instruction is identified as the root disagreement
6. Cannon generates a witness proof containing all information needed to run the instruction onchain
7. The single instruction is executed in MIPS.sol to determine the correct post-state
8. The dispute is resolved based on this definitive result

### Offchain Cannon Components

Cannon's offchain implementation consists of several core components:

#### 1. MIPSEVM State and Memory

- 32-bit addressable memory range [0, 2^32-1]
- Uses monolithic memory structure
- Memory is stored in a binary Merkle tree data structure
- Tree has fixed-depth of 27 levels with leaf values of 32 bytes each
- Spans full 32-bit address space: 2^27 * 32 = 2^32
- Optimizations include caching zeroed-out memory regions and recently used pages

#### 2. ELF Loader

- Loads the Executable and Linkable Format (ELF) binary containing OP-Program
- Parses ELF headers to determine programs to load into memory
- Locates initial values for Program Counter (PC) and NextPC
- Instantiates stack, heap, and data segment pointers
- Patches out incompatible functions (system calls, concurrency features)

#### 3. Memory and State Management

- Maintains the entire 32-bit memory address space
- Uses binary Merkle tree for efficient memory representation
- Provides GetMemory(), ReadMemoryRange(), and SetMemory() functions
- Encodes up to two memory Merkle proofs for onchain verification

#### 4. Witness Proof Generation

- Generates state witness hashes during execution trace bisection
- Creates witness proof for the disputed MIPS instruction
- Encodes VM execution state, memory proofs, and any required pre-image information
- Communicates with OP-Challenger to post pre-image data to PreimageOracle.sol if needed

### Key Technical Details

- **Endianness**: MIPSEVM is Big-Endian, while most host machines are Little-Endian
- **Instruction Set**: Implements MIPS R3000, 32-bit Instruction Set Architecture (ISA)
- **Determinism**: Both onchain and offchain implementations must produce exactly the same results given identical inputs
- **Memory Representation**: Uses binary Merkle tree with nodes combined as: out = keccak256(left ++ right)
- **Statelessness**: MIPS.sol is mostly stateless, requiring only the memory Merkle root and up to two memory proofs

## Recent Updates and Performance

The fault proof system has seen significant development:

- Feature-complete version of OP Stack fault proofs released
- Testing on Sepolia testnet to validate the system
- Integration with the broader OP Stack ecosystem
- Focus on performance optimization and gas efficiency

## Current Limitations

Some limitations of the current implementation include:

- Limited system call support in the FPVM
- No support for concurrency features
- Computational overhead for generating and verifying proofs
- Complexity of the bisection protocol

## Relationship with OP Stack

The fault proof system is a critical component of the OP Stack, providing security guarantees for Layer 2 transactions. It works alongside other components like:

- Rollup contracts
- Sequencers
- Data availability solutions

This research provides a foundation for understanding how Optimism's fault proof system works and how it might be enhanced through integration with other technologies like opML and oppAI.
