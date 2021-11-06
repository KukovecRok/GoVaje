package main

import (
	"github.com/gin-gonic/gin"
	"todorokvaja1/API"
)

type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1/todo
	api := r.engine.Group("/api/v1/todo")

	//Pot /api/v1/todo/opravilo
	opravilo := api.Group("/opravilo")
	r.registerOpraviloRoutes(opravilo)

	return

}

func (r *Router) registerOpraviloRoutes(opravilo *gin.RouterGroup) {
	opravilo.GET("/", r.api.GetOpravilo)
	opravilo.POST("/", r.api.InsertOpravilo)
	opravilo.GET("/:todo_id", r.api.GetOpraviloById)
	opravilo.DELETE("/:todo_id", r.api.RemoveOpravilo)
	opravilo.PUT("/:todo_id", r.api.UpdateOpravilo)
}
