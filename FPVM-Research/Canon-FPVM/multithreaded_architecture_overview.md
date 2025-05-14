## Architecture Analysis of Optimism's Fault Proof System

### Overall Architecture of op-cannon and the Fault Proof VM

#### Core Components
The op-cannon system is an onchain MIPS instruction emulator designed to enable EVM-equivalent fault proofs. It consists of two primary components:

#### 1. Onchain MIPS64.sol:
- An EVM implementation that verifies the execution of a single MIPS instruction
- Implements big-endian 64-bit MIPS instruction execution (MIPS64 Release 1)
- Simulates a minimal subset of the Linux kernel via syscall instructions
- Handles memory operations through a binary Merkle tree structure
#### 2. Offchain mipsevm:
- A Go implementation that produces proofs for MIPS instructions to verify onchain
- Executes instructions one step at a time while maintaining state
- Generates witness data for disputed steps
- Handles delay-slots by isolating individual instructions

#### MT-Cannon Upgrade (64-bit MIPS with Multi-threading)
The MT-Cannon upgrade extends the original Cannon implementation with multi-threading capabilities:

#### 1. Thread Management:
- Implemented in the multithreaded package
- Each thread has its own state (ThreadState struct) with:
    - Thread ID
    - Exit code and exit status
    - CPU state (PC, NextPC, LO, HI registers)
    - 32 general-purpose registers
#### 2. State Structure:
- The State struct in the multithreaded package maintains:
    - Memory (shared between threads)
    - Preimage key and offset for external data
    - Heap management
    - Load-linked (LL) reservation status for atomic operations
    - Thread stacks (left and right)
    - Context switching mechanism
#### 3. Execution Model:
- Uses a deterministic thread scheduling algorithm
- Tracks steps since last context switch
- Maintains thread stacks and traversal direction
- Supports atomic operations through LL/SC (Load-Linked/Store-Conditional) instructions

### Structure and Functionality of op-challenger
The `op-challenger` is a modular challenge agent for dispute games, including fault games. It interacts with op-cannon to resolve disputes.

#### Key Components
1. **Game Monitoring:**
    - Monitors the DisputeGameFactory contract for new games
    - Tracks game status and claims
2. **Trace Providers:**
    - Modular system supporting multiple trace types (Cannon, Asterisc)
    - Each provider implements a common interface for generating and validating execution traces
    - Trace providers are in `optimism/op-challenger/game/fault/trace/`
3. **Claim Management:**
    - Validates claims against locally computed traces
    - Generates counter-claims when invalid claims are detected
    - Implements bisection protocol to narrow down disputed execution steps
4. **Dispute Resolution:**
    - Resolves claims when they reach the base case (single instruction)
    - Executes the disputed instruction onchain to determine truth

#### Interaction with op-cannon
The op-challenger interacts with op-cannon through:

1. **Cannon Provider** (`optimism/op-challenger/game/fault/trace/cannon/provider.go`):
    - Manages the execution of the Cannon VM
    - Generates state hashes at specific execution steps
    - Converts between VM state and dispute game claim format
2. **State Converter** (`optimism/op-challenger/game/fault/trace/cannon/state_converter.go`):
    - Translates between Cannon VM state and the format used in the dispute game


### Current Implementation of the Dispute Game and Bisection Protocol
The dispute game is implemented in the `FaultDisputeGameContract` and related contracts:

#### Game Structure
1. **Claim Tree:**
    - Root claim represents the final state after execution
    - Claims form a tree where each node is challenged by its children
    - Depth of the tree corresponds to the bisection of the execution trace
2. **Bisection Protocol:**
    - When a claim is challenged, the execution trace is bisected
    - The challenger must provide a counter-claim at the midpoint
    - This process continues recursively until a single disputed step is identified
    - The single step is then executed onchain to determine the correct state
3. **Game Resolution:**
    - Games have a maximum duration (maxClockDuration)
    - Claims can be resolved when they can no longer be challenged
    - The game result determines if the original claim was valid

#### Contract Implementation
The dispute game contracts include:

1. **DisputeGameFactory**:
    - Creates new dispute games
    - Registers different game types (Cannon, Asterisc)
2. **FaultDisputeGame:**
    - Manages the claim tree
    - Implements attack/defend moves
    - Handles resolution logic
3. **VM Contracts:**
    - `MIPS64.sol` for Cannon
    - `RISCV.sol` for Asterisc
    - Execute the disputed instruction onchain


### MT-Cannon Upgrade Implementation

The MT-Cannon upgrade adds multi-threading support to the original Cannon implementation:

#### Key Components
1. Thread Management:
    - Thread state structure in `optimism/cannon/mipsevm/multithreaded/thread.go`
    - Thread scheduling and context switching in `state.go`
2. Memory Model:
    - Shared memory between threads
    - Atomic operations support through LL/SC instructions
    - Memory consistency model
3. Witness Data:
    - Extended witness format to include thread information
    - Thread state serialization for proofs
4. Execution Model:
    - Deterministic thread scheduling
    - Context switching based on instruction count
    - Support for thread creation and termination