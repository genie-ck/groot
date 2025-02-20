package routers

import (
	v1 "github.com/crystal/groot/routers/v1"
	"github.com/gin-gonic/gin"
)

const Domain = "hackerone.com"

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiv1 := r.Group("/v1")
	{
		apiv1.GET("/domains", v1.GetDomainsByProject)
	}
	return r
}
