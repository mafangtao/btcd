```cassandraql

(base) ➜  btcd btcd --configfile=./btcd.conf
2019-08-15 23:50:13.974 [INF] BTCD: Version 0.12.0-beta
2019-08-15 23:50:13.974 [INF] BTCD: Loading block database from '/home/fangtaoma/btcd/data/testnet/blocks_ffldb'
2019-08-15 23:50:13.977 [INF] BTCD: Block database loaded
2019-08-15 23:50:13.980 [INF] INDX: Transaction index enabled because it is required by the address index
2019-08-15 23:50:13.980 [INF] INDX: Address index is enabled
2019-08-15 23:50:13.980 [INF] INDX: Committed filter index is enabled
2019-08-15 23:50:13.981 [DBG] INDX: Current internal block ID: 0
2019-08-15 23:50:13.981 [DBG] INDX: Current transaction index tip (height -1, hash 0000000000000000000000000000000000000000000000000000000000000000)
2019-08-15 23:50:13.981 [DBG] INDX: Current address index tip (height -1, hash 0000000000000000000000000000000000000000000000000000000000000000)
2019-08-15 23:50:13.981 [DBG] INDX: Current committed filter index tip (height -1, hash 0000000000000000000000000000000000000000000000000000000000000000)
2019-08-15 23:50:13.981 [INF] INDX: Catching up indexes from height -1 to 0
2019-08-15 23:50:13.981 [INF] INDX: Indexes caught up to height 0
2019-08-15 23:50:13.981 [INF] CHAN: Chain state (height 0, hash 000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943, totaltx 1, work 4295032833)
2019-08-15 23:50:13.993 [INF] RPCS: RPC server listening on 127.0.0.1:18334
2019-08-15 23:50:13.993 [INF] AMGR: Loaded 0 addresses from file '/home/fangtaoma/btcd/data/testnet/peers.json'
2019-08-15 23:50:13.994 [INF] CMGR: Server listening on [::]:18333
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 8: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 9: no valid connect address
2019-08-15 23:50:13.995 [INF] CMGR: Server listening on 0.0.0.0:18333
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 10: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 11: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 12: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 13: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 14: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 15: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 16: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 17: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 18: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 19: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 20: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 4: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 1: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 2: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 3: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 6: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 5: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 7: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 25: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 26: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 27: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 28: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 29: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:13.995 [DBG] CMGR: Failed to connect to reqid 30: no valid connect address
2019-08-15 23:50:13.995 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:13.996 [DBG] CMGR: Failed to connect to reqid 21: no valid connect address
2019-08-15 23:50:13.996 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:13.996 [DBG] CMGR: Failed to connect to reqid 22: no valid connect address
2019-08-15 23:50:13.996 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:13.996 [DBG] CMGR: Failed to connect to reqid 23: no valid connect address
2019-08-15 23:50:13.996 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:13.996 [DBG] CMGR: Failed to connect to reqid 24: no valid connect address
2019-08-15 23:50:13.996 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:13.996 [DBG] CMGR: Failed to connect to reqid 31: no valid connect address
2019-08-15 23:50:13.996 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:13.996 [DBG] CMGR: Failed to connect to reqid 32: no valid connect address
2019-08-15 23:50:13.996 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:14.359 [INF] CMGR: 36 addresses found from DNS seed testnet-seed.bitcoin.jonasschnelli.ch
2019-08-15 23:50:14.398 [INF] CMGR: 5 addresses found from DNS seed testnet-seed.bluematt.me
2019-08-15 23:50:14.436 [INF] CMGR: 38 addresses found from DNS seed seed.tbtc.petertodd.org
2019-08-15 23:50:18.996 [DBG] CMGR: Attempting to connect to 217.172.178.44:18333 (reqid 40)
2019-08-15 23:50:18.996 [DBG] CMGR: Attempting to connect to 188.225.18.57:18333 (reqid 35)
2019-08-15 23:50:18.996 [DBG] CMGR: Attempting to connect to [2001:470:657e:4000:3810:b1ff:fe7d:486a]:18333 (reqid 33)
2019-08-15 23:50:18.996 [DBG] CMGR: Attempting to connect to 54.39.237.72:18333 (reqid 37)
2019-08-15 23:50:18.996 [DBG] CMGR: Failed to connect to [2001:470:657e:4000:3810:b1ff:fe7d:486a]:18333 (reqid 33): dial tcp [2001:470:657e:4000:3810:b1ff:fe7d:486a]:18333: connect: network is unreachable
2019-08-15 23:50:18.996 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:18.996 [DBG] CMGR: Attempting to connect to 182.254.156.155:18333 (reqid 36)
2019-08-15 23:50:18.996 [DBG] CMGR: Attempting to connect to 217.172.178.44:18333 (reqid 39)
2019-08-15 23:50:18.996 [DBG] CMGR: Attempting to connect to 188.225.18.57:18333 (reqid 34)
2019-08-15 23:50:18.996 [DBG] CMGR: Attempting to connect to [2002:4097:6518::4097:6518]:18333 (reqid 38)
2019-08-15 23:50:18.996 [DBG] CMGR: Failed to connect to [2002:4097:6518::4097:6518]:18333 (reqid 38): dial tcp [2002:4097:6518::4097:6518]:18333: connect: network is unreachable
2019-08-15 23:50:18.996 [DBG] CMGR: Max failed connection attempts reached: [25] -- retrying connection in: 5s
2019-08-15 23:50:19.022 [DBG] CMGR: Connected to 182.254.156.155:18333 (reqid 36)
2019-08-15 23:50:19.022 [DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.12.0/, pver 70013, block 0) to 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.035 [DBG] PEER: Received version (agent /Satoshi:0.18.0/, pver 70015, block 1574723) from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.035 [DBG] PEER: Negotiated protocol version 70013 for peer 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.035 [DBG] CHAN: Added time sample of 0s (total: 1)
2019-08-15 23:50:19.035 [DBG] PEER: Connected to 182.254.156.155:18333
2019-08-15 23:50:19.035 [INF] SYNC: New valid peer 182.254.156.155:18333 (outbound) (/Satoshi:0.18.0/)
2019-08-15 23:50:19.035 [INF] SYNC: Syncing to block height 1574723 from peer 182.254.156.155:18333
2019-08-15 23:50:19.035 [INF] SYNC: Downloading headers for blocks 1 to 546 from peer 182.254.156.155:18333
2019-08-15 23:50:19.035 [DBG] SRVR: New peer 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.035 [DBG] PEER: Received verack from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.035 [DBG] PEER: Sending getaddr to 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.035 [DBG] PEER: Sending verack to 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.035 [DBG] PEER: Sending getheaders (locator 000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943, stop 000000002a936ca763904c3c35fce2f3556c559c0214345d31b1bcebf76acb70) to 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.048 [DBG] PEER: Received sendheaders from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.048 [DBG] PEER: Received ping from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.048 [DBG] PEER: Sending pong to 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.048 [DBG] PEER: Received addr (1 addr) from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.048 [DBG] PEER: Received getheaders (locator 000000000000001643bb05adef13b198fbec85bbee6ef47cdf73abb8cd9d7591, stop 0000000000000000000000000000000000000000000000000000000000000000) from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.048 [DBG] PEER: Received feefilter from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.280 [DBG] CMGR: Connected to 217.172.178.44:18333 (reqid 40)
2019-08-15 23:50:19.280 [DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.12.0/, pver 70013, block 0) to 217.172.178.44:18333 (outbound)
2019-08-15 23:50:19.324 [DBG] CMGR: Connected to 188.225.18.57:18333 (reqid 34)
2019-08-15 23:50:19.324 [DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.12.0/, pver 70013, block 0) to 188.225.18.57:18333 (outbound)
2019-08-15 23:50:19.345 [DBG] CMGR: Connected to 188.225.18.57:18333 (reqid 35)
2019-08-15 23:50:19.345 [DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.12.0/, pver 70013, block 0) to 188.225.18.57:18333 (outbound)
2019-08-15 23:50:19.348 [DBG] CMGR: Connected to 217.172.178.44:18333 (reqid 39)
2019-08-15 23:50:19.348 [DBG] PEER: Sending version (agent /btcwire:0.5.0/btcd:0.12.0/, pver 70013, block 0) to 217.172.178.44:18333 (outbound)
2019-08-15 23:50:19.430 [DBG] PEER: Received headers (num 546) from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.431 [INF] SYNC: Verified downloaded block header against checkpoint at height 546/hash 000000002a936ca763904c3c35fce2f3556c559c0214345d31b1bcebf76acb70
2019-08-15 23:50:19.431 [INF] SYNC: Received 546 block headers: Fetching blocks
2019-08-15 23:50:19.434 [DBG] PEER: Sending getdata (size 546) to 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.466 [DBG] PEER: Received block (hash 00000000b873e79784647a6c82962c70d228557d24a747ea4d1b8bbe878e1206, ver 1, 1 tx, 2011-02-03 07:22:08 +0800 CST) from 182.254.156.155:18333 (outbound)
2019-08-15 23:50:19.466 [DBG] CHAN: Accepted block 00000000b873e79784647a6c82962c70d228557d24a747ea4d1b8bbe878e1206
2019-08-15 23:50:19.467 [DBG] PEER: Received block (hash 000000006c02c8ea6e4ff69651f7fcde348fb9d557a06e6957b65552002a7820, ver 1, 1 tx, 2011-02-03 07:22:26 +0800 CST) from 182.254.156.155:18333 (outbound
```