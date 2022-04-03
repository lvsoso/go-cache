### 说明

《分布式缓存-原理、架构及Go语言实现》学习。

#### 集群
高可用性集群（high availability，HA）、负载均衡集群（load balancing）、高性能计算集群（hight performance clusters ，HPC）以及网格计算（grid computing）；

缓存服务不涉及计算，所以缓存服务集群要求的通常是高可用性以及负载均衡；

#### CAP

- 一致性意味着每一次读取要么得到最新的结果，要么得到一个错误；
- 可用性意味着每一次成功的读取都可以得到一个结果，但不保证是最新的；
- 分区容错性意味着在集群节点无法互通的情况下依然能对外提供服务；

#### gossip

- 一对一节点通讯；
- 消息长度有限；
- 节点互动会导致某个节点的状态被更新到至少一个其他节点上；
- 不需要底层网络稳定；
- 通讯频率低，协议开销可忽略；
- 通讯对象的选择具有随机性，可以从全体节点中选出，也可以选自一个较小的相邻节点的集合；
- 消息会在节点间被复制，意味着传递的信息存在冗余；

####  cluster

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

docker cp go-cache n1:/
docker cp go-cache n2:/
docker cp go-cache n3:/

# 分别在容器中启动三个节点
# ./go-cache --node 172.17.0.5
# ./go-cache --node 172.17.0.6 --cluster 172.17.0.5
# ./go-cache --node 172.17.0.7 --cluster 172.17.0.6


# 查询节点信息
(base) lv@lv:go-cache$ docker exec -it n3 curl 172.17.0.5:12345/cluster
["172.17.0.5:7946","172.17.0.6:7946"]

docker cp client n3:/

(base) lv@lv:client$ docker exec -it n3  /client -h 172.17.0.5 -c set -k keya -v a
error: redirect 172.17.0.6

(base) lv@lv:client$ docker exec -it n3  /client -h 172.17.0.6 -c set -k keya -v a
a


```