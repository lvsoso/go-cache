```shell
(base) lv@lv:cache-benchmark$ ./cache-client -h 127.0.0.1 -type http -n 100000 -t set
type is http
server is 127.0.0.1
total 100000 requests
data size is 1000
we have 1 connections
operation is set
keyspacelen is 0
pipeline length is 1
0 records get
0 records miss
100000 records set
5.677702 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
100% requests < 9 ms
55 usec average for each request
throughput is 17.612760 MB/s
rps is 17612.759674
(base) lv@lv:cache-benchmark$ 
(base) lv@lv:cache-benchmark$ 
(base) lv@lv:cache-benchmark$ ./cache-client -h 127.0.0.1 -type redis -n 100000 -t set
type is redis
server is 127.0.0.1
total 100000 requests
data size is 1000
we have 1 connections
operation is set
keyspacelen is 0
pipeline length is 1
0 records get
0 records miss
100000 records set
2.780447 seconds total
99% requests < 1 ms
99% requests < 2 ms
100% requests < 3 ms
26 usec average for each request
throughput is 35.965442 MB/s
rps is 35965.441590

```
