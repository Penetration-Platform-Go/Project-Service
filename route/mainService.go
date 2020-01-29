package route

import (
	"context"
	"time"

	pb "github.com/Penetration-Platform-Go/gRPC-Files/Auth-Service"
	"github.com/gin-gonic/gin"
)

func mainServiceRoute(route *gin.RouterGroup) {

	route.POST("/", func(ctx *gin.Context) {
		client := pb.NewAuthClient(AuthGrpcClient)
		r, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := client.GetUsernameByToken(r, &pb.Token{
			JWT: ctx.Request.Header.Get("Authenticate"),
		})
		if err != nil {
			ctx.Status(403)
		} else {
			ctx.JSON(200, result)
		}

	})
}
