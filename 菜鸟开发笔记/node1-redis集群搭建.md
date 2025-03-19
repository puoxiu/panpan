# 搭建redis集群
基于docker-compose搭建

### 基础知识：
* redis集群共有16384个槽位，默认平均分给每个实例
* 实例节点的失效：是通过集群中超过半数的节点检测失效时才生效
* 当需要在 Redis 集群中放置一个 key-value 时，根据公式HASH_SLOT=CRC16(key) mod 16384的值，决定将一个key放到哪个槽中，查询同理

步骤
1. 配置docker-compose.yaml文件、docker-compose-redis.yaml文件
2. 指定每个redis实例的配置文件，最好放到同目录下
3. docker-compose -f ./docker-compose-redis.yaml up -d
4. 如果没有配置redis的network 就是子网，需要查看每个实例的地址：
    docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-6380
    docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-6381
    docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis-6382

5. 启动工作集群：
    docker exec -it redis-6380 redis-cli --cluster create 172.20.0.2:6380 172.20.0.3:6381 172.20.0.4:6382 --cluster-replicas 0

6. 查看运行状态：
    docker exec -it redis-6380 redis-cli --cluster check 172.20.0.2:6380
    可以查看每个实例负责的槽位情况
