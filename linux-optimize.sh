#!/bin/bash

# Redis 服务调优脚本

# 设置最大内存限制为2GB
redis-cli config set maxmemory 2gb

# 降低系统的swappiness值为10
# 如果绑定 NUMA 亲和性，可以优化性能。但是生产环境千万不要绑定，否则当占用内存超过当前 node 后，会直接 swap，而不是使用其他 node 的内存，掉入`NUMA 陷阱`，严重拖慢 redis 性能。
sudo sysctl vm.swappiness=10

# 禁用THP
sudo sh -c "echo never > /sys/kernel/mm/transparent_hugepage/enabled"

# 安装并启用 NTP 时间同步服务
sudo apt-get install ntp
sudo service ntp start

# 修改 ulimit 参数
# 注意：以下示例是修改当前终端窗口的 ulimit 值，不会永久生效
ulimit -n 65535

# 增加 TCP backlog 大小为511
sudo sysctl -w net.core.somaxconn=511

# 设置 redis 使用 swap 的时机
# 默认情况下，Linux 内核的 swappiness 值为 60。较高的值会导致内核更积极地使用交换空间，而较低的值会减少对交换空间的使用。
# 请注意，这只会修改当前的 swappiness 值，并不会永久生效。如果希望永久更改 swappiness 值，你需要修改 /etc/sysctl.conf 文件，找到 vm.swappiness 行并将其更新为所需的值
sudo sysctl vm.swappiness=<value>
# 让修改生效
sudo sysctl -p

# TCP backlog
# 要配置 Linux 的 TCP backlog 大小，可以通过修改 net.core.somaxconn 参数来实现。TCP backlog 是指 TCP 服务器监听时，允许处于等待连接状态的客户端的数量。较大的 backlog 数值可以提高服务器的并发连接能力。
sudo sysctl -w net.core.somaxconn=<value>
sudo sysctl -p


# 修改io调度算法
# 默认为CFQ, 可以修改为deadline或者noop
echo <algorithm> > /sys/block/<device>/queue/scheduler