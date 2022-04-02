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
2.804455 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 8 ms
100% requests < 9 ms
26 usec average for each request
throughput is 35.657551 MB/s
rps is 35657.550998
(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t get
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 1
63109 records get
36891 records miss
0 records set
2.017625 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 4 ms
100% requests < 9 ms
19 usec average for each request
throughput is 31.278857 MB/s
rps is 49563.227330