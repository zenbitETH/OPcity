## Milestone 6: OPcity Stack OSS repository

The [**OPcity GitHub repository**](https://github.com/zenbitETH/OPcity) serves as the primary workspace for documenting, developing, and reviewing the fault proof systems that power the OPcity prototype. This repository separates the research, design artifacts, and standalone code modules for both the **opML-inspired multi-phase mechanism** and the **DAVE-inspired tournament-based mechanism**, enabling modular development and community peer review.

In contrast, the [**Zenbit/Optimism fork**](https://github.com/zenbitETH/optimism) contains the **integrated implementation** of both fault proof mechanisms within the OP Stack. It is designed to demonstrate functional compatibility with Optimism’s architecture and serves as the candidate repository for submitting an official upgrade proposal to the OP Stack in the coming weeks.

This dual-repository strategy allows for clear separation of concerns:

- zenbitETH/OPcity: Research, specification, and per-mechanism development
- zenbitETH/Optimism: Canonical integration into the OP Stack for validation, benchmarking, and eventual protocol upgrade proposals

### [**Zenbit/Optimism (Integrated OP Stack Implementation)**](https://github.com/zenbitETH/optimism)

1. [**op-challenger/game/Tournament**](https://github.com/zenbitETH/optimism/tree/develop/op-challenger/game/tournament)
    
    Contains the off-chain logic for the DAVE-inspired tournament-based dispute system, including match orchestration (match_monitor.go), automated agents (agent.go, player.go), trace providers (trace_provider.go, riscv_trace_provider.go), and cross-VM execution trace handling.
    
2. [**op-challenger/game/Multiphase**](https://github.com/zenbitETH/optimism/tree/develop/op-challenger/game/multiphase)
    
    Implements the opML-style multi-phase engine, including recursive Phase 1 bisection (phase1_verifier.go), instruction-level Phase 2 logic (phase2_verifier.go), state handling via the Lazy Loading Manager, and game automation (engine.go).
    
3. [**packages/contracts-bedrock/src/dispute**](https://github.com/zenbitETH/optimism/tree/develop/packages/contracts-bedrock/src/dispute)
    
    Contains Solidity contracts for managing disputes across both mechanisms. This includes MultiPhaseResolver.sol, PhaseSelector.sol, Tournament.sol, and TournamentFactory.sol, as well as interfaces and shared utilities for dispute orchestration and protocol upgrades.
    

### [**Zenbit/OPcity (Research & Module Development)**](https://github.com/zenbitETH/OPcity)

1. [**FPVM Research**](https://github.com/zenbitETH/OPcity/tree/main/FPVM-Research)
    
    Documents design experiments on the Canon FPVM, including proposals for trace extensions, memory architecture, and RISC-V compatibility. Contains detailed architecture diagrams and VM modeling experiments that guided OPcity’s modular VM interaction.
    
2. [**Milestone Reports**](https://github.com/zenbitETH/OPcity/tree/main/Milestone-Reports)
    
    Contains published research outputs such as the *OPcity Research Report* and *Development Report*. Each report summarizes findings, design trade-offs, benchmarking results, and links to relevant PRs or modules.
    
3. [**OPstack Research**](https://github.com/zenbitETH/OPcity/tree/main/Op-stack-research)
    
    Aggregates protocol upgrade reviews, version benchmarks, and specification mappings across OPstack releases. Includes the comparative protocol upgrade table and supporting notes used to identify compatibility targets for OPcity.
    
4. [**OPcity Multi-Phase FP**](https://github.com/zenbitETH/OPcity/tree/main/OPcity-Multi-Phase-FP)
    
    Contains the standalone implementation and logic of the opML-inspired multi-phase system. Includes Markdown diagrams, dispute flow descriptions, claim formats, and off-chain strategy notes.
    
5. [**OPcity Tournament-based FP**](https://github.com/zenbitETH/OPcity/tree/main/OPcity-Tournament-based-FP)
    
    Provides the design overview and implementation details for the DAVE-style tournament system, including modular claim handling, game rounds, match resolution flow, and RISC-V integration strategy.
    

### Conclusion

The OPcity Fault Proofs project delivers a dual-mechanism prototype that enhances the OP Stack’s dispute resolution architecture through modularity, scalability, and execution extensibility. By combining the **opML multi-phase flow** and **DAVE tournament model**, OPcity demonstrates how alternative dispute mechanisms can coexist atop a canonical execution engine (Cannon FPVM), while introducing paths toward **RISC-V support**, **cross-architecture verification**, and **private proof systems**.

Through extensive off-chain infrastructure, robust on-chain arbitration, and carefully scoped design trade-offs, OPcity provides a blueprint for the next generation of modular, fault-proof infrastructure in optimistic rollups. With the full system now integrated into the OP Stack via the zenbitETH/Optimism fork, the next milestone is to propose this work as a formal protocol upgrade for evaluation by the Optimism Collective and broader L2 research community.