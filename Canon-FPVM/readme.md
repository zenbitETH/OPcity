<script type="text/javascript" src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=default"></script>

# Introduction
Fraud/fault proof is a scalability mechanism used in blockchain systems to detect and prove that a particular block or transaction is invalid through merkalized data, allowing nodes in the network to receive and verify block states without downloading the entire blockchain, assuming that there is at least one honest node willing to generate them.<sup>[1](https://arxiv.org/abs/1809.09044)</sup> The term was [introduced in 2019 in the _Fraud and Data Availability Proofs_](https://arxiv.org/abs/1809.09044) article to propose the mechanism as a scalability alternative to payment and state channels, and aimed to play a key role in sharding implementation.

This mechanism is also known as interactive game<sup>[2](https://specs.optimism.io/fault-proof/index.html)</sup> where, any participant optimistacally assume that every proposed result is valid by default<sup>[3](https://arxiv.org/abs/2401.17555)</sup> and works as the proof system for the Optimistic Rollups, to determine that an invalid state transition took place according to the protocol rules and enable dispuete games that rewards honest participants for helping to mantain the integrity of the system.<sup>[4](https://l2beat.com/glossary#fraud-proof)</sup>

## Parts of Fault Proof systems
### 1. Program 
The Fault Proof Program defines the verification of claims of the state-transition outputs of the L2 rollup as a pure function of L1 data. The op-program is the reference implementation of the program, based on op-node and op-geth implementations. The program consists of:
- [Prologue](https://specs.optimism.io/fault-proof/index.html#prologue): load the inputs, given minimal bootstrapping, with possible test-overrides.
- [Main content](https://specs.optimism.io/fault-proof/index.html#main-content): process the L2 state-transition, i.e. derive the state changes from the L1 inputs.
- [Epilogue](https://specs.optimism.io/fault-proof/index.html#epilogue): inspect the state changes to verify the claim.

### 2. Fault Proof Virtual Machine
A fault proof VM relies on a fault proof program to provide an interface for fetching any missing pre-images based on hints. The VM emulates the program, as prepared for the VM target architecture, and generates the state-root or instruction proof data as requested through the VM CLI. A fault proof VM implements:
- a smart-contract to verify a single execution-trace step, e.g. a single MIPS instruction.
- a CLI command to generate a proof of a single execution-trace step.
- a CLI command to compute a VM state-root at step N

### 2a. Types of Fault Proofs VM:
1. [Cannon](https://github.com/ethereum-optimism/cannon): big-endian 32-bit MIPS proof, by OP Labs, in active development.
2. [cannon-rs](https://github.com/anton-rs/cannon-rs): Rust implementation of Cannon, by @clabby, in active development.
3. [Asterisc](https://github.com/protolambda/asterisc): little-endian 64-bit RISC-V proof, by @protolambda, in active development.

### 3. Interactive Dispute Games
The interactive dispute game allows actors to resolve a dispute with an onchain challenge-response game that bisects to a disagreed block n→n+1 state transition, and then over the execution trace of the VM which models this state transition, bounded with a base-case that proves a single VM trace step. The game is multi-player: different non-aligned actors may participate when bonded. Response time is allocated based on the remaining time in the branch of the tree and alignment with the claim. The allocated response time is limited by the dispute-game window, and any additional time necessary based on L1 fee changes when bonds are insufficient.

## Canon Fault Proofs
The Cannon FPVM emulates a minimal Linux-based system running on big-endian 32-bit MIPS32 architecture. A lot of its behaviors are copied from Linux/MIPS with a few tweaks made for fault proofs. For the rest of this doc, we refer to the Cannon FPVM as simply the FPVM.

Operationally, the FPVM is a state transition function. This state transition is referred to as a Step, that executes a single instruction. We say the VM is a function _f_, given an input state _S<sub>pre</sub>_, steps on a single instruction encoded in the state to produce a new state _S<sub>post</sub>_
​

$$ f(S_{pre})→ S_{post}$$

​
Thus, the trace of a program executed by the FPVM is an ordered set of VM states.

### State
The virtual machine state highlights the effects of running a Fault Proof Program on the VM. It consists of the following fields:

1. `memRoot` - A bytes32 value representing the merkle root of VM memory.  
2. `preimageKey` - bytes32 value of the last requested pre-image key.  
3. `preimageOffset` - The 32-bit value of the last requested pre-image offset.  
4. `pc` - 32-bit program counter.  
5. `nextPC` - 32-bit next program counter. Note that this value may not always be pc+4 when executing a branch/jump delay slot.  
6. `lo` - 32-bit MIPS LO special register.  
7. `hi` - 32-bit MIPS HI special register.  
8. `heap` - 32-bit base address of the most recent memory allocation via mmap.  
9. `exitCode` - 8-bit exit code.  
10. `exited` - 8-bit indicator that the VM has exited.  
11. `step` - 8-byte step counter.  
12. `registers` - General-purpose MIPS32 registers. Each register is a 32-bit value.

The state is represented by packing the above fields, in order, into a 226-byte buffer.  

## References
<sup>[1](https://arxiv.org/abs/1809.09044)</sup> Al-Bassam, M., Sonnino, A., & Buterin, V. (2019). _Fraud and Data Availability Proofs: Maximising Light Client Security and Scaling Blockchains with Dishonest Majorities._

<sup>[2](https://specs.optimism.io/fault-proof/index.html)</sup> Optimism (2024). _OP Stack Specification_

<sup>[3](https://arxiv.org/abs/2401.17555)</sup> Conway, KD., So, C., Yu, X., & Wong, K. (2024). _opML: Optimistic Machine Learning on Blockchain_


<sup>[4](https://l2beat.com/glossary#fraud-proof)</sup> L2 Beat (2024). _Glossary_
