package main
import (
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
)

func main() {
	r := gin.New()

	customMetrics := []*ginprometheus.Metric{
		&ginprometheus.Metric{
			ID:	"1234",				
			Name:	"test_metric",			
			Description:	"Counter test metric",	
			Type:	"counter",			
		},
		&ginprometheus.Metric{
			ID:	"1235",				
			Name:	"test_metric_2",		
			Description:	"Summary test metric",	
			Type:	"summary", 
		},
	}
	p := ginprometheus.NewPrometheus("gin", customMetrics)
	p.Use(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})
	
	r.Run(":1337") 
}
