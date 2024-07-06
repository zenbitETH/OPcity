### Execution layer conclusions

For geth:

```bash
./geth --datadir "geth-datadir" \
    --http --http.api="engine,eth,web3,net,debug" \
    --ws --ws.api="engine,eth,web3,net,debug" \
    --http.corsdomain "*" \
    --holesky \
    --syncmode=full \
    --authrpc.jwtsecret=/path/to/jwtSecret
```

<figure>
  <img src="./img/geth-EL-holesky.png" alt="State Prune Screenshot">
</figure>

After starting the client, you will be able to interact with the client using the following ports:

````
localhost:8545: Web3 json rpc
localhost:8550: Engine json rpc
localhost:8551: Engine json rpc with json rpc authentication
jwt.hex in the repository where you ran the rpcdaemon: verification key for JWT authentication```
````

> When running it with an L2 node on the same machine make sure you configure ports properly, because geth and op-geth use same ports

```bash
./geth  --datadir "geth-datadir" \
    --port=30309  \
    --http --http.port=10000 --http.api="engine,eth,web3,admin,net,debug"
    --ws --ws.api="engine,eth,web3,net,debug" \
    --http.corsdomain "*" \
    --holesky \
    --syncmode=full \
    --authrpc.jwtsecret=../jwt.hex 
```

#### Pruning the geth chaindata and dBs

```bash
./geth removedb
```

Screenshot:

<figure>
  <img src="./img/state-prune.png" alt="State Prune Screenshot">
</figure>

### Reference and Resources

1. [Etheruem Notes on Holesky](https://notes.ethereum.org/@launchpad/holesky)
2. [Backup and Restore](https://geth.ethereum.org/docs/fundamentals/backup-restore)
