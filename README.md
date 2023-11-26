# DEVTOOLS



```shell
# Activate all services
docker-compose up -d
```


```shell
# Activate some services on demand
docker-compose up -d etcd golang mysql redis
```

---

Taskfile.yml

```shell
wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/Taskfile.yml
```

---

docker-compose.yml

```shell
wget https://github.com/hzhacking/DEVTOOLS/blob/main/docker-compose.yml &&
mkcd deploy && mkcd dtm && wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/deploy/dtm/config.yml && cd - && 
mkcd prometheus && wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/deploy/prometheus/prometheus.yml && wget https://raw.githubusercontent.com/hzhacking/DEVTOOLS/main/deploy/prometheus/rules.yml && cd - &&
mkcd redis && wget https://github.com/hzhacking/DEVTOOLS/blob/main/deploy/redis/redis.conf && cd ...

```



## ref

- [nivin-studio/gonivinck: 一个基于docker的go-zero运行环境。](https://github.com/nivin-studio/gonivinck)