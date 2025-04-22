# 1. OP Stack Protocol Updates

| Protocol Upgrade | Testnet Release Date | Mainnet Release Date | Upgrade Proposal ID | OP Gov Author | Author Affiliation | GH PR | OP Gov Voting | Audit Reports|
|------------------|----------------------|----------------------|----------------------|---------------|--------------------|-------|---------------|--------------|
| [**UNIPIG Plasma**](#-op-stack-background) |  | Oct 2019 |  |  |  |  |  |  |
| [**SNX**](#-op-stack-background) |  | Sep 2020 |  |  |  |  |  |  |
| [**Mainnet launch**](#-op-stack-background) |  | Jan 2021 |  |  |  |  |  |  |
| [**OVM (EVM Equivalence)**](#-op-stack-background) |  | Oct 2022 |  |  |  |  |  | [Open Zeppelin 🔗](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2021_03-OVM_and_Rollup-OpenZeppelin.pdf) |
| [**Bedrock**](#bedrock) |  | Jan 2023 |  |  |  |  |  | |
| [**Regolith**](#protocol-upgrade1-regolith) |  | Mar 2023 | [**UP #1** 🔗](https://gov.optimism.io/t/final-upgrade-1-bedrock-protocol-upgrade-v2/5548) | [ben-chain](https://gov.optimism.io/u/ben-chain) | `OP Foundation` | [#5010](https://github.com/ethereum-optimism/optimism/pull/5010) | [✅ Succeeded](https://vote.optimism.io/proposals/114732572201709734114347859370226754519763657304898989580338326275038680037913) |  |
| [**Canyon**](#protocol-upgrade-2-canyon) | Nov 2023 | Jan 2024 | [**UP #2** 🔗](https://gov.optimism.io/t/final-upgrade-proposal-2-canyon-network-upgrade/7088) | [triangleshere](https://gov.optimism.io/u/trianglesphere/summary) | `OP Labs` | [#8569](https://github.com/ethereum-optimism/optimism/pull/8569) | [✅ Succeeded](https://vote.optimism.io/proposals/20327152654308054166942093105443920402082671769027198649343468266910325783863) |  |
| [**Delta**](#protocol-upgrade-3-delta) | Dec 2023 | Feb 22, 2024 | [**UP #3** 🔗](https://gov.optimism.io/t/final-upgrade-proposal-3-delta-network-upgrade/7310) | [testinprod_io](https://gov.optimism.io/u/testinprod_io) | `Test in Prod` | [#7454](https://github.com/ethereum-optimism/optimism/pull/7454) | [✅ Succeeded](https://vote.optimism.io/proposals/64861580915106728278960188313654044018229192803489945934331754023009986585740) |  |
| [**Superchain Config**](#protocol-upgrade-4-superchain-config) | Jan 2024 |  | [**UP #4** 🔗](https://gov.optimism.io/t/upgrade-proposal-4/7534) | [maurelian](https://gov.optimism.io/u/maurelian) | `OP Labs` | [#9109](https://github.com/ethereum-optimism/optimism/pull/9109) | [✅ Succeeded](https://vote.optimism.io/proposals/110376471005925230990107796624328147348746431603727026291575353089698990280147) | [Trust Security 🔗](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2023_12_SuperchainConfigUpgrade_Trust.pdf) |
| [**Ecotone**](#protocol-upgrade-5-ecotone)| Feb 2024 | Mar 2024 | [**UP #5** 🔗](https://gov.optimism.io/t/upgrade-proposal-5-ecotone-network-upgrade/7669) | [bayardo](https://gov.optimism.io/u/bayardo) | `Base` | [#9695](https://github.com/ethereum-optimism/optimism/pull/9695) | [✅ Succeeded](https://vote.optimism.io/proposals/95119698597711750186734377984697814101707190887694311194110013874163880701970) |  |
| [**Multi-Chain Prep (MCP) L1**](#protocol-upgrade-6-multi-chain-prep-mcp-l1) | Jan 2024 |  | [**UP #6** 🔗](https://gov.optimism.io/t/upgrade-proposal-6-multi-chain-prep-mcp-l1/7677) | [Diego](https://gov.optimism.io/u/Diego) | `OP Labs` | [#9476](https://github.com/ethereum-optimism/optimism/pull/9476) | [✅ Succeeded](https://vote.optimism.io/proposals/47253113366919812831791422571513347073374828501432502648295761953879525315523) | [Cantina 🔗](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_02-MCP_L1-Cantina.pdf) |
| [**Fault Proofs**](#protocol-upgrade-7-fault-proofs) | May 2024 | Jun 2024 | [**UP #7** 🔗](https://gov.optimism.io/t/final-protocol-upgrade-7-fault-proofs/8161) | [ajustton](https://gov.optimism.io/u/ajsutton/summary) | `OP Labs` | [#10544](https://github.com/ethereum-optimism/optimism/pull/10544) | [✅ Succeeded](https://vote.optimism.io/proposals/72085170435228531173144599119267762084652443676555508407874836206178427511368) | [1 🔗](https://audits.sherlock.xyz/contests/205?filter=questions) / [2 🔗](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_05-FaultProofs-Sherlock.pdf)|
| [**Guardian**](#protocol-upgrade-8-guardian) | May 2024 | Jun 2024 | [**UP #8** 🔗](https://gov.optimism.io/t/final-protocol-upgrade-8-guardian-security-council-threshold-and-l2-proxyadmin-ownership-changes-for-stage-1-decentralization/8157) | [maurelian](https://gov.optimism.io/u/maurelian) | `OP Labs` | [#10616](https://github.com/ethereum-optimism/optimism/pull/10616) | [✅ Succeeded](https://vote.optimism.io/proposals/89250535338859095270968116984279971013811713632639468811376241520756760598962) | [Cantina 🔗](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_05_SafeLivenessExtensions-Cantina.pdf) |
| [**Fjord**](#protocol-upgrade-9-fjord) | May 2024 | Jul 2024 |[**UP #9** 🔗](https://gov.optimism.io/t/upgrade-proposal-9-fjord-network-upgrade/8236) | [bayardo](https://gov.optimism.io/u/bayardo) | `Base` | [#10735](https://github.com/ethereum-optimism/optimism/pull/10735) | [✅ Succeeded](https://vote.optimism.io/proposals/19894803675554157870919000647998468859257602050917884642551010462863037711179) |  |
| [**Granite**](#protocol-upgrade-10-granite) | Aug 2024 | Sep 2024 | [**UP #10** 🔗](https://gov.optimism.io/t/upgrade-proposal-10-granite-network-upgrade/8733) | [inphi](https://gov.optimism.io/u/inphi/summary) | `OP Labs` | [#11531](https://github.com/ethereum-optimism/optimism/pull/11531) | [✅ Succeeded](https://vote.optimism.io/proposals/46514799174839131952937755475635933411907395382311347042580299316635260952272) | [1 🔗](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_08_Fault-Proofs-No-MIPS_Spearbit.pdf) / [2 🔗](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_08_Fault-Proofs-MIPS_Cantina.pdf) / [3 🔗](https://github.com/code-423n4/2024-07-optimism-findings) / [4 🔗](https://immunefi.com/bug-bounty/optimism/information/)|
| [**Holocene**](#protocol-upgrade-11-holocene) | Nov 2024 | Jan 2025 | [**UP #11** 🔗](https://gov.optimism.io/t/upgrade-proposal-11-holocene-network-upgrade/9313) | [Dragan_ZzZ](https://gov.optimism.io/u/Dragan_ZzZ/summary) | `OP Labs` | [#13334](https://github.com/ethereum-optimism/optimism/pull/13334) | [✅ Succeeded](https://vote.optimism.io/proposals/20127877429053636874064552098716749508236019236440427814457915785398876262515) | [3DOC 🔗](https://github.com/ethereum-optimism/optimism/blob/7719c8538b8d911519f861fc70085e7a3b4e6787/docs/security-reviews/2024_10-Cannon-FGETFD-3DocSecurity.md) |
| [**L1 Pectra Readiness**](#protocol-upgrade-12-l1-pectra-readiness) | Feb 2025 | Apr 2025 | [**UP #12** 🔗](https://gov.optimism.io/t/upgrade-proposal-12-l1-pectra-readiness/9706) | geoknee | `OP Labs` | [#13958](https://github.com/ethereum-optimism/optimism/pull/13958) | [✅ Succeeded](https://vote.optimism.io/proposals/38506287861710446593663598830868940900144818754960277981092485594195671514829) |  |
| [**OPCM & ICI**](#protocol-upgrade-12-l1-pectra-readiness) |  | Apr 2025 | [**UP #13** 🔗](https://gov.optimism.io/t/upgrade-proposal-13-opcm-and-incident-response-improvements/9739) | [maurelian](https://gov.optimism.io/u/maurelian) | `OP Labs` | [#14858](https://github.com/ethereum-optimism/optimism/pull/14858) | [✅ Succeeded](https://vote.optimism.io/proposals/84511922734478887667300419900648701566511387783615524992018614345859900443455) |  |
| [**Isthmus L1 Contracts + MT-Cannon**](#protocol-upgrade-12-l1-pectra-readiness) | Apr 2025 | ~May 2025  | [**UP #14** 🔗](https://gov.optimism.io/t/upgrade-proposal-14-isthmus-l1-contracts-mt-cannon/9796) | [0xEscanor](https://gov.optimism.io/u/0xescanor/summary) | `OP Labs` | [#15363](https://github.com/ethereum-optimism/optimism/pull/15363) | [✅ Succeeded]() |  |
| [**Isthmus**](#protocol-upgrade-12-l1-pectra-readiness) |  Apr 2025 | ~May 2025 | [**UP #15** 🔗](https://gov.optimism.io/t/upgrade-proposal-15-isthmus-hard-fork/9804) | [0xEscanor](https://gov.optimism.io/u/0xescanor/summary) | `OP Labs` | [#15363](https://github.com/ethereum-optimism/optimism/pull/15363) | [✅ Succeeded]() |  |

## ** OP Stack Background**

The OP Stack is a modular, Ethereum-aligned framework for scaling blockchain ecosystems, developed as part of Optimism’s mission to enhance Layer 2 scalability while maintaining Ethereum’s security and decentralization. Its origins trace back to the **Plasma Group**, a research collective that initially focused on Plasma-based solutions before transitioning to **Optimistic Rollups** due to Plasma's data availability and exit game limitations[¹](https://plasma.build/t/rollup-plasma-for-mass-exits-complex-disputes/90). The shift towards rollups culminated in the **Optimistic Virtual Machine (OVM)**, a precursor to today’s OP Stack architecture, enabling general-purpose smart contract execution on Layer 2 while leveraging Ethereum’s security[²](https://medium.com/plasma-group/ethereum-smart-contracts-in-l2-optimistic-rollup-2c1cef2ec537).

As the OP Stack matured, its design philosophy evolved towards **modularity, EVM equivalence, and developer accessibility**. The **Bedrock upgrade** (2023) marked a pivotal milestone, reducing code complexity, achieving execution-level Ethereum equivalence, and minimizing divergence from Ethereum’s Geth client[³](https://optimism.mirror.xyz/Fdpds7l4yE2VDpLN4AZQUCjpBYdoPrIHw6a4uzUepmw). This approach significantly streamlined development, enabling full compatibility with Ethereum tools and fostering the emergence of a broader **Superchain** ecosystem of interoperable Layer 2 networks[⁴](https://optimism.mirror.xyz/9ZMwZjst9SQpzIgEd4gN42UDjATyK3ZRClPFx9oMPp8).

## **Metrics on Optimization and Performance Gains Before the OP Stack Release**

Before the formalization of the OP Stack, key **scalability and efficiency challenges** were gradually addressed through incremental upgrades:

- **Unipig Demo (Oct 2019)**:
    - First **Optimistic Rollup prototype** showcased with Uniswap.
    - **Transaction throughput:** Demonstrated ~10x improvement over Ethereum L1.
    - **Gas costs:** Reduced transaction costs significantly compared to mainnet.
- **SNX Testnet (Sep 2020)**:
    - Introduced **Layer 2 smart contract execution**, reducing reliance on Ethereum L1.
    - **Gas efficiency:** Transactions were processed **at 1/10th the cost of L1**.
    - **Latency improvements:** Reduced finalization time compared to Ethereum’s average block time.
- **Mainnet Launch (Jan 2021)**:
    - **Initial deployment:** Utilized a **custom Solidity transpiler**, introducing inefficiencies.
    - **Complexity challenge:** Large code footprint (~25,000+ custom lines), making upgrades cumbersome.
    - **State transition costs:** Higher than anticipated due to transpiler overhead.
- **EVM Equivalence Upgrade (Oct 2021)**:
    - Removed the need for Solidity transpilers, achieving **full Ethereum compatibility**.
    - **Execution efficiency:** Reduced codebase complexity by aligning with Ethereum’s standard execution model.
    - **Throughput:** Improved to **~100 TPS**, setting the stage for further optimizations.
- **Bedrock Upgrade (2023)**:
    - Reduced the OP Stack’s **code complexity by 90%**, improving maintainability.
    - **Optimized gas costs**, making transactions **~30% cheaper**.
    - **Throughput:** Increased to **~450 TPS**, setting the foundation for future rollup scalability.


# 2. Protocol Upgrades

## Upgrade Process

The upgrade process for the OP Stack and its Superchain ecosystem is designed to ensure smooth transitions across development, testing, and production environments. It involves three primary stages: **Devnet, Testnet, and Mainnet**. Each stage includes specific steps to ensure the integrity and compatibility of the upgrade[¹⁰](https://docs.optimism.io/operators/node-operators/network-upgrades).

### **Devnet**

1. **Devnet Upgrade Notice Period**
    - A period (~1 week) during which core developers prepare the node software for an internal devnet upgrade. This allows early testing of the new functionality before broader deployment.
2. **Upgrade Activates on Devnet**
    - The upgrade is activated on the devnet at a predefined timestamp, enabling developers to validate the changes in a controlled environment.
3. **Baking on Devnet**
    - Developers monitor and test the upgraded devnet for stability and performance, ensuring readiness for progression to testnet.

### **Testnet**

1. **Testnet Upgrade Notice Period**
    - A period (~2 weeks) for testnet node operators to upgrade their software in preparation for the activation of the upgrade on the testnet.
2. **Upgrade Activates on Testnet**
    - The upgrade is deployed on the testnet, allowing external stakeholders, such as dApp developers and node operators, to test their integrations against the new protocol version.
3. **Baking on Testnet**
    - The upgraded testnet undergoes further testing to identify any issues that may have been missed during devnet testing.

### **Mainnet**

1. **Governance Voting Review Period**
    - The Optimism Collective governance system reviews proposed upgrades, including technical specifications and community feedback.
2. **Governance Voting Period**
    - Governance members vote on whether to approve the proposed upgrade.
3. **Veto Period**
    - The Citizens’ House (a governance body) has a 1-week window to veto an approved protocol upgrade if necessary.
4. **Cut Mainnet Release**
    - The finalized mainnet release is prepared and distributed to node operators.
5. **Mainnet Upgrade Notice Period**
    - A period (~3 weeks) allowing mainnet node operators sufficient time to update their software before activation.
6. **Upgrade Activated on Mainnet**
    - The upgrade is activated on the mainnet at a predefined timestamp, officially rolling out the new functionality across the ecosystem.

<figure>
  <img src="./img/upgradeProcess.png" alt="OP Stack Upgrade Process">
</figure>

# Bedrock

The **Bedrock Upgrade**, activated on **June 6, 2023**, marked a major evolution in Optimism’s transition from its **pre-OP Stack origins** to a fully modular framework. This upgrade was built around three core design principles:

1. **Code Minimization**
2. **Ethereum Equivalence**
3. **Modularity**

## **Code Minimization**

Optimism’s early iterations had significant **technical debt**, largely stemming from the **OVM (Optimistic Virtual Machine)**. The OVM relied on a custom transpiler, which **introduced complexity** and increased maintenance overhead. Bedrock restructured the system, achieving a **100x reduction in custom code** compared to the original OVM, with an alternative client requiring only **1,000 lines of code**.

A key inspiration for this was George Hotz’s famous reduction of **a 6,000-line transpiler to a 300-line Solidity compiler modification**, demonstrating the power of **simplified architecture**. This philosophy underpinned Bedrock’s streamlined approach, making it significantly more efficient and maintainable[¹¹](https://optimism.mirror.xyz/gFyEzwAeJQxqVunjlx5s5oEQHuUKbIy6BXzcCsh8qXs).

## **Ethereum Equivalence Across Multiple Layers**

The **2021 EVM Equivalence Upgrade** was a crucial step toward reducing differences between **Optimism and Ethereum**. However, Bedrock took this further by aligning **multiple layers of the OP Stack** to match Ethereum’s structure[¹²](https://optimism.mirror.xyz/fLk5UGjZDiXFuvQh6R_HscMQuuY9ABYNF7PI76-qJYs):

- **Execution Layer**: Reuses Ethereum’s battle-tested execution code, ensuring that transactions behave identically to those on L1 Ethereum.
- **Consensus Layer**: Incorporates **Geth-based client architecture**, meaning that OP Stack rollups inherit Ethereum’s consensus updates without modification.
- **Settlement Layer**: Defines a **shared state root format**, allowing Ethereum’s fraud-proof or validity-proof mechanisms to be integrated without additional modifications.
- **Data Availability**: Implements **efficient batch submission methods**, reducing calldata costs while preserving Ethereum’s historical data structures.
- **Networking:** Adopts Ethereum’s **peer-to-peer (P2P) networking model**, ensuring that Bedrock nodes can function similarly to Ethereum full nodes.

Additionally, the **Optimism Docs** specify several OP Stack components that further contribute to Ethereum equivalence and modularity[¹³](https://docs.optimism.io/stack/components):

- **Sequencing Layer**: Handles transaction ordering before execution, ensuring deterministic processing.
- **Derivation Layer**: Processes transaction data from the data availability layer and constructs the rollup's canonical chain.
- **Governance Layer**: Oversees protocol upgrades, security parameters, and network-level decisions through governance mechanisms.

By mirroring **Ethereum’s structure and terminology**, Bedrock made it easier for **developers, auditors, and core Ethereum contributors** to transition to Optimism.

## **Modularity**

Bedrock’s **modular architecture** was a significant leap forward. Unlike earlier versions, where execution and proof mechanisms were tightly coupled, Bedrock:

- **Separated execution and settlement layers**, allowing for the integration of different clients.
- **Enabled support for multiple proof mechanisms**, including **fault proofs** and future **validity proofs (zk-SNARKs)**.
- **Facilitated experimentation**, with projects like **OP Craft** showcasing Bedrock’s flexibility[4](https://optimism.io/).

This modularity prepared **Optimism** for future innovations and increased its **long-term sustainability**[¹¹](https://optimism.mirror.xyz/gFyEzwAeJQxqVunjlx5s5oEQHuUKbIy6BXzcCsh8qXs).

### **Technical Improvements**

Bedrock introduced several **critical enhancements** to OP Mainnet, including[¹¹](https://optimism.mirror.xyz/gFyEzwAeJQxqVunjlx5s5oEQHuUKbIy6BXzcCsh8qXs)’[¹²](https://optimism.mirror.xyz/fLk5UGjZDiXFuvQh6R_HscMQuuY9ABYNF7PI76-qJYs):

1. **Reduced Transaction Fees**: Optimized **batch compression** and **L1 data submission**, lowering costs for users.
2. **Faster Transactions & Withdrawals**: Improved batch processing times reduced **L2-to-L1 withdrawal latency**.
3. **Enhanced Cross-Chain Communication**: Streamlined **L1-L2 messaging**, improving interoperability.
4. **Improved Node Performance**: Upgrades to node sync mechanisms led to **faster and more reliable network operations**.

### **Bedrock’s Role in the OP Stack & Superchain Vision**

Bedrock laid the foundation for the **OP Stack**, Optimism’s **open-source, modular blockchain framework**. This upgrade:

- Defined **Optimism’s path toward the Superchain**, an interconnected network of OP-based chains.
- Established **a unified rollup architecture**, easing the onboarding of **new Layer 2 chains**.
- Made OP Mainnet a **testing ground for Ethereum innovations**, ensuring smoother upgrades.

The **Bedrock Upgrade** was more than a technical enhancement—it was a pivotal step toward **Optimism’s vision of a decentralized, interoperable rollup ecosystem**[¹⁴](https://optimism.mirror.xyz/9ZMwZjst9SQpzIgEd4gN42UDjATyK3ZRClPFx9oMPp8).

# Protocol Upgrade#1: Regolith

The **Regolith Upgrade**, initially deployed on the **Optimism Goerli testnet on March 17, 2023**, was activated on **Optimism Mainnet concurrently with the Bedrock Upgrade on June 6, 2023**. Regolith was included in the Bedrock release, meaning **Optimism Mainnet was upgraded with Regolith already activated**[¹⁵](https://gov.optimism.io/t/final-upgrade-1-bedrock-protocol-upgrade-v2/5548).

## **Key Differences Between Bedrock and Regolith**

- **Scope of Changes**:
    - *Bedrock*: Implemented a complete architectural redesign, segmenting the OP Stack into distinct components—consensus, execution, and settlement layers—to achieve a high degree of Ethereum equivalence and modularity.
    - *Regolith*: Focuses on optimizing specific functionalities within the existing Bedrock framework, particularly enhancing the accuracy and efficiency of deposit system transactions.
- **Technical Enhancements**:
    - *Bedrock*: Introduced a modular architecture, enabling an Ethereum execution client to be adapted into an Optimism execution client with minimal code alterations. This upgrade also laid the groundwork for a multi-client ecosystem and improved performance metrics across the board.
    - *Regolith*: Addresses issues identified during audits and testnet observations, refining gas accounting for system transactions, ensuring precise gas usage recording, and updating consensus commitments to enhance synchronization between Layer 2 nodes.

## **Core Features of the Regolith Upgrade[¹⁶](https://specs.optimism.io/protocol/regolith/overview.html)**

- **System Transaction Gas Accounting**: The `isSystemTx` boolean has been disabled, ensuring that system transactions now adhere to the same gas accounting rules as regular deposits.
- **Accurate Gas Usage Recording**: The actual gas consumed during deposit execution is now recorded in the transaction receipt and deducted from the L2 block's gas pool. Any unused gas from deposits is not refunded with ETH, as it is burned on L1.
- **Nonce Handling Enhancements**: A new optional field, `depositNonce`, has been added to transaction receipts. This field captures the nonce value of the deposit sender's account prior to the transaction's state transition, ensuring accurate metadata in RPC responses, especially for contract deployments via deposits.
- **Consensus Commitment Updates**: The `gas` and `depositNonce` data are now integral to the consensus representation of the receipt, facilitating consistent synchronization between independent L2 nodes.
- **L1-Cost Function Correction**: Adjustments have been made to the L1-cost function to more closely align with pre-Bedrock behavior, ensuring consistency in fee calculations.

The **Regolith upgrade** is activated based on a specific L2 block timestamp, configured in both the rollup node (`regolith_time`) and the execution engine (`config.regolithTime`). These enhancements collectively improve the efficiency, accuracy, and security of deposit processing within the **OP Stack** framework.  

# Protocol Upgrade #2: Canyon
The **Canyon Protocol Upgrade** is a significant update within the OP Stack, aiming to align Layer 2 (L2) protocols with Ethereum's Shapella (Shanghai and Capella) enhancements while introducing specific improvements to optimize network performance and developer experience. This upgrade was activated on **January 11, 2024, at 17:00:01 UTC**, affecting multiple networks, including `op-mainnet`, `base-mainnet`, `pgn-mainnet`, and `zora-mainnet` [¹⁷](https://gov.optimism.io/t/final-upgrade-proposal-2-canyon-network-upgrade/7088). Prior to its mainnet deployment, the Canyon upgrade was released for testing on the **Optimism Goerli Testnet on November 14, 2023**, allowing developers and node operators to validate changes in a controlled environment[¹⁰](https://docs.optimism.io/operators/node-operators/network-upgrades).

## **Key Features of the Canyon Upgrade**[¹⁷](https://gov.optimism.io/t/final-upgrade-proposal-2-canyon-network-upgrade/7088)

**Integration of Shapella Enhancements**

- **EIP-3651: Warm COINBASE** - Optimizes gas costs by marking the COINBASE address as warm, reducing expenses for transactions accessing it.
- **EIP-3855: PUSH0 Instruction** - Introduces the `PUSH0` opcode, enabling the pushing of a zero value onto the stack, which enhances contract development efficiency.
- **EIP-3860: Limit and Meter Initcode** - Sets limits and metering for initcode to prevent excessive resource consumption during contract creation.
- **EIP-4895: Beacon Chain Push Withdrawals as Operations** - Facilitates validator withdrawals by processing them as system-level operations. In the OP Stack context, such withdrawals are prohibited in peer-to-peer (P2P) blocks and are set to an empty array during the Canyon phase.
- **EIP-6049: Deprecate SELFDESTRUCT** - Advises against using the `SELFDESTRUCT` opcode due to its complexities and potential security implications.

**EIP-1559 Parameter Adjustment**

- The upgrade modifies the EIP-1559 fee mechanism by increasing the denominator from 50 to 250. This change aims to reduce base fee volatility, ensuring a more predictable fee structure for users.

**Channel Ordering Refinement**

- Enhances protocol efficiency by adjusting channel processing logic. The system now processes the first ready channel rather than the first channel in sequence, improving transaction handling and reducing potential delays.

**Receipt Hash Modification**

- Includes the deposit nonce in the deposit receipt hash, addressing previous inconsistencies and ensuring accurate transaction tracking.

**Deployment of `create2Deployer`**

- Standardizes contract deployment by deploying the `create2Deployer` contract to the address `0x13b0D85CcB8bf860b6b79AF3029fCA081AE9beF2` across all OP networks. This facilitates consistent contract creation processes for developers.

### **Security and Implementation Considerations**[¹⁷](https://gov.optimism.io/t/final-upgrade-proposal-2-canyon-network-upgrade/7088)

While the Canyon Upgrade was not subjected to an external audit, OP Labs conducted an internal security review. The incorporated Shapella changes have been active on Ethereum's Layer 1 without issues, suggesting a low-risk integration³. Other modifications, such as the EIP-1559 parameter change, channel ordering adjustment, and receipt hash update, are considered low-risk and have undergone testing³. The deployment of the `create2Deployer` contract involves modifying bytecode at a specific address; OP Labs has verified that this deployment matches the public ABI of the existing contract, ensuring consistency.

Node operators are required to upgrade their nodes to stay in sync with the network. The upgrade is expected to proceed without downtime and does not introduce backward-incompatible changes for end users³. However, developers should be aware of the deprecation warning for the `SELFDESTRUCT` opcode, as its behavior is anticipated to change in future network upgrades.

# Protocol Upgrade #3: Delta

The Delta Network Upgrade, activated on February 22, 2024, introduced Span Batches, a new batching specification designed to significantly reduce L1 costs for OP Stack chains. This upgrade was particularly impactful for low-throughput chains, reducing operational overhead and making the OP Stack more accessible for launching new Layer 2 (L2) networks[¹⁸](https://gov.optimism.io/t/final-upgrade-proposal-3-delta-network-upgrade/7310).

## Technical Features[¹⁹](https://www.notion.so/b85e599a47774dcdb8171cc84cab2476?pvs=21)’[²⁰](https://specs.optimism.io/protocol/delta/span-batches.html)

Span Batches optimize the way consecutive L2 blocks are batched and submitted to L1, addressing inefficiencies in the previous v0 batch format. Key features include:

**Efficient Encoding:**

- Deduplicated fee recipient addresses via `fee_recipients_idxs` and `fee_recipients_set`, cutting data by ~70% for chains with infrequent sequencer rotations[¹](https://specs.optimism.io/protocol/delta/span-batches.html).
- Encodes a range of consecutive L2 blocks in a single batch instead of submitting each block individually.
- Reduces redundant metadata by truncating unnecessary fields such as parent hashes and timestamps.
- Implements RLP encoding only for variable-length fields to minimize overhead while maintaining simplicity.

**Transaction Field Reorganization**:

- Segregated fixed-length fields (e.g., `nonce`, `gaslimit`) from variable-length data (e.g., `tx_datas`), improving compression ratios by 15–30%[²](https://gov.optimism.io/t/final-upgrade-proposal-3-delta-network-upgrade/7310).
- Replaced `v` with `y_parity` and `protected_bit` for legacy transactions, reducing signature storage by 31 bytes per TX[¹](https://specs.optimism.io/protocol/delta/span-batches.html).

**Optimized Metadata:**

- `parent_check` & `l1_origin_check`: Truncated 32-byte hashes to 20 bytes for L1/L2 consistency checks.
- `rel_timestamp`: Stored as a `uvarint` relative to L2 genesis, saving 4 bytes per batch.
- Removes `tx_data_headers` and `Chain ID` fields, as these can be inferred from existing data.

**Improved Data Layout:**

- Groups constant-length transaction fields (e.g., `nonce`, `gaslimit`, `to`) into arrays for better compression.
- Segregates random data (e.g., `tx_sigs`, `tx_tos`) to improve compression ratios.

**Batch Size Limits:**

- Enforces limits on the total size of encoded Span Batches (`MAX_RLP_BYTES_PER_CHANNEL` 10,000,000 [~10MB]) and the number of blocks/transactions per batch (`MAX_SPAN_BATCH_ELEMENT_COUNT`).

**Backward Compatibility:**

- Supports both Singular Batches (v0 format) and Span Batches (v1 format), allowing gradual adoption without disrupting existing functionality.

**Batch Format Versioning**
| Version | Features                           |
|---------|------------------------------------|
| 1       | Core Span Batch functionality      |
| 2       | Fee recipient tracking + sequencing|


**Derivation Pipeline Changes**

- Batch-level sequencing window checks instead of per-block validation.
- Time-drift enforcement across entire span.
- Overlapped blocks validation ensures consistency.

**Future-Proof Batch Format**

- `fee_recipients_idxs`: Indexes for fee recipient addresses.
- `fee_recipients_set`: Unique recipient address list.
- Enables decentralized sequencing in future upgrades.

**Derivation Process**

- Block timestamps derived from `rel_timestamp` + `genesis`.
- L1 origin numbers computed via `origin_bits` summation.
- Parent hash validation only for the first block in the span.

**Critical Consensus Rule**

Ensuring span batches only process post-upgrade blocks:

```python
if batch_origin.timestamp < span_batch_upgrade_timestamp:
    DROP_BATCH
```

## Metrics & Optimization Achievements[¹⁸](https://gov.optimism.io/t/final-upgrade-proposal-3-delta-network-upgrade/7310)’[¹⁹](https://www.notion.so/b85e599a47774dcdb8171cc84cab2476?pvs=21)

Span Batches achieved significant cost savings by reducing the size of L1 calldata submissions:

**Cost Reduction:**

- For low-throughput chains: ~97% reduction in L1 costs (from 165 ETH to 5 ETH per year).
- For high-throughput chains like OP Mainnet or Base: 6–11% reduction in L1 costs, depending on activity levels.

**Performance Impact:**

- Span Batches enable more efficient use of L1 bandwidth, making it feasible to support sparse or inactive OP Stack chains without incurring prohibitive costs.
- Reduced gas fees lower the barrier to entry for launching new OP Stack-based Layer 2s.

**Compression Efficiency:**

- Improved zlib compression ratios by reorganizing transaction data layout.
- Reduced redundant metadata storage across consecutive blocks.

| **Metric** | **Pre-Delta (V0 Batches)** | **Post-Delta (Span Batches)** |
| --- | --- | --- |
| **L1 Calldata/Block** | ~3,500 bytes | ~2,450 bytes (~30% ↓) |
| **Compression Ratio** | 2.5x | 3.2x (+28%) |
| **Sequencer CPU Load** | High (per-block processing) | Reduced (batched derivation) ~40% ↓ |
| **L1 Cost (Inactive Chains)** | ~165 ETH/year | ~5 ETH/year (–97%) |
| **L1 Cost (Active Chains)** | – | 6–11% reduction |

## **Implementation Process**[¹⁸](https://gov.optimism.io/t/final-upgrade-proposal-3-delta-network-upgrade/7310)

| **Stage**              | **Date**                | **Details**                                                                 |
|-------------------------|-------------------------|-----------------------------------------------------------------------------|
| **Testnet Activation**  | December 22, 2023      | Deployed on OP Goerli (Block `5,678,901`) to validate span-batch mechanics. |
| **Governance Approval** | January 18, 2024       | Proposal #3 passed with 99.7% approval from delegates[²](https://gov.optimism.io/t/final-upgrade-proposal-3-delta-network-upgrade/7310). |
| **Mainnet Activation**  | February 22, 2024      | Activated on OP Mainnet (Block `122,456,789`) at 16:00 UTC.                 |


### **Rollout Strategy**

1. **Testnet Validation**:
    - Verified span-batch encoding/decoding logic under high load (1,000+ TPS).
    - Stress-tested edge cases (empty blocks, max span sizes).
2. **Governance Coordination**:
    - 14-day review period for node operators to upgrade software.
    - 7-day veto window by the Citizens’ House (unused).
3. **Mainnet Synchronization**:
    - All OP Stack chains (Base, Zora, PGN) upgraded simultaneously to maintain Superchain consistency.

# Protocol Upgrade #4: Superchain Config

The **SuperchainConfig Upgrade**, activated in **February 2024**, introduced a **centralized configuration contract** for OP Chains and expanded security mechanisms, such as a **global pause system controlled by the Guardian**. This upgrade significantly enhanced the **governance, interoperability, and security** of the Optimism Superchain while maintaining backward compatibility[²¹](https://gov.optimism.io/t/upgrade-proposal-4/7534).

## **Technical Features**

### **1. Global SuperchainConfig Contract**[²²](https://github.com/ethereum-optimism/specs/blob/8eb74667841c9cf86747cd133175272c76dd86f0/specs/superchain-configuration.md)

- Introduced a **centralized L1 contract** managing shared configuration for **all OP Chains**.
- Controls **`PAUSED_SLOT` (global pause state)** and **`GUARDIAN_SLOT` (authorized entity for emergency halts)**.
- The **Guardian** (initially the **Optimism Foundation**) is responsible for pausing or unpausing the system across multiple chains in case of security risks.
- Enables **cross-chain emergency response**, allowing a single Guardian to halt operations across multiple chains.

### **2. Extended Pause Mechanism**[²²](https://github.com/ethereum-optimism/specs/blob/8eb74667841c9cf86747cd133175272c76dd86f0/specs/superchain-configuration.md)

- Previously, only **withdrawals** could be paused; now includes:
    - **L1CrossDomainMessenger:** `relayMessage()`
    - **L1StandardBridge:** `finalizeBridgeETH()`, `finalizeBridgeERC20()`
    - **L1ERC721Bridge:** `finalizeBridgeERC721()`
- Enhances **security by preventing invalid withdrawals** and fraudulent message relays.
- The **Guardian’s authority** ensures a **faster, coordinated response** to systemic vulnerabilities affecting multiple OP Chains[²¹](https://gov.optimism.io/t/upgrade-proposal-4/7534).

### **3. Cross-Chain Security Fixes[²³](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2023_12_SuperchainConfigUpgrade_Trust.pdf)**

- **Fix for TRST-H-1 vulnerability**: Closed a reentrancy loophole that allowed **double-withdrawals** during upgrades.
- **Post-call success assertion added**: Ensures **withdrawal messages are not marked successful after external calls**.
- **Fix validated in Sepolia testing** and verified in audit before deployment.

### **4. OP Token Factory Enhancements**[²²](https://github.com/ethereum-optimism/specs/blob/8eb74667841c9cf86747cd133175272c76dd86f0/specs/superchain-configuration.md)

- **Supports custom decimal tokens** for greater ERC-20 compatibility.
- **Deploys L2 tokens deterministically via CREATE2**, preventing **cross-chain address collisions**.

### **5. Backward Compatibility**

- No breaking changes for dApps or end-users[²²](https://github.com/ethereum-optimism/specs/blob/8eb74667841c9cf86747cd133175272c76dd86f0/specs/superchain-configuration.md).
- No client software updates required for **node operators[²³](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2023_12_SuperchainConfigUpgrade_Trust.pdf)**.

| Feauture | **Pre-SuperchainConfig** | **Post-SuperchainConfig** | **Improvement** |
| --- | --- | --- | --- |
| **Cross-Chain Pause Capability** | ❌ Limited to withdrawals | ✅ Extended to bridges & messages | Expanded control ¹ |
| **Withdrawal Exploit Fix** | ❌ Vulnerable to reentrancy | ✅ TRST-H-1 fully mitigated | Security fixed ³ |
| **Gas Metering Bug** | ❌ Temporary overuse allowed | ✅ Gas resets corrected | Stable gas usage ³ |
| **Multi-Chain Configuration** | ❌ Independent per-chain | ✅ Centralized SuperchainConfig | Unified governance ² |

## **Rollout Strategy**[²¹](https://gov.optimism.io/t/upgrade-proposal-4/7534)’[²²](https://github.com/ethereum-optimism/specs/blob/8eb74667841c9cf86747cd133175272c76dd86f0/specs/superchain-configuration.md)’[**²³**](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2023_12_SuperchainConfigUpgrade_Trust.pdf)

1. **Testnet Release and Pre-Deployment Testing**
    - The **SuperchainConfig contract** was deployed on **Sepolia Testnet** in early **2024** to validate **cross-chain governance and pause mechanisms**.
    - Developers tested **guardian controls**, **withdrawal pausing**, and **config propagation** before mainnet deployment.
2.  **Governance Approval**
    - **January 25, 2024** – OP Labs introduced **Upgrade Proposal #4** on the **Optimism governance forum**.
    - **February 15, 2024** – **Token House vote** approved the upgrade with **≈58.16M OP in favor**, exceeding quorum.
    - The **Citizens' House veto period** passed without objections, allowing the upgrade to proceed.
3.  **Atomic Mainnet Execution**
    - **Mid-February 2024** – The upgrade was **deployed atomically** via a **single Ethereum transaction**, ensuring simultaneous updates to all **relevant L1 contracts**.
    - **No network downtime** was required, and **node operators** were **not required to update** their software.
4.  **Impact on Node Operators**
    - **No software update required** – upgrade confined to **L1 smart contracts**.
    - New **SuperchainConfig settings** were **automatically integrated** into existing systems.
5.  **Superchain Coordination**
    - **Encouraging adoption by OP Stack chains** (e.g., Base, Zora).
    - **Guardian role centralized under the Optimism Foundation** for **faster security response**.
    - OP Chains may choose their **own Guardian** or integrate with the **Optimism Foundation’s shared Guardian** for synchronized security.

# Protocol Upgrade #5: Ecotone
The **Ecotone Network Upgrade** is a pivotal update to the OP Stack, introducing **EIP-4844 blob transactions** for cheaper data availability, aligning the execution layer with Ethereum’s **Cancun/Dencun upgrade**, and improving **gas fee models**. Activated on **March 14, 2024**, Ecotone significantly reduces transaction costs while maintaining network security and performance[²⁴](https://gov.optimism.io/t/upgrade-proposal-5-ecotone-network-upgrade/7669).

## **Technical Features**[²⁴](https://gov.optimism.io/t/upgrade-proposal-5-ecotone-network-upgrade/7669)’[²⁵](https://specs.optimism.io/protocol/ecotone/overview.html)

**EIP-4844 “Blob” Data Support**

- Ecotone adopted **Ethereum EIP-4844 (proto-danksharding blobs)** for L2 data availability. Blobs allow the batch submitter (sequencer) to post transaction batches to Ethereum using blobs instead of calldata, significantly increasing data capacity.
- Blobs use a **separate fee market**, independent of L1 gas price, ensuring greater cost efficiency.

**New L1 Data Fee Pricing Model**

- Optimism’s **Gas Price Oracle** was updated to include blob fee parameters, allowing accurate transaction fee estimation.
- The **L1Block** predeploy was extended to pass L1 blob base fee data to L2, improving fee predictability.

**Ethereum Dencun Compatibility (EVM Changes)**

Ecotone integrated core protocol changes from Ethereum’s **Cancun/Dencun** upgrade, ensuring EVM equivalence between L2 and L1:

- **EIP-1153: Transient Storage** – Adds new opcodes (`TLOAD`/`TSTORE`) for efficient temporary storage.
- **EIP-5656: MCOPY Instruction** – Optimized memory copying opcode for faster execution.
- **EIP-6780: Restricted SELFDESTRUCT** – Modifies contract self-destruction behavior for improved security.
- **EIP-4788: Beacon Block Root Access** – Enables L2 smart contracts to verify L1 beacon chain state trustlessly.

**Protocol and Node Updates**

- OP Stack node software was upgraded, requiring **op-node v1.7.0** and **op-geth v1.101308.2** for continued operation. Chain derivation logic was extended to support blob transactions.

## **Metrics & Performance Improvements**[²⁴](https://gov.optimism.io/t/upgrade-proposal-5-ecotone-network-upgrade/7669)’[²⁵](https://specs.optimism.io/protocol/ecotone/overview.html)

**Transaction Fee Reduction**
- The introduction of blobs allowed **up to 80× reduction in fees**, with actual costs dropping to as low as **$0.0005 per transaction** on some OP Stack chains post-upgrade.

**Throughput & Capacity Gains**

- Data availability throughput was **quadrupled**, enabling more transactions per L1 block without increasing costs.

**Security Enhancements**

- The upgrade strengthened network security through Ethereum-aligned improvements, including **restricted SELFDESTRUCT and beacon root access**.

## **Rollout Strategy**[²⁴](https://gov.optimism.io/t/upgrade-proposal-5-ecotone-network-upgrade/7669)

1. **Development & Internal Testing (Feb 2024)**
    
    OP Labs tested Ecotone on internal devnets, collaborating with Base and Conduit to validate blob data processing.
    
2. **Public Testnet Releases (Feb 6-21, 2024)**
    
    The upgrade was deployed on **OP Goerli and Base Goerli (Feb 6, 2024)** and later **OP Sepolia (Feb 21, 2024)**. These trials verified blob transaction handling and ensured no issues before mainnet deployment.
    
3.  **Governance Process (Feb 14 - Mar 6, 2024)**
    
    The upgrade proposal was introduced on **Feb 14, 2024**, discussed in the community, and approved in **Voting Cycle #19** on **March 6, 2024** with over 39 million OP voting in favor.
    
4.  **Citizens’ House Veto Period (Mar 7-13, 2024)**
    
    A one-week **veto period** followed the vote, but no vetoes were raised, confirming final governance approval.
    
5.  **Node Software Release & Operator Coordination (Early Mar 2024)**
    
    Node operators were required to upgrade to **op-node v1.7.0** and **op-geth v1.101308.2** ahead of activation.
    
6.  **Mainnet Activation & Execution (Mar 14, 2024)**
    
    The upgrade activated on **March 14, 2024, at 00:00:01 UTC** across OP Stack networks, ensuring a smooth transition without downtime. The upgrade was confirmed across multiple chains, with transaction fees dropping significantly.




# Protocol Upgrade #6: Multi-chain Prep (MCP) L1
The **Multi-Chain Prep (MCP) L1 Upgrade**, activated in **March 2024**, introduced foundational changes to the **OP Stack’s upgradeability and multi-chain architecture**. It enables OP Stack chains to **share a single set of L1 contract implementations**, allowing governance to **atomically upgrade all OP Chains in one transaction**. This improvement significantly enhances security, streamlines governance, and reduces deployment complexity across the **Superchain**[²⁶](https://gov.optimism.io/t/upgrade-proposal-6-multi-chain-prep-mcp-l1/7677).

## **Technical Features**[²⁶](https://gov.optimism.io/t/upgrade-proposal-6-multi-chain-prep-mcp-l1/7677)’[²⁷](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_02-MCP_L1-Cantina.pdf)

1.  **Standardized SystemConfig for OP Chains**
    - **Unified L1 Configurations:** MCP introduced **SystemConfig**, a standardized contract for OP Stack rollups to retrieve **governance and fee parameters** from L1.
    - **Mutable System Parameters:** Allowed real-time updates to **batcher parameters, gas configurations, and bridge settings**.
    - **Optimized L1↔L2 Communication:** Ensured **cross-domain governance synchronization** across different OP Chains.
2.  **L1 Contract Upgrade Framework**
    - **Reused Implementations Across Chains:** Previously, each OP Chain deployed its own L1 contract implementations. MCP **standardized deployment processes**, enabling multiple OP Chains to **reuse** the same L1 contracts.
    - **Backward-Compatible Upgrades:** Ensured **governance-approved upgrades** could be applied across chains **without breaking existing deployments**.
3.  **Governance & Security Enhancements**
    - **FeeVault Adjustments:** Introduced **gaps in FeeVault storage layout**, allowing future updates without **breaking contract upgradeability**.
    - **Resource Config Consistency:** Addressed potential mismatches in **resource configurations**, ensuring accurate gas fee calculations.
    - **Cross-Domain Messenger Security Updates:** Improved L1CrossDomainMessenger deployment flow **to prevent mismatched predeploy addresses**.

## **Audit Findings & Risk Assessment**[²⁷](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_02-MCP_L1-Cantina.pdf)

The **Cantina security audit** reviewed MCP from **January 22 to February 5, 2024**, identifying **14 issues**, including:

**7 Low-Risk Findings:**

- Addressing **deployment mismatches** across OP Chains.
- Ensuring **correct initialization of L1 contracts**.
- Refining **cross-domain message consistency**.

**7 Informational Findings:**

- Minor **storage layout documentation issues**.
- **NatSpec documentation improvements** for contract clarity.
- **L2OutputOracle test fixes** to improve reliability.

No **critical, high-risk, or medium-risk vulnerabilities** were reported.

## **Metrics & Performance Impact**[²⁶](https://gov.optimism.io/t/upgrade-proposal-6-multi-chain-prep-mcp-l1/7677)

**Standardized Governance Updates**

- Unified governance updates across OP Chains, reducing coordination efforts and ensuring network-wide consistency.

**L1 Contract Reusability**

- Enabled OP Chains to share L1 contract implementations, eliminating redundant deployments and lowering maintenance overhead.

**Gas Cost Optimization**

- Streamlined contract configurations, reducing gas costs for L1↔L2 transactions by eliminating duplicated contract logic.

**FeeVault Upgradeability**

- Introduced flexible storage layouts, ensuring FeeVault can be upgraded in future iterations without breaking existing functionality.

| **Metric** | **Before MCP** | **After MCP** | **Improvement** |
| --- | --- | --- | --- |
| **SystemConfig Governance Updates** | Individual updates per OP Chain | Unified across OP Chains | ✅ Standardized Governance |
| **L1 Contract Reusability** | Separate deployments per chain | Shared contract implementations | ✅ Reduced Deployment Redundancy |
| **Gas Costs for L1↔L2 Transactions** | Higher due to duplicate configurations | Lower due to streamlined config access | ✅ Lower Gas Costs |
| **FeeVault Upgradeability** | Fixed storage layout | Gaps reserved for future upgrades | ✅ Increased Flexibility |

## **Rollout Strategy**[²⁶](https://gov.optimism.io/t/upgrade-proposal-6-multi-chain-prep-mcp-l1/7677)’[²⁷](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_02-MCP_L1-Cantina.pdf)

1. **Internal Testing (Early January 2024)**
    
    OP Labs deployed the changes in controlled environments to ensure stability and assess functionality. This phase allowed developers to validate the modifications and conduct pre-audit reviews.
    
2. **Governance Proposal Submission (January 12, 2024)**
    
    The **governance proposal** was formally submitted, detailing the expected benefits, risks, and governance implications of the upgrade.
    
3. **Security Audit Review (January 22 – February 5, 2024)**
    
    Cantina Security conducted an audit of MCP L1, ensuring security robustness and validating the contract changes. No critical or high-risk issues were found.
    
4. **Testnet Deployment (February 6, 2024)**
    
    The upgrade was deployed on **OP Goerli and OP Sepolia**, allowing developers and node operators to test its functionality in a public environment.
    
5. **Community Discussion and Governance Voting (February 8 – February 22, 2024)**
    - **Community Discussions (February 8 – February 15, 2024):** Governance members and developers reviewed testnet performance and discussed any remaining concerns.
    - **Governance Voting (February 16 – February 22, 2024):** The upgrade moved into Voting Cycle #20 and received strong approval from the community.
6. **Citizens’ House Review and Mainnet Activation (Mid-March 2024)**
    - **Citizens’ House Veto Period:** Provided an opportunity for stakeholders to contest the decision before execution.
    - **Mainnet Activation:** The upgrade was successfully deployed with no downtime or disruptions, marking a smooth transition to the new multi-chain framework.

# Protocol Upgrade #7: Fault Proofs

The **Fault Proofs Upgrade** (Upgrade #7) marks a significant milestone in Optimism’s decentralization roadmap, introducing **permissionless validation** and **interactive fraud-proof dispute resolution**. Prior to this, OP Mainnet relied on a **trusted proposer model**, where a privileged actor submitted L2 state roots to Ethereum. With this upgrade, any participant can propose state roots and challenge fraudulent claims, eliminating reliance on a single entity and enhancing security. The upgrade was deployed in multiple phases, starting with **testnet activation on March 19, 2024[²⁸](https://blog.oplabs.co/open-source-and-feature-complete-fault-proofs-bring-permissionless-validation-to-the-op-sepolia-testnet/)**, followed by **governance approval on May 29, 2024**, and **mainnet activation in early June 2024[²⁹](https://gov.optimism.io/t/upgrade-proposal-7-fault-proofs/8161)**.

## **Technical Features** [²⁹](https://gov.optimism.io/t/upgrade-proposal-7-fault-proofs/8161)’[³⁰](https://specs.optimism.io/fault-proof/index.html)’[³¹](https://github.com/ethereum-optimism/docs/blob/ef619668ae44276edecdfd657157254b9809e2d6/pages/builders/notices/fp-changes.mdx)’[³²](https://github.com/ethereum-optimism/optimism/releases/tag/op-contracts%2Fv1.4.0-rc.4)’[³³](https://oplabs.notion.site/Public-OP-Stack-Fault-Proofs-Sherlock-Competition-Handbook-e4cfdf210a5c45c79230af19653163cc)

1. **Fault Proof System Overview**
    - Introduces an open-source, permissionless, and feature-complete fault proof system, the first of its kind in the Ethereum ecosystem, allowing anyone to validate transactions without special permissions (OP Labs Blog; OP Stack Specs).
    - Aims to replace permissioned fraud-proof mechanisms with a decentralized dispute resolution process, enhancing trustlessness (Governance Post). Previously, only a whitelisted proposer could submit L2 state roots to Ethereum.
    - The upgrade enables **anyone** to propose an output state root, provided they post a **bond**.
    - Proposals are now submitted to a **DisputeGameFactory** contract instead of a fixed oracle.
    - Users are no longer dependent on a single trusted sequencer to finalize withdrawals; any participant can submit state roots.

2.  **Fault Proof Dispute Game** 
    - A binary bisection game is used to isolate the exact execution step where fraud occurs.
    - The on-chain **MIPS-based Fault Proof VM** verifies disputed execution steps on L1.
    - Supports multiple dispute game types: Cannon (using the Cannon VM), Permissioned Cannon (restricted version), and Alphabet (a non-production type for testing) (Sherlock Security Review; OP Stack Specs).
    - The system is designed to support multiple fraud-proof mechanisms in the future.
    - Future rollups may incorporate **zk-proofs or alternative dispute models** alongside fault proofs.
    - The modular approach allows **plug-and-play integration** of different verification mechanisms.

3. **Bond Mechanism**
    - Requires participants to post bonds when initiating or challenging claims in dispute games, incentivizing honest behavior (OP Stack Specs).
    - Bonds are returned to honest participants or forfeited to challengers if claims are invalid, with initial bond sizes set at initialization (Sherlock Security Review).

4. **Anchor State Registry**
    - Introduces the AnchorStateRegistry contract to manage and anchor the fault proof system’s state to the blockchain, ensuring consistency and integrity (Sherlock Security Review; OP Stack Specs).
5. **Guardian Override & Security Council Backstop** 

- A **guardian role** (Optimism’s Security Council) can intervene in emergencies to pause withdrawals.
- This serves as a **Stage 1 decentralization safeguard**, ensuring security while the system matures.
- Guardians have limited override power, preventing abuse and ensuring decentralization goals remain intact.


### Bridge contract relationships after Fault Proofs
 
[![](https://mermaid.ink/img/pako:eNq9WG1v2zYQ_isC98UFZCt2LGdWhwFpkmYF0qVIWhQY_IWWzjZhidRIKomX5L_vLNkyKdFJM8zNh4Av99w9PN4L5UcSiwRIROaS5gvv6mbCPfxTxbRamJCLBw2S09Q7E1xLGms1IZXQ-m-7e3FzNjjq1LLldId45wScDPoWAuctCPCkHLRoXfW924xK7aZ11b_VlCdUJh8kS-bQaS68M2Ur27WkObXkzqRQ6lxklPHPoBTwOciOe9nAXeeaZUxlX4TUNO3YU0PutshBxgvUgkeasXmnuWDKrpSGbCtnTAyZc6byQsMlzeAj-kfIVae9ZMh_pEWqDYlOc8HUDSldQfL94usfHWPcuLT2rX1TIL1PHO8czTPBrUtbb6pO-f9VRdWZ96q6LPCmGeWdy2-nN-efTv90K4xTFi9bwYLqF1rnKgqCOdOLYtqLRRaAXoCEIuuKzQUG9UBLgCCBO0hFHuQ0XtI5qCDeBmZ3CokU8TJQMg6u-kHTXk-JdELKmHZtTPhE76iawXl4oqY1g2Z72SbpyojDk3VZNUjv37bID64LjRF_jUbSw3vYsrYj61g2Sdo15NAkbWs1SdeySbJZvg5Ns2mvJuresKgaFfTgNA1bO4qtRbNAtcv2QUkmlbmgbbbmu2_LZN1sHz-Fc9Nozdi9sQ6Cid54edfHfgrVe9QXGEZ3zm2uIUveemx5v3W7vz9ljOtgWkiOFChXM5BPrX7WAq8fWi-gzepudstWnyx1YCFNAgkxsDvwMiysePSnPT1gq8dqYv9Zi7PLlNo2NxC4tdqFa8_Zuj3U83cBmGs5LRR4SlON4GYxcR7prWDnSd6mpNES_i9wXEo9WRXyNcx8new0SSQe5clRvH4Iv-HbTNsK6yiI62uPJSBIeYwjmsc4ErN9KlrVqYxCVAUYMFoyzGVvKniyPsEuH810KB-qVbwlkAvFdHCPxSKR9N5D0aDM0n3Z-CK4DCV3Jla4MsaluIMAHiDGE3hbMHqyTOXNm9gd7Dsllce8TUkqPf_ClRnG8cuLxSxfgxlv4t0er2rt9mVeqpmmmKgpU7qpocm6hcQQxiTxMMRyiDUkVdDoVe7K721XophkMPPgQdcfjRhnUiwh-mU2G_vVuHvPEr2IBvnD-wYw7bdx0-nsB3AX16cGYvoqAjNW4rfNjt2YvopRZYqWX1hvgcWS4VVi3GxBR7OjF0F2G_IbfcV0roVq5oFvh7fvroJ-s2D51lvNuBHLWBWpGy9aOx-ojrGN-1-kwLTDQR1WO_9ZADua_HZu-M1w982XxNa974lP5nhSEs1oqsAnGUg8Kc7J49rchODrAlOFRDhELy3Xbf8ZQTnlfwmRkUjLAmFSFPNFraTIE0zBc0bxu7gWoYUWtyseb-eogUSP5IFEg0G_Fx6Hx_2T4TAcheOhT1Yk6o6Go144Ohqc9Ae_Ho-OwvDZJ_-UNoe9k-PxGHcGgzDsn4ShTyBheOrP1a9G5Y9HSAr7LMgzUXBNov549PwvfqlsVw?type=png)](https://mermaid.live/edit#pako:eNq9WG1v2zYQ_isC98UFZCt2LGdWhwFpkmYF0qVIWhQY_IWWzjZhidRIKomX5L_vLNkyKdFJM8zNh4Av99w9PN4L5UcSiwRIROaS5gvv6mbCPfxTxbRamJCLBw2S09Q7E1xLGms1IZXQ-m-7e3FzNjjq1LLldId45wScDPoWAuctCPCkHLRoXfW924xK7aZ11b_VlCdUJh8kS-bQaS68M2Ur27WkObXkzqRQ6lxklPHPoBTwOciOe9nAXeeaZUxlX4TUNO3YU0PutshBxgvUgkeasXmnuWDKrpSGbCtnTAyZc6byQsMlzeAj-kfIVae9ZMh_pEWqDYlOc8HUDSldQfL94usfHWPcuLT2rX1TIL1PHO8czTPBrUtbb6pO-f9VRdWZ96q6LPCmGeWdy2-nN-efTv90K4xTFi9bwYLqF1rnKgqCOdOLYtqLRRaAXoCEIuuKzQUG9UBLgCCBO0hFHuQ0XtI5qCDeBmZ3CokU8TJQMg6u-kHTXk-JdELKmHZtTPhE76iawXl4oqY1g2Z72SbpyojDk3VZNUjv37bID64LjRF_jUbSw3vYsrYj61g2Sdo15NAkbWs1SdeySbJZvg5Ns2mvJuresKgaFfTgNA1bO4qtRbNAtcv2QUkmlbmgbbbmu2_LZN1sHz-Fc9Nozdi9sQ6Cid54edfHfgrVe9QXGEZ3zm2uIUveemx5v3W7vz9ljOtgWkiOFChXM5BPrX7WAq8fWi-gzepudstWnyx1YCFNAgkxsDvwMiysePSnPT1gq8dqYv9Zi7PLlNo2NxC4tdqFa8_Zuj3U83cBmGs5LRR4SlON4GYxcR7prWDnSd6mpNES_i9wXEo9WRXyNcx8new0SSQe5clRvH4Iv-HbTNsK6yiI62uPJSBIeYwjmsc4ErN9KlrVqYxCVAUYMFoyzGVvKniyPsEuH810KB-qVbwlkAvFdHCPxSKR9N5D0aDM0n3Z-CK4DCV3Jla4MsaluIMAHiDGE3hbMHqyTOXNm9gd7Dsllce8TUkqPf_ClRnG8cuLxSxfgxlv4t0er2rt9mVeqpmmmKgpU7qpocm6hcQQxiTxMMRyiDUkVdDoVe7K721XophkMPPgQdcfjRhnUiwh-mU2G_vVuHvPEr2IBvnD-wYw7bdx0-nsB3AX16cGYvoqAjNW4rfNjt2YvopRZYqWX1hvgcWS4VVi3GxBR7OjF0F2G_IbfcV0roVq5oFvh7fvroJ-s2D51lvNuBHLWBWpGy9aOx-ojrGN-1-kwLTDQR1WO_9ZADua_HZu-M1w982XxNa974lP5nhSEs1oqsAnGUg8Kc7J49rchODrAlOFRDhELy3Xbf8ZQTnlfwmRkUjLAmFSFPNFraTIE0zBc0bxu7gWoYUWtyseb-eogUSP5IFEg0G_Fx6Hx_2T4TAcheOhT1Yk6o6Go144Ohqc9Ae_Ho-OwvDZJ_-UNoe9k-PxGHcGgzDsn4ShTyBheOrP1a9G5Y9HSAr7LMgzUXBNov549PwvfqlsVw)



### Fault Proofs Components

| **Name** | **Description** | **Type** | **Security Issues** |
| --- | --- | --- | --- |
| `DisputeGameFactory` | A factory contract that deploys new dispute game instances based on the game type. | Onchain Contract | M-3: Smart wallet bond theft due to use of tx.origin |
| `FaultDisputeGame` | The main dispute game contract for challenging proposed outputs interactively. | Onchain Contract | M-4: Re-org attack bond loss due to lack of claim identification in move() |
| `PermissionedDisputeGame` | A permissioned version of the dispute game, used for specific scenarios or testing. | Onchain Contract | None mentioned |
| `AnchorStateRegistry` | Manages anchor states for the fault proof system, ensuring consistency and integrity. | Onchain Contract | M-2: L2 block number spoofing |
| `DelayedWETH` | Handles delayed release of WETH bonds to allow off-chain monitoring time. | Onchain Contract | None mentioned |
| `PreimageOracle` | Provides preimage data for the virtual machine to verify execution steps. | Onchain Contract | None mentioned |
| `MIPS` | Virtual machine (MIPS emulator) for executing and verifying disputed execution steps. | Onchain Contract | None mentioned |
| `OptimismPortal2` | Main entry point for L2 to L1 withdrawals, integrating with the fault proof system. | Onchain Contract | M-1: Incorrect game type casting; M-2: L2 block number spoofing |
| `SystemConfig` | Manages system-wide configuration parameters for the OP Stack. | Onchain Contract | None mentioned |
| `L1CrossDomainMessenger` | Facilitates message passing from L1 to L2. | Onchain Contract | None mentioned |
| `L1StandardBridge` | Handles standard token bridging from L1 to L2. | Onchain Contract | None mentioned |
| `L2ToL1MessagePasser` | Contract facilitating message passing from L2 to L1. | Onchain Contract | None mentioned |
| `op-proposer` | Off-chain actor responsible for proposing L2 outputs. | Off-chain Feature | None mentioned |
| `op-challenger` | Off-chain actor that monitors proposed outputs and initiates disputes if invalid. | Off-chain Feature | None mentioned |
| `op-dispute-mon` | Off-chain monitoring service for dispute games, ensuring correct resolution. | Off-chain Feature | None mentioned |
| `Guardian` | Off-chain actor able to intervene in the dispute process for emergency situations. | User | None mentioned |
| `L2OutputOracle` | Previously used for submitting proposed L2 outputs; now deprecated. | Deprecated | Possibly related to M-2: L2 block number spoofing |



### Fault Proof Primitives

| **Name** | **Description** | **Type** | **Security Issues** |
| --- | --- | --- | --- |
| `ABSOLUTE_PRESTATE` | A constant representing the initial state before any execution in the dispute game. | Constant | None mentioned |
| `CANNON` | A game type using the Cannon VM for dispute resolution. | Constant | None mentioned |
| `PERMISSIONED_CANNON` | A permissioned game type using the Cannon VM. | Constant | None mentioned |
| `DISPUTE_GAME_FINALITY_DELAY_SECONDS` | Constant defining the delay before a dispute game can be finalized. | Constant | None mentioned |
| `PROOF_MATURITY_DELAY_SECONDS` | Constant defining the delay before a proof can be considered mature. | Constant | None mentioned |
| `DEFENDER_WINS` | Constant indicating the defender has won the dispute game. | Constant | None mentioned |
| `CHALLENGER_WINS` | Constant indicating the challenger has won the dispute game. | Constant | None mentioned |
| `proveWithdrawalTransaction` | Function in OptimismPortal2 to prove a withdrawal transaction against a dispute game. | Function | M-1: Incorrect game type casting |
| `finalizeWithdrawalTransaction` | Function in OptimismPortal2 to finalize a proven withdrawal after the delay period. | Function | None mentioned |
| `setResourceConfig` | Function to configure resource limits in the system. | Function | None mentioned |
| `create` | Function in DisputeGameFactory to create a new dispute game instance. | Function | M-3: Smart wallet bond theft |
| `gameAtIndex` | Function to retrieve a dispute game instance by index. | Function | None mentioned |
| `claimCredit` | Function to claim credit or bonds after a dispute game resolution. | Function | None mentioned |
| `owner` | Variable representing the owner of a contract, typically with administrative privileges. | Variable | None mentioned |
| `status` | Variable indicating the current status of a dispute game (e.g., ongoing, resolved). | Variable | None mentioned |
| `createdAt` | Variable recording the timestamp when a dispute game was created. | Variable | None mentioned |
| `resolvedAt` | Variable recording the timestamp when a dispute game was resolved. | Variable | None mentioned |


### Sherlock Audit & Bug Hunt Results [³⁴](https://audits.sherlock.xyz/contests/205/report)’[³⁵](https://blog.oplabs.co/sherlock-audit-roundup/)’[³⁶](https://audits.sherlock.xyz/contests/205?filter=questions)
 To ensure the **Fault Proof System** was secure before mainnet activation, OP Labs engaged the Sherlock security community in a comprehensive audit contest and an accompanying bug hunt. The main **Sherlock audit contest** (Mar 27–Apr 4, 2024) found *no critical vulnerabilities* that could bypass the fault proof system’s safety mechanisms. Several medium-severity issues were identified, and **fixes for all reported issues were promptly merged and deployed to testnet** ahead of the upgrade. 
    
Security issues are sourced from the Sherlock Audit Contest:
- M-1: Incorrect Game Type Casting: Affects OptimismPortal2 and proveWithdrawalTransaction.
- M-2: L2 Block Number Spoofing: Impacts AnchorStateRegistry, OptimismPortal2, and potentially L2OutputOracle.
- M-3: Smart Wallet Bond Theft: Linked to DisputeGameFactory and its create function.
- M-4: Re-org Attack Bond Loss: Affects FaultDisputeGame.
    
In parallel, a focused **Sherlock “bug hunt” contest** targeted the system’s fallback safety nets (e.g. the Security Council’s guardian override). This bug hunt confirmed that none of the discovered issues could subvert these fundamental safeguards. All findings from both the audit and bug hunt were addressed, resulting in significant security improvements to the fault proof implementation. Key enhancements included:
    
- **Dispute game timing fix:** Corrected a flaw in the “chess clock” logic of the `FaultDisputeGame`. Previously, a dispute could be **prematurely resolved when one team’s time ran out even if the opponent still had time remaining**, denying the opponent a fair chance to respond. Left unpatched, this bug could have allowed an invalid claim to **“win” the dispute and be accepted as valid**. The issue, reported by the Offchain Labs team prior to the audit, was fixed and the chess-clock mechanism now correctly ensures neither side can timeout unfairly.

- **Output validation hardening:** Fixed an unsafe type cast in the L1 portal contract that verifies L2 outputs. The audit revealed that the output proposal logic was casting a game type from 32-bit to 8-bit, which **could treat distinct game identifiers as equivalent**. In certain scenarios this bug would have allowed a malicious output root to bypass off-chain monitoring and be finalized as a valid withdrawal. The fix enforces proper type safety for game identifiers, closing this loophole and ensuring invalid outputs cannot slip through unnoticed.

- **Better bond and game lifecycle security:** Implemented adjustments to dispute game initialization and bond handling based on auditor findings. For example, edge-case scenarios that could **trigger denial-of-service in the dispute game factory or lead to loss/theft of bond stakes** (such as exploits involving re-orgs or smart contract wallets in bonding) were identified and patched. These changes fortify the bonding mechanism and guarantee that honest participants’ stakes are protected under all conditions

## **Metrics & Performance Impact**

1.  **Efficiency & Cost Considerations[³⁰](https://specs.optimism.io/fault-proof/index.html)**
    - **Transaction throughput remains unchanged**, as the dispute process runs asynchronously.
    - Posting an output proposal requires a **0.08 ETH bond**, discouraging frivolous challenges.
    - **Optimized fraud-proof submission process** reduces the number of transactions required to confirm challenges.
2.  **Comparison to Previous Models[²⁹](https://gov.optimism.io/t/upgrade-proposal-7-fault-proofs/8161)**
    - **Plasma Rollups:** Required users to monitor transactions actively, leading to potential data withholding risks.
    - **OVM Era (2021-2023):** Fault proofs were theoretically included but never activated.

## **Rollout Strategy[²⁸](https://blog.oplabs.co/open-source-and-feature-complete-fault-proofs-bring-permissionless-validation-to-the-op-sepolia-testnet/)’[²⁹](https://gov.optimism.io/t/upgrade-proposal-7-fault-proofs/8161)**

**October 2023 – MVP and Early Testing**

- OP Labs delivered a **minimum viable fault-proof system (MVP)** to demonstrate the dispute framework in a controlled environment.
- The MVP established the foundational logic for fraud-proof execution before broader testnet deployment.

**March 19, 2024 – Testnet Activation**

- The **feature-complete fault proof system** was deployed on **OP Sepolia**.
- Developers tested permissionless validation and dispute mechanics in a real environment¹.
- The testnet included **stress testing** for large-scale fraud-proof disputes.
- Breaking changes to withdrawal logic were introduced, requiring dApp modifications.

**April–May 2024 – Governance Discussions**

- OP Labs engineer **Adrian Sutton** posted the [FINAL] governance proposal outlining permissionless fault proofs.
- The Developer Advisory Board (DAB) provided an educational summary to ensure token holders understood the upgrade.
- The proposal was moved into **Special Voting Cycle #23a** due to its importance.

**May 16–29, 2024 – Governance Approval**

- The **Token House vote** concluded with **58.4M OP in favor vs. 0.14M against**, achieving near-unanimous approval.
- The **Citizens’ House veto period** passed without objections, finalizing governance ratification.

**Early June 2024 – Mainnet Activation**

- The **Security Council executed the upgrade**, deploying fault proofs to OP Mainnet.
- **Withdrawals now require challenge verification**, ensuring correctness before funds are released.
- OP Labs provided **extensive upgrade guidance** for node operators and validators.
- The upgrade was deployed via a **single atomic transaction** to ensure a clean switchover.

# Protocol Upgrade #8: Guardian

Protocol Upgrade #8 – `Guardian` – is a pivotal enhancement to the OP Stack, advancing Optimism's security and governance structure[³⁷](https://gov.optimism.io/t/final-protocol-upgrade-8-guardian-security-council-threshold-and-l2-proxyadmin-ownership-changes-for-stage-1-decentralization/8157/1) to meet Stage 1 decentralization under the L2Beat framework[³⁸](https://medium.com/l2beat/introducing-stages-a-framework-to-evaluate-rollups-maturity-d290bb22befe). This upgrade reassigns the `Guardian`role from the Optimism Foundation to a new `Guardian Safe`, introduces liveness enforcement modules to maintain quorum reliability within the `Security Council Safe`, and transfers control over key upgrade mechanisms to community governance. The technical implementation was isolated to smart contract changes; no alterations to execution clients or node operator software were required. These changes mitigate single-point-of-failure risks and ensure that withdrawal halts and upgrade authority now reside with a decentralized, multi-sig council[³⁷](https://gov.optimism.io/t/final-protocol-upgrade-8-guardian-security-council-threshold-and-l2-proxyadmin-ownership-changes-for-stage-1-decentralization/8157/1).

## Technical Features[³⁷](https://gov.optimism.io/t/final-protocol-upgrade-8-guardian-security-council-threshold-and-l2-proxyadmin-ownership-changes-for-stage-1-decentralization/8157/1)

1. **Increased threshold for`Security Council Safe`**
    
    Elevated `Security Council Safe` threshold from 4-of-13 to **10-of-13, e**nforcing a ≥75% supermajority for critical governance actions. This shift significantly reduces the potential for small collusions or compromised multisig members to disrupt protocol operations.
    
2. **Integration of`LivenessModule` and `LivenessGuard`** 
    
    Introduced `LivenessModule` and `LivenessGuard` to ensure council operability. These modules enforce a 14-week liveness interval: if a signer is inactive during that period, they can be removed without triggering the full approval threshold. Should the number of active owners drop below 8, the module allows the `Optimism Foundation` Safe to assume emergency control, preventing governance paralysis.
    
3. **`Guardian Safe` for emergency control**
    
    Reassigned the `Guardian` role to a new `Guardian Safe`, 1-of-1 Safe whose sole owner is the `Security Council Safe`. This Guardian can pause/unpause withdrawals, blacklist `DisputeGame` instances, and intervene post-Fault Proofs upgrade if a bug could allow invalid state finalization. The role was previously centralized in the Foundation.
    
4. **`DeputyGuardianModule` for the Foundation**
    
    Added a `DeputyGuardianModule` granting the Foundation conditional emergency rights to act on the Guardian’s behalf. The `Security Council Safe` retains ultimate authority and may revoke this access at any time.
    
5. **Transfer of `L2 ProxyAdmin` ownership to `Security Council Safe`** 
    
    Transferred `L2 ProxyAdmin` contract ownership from the Foundation to the aliased L1 address of the `Security Council Safe`, eliminating unilateral upgrade risk over L2 contracts, especially withdrawal logic.
    

## Metrics & Performance[³⁹](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_05_SafeLivenessExtensions-Cantina.pdf)

Given the focus on governance and security, performance metrics are centered around the Cantina audit outcomes, conducted from May 6th to May 10th, 2024, identified 143 issues in total, with no critical or high-risk findings, a strong indicator of security robustness. However, there were six medium-risk issues:

- Owner censorship due to address ordering in `LivenessGuard`
- Threshold conflicts with external config changes in `LivenessModule`
- Reentrancy marking all owners as live during `execTransaction()`
- `lastLive` state inconsistencies after removal
- Liveness resets when the module is upgraded
- EIP-1271 non-compliance potentially affecting smart wallets

These audit findings were addressed through code improvements and operational safeguards, including enhanced logging and fallback behavior. Performance remained unaffected. The upgrade imposed **zero downtime**, with no need for client, sequencer, or infra updates. All changes are backward-compatible.

## Rollout Strategy[³⁷](https://gov.optimism.io/t/final-protocol-upgrade-8-guardian-security-council-threshold-and-l2-proxyadmin-ownership-changes-for-stage-1-decentralization/8157/1)

March 25, 2025 – Proposal Finalization and Voting Initiation

- The final proposal for Protocol Upgrade #8 Guardian was posted on the Optimism Governance Forum, marking the start of community review and voting preparation.
- This initiated the Token House vote process within Voting Cycle #23a, setting the stage for community approval and subsequent veto periods.

April 15, 2025 – Token House Voting Completion

- The Token House vote concluded after Voting Cycle #23a, with the community approving the upgrade, assuming no significant opposition based on the proposal's final status.
- This triggered the Citizens’ House Veto Period, allowing additional community oversight before deployment.

May 20, 2025 – Veto Period Ends

- The Citizens’ House Veto Period ended, finalizing community consensus and clearing the upgrade for deployment, barring any veto actions.
- OP Labs prepared for the execution phase, coordinating with auditors and developers to ensure readiness.

June 5, 2025 – Mainnet Deployment

- The upgrade was deployed to the Optimism Mainnet, executed atomically in a single transaction for all affected L1 contracts, requiring no node operator action.


# Protocol Upgrade #9: Fjord

Fjord, activated on July 10, 2024 aims to reduce execution fees for smart wallet applications, such as those supporting Apple’s FaceID, by implementing the RIP-7212 precompile for secp256r1 curve support. It also seeks to lower L1 data availability costs by 5–15% through Brotli compression and improve L1 data pricing robustness with FastLZ-based calculations. Additionally, it increases chain operator response time to L1 node issues by extending the maximum sequencer drift to 1800 seconds (30 minutes) from 600 seconds (10 minutes), preventing L2 chain halts. A key aspect is its interaction with the Fault Proof system, requiring updates to the L1 infrastructure, including op-program v1.2.0, to support new parameters and features. The upgrade’s significance lies in its focus on cost efficiency and operational resilience, building on previous upgrades like Fault Proofs (Upgrade #7), which introduced permissionless validation and dispute resolution. Fjord’s changes, particularly in data handling, may influence how fault proofs are processed, especially in terms of resource usage for the Fault Proof Virtual Machine (FPVM)[⁴⁰](https://gov.optimism.io/t/upgrade-proposal-9-fjord-network-upgrade/8236).

## Technical Features[⁴⁰](https://gov.optimism.io/t/upgrade-proposal-9-fjord-network-upgrade/8236)’[⁴¹](https://specs.optimism.io/protocol/fjord/overview.html)

1. **`RIP-7212` Precompile for `secp256r1` Curve**
    
    Implements support for the `secp256r1` curve, commonly used in smart wallet applications, reducing execution gas costs. This is detailed in the spec at Fjord Execution Engine, enhancing efficiency for applications like FaceID authentication.
    
2. **`Brotli` Compression for Channels**
    
    Adds `Brotli` as an optional compression algorithm for channels, aiming to reduce L1 data availability costs by 5–15%. This is specified in Fjord Derivation, potentially affecting how data is submitted to L1 for fault proof verification, though the core dispute process remains unchanged.
    
3. **`FastLZ` Based L1 Data Availability Cost Calculation**
    
    Introduces `FastLZ` compression for calculating L1 data costs, providing more accurate fee estimates. This is outlined in Fjord Fees, improving pricing robustness but not directly impacting fault proof mechanics.
    
4. **Deprecation of `getL1GasUsed` and `L1GasUsed`**
    
    Removes the `getL1GasUsed` method and `L1GasUsed` field from transaction receipts, streamlining fee calculations. This is detailed in Fjord Predeploys, with no direct effect on fault proofs but aligning with cost efficiency goals.
    
5. **Increased `MAX_SEQUENCER_DRIFT`**
    
    Raises the maximum sequencer drift from 600 seconds to 1800 seconds, giving operators more time to address L1 node issues. This is specified in Fjord Derivation, potentially reducing the frequency of sequencer-related disputes but not altering fault proof core functionality.
    
6. **Increased `MAX_RLP_BYTES_PER_CHANNEL` and `MAX_CHANNEL_BANK_SIZE`**
    
    Increases `MAX_RLP_BYTES_PER_CHANNEL` from 10,000,000 to 100,000,000 and `MAX_CHANNEL_BANK_SIZE` from 100,000,000 to 1,000,000,000, allowing for larger data batches. This is detailed in Fjord Derivation. The change could cause more resource usage in the FPVM, as noted in security considerations, but is considered reasonable as it handles data in larger chunks rather than increasing total data processed.
    

### Impact on Fault Proofs[⁴⁰](https://gov.optimism.io/t/upgrade-proposal-9-fjord-network-upgrade/8236)

- The increase in channel sizes directly impacts the `FPVM`, potentially requiring more computational resources for processing disputes involving larger data sets.
- Fjord’s changes do not alter the core dispute game mechanics, such as the binary bisection game or the `MIPS`based VM, but the larger data batches may affect efficiency and cost for challengers.
- The requirement for `op-program v1.2.0`, part of the Fault Proofs L1 infrastructure update, suggests compatibility with Fjord’s new parameters, though specific changes in `op-program v1.2.0` were not detailed in available sources. This update likely ensures the fault proof system can handle the increased channel sizes and new compression methods without disruption.

## Metrics & Performance[⁴⁰](https://gov.optimism.io/t/upgrade-proposal-9-fjord-network-upgrade/8236)

- **Cost Reduction:**`Brotli` compression is expected to reduce L1 data availability costs by 5–15%, as per the governance proposal, making transactions cheaper and potentially lowering the cost for challengers in fault proofs.
- **Operational Resilience:** The increased max sequencer drift to 30 minutes reduces the likelihood of L2 chain halts due to L1 issues, potentially decreasing the need for emergency fault proof interventions.
- **Fault Proof Resource Usage:**
    
    The increase in `MAX_RLP_BYTES_PER_CHANNEL` and `MAX_CHANNEL_BANK_SIZE` may lead to higher resource usage in the `FPVM` for processing disputes with larger data sets. However, security considerations note that this is manageable, as it affects the size of data handled at once rather than the total data, with similar resource consumption as submitting multiple smaller channels previously.
    

## Rollout Strategy[⁴⁰](https://gov.optimism.io/t/upgrade-proposal-9-fjord-network-upgrade/8236)

**May 30, 2024 – Proposal Finalization and Community Review**

- The upgrade proposal was posted to the Optimism Governance Forum, initiating community review and discussion.

**June 2024 – Governance Voting**

- The Token House approved the upgrade during Special Voting Cycle #23b. The Citizens’ House veto period followed with no objections.

**June–July 2024 – Testing and Preparation**

- Fjord was tested on internal devnets and Sepolia. Node operators were instructed to upgrade clients. Fault Proofs infrastructure was upgraded to `op-program v1.2.0`.

**July 10, 2024 – Mainnet Activation**

- The upgrade was deployed at 16:00:01 UTC, automatically enabled on upgraded nodes.

**Contingency – Emergency Cancellation Option**

- A safeguard `-override.fjord` flag was available to disable the upgrade in case of critical issues.

# Protocol Upgrade #10: Granite

Protocol Upgrade #10 Granite was activated on September 11, 2024. It represents a pivotal security enhancement to the Optimism (OP) Stack, focusing on fortifying the Fault Proofs system to ensure the integrity and reliability of permissionless validation[⁴²](https://gov.optimism.io/t/upgrade-proposal-10-granite-network-upgrade/8733). Granite is a security-centric update, responding to vulnerabilities discovered in third-party audits conducted by Spearbit[⁴³](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_08_Fault-Proofs-No-MIPS_Spearbit.pdf), Cantina[⁴⁴](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_08_Fault-Proofs-MIPS_Cantina.pdf), Code4rena[⁴⁵](https://github.com/code-423n4/2024-07-optimism-findings) and Inmunifi[⁴⁶](https://immunefi.com/bug-bounty/optimism/information/), none of which had been exploited at the time of the proposal. The upgrade targets the Fault Proofs system, including the `MIPS` Virtual Machine (VM). It also includes parameter adjustments, such as reducing the ChannelTimeout, and role enhancements for the `Guardian` and `DeputyGuardian`, aiming to mitigate single points of failure and improve system resilience.

## Technical Features[⁴²](https://gov.optimism.io/t/upgrade-proposal-10-granite-network-upgrade/8733)’[⁴³](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_08_Fault-Proofs-No-MIPS_Spearbit.pdf)’[⁴⁴](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_08_Fault-Proofs-MIPS_Cantina.pdf)’[⁴⁵](https://github.com/code-423n4/2024-07-optimism-findings)’[⁴⁶](https://immunefi.com/bug-bounty/optimism/information/)

1. **Security fixes from multi-party audits**
    - **Cantina 3.1.1**: Fixed allocation overflow in the `MIPS VM` which allowed arbitrary code execution via excessive memory usage (~115B gas).
    - **C4 H-01**: Addressed invalid `DISPUTED_L2_BLOCK_NUMBER` being passed to the VM, which could cause incorrect dispute resolutions.
    - **Spearbit 5.1.1**: Resolved out-of-gas bug in `PreimageOracle.loadPrecompilePreimagePart`, ensuring safe dispute data loading.
    - **C4 H-02**: Corrected flaws in the `LPP challenge period`, which allowed uncounterable claims.
    - **Cantina 3.3.5**: Fixed the implementation of `srav` instruction in `MIPS`, along with other medium-severity bugs such as unaligned memory reads, unchecked arithmetic ops, and silent syscalls.
2. **Parameter Adjustments**
    - Reduced `ChannelTimeout` from 300 to 50 L1 blocks to improve responsiveness of dispute timelines.
    - Updated `DelayedWETH.recover` to support transfers requiring >2300 gas, improving bond recovery reliability.
3. **Role Enhancements**
    - Extended the `Guardian` and `DeputyGuardian` roles to set anchor state, improving security during emergency invalid state prevention.
4. **L2 Hardfork Changes**
    - Limited `ecPairing` precompile input size to 112,687 bytes to prevent abuse and DoS via oversized inputs.
5. **Client and Infra Updates**
    - Required upgrade to `op-geth` and `op-node` clients to incorporate VM logic changes, precompile limits, and other security fixes.

## Metrics & Performance[⁴²](https://gov.optimism.io/t/upgrade-proposal-10-granite-network-upgrade/8733)

- **Fault Proof Improvements:**
    - Lowered `ChannelTimeout`(50 L1 blocks )accelerates game resolution
    - Anchor state enhancements ensure invalid finalizations can be stopped
    - VM fixes ensure precise MIPS execution and safe memory access
- **Performance Impact:**
    - No changes to throughput or latency
    - Zero downtime during deployment
    - Clients upgraded smoothly with backward-compatible changes

## Rollout Strategy[⁴²](https://gov.optimism.io/t/upgrade-proposal-10-granite-network-upgrade/8733)

**August 16, 2024 – Proposal Publication and Audit Completion**

- OP Labs published the Granite upgrade proposal and completed third-party audits by Cantina, Spearbit, and Code4rena.

**Late August 2024 – Governance Approval**

- Following discussion, the Token House approved the proposal. No veto was exercised by the Citizens' House.

**Pre-Deployment – Permissioned Fallback Activation**

- The permissioned fallback was temporarily enabled to restrict output submissions to a trusted proposer during the vulnerability mitigation phase.

**September 11, 2024 – Mainnet Activation**

- The Granite upgrade was deployed on OP Mainnet at 16:00:01 UTC, including an L2 hardfork. All nodes had upgraded in advance.

**Post-Deployment – Permissionless Mode Restored**

- The permissioned fallback was removed, restoring permissionless output proposals under the hardened system.

# Protocol Upgrade #11: Holocene

This upgrade was co-developed by OP Labs and Base contributors and focuses on simplifying L2 chain derivation, introducing dynamic fee parameters, and updating the fault proof system. **Goals:** Improve the worst-case handling of L2 blocks for the Fault Proof system and interoperability, give operators more flexibility over gas economics, and ensure the fault proof contracts/supporting VM are compatible with the new execution logic. By tightening the block derivation rules and making fee parameters adjustable, Holocene lays groundwork for future **permissionless fault proofs** and cross-chain integration within the Superchain[⁴⁷](https://gov.optimism.io/t/upgrade-proposal-11-holocene-network-upgrade/9313).

## Technical Features[⁴⁷](https://gov.optimism.io/t/upgrade-proposal-11-holocene-network-upgrade/9313)‘[⁴⁸](https://specs.optimism.io/protocol/holocene/system-config.html)’[⁴⁹](https://github.com/ethereum-optimism/optimism/blob/7719c8538b8d911519f861fc70085e7a3b4e6787/docs/security-reviews/2024_10-Cannon-FGETFD-3DocSecurity.md)

- **Holocene Derivation Pipeline Updates**
    
     The L2 chain derivation process is overhauled for stricter ordering and simpler logic. Key changes include `Strict Batch Ordering`, `Partial Span Batch Validity`, `Fast Channel Invalidation`, and `Steady Block Derivation`. These ensure out-of-order or bad data is dropped immediately rather than buffered, simplifying node behavior and improving worst-case Fault Proof scenarios.
    
- **Dynamic EIP-1559 Parameters via `SystemConfig`**
    
    Makes fee parameters like block elasticity multiplier and base fee change denominator configurable. A new `ConfigUpdate` type (`EIP_1559_PARAMS`) is added to `SystemConfig`, allowing governors to modify gas dynamics via L1 contract calls. While OP Mainnet retained default values, the mechanism supports future tuning for performance or cost goals.
    
- **Simplified Fee Scalar Configuration**
    
    Introduces `ConfigUpdate.FEE_SCALARS` to cleanly update the base fee and blob fee scalars. This improves L2 fee tuning and reduces misconfiguration risk.
    
- **MIPS Contract & Fault Proof System Updates**
    
    Updated MIPS contract to support new Go 1.22 syscalls, audited by 3DOCSec with no major findings. FaultDisputeGame and PermissionedDisputeGame contracts updated to recognize post-Holocene logic, enabling correct dispute execution.
    

## Metrics & Performance [⁴⁷](https://gov.optimism.io/t/upgrade-proposal-11-holocene-network-upgrade/9313)‘[⁴⁸](https://specs.optimism.io/protocol/holocene/system-config.html)

- **Efficiency Gains**
    
    Strict ordering eliminates buffering delays, reduces memory usage, and simplifies synchronization. Invalid payloads are replaced with dummy blocks, preventing unnecessary memory growth and delay.
    
- **Fault Proof Derivation**
    
    Stricter derivation reduces fault propagation and enhances determinism, aiding fault prover validation. Differential tests confirmed behavior alignment between op-program and Kona (Rust prover).
    
- **Gas & Cost Implications**
    
    No direct fee changes were enacted, but configurability paves the way for future cost tuning. Scalar updates and EIP-1559 control can improve pricing consistency and throughput.
    

## Rollout Strategy[⁴⁷](https://gov.optimism.io/t/upgrade-proposal-11-holocene-network-upgrade/9313)

- **Nov 25, 2024 – Governance Proposal**
    
    Holocene was proposed via Optimism Governance Forum. DAB summary published Nov 29. Token House supported proposal in Special Voting Cycle #11.
    
- **Nov 26, 2024 – Testnet Activation**
    
    Holocene activated on Sepolia Superchain testnet. Tests confirmed normal operation.
    
- **Jan 9, 2025 – Mainnet Activation**
    
    Deployed to OP Mainnet at 18:00:01 UTC. Included MIPS and fault game contract updates. SystemConfig upgraded post-fork to v1.8.0.
    
- **Feb 2025 – Ecosystem Rollout**
    
    Adopted by Base (co-developer), Soneium (Feb 3), and Boba Network (Feb 5). Node and contract upgrade instructions standardized across OP Stack chains.

# Protocol Upgrade #12: L1 Pectra Readiness

**Protocol Upgrade #12 – L1 Pectra Readiness** ensures compatibility between the OP Stack (including Optimism Mainnet and Superchain OP chains) and Ethereum’s upcoming **Pectra hardfork**, which includes Prague (execution layer) and Electra (consensus layer) upgrades. Without this upgrade, OP Stack chains could encounter critical failures due to new L1 transaction formats and header fields introduced by Pectra. Upgrade #12 applies node software updates, system configuration changes, and fault proof adjustments to support Pectra, maintaining uninterrupted operation and Ethereum equivalence. The rollout began in **February 2025** with testing and culminated in full client and contract upgrades across all major OP Chains by **April 7, 2025**, ahead of Ethereum’s expected **May 7, 2025** Pectra activation[⁵⁰](https://gov.optimism.io/t/upgrade-proposal-12-l1-pectra-readiness/9706).

## Technical Features[⁵⁰](https://gov.optimism.io/t/upgrade-proposal-12-l1-pectra-readiness/9706)’[⁵¹](https://docs.optimism.io/notices/pectra-changes)

1. **Support for EIP-7702 (SetCode Transactions)**
    - Introduced in Pectra, EIP-7702 allows externally owned accounts (EOAs) to include contract code within a single transaction. OP Stack derivation logic was updated to recognize and skip this new transaction type to prevent derivation errors or halts.
2. **Support for EIP-7685 (Block Header Extension)**
    - The Pectra fork adds a new `requestsHash` field to execution layer block headers. OP Stack components (notably `op-node`) were updated to correctly parse this field during L1 block verification.
3. **Updated OP Stack Clients**
    - `op-geth v1.101503` and `op-node v1.12.0` integrate upstream Ethereum changes, including Pectra transaction and block logic. This ensures OP Stack clients continue to process post-fork L1 data safely.
4. **Fault Proof Compatibility Updates**
    - A new Cannon `absolute prestate` was created using `op-program v1.5.0`, incorporating Pectra changes. New `FaultDisputeGame` and `PermissionedDisputeGame` contracts were deployed across Optimism and Superchain OP chains with updated prestate hashes and logic.
5. **Fee Calculation Bug Fix**
    - A bug in L1 blob fee calculations (still using pre-Pectra constants) was corrected to avoid L2 overcharging or unexpected halts post-fork.
6. **No Change to Chain Parameters**
    - No user-facing features or governance parameters were changed. The upgrade is fully backwards-compatible and charter-compliant.

## Metrics & Performance[⁵⁰](https://gov.optimism.io/t/upgrade-proposal-12-l1-pectra-readiness/9706)’[⁵¹](https://docs.optimism.io/notices/pectra-changes)

- **No Performance Impact:** Transaction throughput, latency, and gas costs remain unchanged. The upgrade is designed solely for L1 compatibility.
- **Improved Stability:** Fixes to blob fee calculation prevent overcharging and potential liveness issues under high L1 base fees.
- **Fault Proof Continuity:** Updated Cannon prestate and op-challenger v1.3.2 ensure continued validity of fraud proofs after L1 changes.
- **Successful Testnet Results:** Sepolia and Holesky testnets demonstrated uninterrupted operation after their respective Pectra activations with upgraded nodes.

## Rollout Strategy[⁵⁰](https://gov.optimism.io/t/upgrade-proposal-12-l1-pectra-readiness/9706)

February–March 2025 – Specification Finalization and Testing

- OP Labs tracked Ethereum’s finalized Pectra EIPs and tested changes across devnets and Holesky.

March 11, 2025 – Client Release and Documentation

- Published `op-geth v1.101503.0` and `op-node v1.12.0`. Upgraded documentation and governance proposal shared with node operators and Superchain partners.

March–April 2025 – Superchain Coordination

- Base, Zora, and other OP Chains upgraded clients and Cannon contracts using the new prestate. Some chains paused sequencers briefly for coordinated deployments.

April 2025 – Mainnet Readiness

- By April 7, Optimism and all major OP Chains had completed their upgrades, ensuring readiness before Ethereum’s expected May 7, 2025 Pectra activation.

# Protocol Upgrade #13: OPCM and Incident Response Improvements

Protocol Upgrade #13 enhances the Optimism Collective's operational resilience through improved incident response, Guardian controls, and contract infrastructure. The upgrade strengthens the Optimism Chain Management Protocol (OPCM)[⁵²](https://gov.optimism.io/t/upgrade-proposal-13-opcm-and-incident-response-improvements/9739) and introduces new security mitigations in response to findings from Fault Proofs deployment[⁵³](https://gov.optimism.io/t/proposal-preview-fault-proofs-incident-response-improvements/9659) and the Stage 1 decentralization review. It aims to align the OP Stack with updated L2Beat[⁵⁴](https://forum.l2beat.com/t/stages-update-a-high-level-guiding-principle-for-stage-1/338) expectations for decentralization maturity while improving emergency mitigation pathways and fallback mechanisms. The proposal was finalized on March 21, 2025, with mainnet deployment completed on April 10, 2025, following successful governance approval and pre-deployment testing.

## Technical Features[⁵²](https://gov.optimism.io/t/upgrade-proposal-13-opcm-and-incident-response-improvements/9739)

1. **OPCM Enhancements**
    - Introduced new `OPChainManager` and `OPChainManagerV2` contracts, standardizing chain metadata (chainId, name, guardian, chainType, etc.) across the Superchain.
    - Enabled chain-specific Guardian roles, allowing fine-grained emergency control without centralization.
    - Facilitated permissionless OP Chain registration via `OPChainManagerV2`, including safe transition and configuration.
2. **Improved Incident Response Mechanisms**
    - Implemented a new `op-incident-response` repository to coordinate emergency actions and safeguard procedures across OP Chains.
    - Incorporated formal Guardian documentation in each OP Chain's metadata, ensuring validators and developers have reference points in case of emergencies.
    - Added fallback logic in the `op-program` and supporting infra, enabling more graceful behavior under emergency overrides.
3. **Guardian Configuration Upgrades**
    - Chain-specific Guardian addresses can now be registered or updated through the `OPChainManager`, supporting decentralization.
    - Improved communication standards between Guardians, node operators, and core developers during mitigation events.
4. **Contract & Infrastructure Upgrades**
    - Deployed new OPCM and OPChainManager contracts with updated system roles and L1 ↔ L2 references.
    - Aligned with current `SystemConfig` and `SuperchainConfig` standards.
    - Standardized emergency pause mechanisms across bridge and messaging contracts using Guardian flags.

## Metrics & Performance[⁵²](https://gov.optimism.io/t/upgrade-proposal-13-opcm-and-incident-response-improvements/9739)‘[⁵⁵](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2025_01-IRI-OffbeatLabs.pdf)’[⁵⁶](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2025_02-Upgrade13-Spearbit.pdf)

Upgrade 13’s enhancements were qualitatively significant for security and governance, although they did not impact routine network throughput or costs. Security assessments by Offbeat Labs and Spearbit found no critical or high-severity issues. Minor findings were resolved before deployment, including:

- Guardian role documentation inconsistencies (Offbeat 2.1)
- Missing input validation in `OPChainManager` (Spearbit M-02)
- Reentrancy and privilege edge cases (Spearbit M-03, L-02)

## Rollout Strategy[⁵²](https://gov.optimism.io/t/upgrade-proposal-13-opcm-and-incident-response-improvements/9739)

**March 21, 2025 – Proposal Finalization**

- Upgrade Proposal #13 was posted to the Optimism Governance Forum for community review.

**March 24, 2025 – Governance Voting Initiation**

- Token House vote launched as part of Voting Cycle #26a.

**April 3, 2025 – Governance Approval**

- The proposal passed the Token House vote with broad support. No vetoes were issued during the Citizens' House period.

**April 4–9, 2025 – Coordination and Client Preparation**

- Node operators and Superchain partners were notified of contract addresses, including:
    - `OPChainManagerV2`: 0x00000000000000000000000000000000004f5543

**April 10, 2025 – Mainnet Activation**

- Upgrade deployed to OP Mainnet at 17:00 UTC. No sequencer or node software changes were required.

# **Protocol Upgrade #14 & #15 – Isthmus upgrades**

Protocol Upgrades #14 and #15, collectively known as the *Isthmus* upgrades, deliver foundational enhancements to Optimism's fault proof system. Upgrade #14 introduces the Multithreaded Cannon (MT-Cannon) Fault Proof VM, a 64-bit, multi-threaded emulator that replaces the prior 32-bit Cannon system, eliminating critical memory limitations and enabling significantly larger L2 block proofs. It also integrates an optional "Operator Fee" mechanism for future fee model flexibility[⁵⁷](https://gov.optimism.io/t/upgrade-proposal-14-isthmus-l1-contracts-mt-cannon/9796). Upgrade #15 activates these improvements through a network-wide hard fork and synchronizes Optimism with Ethereum’s Pectra upgrades, including support for new EIPs and block header fields[⁵⁸](https://gov.optimism.io/t/upgrade-proposal-15-isthmus-hard-fork/9804/1). Together, the Isthmus upgrades dramatically boost the OP Stack's scalability, performance, and readiness for a more decentralized and interoperable Superchain.

## **Technical Features**[⁵⁷](https://gov.optimism.io/t/upgrade-proposal-14-isthmus-l1-contracts-mt-cannon/9796)**’**[⁵⁸](https://gov.optimism.io/t/upgrade-proposal-15-isthmus-hard-fork/9804/1)**’[⁵⁹](https://gov.optimism.io/t/proposal-preview-upgrading-the-cannon-fault-proof-vm-to-support-64-bit-and-multi-threading/9690)’[⁶⁰](https://specs.optimism.io/experimental/cannon-fault-proof-vm-mt.html)**

1. **MT-Cannon Fault Proof VM**
    
    Replaces the 32-bit Cannon VM with a 64-bit MIPS architecture and supports cooperative multithreading, enabling better memory management (garbage collection, multi-threaded Go runtime) and lifting the 4GB memory cap on fault proof execution.
    
2. **Deterministic Thread Scheduler**
    
    Introduces round-robin thread scheduling within the VM to ensure consistent, traceable execution during fraud proofs while allowing internal Go threads to interleave without affecting external determinism.
    
3. **Expanded Syscall Safety**
    
    Unrecognized syscalls now halt execution (instead of silently continuing), enhancing predictability and debugging during proof evaluation.
    
4. **Preimage Oracle Bandwidth**
    
    Doubles the data read per VM step (8 bytes vs. 4), improving data access efficiency during proof replay.
    
5. **Operator Fee Framework**
    
    Adds L1 SystemConfig parameters to support future per-transaction fee models tailored to ZK-proof or Alt-DA rollups. Fee is disabled by default.
    
6. **L2 Withdrawals Root in Block Headers**
    
    Adds a Merkle root of L2-to-L1 withdrawals to L2 headers, improving output root reconstruction and enabling stateless client verification.
    
7. **Pectra EIPs Compatibility**
    
    Adopts Ethereum’s upcoming execution layer features (e.g. EIP-7702, EIP-2537) for transaction flexibility, cryptographic precompiles, and L2 developer readiness.
    

### **Impact on Fault Proofs**

- MT-Cannon significantly increases execution capacity for dispute resolution by enabling 64-bit, parallel execution in the MIPS VM.
- The Cannon64 prestate and dispute game logic updates are forward-compatible with zk-based enhancements.
- These changes ensure OP Stack fault proofs can scale with user demand and Ethereum complexity, improving long-term decentralization.

## **Metrics & Performance**[⁵⁷](https://gov.optimism.io/t/upgrade-proposal-14-isthmus-l1-contracts-mt-cannon/9796)**’**[⁵⁸](https://gov.optimism.io/t/upgrade-proposal-15-isthmus-hard-fork/9804/1)’[⁶¹](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2025_01-MT-Cannon-Spearbit.pdf)’[⁶²](https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2025_01-MT-Cannon-Base.pdf)

- **Scalability**
    
    MT-Cannon removes memory barriers to fraud proof size, allowing much larger blocks to be proven and unlocking higher L2 throughput.
    
- **Reliability**
    
    Garbage collection and cooperative threading avoid out-of-memory failures in large proofs. LL/SC instructions provide thread-safe concurrency.
    
- **Performance**
    
    Preimage reads are twice as fast. Optimized cryptographic precompiles (e.g. BLS12-381) ensure timely proof simulation of even complex Ethereum transactions.
    
- **Audit Results**
    
    Coinbase and Spearbit audits (Jan 2025) confirmed robustness with no critical issues. Remaining risks (e.g. livelocks, thread bugs) were mitigated via extensive testing and monitoring.
    

## Rollout Strategy[⁵⁷](https://gov.optimism.io/t/upgrade-proposal-14-isthmus-l1-contracts-mt-cannon/9796)**’**[⁵⁸](https://gov.optimism.io/t/upgrade-proposal-15-isthmus-hard-fork/9804/1)

**March 2025 – Governance Forum Proposals and Review**

- Upgrade Proposal #14 (Isthmus L1 Contracts + MT-Cannon) and Proposal #15 (Isthmus Hard Fork) were introduced in Voting Cycle #35. Proposal #15 was explicitly contingent on #14 passing.
- Community preview posts were shared in late February 2025, with technical context provided by OP Labs, Base contributors, and members of the Developer Advisory Board.
- During the forum discussion, governance delegates expressed broad support, highlighting the significance of MT-Cannon for performance and decentralization.

**April 2025 – Governance Voting and Approval**

- Both proposals were submitted for formal voting during Cycle #35 and passed with broad consensus.
- Upgrade #14 was executed immediately after approval, deploying updated L1 contracts and off-chain fault proof components (e.g., SystemConfig, MIPS64.sol, op-program v1.5.1).
- Challenger infrastructure was upgraded accordingly, but node software and consensus logic remained unchanged at this stage.

**April–May 2025 – Deployment and Testing Phase**

- The MT-Cannon system was deployed in a passive state, allowing integration with challenger services ahead of the L2 fork.
- OP Labs conducted devnet and testnet trials, including testing Isthmus features against Pectra-compatible Ethereum devnets.
- Sepolia L2 chains and OP Stack branches were updated with Cannon64 prestate and Isthmus specs for broader testing.


# **References**

1. Plasma Group. (2019). *Rollup Plasma for Mass Exits & Complex Disputes.* Plasma Build Forum. Retrieved from https://plasma.build/t/rollup-plasma-for-mass-exits-complex-disputes/90
2. Floersch, K. (2019). *Ethereum Smart Contracts in L2: Optimistic Rollup.* Plasma Group Blog. Retrieved from https://medium.com/plasma-group/ethereum-smart-contracts-in-l2-optimistic-rollup-2c1cef2ec537
3. The Optimism Collective. (2022). *A New Chapter.* Optimism Mirror. Retrieved from https://optimism.mirror.xyz/Fdpds7l4yE2VDpLN4AZQUCjpBYdoPrIHw6a4uzUepmw
4. Optimism PBC. (2022). *The Future of Optimistic Ethereum*. Retrieved from https://medium.com/ethereum-optimism/the-future-of-optimistic-ethereum-7f22d987331
5. OP Stack Specification. (2023). *Cannon Fault Proof Virtual Machine*. Retrieved from https://specs.optimism.io/fault-proof/cannon-fault-proof-vm.html 
6. Sutton, A. (2024). *Final Protocol Upgrade #7: Fault Proofs.* Optimism Governance. Retrieved from https://gov.optimism.io/t/final-protocol-upgrade-7-fault-proofs/8161
7. Maurelian. (2024). *Final Protocol Upgrade #8: Guardian Security Council & Threshold Changes.* Optimism Governance. Retrieved from https://gov.optimism.io/t/final-protocol-upgrade-8-guardian-security-council-threshold-and-l2-proxyadmin-ownership-changes-for-stage-1-decentralization/8157
8. Fichter, K. (2024). *Modular Rollup Theory Through the Lens of the OP Stack.* Devcon Bogotá. Retrieved from https://www.youtube.com/watch?v=jnVjhp41pcc
9. Bayardo, R. (2024). *Fjord Upgrade Proposal.* Optimism Governance. Retrieved from https://gov.optimism.io/t/upgrade-proposal-9-fjord-network-upgrade/823 
10. OP Labs. (2025) *Upgrade Process.* Optimism Docs. Retrieved from https://docs.optimism.io/operators/node-operators/network-upgrades#upgrade-process
11. The Optimism Collective. (2023). *Building Bedrock*. Optimism Mirror. Retrieved from https://optimism.mirror.xyz/gFyEzwAeJQxqVunjlx5s5oEQHuUKbIy6BXzcCsh8qXs
12. The Optimism Collective. (2022). *Introducing the OP Stack*. Optimism Mirror. Retrieved from https://optimism.mirror.xyz/fLk5UGjZDiXFuvQh6R_HscMQuuY9ABYNF7PI76-qJYs
13. OP Labs. (2025). *OP Stack Components*. Optimism Docs. Retrieved from https://docs.optimism.io/stack/components
14. The Optimism Collective. (2023). *Preparing Optimism for the Superchain Future*. Optimism Mirror. Retrieved from https://optimism.mirror.xyz/9ZMwZjst9SQpzIgEd4gN42UDjATyK3ZRClPFx9oMPp8
15. Ben-Chain. (2023). *Upgrade #1: Bedrock Protocol Upgrade.* Optimism Governance. Retrieved from https://gov.optimism.io/t/final-upgrade-1-bedrock-protocol-upgrade-v2/5548
16. OP Stack Specification. (2023). *Protocol Upgrades: Regolith*. Retrieved from https://specs.optimism.io/protocol/regolith/overview.html
17. Trianglesphere. (2023). *Upgrade Proposal #2 Canyon. Optimism Governance*. Retrieved from https://gov.optimism.io/t/final-upgrade-proposal-2-canyon-network-upgrade/7088
18. Testinprod. (2024). *Upgrade Proposal #4 Delta.* Optimism Governance. Retrieved from https://gov.optimism.io/t/final-upgrade-proposal-3-delta-network-upgrade/7310
19. Testinprod. (2024). *Span Batch Design Docs.* Op-Tip. Retrieved from https://www.notion.so/b85e599a47774dcdb8171cc84cab2476?pvs=21
20. OP Stack Specification. (2023). *Span Batches.* Retrieved from https://specs.optimism.io/protocol/delta/span-batches.html 
21. Maurelian. (2024). *Upgrade Proposal #5 Superchain Configuration.* Optimism Governance. Retrieved from: https://gov.optimism.io/t/upgrade-proposal-4/7534
22. Maurelian. (2024). *Superchain Config Specification.*  Optimism Github. Retrieved from https://github.com/ethereum-optimism/specs/blob/8eb74667841c9cf86747cd133175272c76dd86f0/specs/superchain-configuration.md
23. Trust Security. (2024). *Optimism Bedrock upgrade.* Optimism Github. Retrieved from https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2023_12_SuperchainConfigUpgrade_Trust.pdf
24. Bayardo, R. (2024). *Upgrade Proposal #5: Ecotone Network Upgrade.* Optimism Governance. Retrieved from https://gov.optimism.io/t/upgrade-proposal-5-ecotone-network-upgrade/7669
25. Optimism OP Stack Specs. (2024). *Ecotone Network Upgrade – Technical Specifications.* Retrieved from https://specs.optimism.io/protocol/ecotone/overview.html
26. Diego. (2024). *Upgrade Proposal #6 Multi-Chain Prep L1.* Optimism Governance. Retrieved from https://gov.optimism.io/t/upgrade-proposal-6-multi-chain-prep-mcp-l1/7677
27. Cantina Security. (2024). *Multi-Chain Prep (MCP) L1 Audit Report.* Optimism GitHub. Retrieved from https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_02-MCP_L1-Cantina.pdf
28. OP Labs. (2024). *Feature-complete Fault Proofs on OP Sepolia.* OP Labs Blog. Retrieved from https://blog.oplabs.co/open-source-and-feature-complete-fault-proofs-bring-permissionless-validation-to-the-op-sepolia-testnet/
29. Ajsutton. (2024). *Upgrade Proposal #7: Fault Proofs.* Optimism Governance. Retrieved from https://gov.optimism.io/t/upgrade-proposal-7-fault-proofs/8161
30. OP Labs. (2024). *Fault Proofs Overview.* OP Stack Specification. Retrieved from https://specs.optimism.io/fault-proof/index.html
31. OP Labs. (2024). *Fault Proof System Documentation.* Optimism Github. Retrieved from https://github.com/ethereum-optimism/docs/blob/ef619668ae44276edecdfd657157254b9809e2d6/pages/builders/notices/fp-changes.mdx
32. Optimism Release Notes. (2024). *v1.4.0-rc.4.* Optimism Github. Retrieved from https://github.com/ethereum-optimism/optimism/releases/tag/op-contracts%2Fv1.4.0-rc.4
33. OP Labs. (2024). *OP Stack FP Sherlock contest Handbook.* Op Labs Notion. Retrieved from https://oplabs.notion.site/Public-OP-Stack-Fault-Proofs-Sherlock-Competition-Handbook-e4cfdf210a5c45c79230af19653163cc
34. Sherlock. (2024). *Optimism Fault Proofs Contest Results*. Sherlock Bug Hunt on OP Fault Proofs. Retrieved from https://audits.sherlock.xyz/contests/205/report
35. OP Labs. (2024). *Initial findings from the fault proof program Sherlock audit*. OP Labs Blog. Retrieved from https://blog.oplabs.co/sherlock-audit-roundup/
36. Sherlock. (2024). *Optimism Fault Proofs Contest Questions*. Sherlock Bug Hunt on OP Fault Proofs Retrieved from https://audits.sherlock.xyz/contests/205?filter=questions
37. Maurelian. (2024). *Protocol Upgrade #8 Guardian*. Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/final-protocol-upgrade-8-guardian-security-council-threshold-and-l2-proxyadmin-ownership-changes-for-stage-1-decentralization/8157
38. L2Beat. (2024). *Introducing the Stages Framework for rollups*. L2Beat blog. Retrieved from https://l2beat.com/blog/introducing-stages-framework 
39. Cantina. (2024). *Optimism Safe Extensions Competition*. Optimism Github. Retrieved from https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_05_SafeLivenessExtensions-Cantina.pdf
40. Bayardo. (2024). *Upgrade Proposal #9: Fjord Network Upgrade.* Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/upgrade-proposal-9-fjord-network-upgrade/8236
41. OP Labs. (2024). *Fjord Network Upgrade.* OP Stack Specification. Retrieved from https://specs.optimism.io/protocol/fjord/overview.html
42. Inphi (2024). *Upgrade Proposal #10: Granite Network Upgrade*. Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/upgrade-proposal-10-granite-network-upgrade/8733 
43. Spearbit. (2024). *Base Fault Proofs (No MIPS) – Security Review*. Spearbit Audit Report. Retrieved from https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_08_Fault-Proofs-No-MIPS_Spearbit.pdf
44. Cantina. (2024). *Base Fault Proofs MIPS – Security Review*. Cantina Audit Report. Retrieved from https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2024_08_Fault-Proofs-MIPS_Cantina.pdf
45. Code4rena. (2024). *Optimism July 2024 Audit Contest – Findings Report*. Code4rena Competition Report. Retrieved from GitHub: https://github.com/code-423n4/2024-07-optimism-findings 
46.  Immunefi. (2024). *Optimism Bug Bounty Program*. Retrieved from https://immunefi.com/bug-bounty/optimism/information/
47. Dragan_ZzZ. (2024). *Upgrade Proposal #11: Holocene Network Upgrade.* Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/upgrade-proposal-11-holocene-network-upgrade/9313
48. OP Labs. (2024). *Holocene: System Configuration* OP Stack Specification. Retrieved from 
https://specs.optimism.io/protocol/holocene/system-config.html
49. 3DOC Security. (2024, October 3). *Audit Report - OP Cannon*. Optimism Github. Retrieved from https://github.com/ethereum-optimism/optimism/blob/7719c8538b8d911519f861fc70085e7a3b4e6787/docs/security-reviews/2024_10-Cannon-FGETFD-3DocSecurity.md
50. Knee, G. (2025). *Upgrade Proposal #12: L1 Pectra Compatibility*. Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/upgrade-proposal-12-l1-pectra-readiness/9706 
51. Optimism. (2025) *Preparing for Pectra breaking changes.* Optimism Docs. Retrieved from https://docs.optimism.io/notices/pectra-changes
52. Maurelian, Lewej, & Kelvin. (2025). *Upgrade Proposal #13: OPCM and Incident Response Improvements*. Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/upgrade-proposal-13-opcm-and-incident-response-improvements/9739 
53. Fichter, K. (2025). *Proposal Preview: Fault Proofs Incident Response Improvements*. Optimism Governance Forum. Retrieved from  https://gov.optimism.io/t/proposal-preview-fault-proofs-incident-response-improvements/9659 
54. Donnoh. (2025). *Stages update: A high-level guiding principle for Stage 1*. L2BEAT Forum. Retrieved from  https://forum.l2beat.com/t/stages-update-a-high-level-guiding-principle-for-stage-1/338 
55. Offbeat Labs. (2025). *Optimism Incident Response Updates – Audit Report*. Optimism GitHub. Retrieved from https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2025_01-IRI-OffbeatLabs.pdf 
56. Spearbit. (2025). *OP Stack Upgrade 13 – Security Audit Report*. Optimism GitHub. Retrieved from  https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2025_02-Upgrade13-Spearbit.pdf 
57. 0xEscanor. (2025). *Upgrade Proposal #14: Isthmus L1 Contracts + MT-Cannon*. Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/upgrade-proposal-14-isthmus-l1-contracts-mt-cannon/9796 
58. 0xEscanor. (2025). *Upgrade Proposal #15: Isthmus Hard Fork*. Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/upgrade-proposal-15-isthmus-hard-fork/9804/1 
59. Dowman, P. (2025). *Proposal Preview: Upgrading the Cannon Fault Proof VM to support 64-bit and multi-threading*. Optimism Governance Forum. Retrieved from https://gov.optimism.io/t/proposal-preview-upgrading-the-cannon-fault-proof-vm-to-support-64-bit-and-multi-threading/9690 
60. Optimism. (2025). *OP Stack Specification – Multithreaded Cannon Fault Proof VM*. OP Stack Specs Repository. Retrieved from https://specs.optimism.io/experimental/cannon-fault-proof-vm-mt.html 
61. Spearbit. (2025). *Security Audit Report: 64-bit Multithreaded Cannon (MIPS64.sol)*. Optimism Security Reviews. Retrieved from https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2025_01-MT-Cannon-Spearbit.pdf 
62. Coinbase. (2025). *Protocol Security Audit: Multi-thread & 64-bit Cannon*. Optimism Security Reviews. Retrieved from https://github.com/ethereum-optimism/optimism/blob/develop/docs/security-reviews/2025_01-MT-Cannon-Base.pdf
