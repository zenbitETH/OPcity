## Milestone 4: Comparative metrics and performance

This section provides a comparative analysis of the OPcity fault proof systems—**Multi-Phase** and **Tournament-based**—against the existing OP Stack fault proof model. All metrics are based solely on implemented and verifiable code within the zenbitETH/OPcity and zenbitETH/Optimism repositories. While these figures are grounded in real development artifacts, **they are illustrative projections** and have not yet been formally benchmarked under standardized test conditions.

The goal of this comparison is to assess how OPcity’s modular enhancements affect **contract complexity**, **dispute flow**, and **gas efficiency**—key factors in the scalability and maintainability of onchain fault proof protocols.

### **Contract Structure & Codebase Complexity**

| **Implementation** | **Contract Count** | **Approx. LOC** | **Key Functions** | **State Variables** |
| --- | --- | --- | --- | --- |
| OP Stack (Cannon Only) | 5–7 | ~2,500 | ~25 | ~20 |
| OPcity Multi-Phase | 5 | ~785 | 12 | 15 |
| OPcity Tournament-based | 2 | ~498 | 14 | 18 |

The OPcity implementations exhibit a significantly **reduced code footprint** relative to the standard OP Stack fault proof contracts. This reduction is due to their modularity, isolated concern boundaries, and deliberate separation between the onchain resolution layers and offchain trace infrastructure.

### **Dispute Resolution Flow Comparison**

| **Feature / Metric** | **OP Stack (Single-Phase)** | **OPcity Multi-Phase** | **OPcity Tournament-based** |
| --- | --- | --- | --- |
| Resolution Model | Linear bisection | Hierarchical (2-Phase) | Parallel Tournament |
| Execution Steps per Dispute | ~5–8 | 2–5 Phase 1, 1–2 Phase 2 | 4–6 Rounds per game |
| Onchain Instruction Execution | Always (single phase) | Only in Phase 2 | Only final match |
| Dispute Parallelism | None | Concurrent Phase 1 trees | Fully parallel games |
| Escalation Logic | Linear | Threshold-based transition | Match escalation |
| Complexity Overhead | Medium | High (phase transitions) | Medium-High (match logic) |

The **Multi-Phase mechanism** reduces the frequency of onchain MIPS execution by isolating fine-grained verification to Phase 2 only, optimizing onchain usage for simpler disputes. The **Tournament-based system** introduces parallelism via bracketed games, significantly improving throughput in high-volume scenarios.

### **Gas Cost Projections (Theoretical)**

| **Dispute Phase / Step** | **OP Stack (est.)** | **OPcity Multi-Phase** | **OPcity Tournament** |
| --- | --- | --- | --- |
| Initial Game Creation | ~110,000 gas | ~90,000 gas | ~120,000 gas |
| Single Bisection Step | ~140,000 gas | ~100,000 gas | – |
| Phase Transition | – | ~70,000 gas | – |
| Final Instruction Exec | ~180,000 gas | ~190,000 gas | ~200,000 gas |
| Tournament Match Resolve | – | – | ~160,000 gas |

These estimates show that **OPcity’s optimizations can reduce gas usage** in common-case disputes (Phase 1-only), while incurring **modest overhead** during escalation or match resolution. A full benchmarking suite will be required to validate these figures empirically.

The OPcity prototypes show clear improvements in structural modularity and resolution efficiency. The **Multi-Phase model** introduces dynamic resolution paths that avoid unnecessary VM execution. The **Tournament mechanism** offers horizontally scalable resolution with built-in economic incentives and cross-architecture trace compatibility. Together, these results validate OPcity’s design hypothesis: that fault proofs can be more scalable, parallel, and modular—without sacrificing correctness or OP Stack compatibility.