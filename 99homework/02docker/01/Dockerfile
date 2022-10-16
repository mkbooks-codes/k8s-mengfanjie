# 编译阶段
FROM golang:1.17-alpine AS build

# GO111MODULE=on: 强制开启gomod
# GOPROXY="https://goproxy.cn,direct": 使用国内七牛云包代理
# CGO_ENABLED=0: 关闭cgo，否则在alpine不可运行。
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

# 拷贝源码
WORKDIR /go/src/
COPY main.go .

# 预下载依赖 modules 到容器本地 cache
RUN go mod init example.com/httpserver
RUN go mod tidy

# 编译go程序
RUN go build -o /tmp/httpserver


# 运行阶段
FROM scratch
# 拷贝编译阶段编译好的程序包
COPY --from=build /tmp/httpserver /bin/httpserver
# 开放的端口
EXPOSE 8080
# 容器启动执行命令，启动httpserver服务
ENTRYPOINT ["/bin/httpserver"]
