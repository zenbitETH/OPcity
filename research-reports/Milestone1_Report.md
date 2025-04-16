Ethereum has long envisioned itself as the World Computer—a shared global platform for registering and verifying the state of public goods, collective decisions, and decentralized markets. By embedding trust directly into code and consensus, Ethereum enables a new kind of programmable legitimacy. But for these values to reach governments, organizations, and everyday people, the technology must scale.

Layer 2 (L2) rollups have emerged as a practical solution to Ethereum’s scalability challenges, addressing issues like cost, capacity, and transaction throughput. Among these, optimistic rollups allow transactions to be executed off-chain while posting minimal data on-chain, preserving decentralization and network security.

The OP Stack, developed by OP Labs and coordinated by the Optimism Foundation, is one of the most widely adopted open-source frameworks for launching dedicated L2 rollups. Its minimal, modular design allows chains to post periodic state roots to Ethereum (or alternative L1s), relying on optimistic validation. These state roots can be challenged within a seven-day window using Fault Proof mechanisms. A robust and permissionless implementation of these mechanisms is essential to meet Stage 1 requirements in L2BEAT’s state validation maturity framework.

This research report explores the evolution and implementation of fault proof systems in the OP Stack. It builds on our earlier work, [OP City: Research and Optimization of OP Stack Deployments and Canon Fault Proofs VM](https://mirror.xyz/zenbit.eth/atHQ_Nz1--bQbY6Vx7NzYO-aOtXibcQi70aLl_WhTzY), and presents findings across three milestones: a comprehensive review of OP Stack upgrades and their impact on fault proofs; a technical analysis of emerging mechanisms like Canon FPVM, opML, oppAI, and DAVE; and a comparative benchmark to assess readiness, performance, and integration paths. Together, these contributions aim to advance the decentralization and resilience of the OP Stack—and by extension, Ethereum as a truly global computational engine.

# Milestone 0: OP stack research

## OP stack background

Before the OP Stack was formalized as a unified, modular framework, its architecture evolved through iterative stages that laid the groundwork for scalable Layer 2 execution. The journey began with the **Unipig Demo** in October 2019, a proof-of-concept built in collaboration with Uniswap that showcased the power of Optimistic Rollups. It demonstrated a ~10× increase in throughput and significant reductions in gas fees compared to Ethereum mainnet, proving the concept’s viability. 

This was followed by the **SNX Testnet** in September 2020, which introduced general-purpose smart contract execution on Layer 2. Transactions ran at roughly a tenth of L1’s cost, with improved latency and faster confirmation times. In January 2021, the network transitioned into its first mainnet phase—often referred to as the **OVM Era**—with a version of the Optimistic Virtual Machine that relied on a custom Solidity transpiler. 

While functional, this setup introduced significant inefficiencies: over 25,000 lines of bespoke code, complex tooling, and inflated state transition costs. By October 2021, Optimism rolled out the **EVM Equivalence Upgrade**, a major step forward that removed the transpiler, drastically simplified the execution layer, and brought the network into full compatibility with Ethereum tooling. Throughput increased to ~100 TPS, and developer experience improved significantly. Just two months later, in December 2021, Optimism launched an **Open Mainnet**, allowing public contract deployment and catalyzing adoption by removing whitelists. This open access marked a turning point in Optimism’s growth. 

The culmination of these early experiments arrived with the **Bedrock Upgrade** in 2023, which replaced legacy OVM components with a leaner, modular, and Ethereum-equivalent design. Bedrock slashed code complexity by ~90%, cut transaction costs by ~30%, and boosted throughput to ~450 TPS—ushering in the era of the OP Stack and laying the foundation for a Superchain of standardized, interoperable rollups.

<figure>
  <img src="./img/prebedrock.webp" alt="Evolution of the OP stack">
</figure>


From Optimism Collective Mirror

## OP stack protocol upgrades

Since the formalization of the OP Stack, the protocol has undergone **15 official upgrades**, ranging from foundational architectural transitions to fine-grained feature releases across the OP Stack’s modular layers. Our review spanned **62 curated sources**, including governance forum proposals, developer documentation, audit reports, and OP Stack specification entries. These materials were analyzed to trace the evolution of the OP Stack’s dispute system—from the launch of permissionless Cannon-based fault proofs in Protocol Upgrade #7, to the infrastructure pre-requisites for multi-threaded MIPS64 fault games in Upgrades #14 and #15. Sources such as governance threads, public audit PDFs, and design documents not only clarified technical intent, but also allowed us to assess how each upgrade contributed to meeting **Stage 1 decentralization criteria** from the L2BEAT framework.

In our review, we identified **13 upgrades** with direct impact on the **settlement layer**, and notably, **7 of these upgrades** introduced substantial improvements to the **fault proof system**—transforming it from a trusted fallback into a fully modular, permissionless verification layer.

*All referenced documents are available in our open repository for transparency and further review:*

🔗 [OP Stack Protocol Upgrades Review – Zenbit GitHub](https://github.com/zenbitETH/OPcity/tree/main/Op-stack-research/Protocol-Updates)

### Protocol Upgrade #7 Fault Proof

The state validation improvements began with **Protocol Upgrade #7: Fault Proofs**, which marked a foundational shift in the OP Stack’s security model by replacing the previously trusted state root proposer mechanism with a permissionless fault proof system. This upgrade introduced a modular, on-chain binary bisection game powered by the Cannon VM, allowing any participant to propose or challenge state roots. Supporting this dispute flow, key components such as the DisputeGameFactory, AnchorStateRegistry, and bonding mechanisms were deployed. Together, these features enabled the OP Stack to meet the **Stage 1 decentralization milestone** set by L2BEAT, especially through the integration of a **Security Council override** to pause withdrawals in emergencies.

### Protocol Upgrade #8 Guardian

To solidify this backstop, **Protocol Upgrade #8: Guardian** introduced formal guardian governance. This upgrade extended the Security Council’s role by granting them authority over L2 ProxyAdmin control and emergency withdrawal pausing, providing a critical safeguard during fault proof escalation periods. It was a vital step in completing the governance infrastructure required for fault proof decentralization, ensuring that challenge games could proceed trustlessly without introducing systemic risk.

### Protocol Upgrade #10 Granite

Building on this foundation, **Protocol Upgrade #10: Granite** focused on enhancing the protocol infrastructure to support more advanced proof systems. Updates to SystemConfig and dispute game interfaces increased flexibility for challenger clients and made space for introducing alternative VMs, such as MIPS64. These changes future-proofed the OP Stack’s dispute layer, ensuring compatibility with evolving fault proof implementations like MT-Cannon and potential zk variants.

### Protocol Upgrade #11 Holocene

As fault proofs moved into production, **Protocol Upgrade #11: Holocene** standardized the format for L2-to-L1 outputs, refining how withdrawals and challenge games would be validated across chains. These improvements enhanced compatibility with new proof types and helped streamline the verification logic for post-upgrade withdrawals, an essential update as the system transitioned away from trusted third parties.

### Protocol Upgrade #12 Pre pectra

In preparation for Ethereum’s Pectra hard fork, **Protocol Upgrade #12: Pre-Pectra Readiness** introduced important features to support hybrid and forward-compatible proofs. This included access to the **L1 Beacon Root**, enabling withdrawal verification against Ethereum’s consensus layer, and introduced precompiles like BLS and timestamp access that are critical for next-gen dispute VMs. These changes ensured that fault proofs could continue to operate in a post-Dencun Ethereum environment and opened the door to zk-fault proof hybrids.

### Protocol Upgrade #14 Pre Isthmus

This set the stage for **Protocol Upgrade #14: Pre-Isthmus**, which deployed the **L1 infrastructure for MT-Cannon**, a multithreaded, 64-bit evolution of the Cannon VM. Without requiring a hard fork, this upgrade delivered new contracts like OPChainManagerV2, deployed MIPS64.sol, and restructured the fault game logic to support more scalable dispute execution. MT-Cannon could now be tested in production environments in parallel with legacy Cannon, serving as a trial run for the more performant and deterministic VM architecture.

### Protocol Upgrade #15 Isthmus

Finally, **Protocol Upgrade #15: Isthmus** completed the transition by activating MT-Cannon as the canonical fault proof VM. This hard fork integrated Pectra-compatible features, such as support for new precompiles and the inclusion of withdrawalsRoot in L2 block headers, solidifying the OP Stack’s ability to anchor dispute data efficiently. It also introduced operator fee fields and formally recognized MIPS64 in dispute logic, ensuring that fault proofs could now run faster, more securely, and with greater parallelism.

<figure>
  <img src="./img/PUOPstack.png" alt="Evolution of the OP stack">
</figure>

# Milestone 1: Fault Proofs mechanisms research

Fault proofs are a foundational component of optimistic rollups, enabling trustless verification of off-chain state transitions. As the OP Stack advances toward greater decentralization and modularity, the design and implementation of fault proof mechanisms have evolved significantly. This milestone documents our research into the key architectures shaping this landscape—beginning with Optimism’s default Canon FPVM, and expanding into cutting-edge systems like opML, oppAI, and Cartesi’s DAVE. Each mechanism offers unique trade-offs between performance, verifiability, and scalability, collectively enriching the design space for secure and efficient Layer 2 systems.

*All referenced documents are available in our open repository for transparency and further review:*

🔗 [FP research reports – Zenbit GitHub](https://github.com/zenbitETH/OPcity/tree/812bf40acc013b60b96fd54c3b3a2c932104e8ba/research-reports)

## **Optimism's Fault Proofs & Canon FPVM**

At the core of the OP Stack’s security model lies the **Fault Proof system**, a cryptoeconomic mechanism that ensures the correctness of off-chain state transitions by enabling anyone to challenge invalid claims posted to Ethereum. This milestone focused on understanding and documenting the evolution of this mechanism, beginning with **Optimism’s Canon Fault Proof Virtual Machine (Canon FPVM)**—the default system currently securing Layer 2 outputs.

Canon implements a **dispute resolution game**, where participants engage in an interactive bisection protocol to isolate a single disputed MIPS instruction. The system combines **onchain and offchain components**: MIPS.sol, a lightweight smart contract that deterministically executes a single instruction, and **Cannon**, a Go-based offchain VM that computes the state transition trace and generates verifiable proofs. Key elements like the **DisputeGameFactory**, **AnchorStateRegistry**, and **OP-Challenger** infrastructure coordinate to handle challenges, execute proofs, and enforce results. Memory in Cannon is modeled as a **Merkleized 32-bit address space**, enabling stateless onchain verification of computation with minimal inputs.

This architecture was tested in production during the release of feature-complete fault proofs on OP Sepolia, and subsequently evolved through multiple protocol upgrades to enhance modularity, precision, and compatibility with Ethereum’s Pectra roadmap. Our technical deep dive revealed strengths—such as **deterministic execution**, **statelessness**, and modular game design—but also current limitations, including **limited syscall support** and high computational overhead.

## opML

One of the most forward-looking proposals in fault proof innovation is **opML**, an optimistic verification system designed to support scalable and verifiable onchain machine learning (ML) inference. Inspired by optimistic rollups, opML replaces expensive zero-knowledge proofs with an interactive fraud-proof process, allowing ML computations to be challenged only when disputed. This significantly reduces overhead, making it feasible to run complex models onchain without compromising trust. At its core, opML features a **Fraud Proof Virtual Machine (FPVM)** tailored for ML workloads, equipped with Merkleized memory and a segmented layout that separates code, input/output buffers, oracle data, and model parameters. This structure supports a full 32-bit address space and deterministic execution, essential for cryptographic verifiability.

The **opML architecture** consists of four integrated components: the FPVM, a dual-compilation **Machine Learning Engine (MLE)**, an interactive dispute game, and the protocol layer. The MLE is particularly innovative, compiling the same ML source code into two targets: one for high-speed native execution (with GPU/CUDA support), and another for verifiable execution on the FPVM, using fixed-point arithmetic to preserve determinism. In typical use, trusted actors submit computation results offchain, while verifiers can trigger a **two-phase interactive dispute**. The first phase isolates disputed computation nodes at the graph level with semi-native execution, while the second drills into low-level FPVM instructions to prove or disprove correctness.

This **multi-phase bisection protocol** improves on Optimism’s single-phase approach by tailoring fraud proof resolution to ML-specific workflows, striking a balance between performance and trust minimization. Disputes are resolved onchain using MIPS-like instructions executed within the FPVM, with proofs derived offchain using opML’s Merkleized memory model. Security is guaranteed through an **AnyTrust model**, where only one honest validator is needed to contest false claims. Submitters and challengers are incentivized with crypto-economic stakes and penalties to align behavior.

To further enhance scalability, opML incorporates **performance optimizations** like lazy loading—loading only required model segments into memory—and **semi-native execution**, where only disputed segments fall back to verifiable computation. This approach enables opML to support large-scale models (e.g. 7B LLaMA) while maintaining verifiability when needed. On the protocol side, smart contract interfaces manage submission, challenge periods, and dispute resolution workflows, enabling seamless integration with onchain systems. While opML still faces constraints—like fixed finality windows and memory limitations—it presents a robust pathway for integrating verifiable ML into rollup environments, aligning with the broader goals of modular and decentralized computation.

## oppAI

The **oppAI** framework introduces a novel hybrid architecture that merges the computational efficiency of **Optimistic Machine Learning (opML)** with the strong privacy guarantees of **Zero-Knowledge Machine Learning (zkML)**. Designed to enable secure, verifiable, and privacy-preserving AI inference on blockchain systems, oppAI addresses the critical tradeoff in onchain AI: maintaining both performance and confidentiality. By allowing developers to selectively partition AI models into components executed via optimistic or zero-knowledge verification, oppAI provides a flexible system tailored to diverse privacy and cost constraints.

At the core of oppAI lies a **dual execution model**. Model components deemed non-sensitive or performance-critical are processed via opML’s Fraud Proof Virtual Machine (FPVM), while sensitive logic or proprietary model weights are processed under zkML, using zero-knowledge proofs such as zk-SNARKs. This structure enables the same AI model to achieve high performance where needed, and high privacy where necessary. The **execution workflow** begins with a partitioned AI model—e.g., f₁ᵒᵖ, f₁ᶻᵏ, ..., fₙᵒᵖ, fₙᶻᵏ—where each submodel is designated for either optimistic or zk execution. After processing, the results are submitted to an onchain verifier, and a challenge-response game can be initiated if needed.

The **oppAI system architecture** builds on opML’s foundation and extends it with additional zkML capabilities. It includes components such as:

- The **FPVM** (adapted from opML) for step-by-step verifiability in optimistic disputes;
- A dual-compiled **Machine Learning Engine**, supporting native and FPVM-compatible execution;
- An **interactive dispute resolution game** for verifying opML results;
- A zkML **Prover** that generates ZKPs for private components, alongside an onchain **Verifier** contract to check them.

This design offers direct applications to **fault proof systems** like those used in the OP Stack. First, oppAI enables **privacy-preserving fault proofs**, allowing sensitive data (e.g. user inputs or proprietary models) to be verified without full exposure. Second, it proposes a **hybrid verification scheme** where only the most privacy-sensitive operations incur zk costs, while the rest benefit from faster optimistic execution. This creates a more efficient yet secure fraud proof system, especially relevant for increasingly complex computations such as AI oracles and ML-based coordination mechanisms in DAOs and L2 governance.

From a **security standpoint**, oppAI adopts an **AnyTrust model**, requiring only a single honest verifier to guarantee system integrity. Its **economic security model** disincentivizes attacks through dynamic inference pricing, calculated to keep the cost of reconstructing private models prohibitively high. Notably, the system allows tuning the ratio between zkML and opML (represented by p) to modulate the tradeoff between privacy and efficiency. Performance benchmarks confirm that reducing zkML usage (lower p) significantly lowers computational overhead while still enabling verification of critical model segments like attention layers in large language models (e.g. Stable Diffusion, LLaMA).

In practical terms, **oppAI’s capabilities could expand the scope of OP Stack-based chains** by allowing:

- **Secure AI oracles**, which deliver verifiable predictions to smart contracts without revealing internal logic;
- **Privacy-enhanced onchain inference**, protecting model IP while maintaining fraud-proof guarantees;
- **Support for high-complexity computation**, such as federated ML, game theory agents, and dynamic governance models.

Despite its advantages, oppAI introduces new **design and implementation challenges**. Partitioning models and determining which parts require privacy demands interdisciplinary knowledge of ML and blockchain. Furthermore, managing performance tradeoffs and adapting to rapidly evolving zkML and opML ecosystems require continual iteration. Nonetheless, oppAI sets a strong precedent for modular, privacy-aware fault proof design—bridging efficient onchain computation with the next generation of decentralized AI applications.

## DAVE

**DAVE** (a name, not an acronym) is Cartesi’s advanced fraud-proof protocol, designed to strengthen decentralization, liveness, and efficiency in blockchain verification. It introduces a permissionless, tournament-style dispute mechanism that offers a compelling alternative to Optimism’s Fault Proofs and Arbitrum’s BoLD.

At its core, DAVE leverages **Permissionless Refereed Tournaments (PRT)**—a novel framework that resists Sybil attacks, maintains constant hardware and bond requirements, and provides exponential resource advantages to honest actors. This design allows a single honest validator to defend the network—even against coordinated adversaries—at minimal cost.

DAVE runs atop the **Cartesi Machine**, a deterministic RISC-V emulator that separates logic into two layers: a micro-architecture implemented in Solidity and a high-performance computation layer offchain. This architecture ensures execution integrity while enabling support for complex instructions without bloating the onchain footprint.

Unlike traditional fault proof systems, DAVE uses a **tournament-style challenge mechanism** where validators engage in logarithmically-scaling dispute games. These tournaments resolve computation disagreements in **2–5 challenge periods**, significantly reducing the overhead seen in earlier systems like Cartesi’s own PRT (which could take up to 20 weeks).

Compared to other systems, DAVE offers notable benefits:

- **Against Optimism’s OPFP**, it brings better Sybil resistance, faster dispute finality, and lower validator costs.
- **Versus Arbitrum’s BoLD**, DAVE maintains similar security guarantees with much lower bond requirements (~3 ETH vs. 3,600 ETH).
- **Internally**, it improves liveness and reduces delay over Cartesi’s prior models while keeping decentralization intact.

DAVE’s flexible architecture allows it to serve both **rollups with continuous input processing** and **compute-focused systems**. While it currently relies on the Cartesi Machine, its **execution environment agnosticism** opens paths for integration with other rollup ecosystems.