package Vaja1

import (
	"github.com/gin-gonic/gin"
	"todorokvaja1/API"
)

type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1
	api := r.engine.Group("/api/v1")

	//Pot /api/v1/opravilo
	opravilo := api.Group("/opravilo")
	r.registerOpraviloRoutes(opravilo)

	return

}

func (r *Router) registerOpraviloRoutes(opravilo *gin.RouterGroup) {
	opravilo.GET("/", r.api.GetOpravilo)
}
