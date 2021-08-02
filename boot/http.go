package boot

import (
	"ihtest/library/global"
	"ihtest/router"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	middleware "ihtest/library/http/meddleware"
	"ihtest/library/http/validator"
)

func InitHttpServer() error {
	binding.Validator = new(validator.DefaultValidator)

	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.RegisterGlobalParams())
	r.Use(middleware.Log())
	//r.Use(middleware.Transform())

	//RegisterSwagger(r)

	router.Register(r)

	global.Log.Info("start http server")
	err := r.Run(global.Config.Service.LocalHttpAddr)
	return err
}
