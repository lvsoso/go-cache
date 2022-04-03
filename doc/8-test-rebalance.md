```shell
# 分别启动三个容器
docker  run --rm -itd --name n1 ubuntu:latest /bin/bash
docker  run --rm -itd --name n2 ubuntu:latest /bin/bash
docker  run --rm -itd --name n3 ubuntu:latest /bin/bash

docker exec -it n1 apt-get update
docker exec -it n2 apt-get update
docker exec -it n3 apt-get update

docker exec -it n1 apt-get install make g++ libz-dev libsnappy-dev -y
docker exec -it n2 apt-get install make g++ libz-dev libsnappy-dev -y
docker exec -it n3 apt-get install make g++ libz-dev libsnappy-dev -y

# 拷贝可执行程序
docker cp go-cache n1:/
docker cp go-cache n2:/
docker cp go-cache n3:/

docker cp client/client n3:/
docker cp cache-benchmark/cache-client n3:/

# 启动第一个节点
# ./go-cache --node 172.17.0.5

# 设置值
(base) lv@lv:go-cache$ docker exec -it n3 /cache-client -type tcp -n 10000 -d 1 -h 172.17.0.5
type is tcp
server is 172.17.0.5
total 10000 requests
data size is 1
we have 1 connections
operation is set
keyspacelen is 0
pipeline length is 1
0 records get
0 records miss
10000 records set
0.251928 seconds total
100% requests < 1 ms
24 usec average for each request
throughput is 0.039694 MB/s
rps is 39693.824070

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/status
{"Count":10000,"KeySize":38890,"ValueSize":48890}

# 启动第二个节点
# ./go-cache --node 172.17.0.6 --cluster 172.17.0.5

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/status
{"Count":10000,"KeySize":38890,"ValueSize":48890}(base) lv@lv:go-cache$ 

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.6:12345/status
{"Count":0,"KeySize":0,"ValueSize":0}(base) lv@lv:go-cache$ 

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/rebalance -XPOST

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/status
{"Count":5187,"KeySize":20178,"ValueSize":25365}(base)

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.6:12345/status
{"Count":4813,"KeySize":18712,"ValueSize":23525}(base) 

# ./go-cache --node 172.17.0.7 --cluster 172.17.0.6

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/rebalance -XPOST
(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.6:12345/rebalance -XPOST

(base) lv@lv:go-cache$  docker exec -it n3 curl 172.17.0.5:12345/status
{"Count":3084,"KeySize":12005,"ValueSize":15089}

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.6:12345/status
{"Count":3166,"KeySize":12301,"ValueSize":15467}
(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.7:12345/status
{"Count":3750,"KeySize":14584,"ValueSize":18334}
```

(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/status
{"Count":10000,"KeySize":38890,"ValueSize":48890}(base) lv@lv:go-cache$ 
(base) lv@lv:go-cache$ 
(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.6:12345/status
{"Count":0,"KeySize":0,"ValueSize":0}(base) lv@lv:go-cache$ 
(base) lv@lv:go-cache$ 
(base) lv@lv:go-cache$ 
(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/rebalance -XPOST
(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/status
{"Count":10000,"KeySize":38890,"ValueSize":48890}(base) lv@lv:go-cache$ 
(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.6:12345/status
{"Count":0,"KeySize":0,"ValueSize":0}(base) lv@lv:go-cache$ 
(base) lv@lv:go-cache$ 
