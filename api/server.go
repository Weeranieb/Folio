package api

import (
	"Folio/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func registerRouter(r *gin.RouterGroup) {
	r.GET("/realtime", handler.GetRealTime())
	r.GET("/auth", handler.GetAuth())
	r.GET("/folio", handler.GetFolio())
	r.GET("/secure", handler.GetSecure())
	r.GET("/account", handler.GetAccount())

	r.POST("/realtime", handler.PostRealTime())
	r.POST("/auth", handler.PostAuth())
	r.POST("/folio", handler.PostFolio())
	r.POST("/secure", handler.PostSecure())
	r.POST("/account", handler.PostAccount())
}

func NewServer() *gin.Engine {
	router := gin.Default()

	r := router.Group("/api")

	registerRouter(r)

	return router
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
