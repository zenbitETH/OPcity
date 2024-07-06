## Know before you go

Before I kick off, note that this is a relatively long tutorial! You should prepare to set aside an hour or two to get everything running. Here’s an itemized list of what we’re about to do:

1. Install dependencies
2. Build the source code
3. Generate and fund accounts and private keys
4. Configure your network
5. Deploy the L1 contracts
6. Initialize op-geth
7. Run op-geth
8. Run op-node
9. Get some Goerli ETH on your L2
10. Send some test transactions
11. Celebrate!

## Prerequisites

You’ll need the following software installed to follow this tutorial:

- [Git](https://git-scm.com/)
- [Go](https://go.dev/)
- [Node](https://nodejs.org/en/)
- [Pnpm](https://classic.yarnpkg.com/lang/en/docs/install/)
- [Foundry](https://github.com/foundry-rs/foundry#installation)
- [Make](https://linux.die.net/man/1/make)
- [jq](https://github.com/jqlang/jq)
- [direnv](https://direnv.net)

This tutorial was checked on:

| Software | Version    | Installation command(s) |
| -------- | ---------- | - |
| Ubuntu   | 20.04 LTS  | |
| git, curl, jq, and make | OS default | `sudo apt update` <br> `sudo apt install -y git curl make jq direnv` |
| Go       | 1.20       | `wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz` <br> `tar xvzf go1.22.4.linux-amd64.tar.gz` <br> `sudo cp go/bin/go /usr/bin/go` <br> `sudo mv go /usr/lib` <br> `echo export GOROOT=/usr/lib/go >> ~/.bashrc`
| Node     | 16.19.0    | `curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh \| bash` <br> `export NVM_DIR="$HOME/.nvm"` <br> `[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm` <br> `[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion` <br> `nvm install 22`
| pnpm     | 8.5.6      | `npm install -g pnpm`
| Foundry  | 0.2.0      | `curl -L https://foundry.paradigm.xyz \| bash` <br> `source /root/.bashrc` <br> `foundryup`


## Build the Source Code

We’re going to be spinning up an EVM Rollup from the OP Stack source code.  We can use docker images as well, but this way we can modify component behavior if we need to do so. The OP Stack source code is split between two repositories, the [Optimism Monorepo](https://github.com/ethereum-optimism/optimism) and the [`op-geth`](https://github.com/ethereum-optimism/op-geth) repository.

### Build the Optimism Monorepo

1. Clone the [Optimism Monorepo](https://github.com/ethereum-optimism/optimism).

    ```bash
    cd ~
    git clone https://github.com/ethereum-optimism/optimism.git
    ```

2. Enter the Optimism Monorepo.

    ```bash
    cd optimism
    ```

3. Install required modules. This is a slow process, while it is running you can already start building `op-geth`, as shown below.

    ```bash
    pnpm install
    ```

4. Build the various packages inside of the Optimism Monorepo.

    ```bash
    make op-node op-batcher op-proposer
    pnpm build
    ```

    Please note that this a considerably longer step (about 15 mins) and §solc can take up 

### Build op-geth

1. Clone [`op-geth`](https://github.com/ethereum-optimism/op-geth):

    ```bash
    cd ~
    git clone https://github.com/ethereum-optimism/op-geth.git
    ```

1. Enter `op-geth`:

    ```bash
    cd op-geth
    ```

1. Build `op-geth`:

    ```bash
    make geth
    ```

## Get access to a Holesky node

Since we’re deploying our OP Stack chain to Goerli, we’ll need to have access to a Goerli L1 node. We can either use a node provider like [Alchemy](https://www.alchemy.com/) (easier) or [run your own Holesky node](https://notes.ethereum.org/@launchpad/holesky) (harder).

## Generate some keys

We’ll need four accounts and their private keys when setting up the chain:

- The `Admin` account which has the ability to upgrade contracts.
- The `Batcher` account which publishes Sequencer transaction data to L1.
- The `Proposer` account which publishes L2 transaction results to L1.
- The `Sequencer` account which signs blocks on the p2p network.

You can generate all of these keys with the `rekey` tool in the `contracts-bedrock` package.


1. Enter the Optimism Monorepo:

    ```bash
    cd ~
    cd optimism
    ```

1. Move into the `contracts-bedrock` package:

    ```bash
    cd packages/contracts-bedrock
    ```

1. Use `cast wallet` to generate new accounts

    ```bash
    echo "Admin:"
    cast wallet new
    echo "Proposer:"
    cast wallet new
    echo "Batcher:"
    cast wallet new
    echo "Sequencer:"
    cast wallet new
    ```

We get an output like the following:

```Admin:
Successfully created new keypair.
Address:     0xC4481aa21AeAcAD3cCFe6252c6fe2f161A47A771
Private key: 0x3dd6681bda6458773e9e8aa6b45574b73a953eaea67ff6d9d00795da73e39177
Proposer:
Successfully created new keypair.
Address:     0x0F10759D26F07a5D9e7F4BC0b54d1e4CAAEd15C9
Private key: 0x723fcb35ce979dea9f269c26d0cd685cfd45cd90681900c1e69fa9c9ba0479d8
Batcher:
Successfully created new keypair.
Address:     0xab110dA2064AC0B44c08D71A3D8148BBB0C3aD1F
Private key: 0x2e9c494f5a6c0325ea5bd2265132f73dd040a6d75bd4303866c10ff1ec51b356
Sequencer:
Successfully created new keypair.
Address:     0x50987039A15c83C4090eD5ecfda9E7F07160D4a0
Private key: 0x2c65361c17805b0b8da483cbc7fd462608f1c3fd73f4ac5e809082a2c52cf879```

I'll save these accounts and their respective private keys somewhere for now. Fund the `Admin` address with a small amount of Holesky ETH as we’ll use that account to deploy our smart contracts. You’ll also need to fund the `Proposer` and `Batcher` address — note that the `Batcher` burns through the most ETH because it publishes transaction data to L1.

Recommended funding amounts are as follows:

- `Admin` — 2 ETH
- `Proposer` — 5 ETH
- `Batcher` — 10 ETH

::: danger Not for production deployments

The `cast wallet new` tool is *not* designed for production deployments. If we are deploying an OP Stack based chain into production, we should likely be using a combination of hardware security modules and hardware wallets.

:::

## Configure your network

Once you’ve built both repositories, you’ll need head back to the Optimism Monorepo to set up the configuration for your chain. Currently, chain configuration lives inside of the [`contracts-bedrock`](https://github.com/ethereum-optimism/optimism/tree/129032f15b76b0d2a940443a39433de931a97a44/packages/contracts-bedrock) package.

1. Enter the Optimism Monorepo:

   ```bash
   cd ~/optimism
   ```

2. Inside of `optimism` repo, copy the environment file into the `.envrc` path

   ```sh
   cp .envrc.example .envrc
   ```

_Note: If you think there any env variable missing just add them inside the .envrc file_

3. Fill out the environment variables inside of that file:

    - `L1_RPC_URL` or `ETH_RPC_URL` — URL for your L1 node.
    - `PRIVATE_KEY` — Private key of the `Admin` account.
    - `DEPLOYMENT_CONTEXT` - Name of the network, should be anything like "lumino-test"

4. Pull the environment variables into context using `direnv`

    ```bash
    direnv allow .
    ```

If you need to install `direnv`, [make sure you also modify the shell configuration](https://direnv.net/docs/hook.html).

5. Before we can create our configuration file, we’ll need to pick an L1 block to serve as the starting point for our Rollup. It’s best to use a finalized L1 block as our starting block. You can use the `cast` command provided by Foundry to grab all of the necessary information:

    ```bash
    cast block finalized --rpc-url $ETH_RPC_URL | grep -E "(timestamp|hash|number)"
    ```

    You’ll get back something that looks like the following:

    ```
    hash                 0xd3091ed8a750029c1071693b8974ec1156bf5685d92adf5620ed857821869e02
    number               1717366
    timestamp            1718142912
    ```

1. Fill out the remainder of the `pre-populated` .envrc variable. Use the default values in the config file and make following modifications:

- Populate variable with their respective addresses and private keys that we generated earlier.
- Replace `BLOCKHASH` with the blockhash you got from the `cast` command.
- Replace `TIMESTAMP` with the timestamp you got from the `cast` command. Note that although all the other fields are strings, this field is a number! Don’t include the quotation marks.

## Generate the configuration file

Once you've built both repositories, you'll need to head back to the Optimism Monorepo to set up the configuration file for your chain. Currently, chain configuration lives inside of the [contracts-bedrock](https://github.com/ethereum-optimism/optimism/tree/v1.1.4/packages/contracts-bedrock) package in the form of a JSON file.

1. Enter the Optimism Monorepo

```bash
cd ~/optimism
```

2. Move into the contracts-bedrock package

```bash
cd packages/contracts-bedrock
```

3. Generate the configuration file
Run the following script to generate the getting-started.json configuration file inside of the deploy-config directory.

_Note: change the name of the output at the bottom/last line of the script if you want a custom named config file_

```bash
./scripts/getting-started/config.sh
```

After running the script correctly we should see a file that is named something like <custom-name>.json

4. Add this line to your .envrc

```bash
export DEPLOY_CONFIG_PATH=./deploy-config/<custom-name>.json
```

## Deploy the Create2 Factory (Optional)

If you're deploying an OP Stack chain to a network other than Sepolia/Holesky, you may need to deploy a Create2 factory contract to the L1 chain. This factory contract is used to deploy OP Stack smart contracts in a deterministic fashion.

_This step is typically only necessary if you are deploying your OP Stack chain to custom L1 chain. If you are deploying your OP Stack chain to Sepolia, you can safely skip this step._

#### 1. Check if the factory exists
The Create2 factory contract will be deployed at the address `0x4e59b44847b379578588920cA78FbF26c0B4956C`. You can check if this contract has already been deployed to your L1 network with a block explorer or by running the following command:

```bash
cast codesize 0x4e59b44847b379578588920cA78FbF26c0B4956C --rpc-url $L1_RPC_URL
```

If the command returns `0` then the contract has not been deployed yet. If the command returns `69` then the contract has been deployed and you can safely skip this section.

#### 2. Fund the factory deployer
You will need to send some ETH to the address that will be used to deploy the factory contract, `0x3fAB184622Dc19b6109349B94811493BF2a45362`. This address can only be used to deploy the factory contract and will not be used for anything else. Send at least 1 ETH to this address on your L1 chain.

#### 3. Deploy the factory
Using cast, deploy the factory contract to your L1 chain:

```bash
cast publish --rpc-url $L1_RPC_URL 0xf8a58085174876e800830186a08080b853604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf31ba02222222222222222222222222222222222222222222222222222222222222222a02222222222222222222222222222222222222222222222222222222222222222
``` 

Wait for the transaction to be mined
Make sure that the transaction is included in a block on your L1 chain before continuing.

#### 4. Verify that the factory was deployed

Run the code size check again to make sure that the factory was properly deployed:

```bash
cast codesize 0x4e59b44847b379578588920cA78FbF26c0B4956C --rpc-url $L1_RPC_URL
```


## Deploy the L1 contracts

Once you’ve configured your network, it’s time to deploy the L1 smart contracts necessary for the functionality of the chain.

1. Create a `getting-started` deployment directory.

   ```bash
   mkdir deployments/lumino-test
   ```


2. Once you’re ready, deploy the L1 smart contracts.

    ```bash
    forge script scripts/Deploy.s.sol:Deploy --private-key $GS_ADMIN_PRIVATE_KEY --broadcast --rpc-url $L1_RPC_URL --slow
    ```

_If you see a nondescript error that includes `EvmError: Revert` and `Script failed` then you likely need to change the `IMPL_SALT` environment variable. This variable determines the addresses of various smart contracts that are deployed via [CREATE2](https://eips.ethereum.org/EIPS/eip-1014). If the same `IMPL_SALT` is used to deploy the same contracts twice, the second deployment will fail. You can generate a new `IMPL_SALT` by running direnv allow anywhere in the Optimism Monorepo._

<!-- 3. Generate contract artifacts:
    
    ```bash
    forge script scripts/Deploy.s.sol:Deploy --sig 'sync()' --rpc-url $L1_RPC_URL
    ``` -->

Contract deployment can take up to 15-30 minutes. Please wait for all smart contracts to be fully deployed before continuing to the next step.

## Generate the L2 config files

We’ve set up the L1 side of things, but now we need to set up the L2 side of things. We do this by generating three important files, a `genesis.json` file, a `rollup.json` configuration file, and a `jwt.txt` [JSON Web Token](https://jwt.io/introduction) that allows the `op-node` and `op-geth` to communicate securely.

1. Generate `state-dump` and `l2-allocs`

```bash
cd ~/optimism/packages/contracts-bedrock

CONTRACT_ADDRESSES_PATH=./deployments/17000-deploy.json DEPLOY_CONFIG_PATH=./deploy-config/lumino-test.json STATE_DUMP_PATH=./deployments/l2-allocs.json  forge script scripts/L2Genesis.s.sol:L2Genesis  --sig 'runWithStateDump()'
```

2. Head over to the `op-node` package.

    ```bash
    cd ~/optimism/op-node
    ```

1. Run the following command, and make sure to replace `<RPC>` with your L1 RPC URL:

    ```bash
    go run cmd/main.go genesis l2 
    --l1-rpc <l1-rpc-value>
    --deploy-config ../packages/contracts-bedrock/deploy-config/lumino-test.json 
    --l2-allocs ../packages/contracts-bedrock/deployments/l2-allocs.json 
    --l1-deployments ../packages/contracts-bedrock/deployments/17000-deploy.json 
    --outfile.l2 genesis.json  
    --outfile.rollup rollup.json
    ```

    You should then see the `genesis.json` and `rollup.json` files inside the `op-node` package.

 4. Next, generate the `jwt.txt` file with the following command:

    ```bash
    openssl rand -hex 32 > jwt.txt
    ```

5. Finally, we’ll need to copy the `genesis.json` file and `jwt.txt` file into `op-geth` so we can use it to initialize and run `op-geth` in just a minute:

    ```bash
    cp genesis.json ~/op-geth
    cp jwt.txt ~/op-geth
    ```

    ## Initialize op-geth

We’re almost ready to run our chain! Now we just need to run a few commands to initialize `op-geth`. We’re going to be running a Sequencer node, so we’ll need to import the `Sequencer` private key that we generated earlier. This private key is what our Sequencer will use to sign new blocks.

1. Head over to the `op-geth` repository:

    ```bash
    cd ~/op-geth
    ```

2. Create a data directory folder:

    ```bash
    mkdir datadir
    ```

1. Next we need to initialize `op-geth` with the genesis file we generated and copied earlier:

    ```bash
    build/bin/geth init --datadir=datadir genesis.json
    ```

Everything is now initialized and ready to go!

## Run the node software

There are four components that need to run for a rollup.
The first two, `op-geth` and `op-node`, have to run on every node.
The other two, `op-batcher` and `op-proposer`, run only in one place, the sequencer that accepts transactions.

Set these environment variables for the configuration

| Variable       | Value |
| -------------- | -
| `SEQ_KEY`      | Private key of the `Sequencer` account
| `BATCHER_KEY`  | Private key of the `Batcher` accounts, which should have at least 1 ETH
| `PROPOSER_KEY` | Private key of the `Proposer` account
| `L1_RPC`       | URL for the L1 (such as Goerli) you're using
| `RPC_KIND`     | The type of L1 server to which you connect, which can optimize requests. Available options are `alchemy`, `quicknode`, `parity`, `nethermind`, `debug_geth`, `erigon`, `basic`, and `any`
| `L2OO_ADDR`    | The address of the `L2OutputOracleProxy`, available at `~/optimism/packages/contracts-bedrock/deployments/getting-started/L2OutputOracleProxy.json`

## Start `op-geth`

#### 1. Run `op-geth` with the following commands.

_You're using `--gcmode=archive` to run `op-geth` here because this node will act as your Sequencer. It's useful to run the Sequencer in archive mode because the `op-proposer` requires access to the full state. Feel free to run other (non-Sequencer) nodes in full mode if you'd like to save disk space._

It's important that you've already initialized the geth node at this point as per the previous section. Failure to do this will cause startup issues between `op-geth` and `op-node`.


```bash
cd ~/op-geth

./build/bin/geth \
  --datadir ./datadir \
  --http \
  --http.corsdomain="*" \
  --http.vhosts="*" \
  --http.addr=0.0.0.0 \
  --http.api=web3,debug,eth,txpool,net,engine \
  --ws \
  --ws.addr=0.0.0.0 \
  --ws.port=8546 \
  --ws.origins="*" \
  --ws.api=debug,eth,txpool,net,engine \
  --syncmode=full \
  --gcmode=archive \
  --nodiscover \
  --maxpeers=0 \
  --networkid=42069 \
  --authrpc.vhosts="*" \
  --authrpc.addr=0.0.0.0 \
  --authrpc.port=8551 \
  --authrpc.jwtsecret=./jwt.txt \
  --rollup.disabletxpoolgossip=true
```

And `op-geth` should be running! You should see some output, but you won’t see any blocks being created yet because `op-geth` is driven by the `op-node`. We’ll need to get that running next.

## Start op-node

Once you've got op-geth running you'll need to run op-node. Like Ethereum, the OP Stack has a Consensus Client (op-node) and an Execution Client (op-geth). The Consensus Client "drives" the Execution Client over the Engine API.

#### 1. Open up a new terminal

You'll need a terminal window to run the op-node in. Navigate to the op-node directory

```bash
cd ~/optimism/op-node
```

#### 2. Run op-node
```bash
    ./bin/op-node
    --l2=http://localhost:8551
    --l2.jwt-secret=./jwt.txt --sequencer.enabled
    --sequencer.l1-confs=5
    --verifier.l1-confs=4
    --rollup.config=./rollup.json  --rpc.addr=0.0.0.0 
    --rpc.port=8547 
    --p2p.disable  
    --rpc.enable-admin  
    --p2p.sequencer.key=$GS_SEQUENCER_PRIVATE_KEY 
    --l1=$L1_RPC_URL  
    --l1.rpckind=$L1_RPC_KIND
    --l1.trustrpc=true
```

Once you run this command, you should start seeing the op-node begin to sync L2 blocks from the L1 chain. Once the op-node has caught up to the tip of the L1 chain, it'll begin to send blocks to op-geth for execution. At that point, you'll start to see blocks being created inside of op-geth


## Start op-batcher

The op-batcher takes transactions from the Sequencer and publishes those transactions to L1. Once these Sequencer transactions are included in a finalized L1 block, they're officially part of the canonical chain. The op-batcher is critical!

It's best to give the Batcher address at least 1 ETH in native to ensure that it can continue operating without running out of ETH for gas. Keep an eye on the balance of the Batcher address because it can expend ETH quickly if there are a lot of transactions to publish.

#### 1. Open up a new terminal

You'll need a terminal window to run the op-batcher in. Navigate to the op-batcher directory 

```bash
cd ~/optimism/op-batcher
```

#### 2. Run op-batcher
```bash
./bin/op-batcher \
  --l2-eth-rpc=http://localhost:8545 \
  --rollup-rpc=http://localhost:8547 \
  --poll-interval=1s \
  --sub-safety-margin=6 \
  --num-confirmations=1 \
  --safe-abort-nonce-too-low-count=3 \
  --resubmission-timeout=30s \
  --rpc.addr=0.0.0.0 \
  --rpc.port=8548 \
  --rpc.enable-admin \
  --max-channel-duration=4 \
  --l1-eth-rpc=$L1_RPC_URL \
  --private-key=$GS_BATCHER_PRIVATE_KEY
```

_The `--max-channel-duration=n` setting tells the batcher to write all the data to L1 every `n` L1 blocks. When it is low, transactions are written to L1 frequently and other nodes can synchronize from L1 quickly. When it is high, transactions are written to L1 less frequently and the batcher spends less ETH. If you want to reduce costs, either set this value to 0 to disable it or increase it to a higher value._


## Start op-proposer

Now start op-proposer, which proposes new state roots.

#### 1. Open up a new terminal

You'll need a terminal window to run the op-proposer in. Navigate to the op-proposer directory

```bash
cd ~/optimism/op-proposer
```

#### 2. Run op-proposer
```bash
./bin/op-proposer \  --poll-interval=12s --rpc.port=8560  --rollup-rpc=http://localhost:8547  --l2oo-address=<L2OutputOracleProxyAddress>   --private-key=$GS_PROPOSER_PRIVATE_KEY  --l1-eth-rpc=$L1_RPC_URL
```

you'll be able to get L2OutputOracleProxyAddress from the path `/packages/contracts-bedrock/deployments/<l1-chain-id>-deploy.json`

In our the L1 chainId is `17000` which belongs to Holesky

## Run Transaction Explorer (Blockscout)

Runs Blockscout locally in Docker containers with [docker-compose](https://github.com/docker/compose).

## Prerequisites

- Docker v20.10+
- Docker-compose 2.x.x+
- Running Ethereum JSON RPC client

## Building Docker containers from source

```bash
cd ./docker-compose
docker-compose up --build
```

**Note**: if you don't need to make backend customizations, you can run `docker-compose up` in order to launch from pre-build backend Docker image. This will be much faster.

This command uses `docker-compose.yml` by-default, which builds the backend of the explorer into the Docker image and runs 9 Docker containers:

- Postgres 14.x database, which will be available at port 7432 on the host machine.
- Redis database of the latest version.
- Blockscout backend with api at /api path.
- Nginx proxy to bind backend, frontend and microservices.
- Blockscout explorer at http://localhost.

and 5 containers for microservices (written in Rust):

- [Stats](https://github.com/blockscout/blockscout-rs/tree/main/stats) service with a separate Postgres 14 DB.
- [Sol2UML visualizer](https://github.com/blockscout/blockscout-rs/tree/main/visualizer) service.
- [Sig-provider](https://github.com/blockscout/blockscout-rs/tree/main/sig-provider) service.
- [User-ops-indexer](https://github.com/blockscout/blockscout-rs/tree/main/user-ops-indexer) service.

**Note for Linux users**: Linux users need to run the local node on http://0.0.0.0/ rather than http://127.0.0.1/

## Configs for different Ethereum clients

The repo contains built-in configs for different JSON RPC clients without need to build the image.

**Note**: in all below examples, you can use `docker compose` instead of `docker-compose`, if compose v2 plugin is installed in Docker.

| __JSON RPC Client__    | __Docker compose launch command__ |
| -------- | ------- |
| Erigon  | `docker-compose -f erigon.yml up -d`    |
| Geth (suitable for Reth as well) | `docker-compose -f geth.yml up -d`     |
| Geth Clique    | `docker-compose -f geth-clique-consensus.yml up -d`    |
| Nethermind, OpenEthereum    | `docker-compose -f nethermind up -d`    |
| Ganache    | `docker-compose -f ganache.yml up -d`    |
| HardHat network    | `docker-compose -f hardhat-network.yml up -d`    |

- Running only explorer without DB: `docker-compose -f external-db.yml up -d`. In this case, no db container is created. And it assumes that the DB credentials are provided through `DATABASE_URL` environment variable on the backend container.
- Running explorer with external backend: `docker-compose -f external-backend.yml up -d`
- Running explorer with external frontend: `docker-compose -f external-frontend.yml up -d`
- Running all microservices: `docker-compose -f microservices.yml up -d`

All of the configs assume the Ethereum JSON RPC is running at http://localhost:8545.

In order to stop launched containers, run `docker-compose -d -f config_file.yml down`, replacing `config_file.yml` with the file name of the config which was previously launched.

You can adjust BlockScout environment variables:

- for backend in `./envs/common-blockscout.env`
- for frontend in `./envs/common-frontend.env`
- for stats service in `./envs/common-stats.env`
- for visualizer in `./envs/common-visualizer.env`
- for user-ops-indexer in `./envs/common-user-ops-indexer.env`

Descriptions of the ENVs are available

- for [backend](https://docs.blockscout.com/for-developers/information-and-settings/env-variables)
- for [frontend](https://github.com/blockscout/frontend/blob/main/docs/ENVS.md).

## Running Docker containers via Makefile

Prerequisites are the same, as for docker-compose setup.

Start all containers:

```bash
cd ./docker
make start
```

Stop all containers:

```bash
cd ./docker
make stop
```

***Note***: Makefile uses the same .env files since it is running docker-compose services inside.