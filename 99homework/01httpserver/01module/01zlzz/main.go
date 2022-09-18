package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
)

func main() {
	r := gin.Default()
	r.GET("/healthz", healthHandler)
	// listen and serve on 0.0.0.0:8080
	if err := r.Run(); err != nil {
		panic(err)
	}
}

// healthHandler healthcheck实现, curl -v localhost:8080/healthz -H "ttt:yyy"
func healthHandler(c *gin.Context) {

	// 遍历request的header并循环写入response header
	for k, v := range c.Request.Header {
		//fmt.Printf("key : %s\n", k)
		//fmt.Printf("value : %v\n", v)
		c.Writer.Header().Set(k, v[0])
	}

	// 获取环境变量VERSION并写入reponse header
	c.Writer.Header().Set("version", os.Getenv("VERSION"))

	// 日志记录客户端ip, http返回码, 输出到标准输出
	ip := exnet.ClientPublicIP(c.Request)
	if ip == "" {
		ip = exnet.ClientIP(c.Request)
	}
	fmt.Printf("客户端ip:%s, http返回码:%d\n", ip, 200)

	// 响应healthcheck信息, 状态码为200
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
