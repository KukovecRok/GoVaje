package main

// obvezno package main!!

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
	"todorokvaja1/API"
	"todorokvaja1/DataStructures"
)

type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	//Pot /api/v1/to_do
	api := r.engine.Group("/api/v1/todo")

	authorized := api.Group("/admin", gin.BasicAuth(gin.Accounts{
		"testko": "$2a$14$hgha5kOBqlmH3ikUn/6E3OQmjAOxTcUY/dz4oMH7JaurrwBE.8/qK",
		"rok":    "123456",
	}))

	//Pot /api/v1/"to_do/opravilo
	opravilo := api.Group("/opravilo")
	r.registerOpraviloRoutes(opravilo)

	login := api.Group("/login")
	r.registerLoginRoutes(login)

	health := api.Group("/health")
	r.registerHealthRoutes(health)

	authorized.GET("/prijava", func(c *gin.Context) {
		// GET user, nastavljen by BasicAuth
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Print(user + " Prepoznan")
		c.JSON(http.StatusOK, "User prepoznan")
	})

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
	login.POST("/", Login)
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var user = DataStructures.User{
	Username: "username",
	Password: "password",
}

func Login(c *gin.Context) {
	var u DataStructures.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(u.Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
func CreateToken(userid primitive.ObjectID) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
