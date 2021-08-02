package router

import (
	"ihtest/app/medchan"
	"ihtest/app/store"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	registerMedChan(router)
	registerStore(router)
}

func registerMedChan(router *gin.Engine) {
	api := new(medchan.Controller)
	v1 := router.Group("/api/v1/channel")
	{
		//
		v1.POST("/list/", api.List)
	}
}

func registerStore(router *gin.Engine) {
	api := new(store.Controller)
	v1 := router.Group("/api/v1/store")
	{
		v1.POST("/list/", api.List)
	}
}
