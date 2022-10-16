# 代码运行测试
```
export VERSION=v0.0.3
go run main.go
```
```
curl -v localhost:8080/healthz -H "ttt:yyy"
```
# 镜像测试
```
docker build -t cncamp/httpserver:v0.0.3 .
docker run -d --name httpserver -p 9090:8080 -e VERSION=v0.0.3 cncamp/httpserver:v0.0.3
docker logs -f httpserver
```
```
curl -v localhost:9090/healthz -H "ttt:yyy"
```

# 镜像比较
```
docker images|grep httpserver
cncamp/httpserver   v0.0.3        e6ccaa261248   29 seconds ago   15.5MB
cncamp/httpserver   v0.0.2        1fc9919e6afa   4 minutes ago    9.93MB
cncamp/httpserver   v0.0.1        22fc44df77a6   10 days ago      558MB
```