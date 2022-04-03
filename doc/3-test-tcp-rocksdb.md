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
3.639070 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 7 ms
99% requests < 8 ms
100% requests < 9 ms
35 usec average for each request
throughput is 27.479552 MB/s
rps is 27479.551654
(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t get
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 1
63372 records get
36628 records miss
0 records set
5.248079 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
100% requests < 7 ms
50 usec average for each request
throughput is 12.075277 MB/s
rps is 19054.592818
```