package routers

import (
	"github.com/gin-gonic/gin"

	"go/gin/api/routers/api/v1"

	"go/gin/api/pkg/setting"
)

func InitRouter() *gin.Engine{
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunModule)

	apiV1 := r.Group("/api/v1")

	{
		apiV1.GET("/tag", v1.GetTags)
		apiV1.POST("/tag", v1.AddTag)
		apiV1.PUT("/tag", v1.EditTag)
		apiV1.DELETE("/tag", v1.DeleteTag)
	}

	return r
}