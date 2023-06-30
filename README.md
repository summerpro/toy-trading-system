# toy-trading-system
This project accepts batch txs requests from external sources and then executes txs, theoretically capable of facing a large number of transaction requests, but lacking more testing.
## module
* `execution` module will execute transactions, complete user-to-user transfers, and update changes in user balances to the database.
* `txpool` module is used to cache received txs requests, ensure that transactions in the txs pool are arranged in order of fee income. The txs pool use the `heap` data structure, ensure txs in order, high performance of transaction insert, and high performance of obtaining topN txs
* `schedule` module is used to manage other modules to complete system functions. This module will specially start a goroutine to monitor external transaction requests. Once the request is received, it will verify it, and the verified txs will be sent to txspool. This module also has a goroutine to periodically process unprocessed txs from the txspool, sending these txs to the exection module for processing.
* `database` module is used to manage the data saved in the system, the main storage is the user's account data, through the interface to ensure future scalability, currently only memory database, but it can also be easily expanded to use leveldb this disk database.
* `types` module defines the data types used in the system.
* `config` module is used to manage the parameters used in the system.

## test
* `test` module is only used as a test to realize the most common situation in the system. There are a certain amount of concurrent txs requests on a regular basis, and then txs in the cache will be processed regularly. The running process of the system will be displayed in the log.
* For more test examples, see schedule_test in the schedule module.

## future work
* more test
* cmd
* network

