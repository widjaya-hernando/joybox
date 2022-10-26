package server

import (
	"fmt"
	"io/ioutil"
	"sleekflow/app/api/middleware"
	"sleekflow/lib/validators"

	// "sleekflow/lib/validators"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Handler interface {
	Register(r *gin.Engine, m *middleware.Middleware)
}

func BuildHandler(middleware *middleware.Middleware, handlers ...Handler) http.Handler {
	if os.Getenv("APP_ENV") == "production" {
		log.Println("on production=================================")
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// register all custom validator here
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("unique", validators.CustomValidator.Unique())
		if err != nil {
			log.Println("error when applying unique validator")
		}
		err = v.RegisterValidation("exist", validators.CustomValidator.Exist())
		if err != nil {
			log.Println("error when applying exist validator")
		}
		err = v.RegisterValidation("exist_multiple", validators.CustomValidator.ExistMultiple())
		if err != nil {
			log.Println("error when applying exist_multiple validator")
		}
		err = v.RegisterValidation("value", validators.CustomValidator.Value())
		if err != nil {
			log.Println("error when applying value validator")
		}
	}

	router.Use(middleware.ErrorHandle())

	// set max upload file size
	//router.MaxMultipartMemory = 8 << 20  // 8 MiB

	// router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// router.GET("/healthz", healthz)
	//router.GET("/handling-midtrans", test)
	//router.GET("/test",test)

	// start registering routes from all handlers
	for _, reg := range handlers {
		reg.Register(router, middleware)
	}

	// 404 not found function
	router.NoRoute(notFound)

	return router
}

func healthz(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprint("ok:", runtime.NumGoroutine()))
}

func test(c *gin.Context) {
	x, _ := ioutil.ReadAll(c.Request.Body)
	log.Println("===========")
	log.Printf("%s", string(x))
}

func notFound(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}
