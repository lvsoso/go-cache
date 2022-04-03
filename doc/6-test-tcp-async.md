

```shell

(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t get -P 3
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 3
99314 records get
686 records miss
0 records set
2.359713 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 7 ms
100% requests < 10 ms
67 usec average for each request
throughput is 42.087329 MB/s
rps is 42378.041906


(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t get -P 10
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 10
99357 records get
643 records miss
0 records set
1.592218 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 7 ms
99% requests < 8 ms
99% requests < 9 ms
100% requests < 12 ms
151 usec average for each request
throughput is 62.401636 MB/s
rps is 62805.475194

(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t get -P 100
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 100
99308 records get
692 records miss
0 records set
1.008744 seconds total
86% requests < 1 ms
97% requests < 2 ms
98% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 7 ms
99% requests < 8 ms
99% requests < 9 ms
99% requests < 10 ms
100% requests < 12 ms
950 usec average for each request
throughput is 98.447191 MB/s
rps is 99133.192942
```