# DEVTOOLS



```shell
# Activate all services
docker-compose up -d
```


```shell
# Activate some services on demand
docker-compose up -d etcd golang mysql redis
```



```shell
# Taskfile.yml
wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/Taskfile.yml &&
# modd.conf
wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/modd.conf &&
# docker-compose.yml
wget https://github.com/hzhacking/DEVTOOLS/blob/main/docker-compose.yml &&
mkcd deploy && mkcd dtm && wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/deploy/dtm/config.yml && cd - && 
mkcd prometheus && wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/deploy/prometheus/prometheus.yml && wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/deploy/prometheus/rules.yml && cd - &&
mkcd filebeat && wget https://github.com/hzhacking/DEVTOOLS/blob/main/deploy/filebeat/filebeat.yml && cd - &&
mkcd go-stash && wget https://github.com/hzhacking/DEVTOOLS/blob/main/deploy/go-stash/config.yml && cd - &&
# nginx
mkcd nginx && wget https://github.com/hzhacking/DEVTOOLS/blob/main/deploy/nginx/nginx.conf && cd - &&
mkcd redis && wget https://github.com/hzhacking/DEVTOOLS/blob/main/deploy/redis/redis.conf && cd ...

```



## ref

- [nivin-studio/gonivinck: 一个基于docker的go-zero运行环境。](https://github.com/nivin-studio/gonivinck)