package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/therusetiawan/http-rate-limiter/pkg/redis"
)

func middleware(rateLimit int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get user token
		userToken := ctx.Query("user_token")
		if userToken == "" {
			httpValidationErrorResponse(ctx, "user_token is required")
			ctx.Abort()
			return
		}

		rateLimitBasedOnToken(ctx, rateLimit, userToken)
	}
}

func rateLimitBasedOnToken(ctx *gin.Context, rateLimit int64, userToken string) {
	start := time.Now()
	expTime := start.Add(time.Second * 1) // key will expired in 1 second

	requestCount, err := redis.LLen(userToken)
	if err != nil {
		fmt.Println("Can't get data from redis")
	}

	if requestCount >= rateLimit { // request reach the limit
		httpTooManyRequestsResponse(ctx)
		ctx.Abort()
		return
	} else {
		if exist := redis.Exists(userToken); exist == 0 {
			// TODO : must be run in transaction mode
			redis.RPush(userToken, userToken)
			redis.ExpireAt(userToken, expTime)
		} else {
			// insert if only key already exist
			// handling race condition using RPushX
			redis.RPushX(userToken, userToken)
		}
	}

}
