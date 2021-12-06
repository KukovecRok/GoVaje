package main

// obvezno package main!!

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todorokvaja1/API"
)

type Router struct {
	engine *gin.Engine
	api    API.Controller
	secret []byte
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1/to.do
	api := r.engine.Group("/api/v1/todo")

	//Pot /api/v1/"to_do/opravilo
	opravilo := api.Group("/opravilo")
	opravilo.Use(r.CheckPermission())
	r.registerOpraviloRoutes(opravilo)

	login := api.Group("/login")
	r.registerLoginRoutes(login)

	health := api.Group("/health")
	health.Use(r.CheckPermission())
	r.registerHealthRoutes(health)

	return
}

func (r *Router) registerOpraviloRoutes(opravilo *gin.RouterGroup) {
	opravilo.GET("/", r.api.GetOpravilo)
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

func (r *Router) CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		tokenSplit := strings.Split(tokenString, " ")
		if len(tokenSplit) != 2 {
			c.AbortWithError(500, errors.New("Nepravilno splitanje tokena"))
			return
		}
		token := tokenSplit[1]

		tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				sentry.CaptureException(errors.New("error checking JWT token"))
				c.AbortWithStatus(http.StatusUnauthorized)
				return nil, errors.New("there was an error")
			}
			return r.secret, nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok && tokenParsed.Valid {
			c.Set("user_id", claims["user_id"])
		}
	}
}
