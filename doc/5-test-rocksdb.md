```shell
(base) lv@lv:rocksdb_performance$ ./test_batch_write -t 100000
total record number is 100000
value size is 1000
batch size is 1
100000 records batch put in 375255 usec, 3.75255 usec average, throughput is 266.485 MB/s, rps is 266485

(base) lv@lv:rocksdb_performance$ ./test_batch_write -t 100000 -b 100
total record number is 100000
value size is 1000
batch size is 100
100000 records batch put in 131724 usec, 1.31724 usec average, throughput is 759.163 MB/s, rps is 759163

```