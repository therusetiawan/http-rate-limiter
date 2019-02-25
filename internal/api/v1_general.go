package api

import (
	"github.com/gin-gonic/gin"
)

func getPing(c *gin.Context) {

	result := map[string]interface{}{
		"ping": "pong",
	}

	httpOkResponse(c, result)
}
