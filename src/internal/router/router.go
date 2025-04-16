package router

import (
	"github.com/LacirJR/psygrow-api/src/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{

			users := v1.Group("/users")
			{
				users.POST("", handler.RegisterUser)
			}
		}

	}

}
