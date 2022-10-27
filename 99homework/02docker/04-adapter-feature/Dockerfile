FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /root/cncamp
COPY main.go .

RUN go mod init example.com/httpserver
RUN go mod tidy
RUN go build -o httpserver .

WORKDIR /opt/modules/httpserver
RUN mkdir src .

RUN cp /root/cncamp/httpserver ./src
EXPOSE 8080

CMD ["/opt/modules/httpserver/src/httpserver"]