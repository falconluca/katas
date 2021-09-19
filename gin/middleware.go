package gin

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("greeting", "Luca")

		c.Next()

		latency := time.Since(t)
		log.Println(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}

func MiddlewareEntry() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/greetings", func(c *gin.Context) {
		name := c.MustGet("greeting").(string)
		log.Println(name)
	})

	r.Run(":8080")
}
