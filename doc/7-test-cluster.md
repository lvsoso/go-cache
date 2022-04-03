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