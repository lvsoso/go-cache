```shell
root@lv:/data# redis-benchmark -c 50 -n 100000 -d 1000 -t set -r 100000 -P 1
====== SET ======
  100000 requests completed in 0.73 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1

99.52% <= 1 milliseconds
99.83% <= 2 milliseconds
99.93% <= 3 milliseconds
99.95% <= 4 milliseconds
99.96% <= 5 milliseconds
99.97% <= 6 milliseconds
99.99% <= 7 milliseconds
100.00% <= 7 milliseconds
136425.66 requests per second

```