package api

import (
	"Folio/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := Server{}
	router := gin.Default()

	router.GET("/realtime", handler.GetRealTime())
	router.GET("/auth", handler.GetAuth())
	router.GET("/folio", handler.GetFolio())
	router.GET("/secure", handler.GetSecure())
	router.GET("/account", handler.GetAccount())

	router.POST("/realtime", handler.PostRealTime())
	router.POST("/auth", handler.PostAuth())
	router.POST("/folio", handler.PostFolio())
	router.POST("/secure", handler.PostSecure())
	router.POST("/account", handler.PostAccount())
	server.router = router
	return &server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
