package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.String(200, "healthy")
	})

	r.GET("/", func(c *gin.Context) {
		envVar := os.Getenv("MY_ENV_VAR")
		c.String(200, "ENV VARIABLE: "+envVar)
	})

	log.Printf("Starting server on port 8081")
	err := r.Run(":8081")
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
