package main
import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"fmt"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	fmt.Println("!! Hello world from The New Version 5.0 of test app!! It works!")

	customMetrics := []*ginprometheus.Metric{
		&ginprometheus.Metric{
			ID:	"31337",
			Name:	"test_metric",
			Description:	"Counter test metric",
			Type:	"counter",
		},
	}
	p := ginprometheus.NewPrometheus("gin", customMetrics)
	p.Use(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "It's new version!! 5.0 Hello world!")
	})

	r.Run(":1337")
}
