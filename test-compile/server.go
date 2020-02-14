package main
import (
        "github.com/gin-gonic/gin"
)

func main() {
        r := gin.New()

        r.GET("/", func(c *gin.Context) {
                c.JSON(200, "Hello world!")
        })

	r.GET("/event", func(c *gin.Context) {
                c.JSON(200)                   
        })

        r.Run(":31337")
}
