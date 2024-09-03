# Node setting for rollup deployment
This section provides a comprehensive account of our hands-on experience setting up and running a node and deploying a rollup using the OP Stack. We conducted the setup process using two Intel NUC 13 PRO NUC13ANHi7 Arena Canyon devices, each equipped with a 13th-generation Intel Core i7-1360P CPU, 32GB RAM, and a 4TB SSD running Linux (Ubuntu).

We relied on a remote VM and a third-party RPC service in the December tests. However, this setup presented significant challenges, particularly regarding the limitations of RPC calls and the restrictions imposed by proprietary hardware. The use of third-party RPC services introduced latency and potential security concerns, which hindered the overall performance of our test environment. Furthermore, the reliance on a remote VM limited our control over the hardware, leading to issues in scalability and reliability during the testing phase.
 
To overcome these challenges, we transitioned to a home environment setup using the Intel NUC devices. This shift allowed us to bypass the limitations of third-party RPC services and gain full control over the hardware, leading to a more reliable and efficient testing environment. We documented the setup process in our GitHub repository, including all commands and screenshots, to guide others through the deployment process.

We documented the setup process into two main stages required to deploy a rollup on the Holesky testnet:

## A. Spin up an L1 Node (Holesky Testnet):

This stage outlines the steps to configure an L1 node, starting with acquiring the necessary hardware. The Intel NUC devices provided the computing power required to handle the installation of Linux (Ubuntu), followed by the setup of Geth and Prysm dependencies. We then installed and started Geth, forming the Ethereum network's base layer, followed by the installation and initialization of Prysm, which facilitated the execution of the proof-of-stake consensus mechanism. This configuration ensured a robust L1 node capable of supporting subsequent rollup deployments.

1.  [Geth (Execution Layer) setup instructions](https://github.com/zenbitETH/OPcity/blob/main/node-setup/Holesky/geth.md)
2.  [Prysm (Consensus Layer) setup instructions](https://github.com/zenbitETH/OPcity/blob/main/node-setup/Holesky/prysm.md)

### General process
1. Get Node Hardware
2. Install Linux (Ubuntu)
3. Install Geth dependencies
4. Install and start Geth
5. Install Prysm dependencies
6. Install and start Prysm

## B. Deploy a L2 rollup from the OP stack

Building on the foundation of the L1 node, this stage documents the deployment of an L2 rollup using the OP Stack. We began by installing the necessary Optimism dependencies and building the source code. We then proceeded to deploy the L1 contracts and initialize the OP-GETH instance, which involved setting up and executing key components such as op-geth and op-node. The process culminated in obtaining Holesky ETH on the L2 network and sending test transactions, successfully demonstrating the deployment of a rollup in a controlled home environment

1. [OPstack setup and deployment instructions](https://github.com/zenbitETH/OPcity/blob/main/node-setup/Op-stack/op-setup.md)
2. [OPstack Data Visualisation Dashboard setup](https://github.com/zenbitETH/OPcity/blob/main/node-setup/Op-stack/dashboard-setup.md)

### General process
1. Install Optimism dependencies
2. Build the source code
3. Generate and fund accounts and private keys
4. Configure your network
5. Deploy the L1 contracts
6. Initialize op-geth
7. Run op-geth
8. Run op-node
9. Get some Holesky ETH on your L2
10. Send some test transactions
