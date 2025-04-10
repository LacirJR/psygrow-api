package router

import (
	"github.com/LacirJR/psygrow-api/src/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/ping", handler.Ping)
		}
	}

}
