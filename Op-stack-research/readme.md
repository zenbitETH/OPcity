# OP stack theoretical framework

## Introduction
Optimistic rollups, a type of L2 solution, aggregate multiple transactions into batches and post them on the Ethereum mainnet, significantly reducing the data and computation required on Layer 1[[1]](https://www.notion.so/OP-city-stack-report-fff3905fdc3a800b8b07e6a55b59d456?pvs=21).. This approach enables faster and more cost-effective scalability by assuming transactions are valid by default[[2]](https://www.notion.so/OP-city-stack-report-fff3905fdc3a800b8b07e6a55b59d456?pvs=21).

Optimism is a Layer 2 scaling solution for Ethereum, utilizing Optimistic Rollups to improve scalability, reduce transaction speed and costs, and maintain robust security[[3]](https://www.notion.so/OP-city-stack-report-fff3905fdc3a800b8b07e6a55b59d456?pvs=21). The core of this ecosystem is OP Mainnet, an EVM-equivalent Layer 2 blockchain directly connected to Ethereum, very similar to the Ethereum mainnet, with only minor differences, primarily involving specifying different network endpoints[[4]](https://www.notion.so/OP-city-stack-report-fff3905fdc3a800b8b07e6a55b59d456?pvs=21). OP Mainnet leverages the OP Stack, a modular, open-source software development stack maintained by the Optimism Collective. The OP stack facilitates the creation of Optimistic Rollups, production-ready Layer 2 blockchains, and allows developers to independently customize and upgrade components like consensus, execution, and settlement layers, fostering innovation and ensuring long-term adaptability.

The OP Stack also supports the Optimism Superchain, a network of interoperable OP Stack chains with common standards and protocols. This modular design allows for scalability, customization, interoperability and innovation, helping avoid the need for repeatedly building similar software from scratch[[4]](https://www.notion.so/OP-city-stack-report-fff3905fdc3a800b8b07e6a55b59d456?pvs=21).  The introduction of a permissionless Fault Proof System on OP Mainnet marks a significant step toward full decentralization, removing dependence on privileged roles for withdrawals and allowing any user to challenge and verify transactions, thus reducing reliance on centralized authorities. As the OP Stack evolves, it will incorporate additional functionalities, such as more sophisticated fault-proof systems and interoperability features, and paves the way for a robust, interconnected ecosystem of decentralized applications on Ethereum. The current iteration of the OP Stack, known as Optimism Bedrock, not only powers OP Mainnet but also simplifies deploying Optimistic Rollups and supports the Superchain concept.

This vision aims to create a scalable and interconnected blockchain ecosystem, where multiple L2 chains can seamlessly interact while benefiting from shared security and development standards. As the technology proves its security and reliability, the plan is to extend these capabilities, including fraud proofs, to other networks within the Optimism ecosystem, such as Base, Metal, Mode, and Zora, further solidifying the Superchain concept and pushing the boundaries of Layer 2 scalability and interoperability [[5]](https://www.notion.so/OP-city-stack-report-fff3905fdc3a800b8b07e6a55b59d456?pvs=21)..

## OP Stack

_Copy from [Optimism Docs](https://docs.optimism.io/)_

> ***The OP Stack is a common development stack for building L2 blockchain ecosystems, built by the Optimism Collective to power Optimism.***
> *The OP Stack is best thought of as a collection of software components maintained by the Optimism Collective that either help to define new layers of the stack or fit in as modules within the stack.*


- 1 Data Availability
    
    > *The Data Availability Layer defines where the raw inputs to an OP Stack based chain are published. An OP Stack chain can use one or more Data Availability module to source its input data. Because an OP Stack chain is derived from the Data Availability Layer, the Data Availability module(s) used have a significant impact on the security model of a system*
    > 
    
    > *Ethereum DA is currently the most widely used Data Availability module for the OP Stack. When using the Ethereum DA module, source data can be derived from any piece of information accessible on the Ethereum blockchain. This includes Ethereum calldata, events, and 4844 data blobs.*
    > 
- 2 Sequencing Layer
    
    > *The Sequencing Layer determines how user transactions on an OP Stack chain are collected and published to the Data Availability Layer module(s) in use. In the default Rollup configuration of the OP Stack, Sequencing is typically handled by a single dedicated Sequencer. Rules defined in the Derivation Layer generally restrict the Sequencer's ability to withhold transactions for more than a specific period of time*
    > 
- 3 Derivation Layer (Rollup / OP node)
    
    > The Derivation Layer defines how the raw data in the Data Availability Layer is processed to form the processed inputs that are sent to the Execution Layer via the standard [Ethereum Engine API(opens in a new tab)](https://github.com/ethereum/execution-apis/blob/94164851c1630ff0a9c31d8d7d3d4fb886e196c0/src/engine/README.md). The Derivation Layer may also use the current system state, as defined by the Execution Layer, to inform the parsing of raw input data. The Derivation Layer can be modified to derive Engine API inputs from many different data sources. The Derivation Layer is typically tied closely to the Data Availability Layer because it must understand how to parse any raw input data.
    > 
- 4 Execution Layer
    
    > *The Execution Layer defines the structure of state within an OP Stack system and defines the state transition function that mutates this state. State transitions are triggered when inputs are received from the Derivation Layer via the Engine API. The Execution Layer abstraction opens up the door to EVM modifications or different underlying VMs entirely.*
    > 
    > 
    > ### ***EVM***
    > 
    > *The EVM is an Execution Layer module that uses the same state representation and state transition function as the Ethereum Virtual Machine. The EVM module in the Ethereum Rollup configuration of the OP Stack is a [lightly modified(opens in a new tab)](https://op-geth.optimism.io/) version of the EVM that adds support for L2 transactions initiated on Ethereum and adds an extra L1 Data Fee to each transaction to account for the cost of publishing transactions to Ethereum.*
    > 
- 5 Settlement Layer
    
    > *The Settlement Layer is a mechanism on external blockchains that establish a **view** of the state of an OP Stack chain on those external chains (including other OP Stack chains). For each OP Stack chain, there may be one or more Settlement mechanisms on one or more external chains. Settlement Layer mechanisms are **read-only** and allow parties external to the blockchain to make decisions based on the state of an OP Stack chain.*
    > 
    > 
    > *The term "Settlement Layer" has its origins in the fact that Settlement Layer mechanisms are often used to handle withdrawals of ETH and tokens out of a blockchain. This sort of withdrawal system first involves proving the state of the target blockchain to some third-party chain and then processing a withdrawal based on that state. The Settlement Layer, at its core, simply allows a third-party chain to become aware of the state of the target chain.*
    > 
    > An Attestation-based Fault Proof mechanism uses an optimistic protocol to establish a view of an OP Stack chain. In optimistic settlement mechanisms generally, **Proposer** entities can propose what they believe to be the current valid state of the OP Stack chain. If these proposals are not invalidated within a certain period of time (the "challenge period"), then the proposals are assumed by the mechanism to be correct. In the Attestation Proof mechanism in particular, a proposal can be invalidated if some threshold of pre-defined parties provide attestations to a valid state that is different than the state in the proposal. This places a trust assumption on the honesty of at least a threshold number of the pre-defined participants.
    > 
- 6 Governance Layer
    
    > The Governance Layer refers to the general set of tools and processes used to manage system configuration, upgrades, and design decisions. This is a relatively abstract layer that can contain a wide range of mechanisms on a target OP Stack chain and on third-party chains that impact many of the other layers of the OP Stack.
    >

    ![Op stack diagram](https://www.notion.so/image/https%3A%2F%2Fprod-files-secure.s3.us-west-2.amazonaws.com%2Fb93d7c62-4ce7-46c6-a8c7-8f9ce471d508%2F8cc7f860-ec39-4102-a455-0dd1d9e1e430%2FScreenshot_2024-01-04_at_10.47.03_a.m..png?table=block&id=e8e00522-7b8e-41e6-81fa-fbcf87929f4b&spaceId=b93d7c62-4ce7-46c6-a8c7-8f9ce471d508&width=2000&userId=9ec5d830-7a6e-4543-99a2-3a110d7fec88&cache=v2)


    ## References

    [[1]](https://plasma.io/plasma.pdf) Poon, J., & Buterin, V. (2016). Plasma: Scalable Autonomous Smart Contracts.

    [[2]](https://vitalik.eth.limo/general/2021/01/05/rollup.html) Buterin, v. (2021) An Incomplete Guide to Rollups.

    [[3]](https://docs.optimism.io/connect/resources/glossary#optimism-collective) Optimism documents: Glossary. 

    [[4]](https://docs.optimism.io/chain/getting-started) Getting started: OP Mainnet.

    [[5]](https://blog.oplabs.co/the-fault-proof-system-is-available-for-the-op-stack/) Optimism Developer Blog. Protocol Development: The Fault Proof System is available for the OP Stack.