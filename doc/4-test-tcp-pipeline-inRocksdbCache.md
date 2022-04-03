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
4.154876 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 7 ms
99% requests < 8 ms
99% requests < 9 ms
100% requests < 12 ms
40 usec average for each request
throughput is 24.068106 MB/s
rps is 24068.105866
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
2.973801 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 7 ms
99% requests < 8 ms
99% requests < 9 ms
99% requests < 10 ms
100% requests < 11 ms
85 usec average for each request
throughput is 33.626993 MB/s
rps is 33626.993417

(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t get -P 3
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 3
86546 records get
13454 records miss
0 records set
2.477584 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 6 ms
100% requests < 7 ms
71 usec average for each request
throughput is 34.931612 MB/s
rps is 40361.901828
```