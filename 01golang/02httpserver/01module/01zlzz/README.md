# 代码运行测试
```
export VERSION=v0.0.0
go run main.go
```
```
curl -v localhost:8080/healthz -H "ttt:yyy"
```
# 镜像测试
```
docker build -t cncamp/httpserver:v0.0.1 .
docker run -d --name httpserver -p 9090:8080 -e VERSION=v0.0.1 cncamp/httpserver:v0.0.1
docker logs -f httpserver
```
```
curl -v localhost:9090/healthz -H "ttt:yyy"
```
