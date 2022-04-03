不使用pipeline set

```shell
root@lv:/data# redis-benchmark -c 1 -n 100000 -d 1000 -t set,get  -r 100000     
====== SET ======
  100000 requests completed in 1.88 seconds
  1 parallel clients
  1000 bytes payload
  keep alive: 1

99.99% <= 1 milliseconds
100.00% <= 2 milliseconds
100.00% <= 3 milliseconds
100.00% <= 6 milliseconds
53248.14 requests per second

====== GET ======
  100000 requests completed in 1.79 seconds
  1 parallel clients
  1000 bytes payload
  keep alive: 1

100.00% <= 1 milliseconds
100.00% <= 2 milliseconds
55959.71 requests per second

```

使用 pipeline set

```shell
root@lv:/data# redis-benchmark -c 1 -n 100000 -d 1000 -t set,get -r 100000 -P 3
====== SET ======
  100000 requests completed in 1.05 seconds
  1 parallel clients
  1000 bytes payload
  keep alive: 1

100.00% <= 1 milliseconds
100.00% <= 1 milliseconds
95147.48 requests per second

====== GET ======
  100000 requests completed in 0.74 seconds
  1 parallel clients
  1000 bytes payload
  keep alive: 1

100.00% <= 0 milliseconds
134770.89 requests per second

```