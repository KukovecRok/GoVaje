package main

// obvezno package main!!

import (
	"github.com/gin-gonic/gin"
	"todorokvaja1/API"
)

type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1/to_do
	api := r.engine.Group("/api/v1/todo")

	//Pot /api/v1/"to_do/opravilo
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

/*
GET ALL http://localhost:8000/api/v1/to do/opravilo/
POST ONE http://localhost:8000/api/v1/to do/opravilo/ - ID je generiran avtomatsko, Postman - Body, raw, JSON, {
  "naslov": "Primer",
  "opis": "Dodajanja",
  "datum_dodajanja": "2021-08-20T11:42:27Z",
  "predviden_datum_dela": "2021-06-20T04:13:45Z"
}
GET BY ID http://localhost:8000/api/v1/to do/opravilo/ID
DELETE BY ID http://localhost:8000/api/v1/to do/opravilo/ID
PUT http://localhost:8000/api/v1/to do/opravilo/ID + JSON, enako kot POST
*/
