<script type="text/javascript" src="http://cdn.mathjax.org/mathjax/latest/MathJax.js?config=default"></script>

# OPcity
_Proposed during the [OP Governance Season 5](https://app.charmverse.io/op-grants/page-1306815702055122) and grant finalist during cycle 22_

We will research the compatibility of the OP stack's Canon Fault Proof VM with the opML's Multi-Phase Fault Proof protocol. The goal is to implement a custom Fault Dispute Game that manages the challenges related to data availability states from the L2 rollup and the computation results from the Deep Neural Networks (DNN) of the multi-phase opML. This could be possible with an incentive mechanism for node verifiers that resolve disputes for both technologies.

## Proposed Modifications:

OpML uses a multi-phase fraud-proof to ensure the accuracy of machine learning results onchain. This mechanism is similar to experimental Canon Fault Proofs from the OP stack. Both technologies use a Fault Dispute Game to allow verifiers to resolve challenges on a game tree. Within this process, we aim to expand the merklized data on the OP stack FPVM to include the state transition in the opML Multi-Phase dispute game. By doing this, we aim to potentially achieve a unified framework capable of natively processing machine learning inferences onchain. By digging a bit deeper, we intend to find out if this framework’s implementation can help with the current specs or be an alternative version to the existing specs of FPVM.

### 1. State Transition Function Modeling

The FPVM functions as a state transition system where a function _f_ maps a pre-state _S<sub>pre</sub>_ to a post-state _S<sub>post</sub>_ based on an executed instruction:
$$ f(S_{pre})→ S_{post}$$

For integration:

- Proposed Framework Modification: Introduce an additional layer that handles complex decision trees or neural network outputs, which adjusts how the state transitions are computed, especially in handling error states or exceptions.

- Consider a modified state transition function $$ f(S_{pre}, D)$$ where _D_ represents data or decisions derived from opML processes, impacting the transition to S_post.

- Modified Function:

$$ f(S_{pre}, D)→ S_{post}$$ 

- Define a new state component that includes neural network inference results, which influences the transition process, particularly in how exceptions are handled.

### 2. Memory Management Analysis

Given the detailed memory specifications:

Heap and Memory Operations: Analyze the implications of integrating a mechanism for handling large datasets required by machine learning models directly within the memory structure of FPVM.

Suppose _M(S)_ is the memory utilization state function. Introduce _M'(S, D)_ to handle additional data structures or caching mechanisms to optimize ML data handling.

Memory Function: 

$$ M(S)→ M'(S, D) $$

### 3. Syscalls and I/O Enhancements

The proposed framework could potentially extend the `syscall` and `I/O` functionalities to better support ML-driven data processing:

Extended Syscalls for ML: Introduce new `syscalls` specific to ML operations, such as data batching or model loading.

I/O Modeling: Adjust the I/O model to handle larger data streams efficiently, crucial for ML processes. Propose modifications like enhanced buffer management or asynchronous I/O operations.

### 4. Formal Verification and Error Analysis

Given the complexity of ML integrations and the critical role of fault-proofing:

Model how errors in the ML phase could propagate through the system, influencing state transitions and memory states.

Utilize formal methods to verify the correctness of the integrated system under various operational conditions, ensuring that the modifications do not introduce new vulnerabilities.

### 5. Simulation and Evaluation Metrics

Develop simulations that mimic real-world operational conditions to evaluate the effectiveness of the proposed modifications:

Create scenarios where traditional and ML-modified FPVMs are subjected to typical and atypical loads, measuring performance metrics like throughput, error rate, and response time.

Define specific metrics to evaluate improvements or regressions in system behavior due to the integration, such as memory efficiency, fault detection accuracy, and computational overhead.

Further, use this framework implementation to provide the infrastructure required to process onchain the high volumes of public data generated in cities and use onchain models to make ML inferences from those datasets. To protect the privacy of citizens' sensitive data, we are exploring the implementation of a zkML through the ORA's oppAI framework. This extension strategically balances the trade-offs between privacy and computational efficiency. By leveraging the strengths of zkML's privacy-preserving techniques and opML's computational efficiency, oppAI enables a hybrid model to optimize both aspects for onchain AI applications.

## 

We aim to partner with builders worldwide to create OPcity through a collaborative R&D approach. We will document the entire process on an open-source repository and offer incentives for developers/builders interested in contributing during OPcity development, testing, and release. 