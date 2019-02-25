package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (c *Config) Start() {
	router := gin.New()

	router.Use(middleware(c.RateLimit))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", getPing)
	}

	listenPort := fmt.Sprintf(":%s", c.ListenPort)
	router.Run(listenPort)
}
