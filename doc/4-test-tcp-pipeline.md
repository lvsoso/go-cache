不使用 pipeline

```shell
(base) lv@lv:cache-benchmark$ cache-client -type tcp -n 100000 -r 100000 -t set
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
1.915990 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
100% requests < 7 ms
18 usec average for each request
throughput is 52.192341 MB/s
rps is 52192.341034
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
1.148432 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 4 ms
99% requests < 5 ms
100% requests < 11 ms
32 usec average for each request
throughput is 87.075277 MB/s
rps is 87075.277122

(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t get -P 3
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 3
63391 records get
36609 records miss
0 records set
1.221070 seconds total
99% requests < 1 ms
99% requests < 3 ms
99% requests < 6 ms
100% requests < 7 ms
33 usec average for each request
throughput is 51.914306 MB/s
rps is 81895.389113
```