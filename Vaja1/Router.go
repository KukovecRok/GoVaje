package main

// obvezno package main!!

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todorokvaja1/API"
)

type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1/to.do
	api := r.engine.Group("/api/v1/todo")

	test := api.Group("/test")
	test.GET("/", r.CheckPermission(101), r.api.GetOpravilo)

	//Pot /api/v1/"to_do/opravilo
	opravilo := api.Group("/opravilo")
	r.registerOpraviloRoutes(opravilo)

	login := api.Group("/login")
	r.registerLoginRoutes(login)

	health := api.Group("/health")
	r.registerHealthRoutes(health)

	return
}

func (r *Router) registerOpraviloRoutes(opravilo *gin.RouterGroup) {
	opravilo.GET("/", r.CheckPermission(101), r.api.GetOpravilo)
	opravilo.POST("/", r.api.InsertOpravilo)
	opravilo.GET("/:todo_id", r.api.GetOpraviloById)
	opravilo.DELETE("/:todo_id", r.api.RemoveOpravilo)
	opravilo.PUT("/:todo_id", r.api.UpdateOpravilo)
}
func (r *Router) registerLoginRoutes(login *gin.RouterGroup) {
	login.POST("/register", r.api.InsertUser)
	login.GET("/:username", r.api.GetUserByName)
	login.POST("/", r.api.Login)
}

func (r *Router) registerHealthRoutes(health *gin.RouterGroup) {
	health.GET("/", r.api.GetHealth)
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

func (r *Router) basicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Ta del se izvede pred nadaljnjimi stvarmi
		fmt.Println("Before")

		//Z tem nadaljujemo zahtevo
		c.Next()

		//Ta del se izvede za nadaljnjimi stvarmi
		fmt.Println("After")

	}
}

func (r *Router) CheckPermission(permissionNumber int) gin.HandlerFunc {
	return func(c *gin.Context) {

		/*
			Izvedemo kodo za preverjanje pravic in spustimo zahtevo naprej ali jo zavrnemo
		*/

	}
}
