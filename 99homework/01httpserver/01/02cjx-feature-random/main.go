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
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)

	x_data := map[string]interface{}{}
	values := list{}
	for i := 0; i < 10; i++ {
		values
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}

	data := map[string]interface{}{
		"data": x_data,
		"code": 200,
	}
	c.JSON(200, data)
}

// return {f"x{index}": round(random.uniform(-5, 5), 6) for index in range(10)}
