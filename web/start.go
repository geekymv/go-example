package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func fun1() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		log.Println("func1")
	}
}

func fun2() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		log.Println("func2")
	}
}

func main() {

	//r := gin.Default()
	r := gin.New()
	//r.RouterGroup
	apiGroup := r.Group("/api", fun1())

	g1 := apiGroup.Group("/group1", fun2())
	{
		g1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong1",
			})
		})
	}

	g2 := apiGroup.Group("/group2")
	{
		g2.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong2",
			})
		})
	}

	//handlers := make([]gin.HandlerFunc, 63)
	//r.Handle("GET", "/hello", handlers...)

	r.Run(":9090")
}
