#!/bin/bash

geth --dev --dev.period 5 --http --http.addr localhost \
    --http.port 8545 \
    --http.api 'admin,debug,web3,eth,txpool,personal,clique,miner,net' \
    --verbosity 5 --rpc.gascap 50000000  --rpc.txfeecap 0 \
    --miner.gaslimit  10 --miner.gasprice 1 --gpo.blocks 1 \
    --gpo.percentile 1 --gpo.maxprice 10 --gpo.ignoreprice 2 \
    --dev.gaslimit 50000000
