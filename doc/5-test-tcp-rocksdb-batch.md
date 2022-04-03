不使用 pipeline
```shell
(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t set
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is set
keyspacelen is 100000
pipeline length is 1
0 records get
0 records miss
100000 records set
2.466963 seconds total
99% requests < 1 ms
99% requests < 2 ms
100% requests < 6 ms
23 usec average for each request
throughput is 40.535669 MB/s
rps is 40535.668728
```
使用 pipeline
```shell
(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t set -P 3
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is set
keyspacelen is 100000
pipeline length is 3
0 records get
0 records miss
100000 records set
1.423694 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
100% requests < 8 ms
39 usec average for each request
throughput is 70.239804 MB/s
rps is 70239.804431

```

使用 50 个链接

```shell
(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t set -c 50
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 50 connections
operation is set
keyspacelen is 100000
pipeline length is 1
0 records get
0 records miss
100000 records set
0.794791 seconds total
94% requests < 1 ms
97% requests < 2 ms
98% requests < 3 ms
98% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 7 ms
99% requests < 8 ms
99% requests < 9 ms
99% requests < 10 ms
99% requests < 11 ms
99% requests < 12 ms
99% requests < 13 ms
99% requests < 15 ms
99% requests < 16 ms
99% requests < 18 ms
99% requests < 19 ms
99% requests < 20 ms
99% requests < 21 ms
99% requests < 22 ms
99% requests < 24 ms
100% requests < 26 ms
364 usec average for each request
throughput is 125.819231 MB/s
rps is 125819.231348
```