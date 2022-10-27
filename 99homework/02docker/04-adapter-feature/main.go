package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/feature", FeatureHandler)
	if err := r.Run(); err != nil {
		panic(err)
	}
}

// curl -v localhost:8080/feature
func FeatureHandler(c *gin.Context) {
	x_data := map[string]interface{}{}
	rand.Seed(time.Now().UnixNano())
	floats_list := randFloats(-5.0, 5.0, 10)
	for i := range floats_list {
		value := fmt.Sprintf("%.6f", floats_list[i])
		key := fmt.Sprintf("x%d", i)
		x_data[key] = value
	}

	data := map[string]interface{}{
		"data": x_data,
		"code": 200,
	}
	c.JSON(200, data)
}

// 根据区间产生随机数
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}
