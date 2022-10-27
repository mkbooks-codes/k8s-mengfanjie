FROM golang:1.17-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /go/src/
COPY main.go .

RUN go mod init klb.com/adapter && go mod tidy && go build -o /tmp/adapter


FROM scratch
COPY --from=build /tmp/adapter /bin/adapter
EXPOSE 8080
ENTRYPOINT ["/bin/adapter"]
