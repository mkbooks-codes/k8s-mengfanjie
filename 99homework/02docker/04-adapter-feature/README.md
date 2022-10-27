# 代码运行测试
```
go run main.go
```
```
curl -v localhost:8080/feature -H 'Content-Type: application/json' -d '{}'
```
# 镜像测试
```
docker build -t klb/adapter:v1 .
docker run -d --name adapter -p 20000:8080 klb/adapter:v1
docker logs -f adapter
```
```
curl -v localhost:20000/feature -H 'Content-Type: application/json' -d '{}'
```
