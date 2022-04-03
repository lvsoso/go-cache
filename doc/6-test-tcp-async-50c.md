```shell
(base) lv@lv:cache-benchmark$ ./cache-client -type tcp -n 100000 -r 100000 -t get -c 50
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 50 connections
operation is get
keyspacelen is 100000
pipeline length is 1
99289 records get
711 records miss
0 records set
0.662172 seconds total
96% requests < 1 ms
98% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
99% requests < 5 ms
99% requests < 6 ms
99% requests < 7 ms
99% requests < 8 ms
99% requests < 9 ms
99% requests < 10 ms
99% requests < 14 ms
99% requests < 22 ms
100% requests < 23 ms
323 usec average for each request
throughput is 149.944325 MB/s
rps is 151018.063432
```