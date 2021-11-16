package API

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"os"
	"time"
	"todorokvaja1/DataStructures"
	"todorokvaja1/Token"
)

func (a *Controller) Login(c *gin.Context) {

	var user DataStructures.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	pravilno, err := a.c.Login(c.Request.Context(), user)

	if pravilno != true {
		//Vrne error 401 - Unauthorized
		c.String(http.StatusUnauthorized, "Prijava uspešna: %t", pravilno)
		err = errors.New("Unauthorized")
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	token, err := CreateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	maker, err := Token.NewJWTMaker("98zgbnmkiutfvbnjuztreertgbnjzdfbnjkKIUZTFCVBJUZFCVBNJKIUZGFCVBNMKiuztrdcvbnmkiuztfvb")

	username := user.Username
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	fmt.Print(expiredAt)

	token, err = maker.CreateToken(username, duration)

	payload, err := maker.VerifyToken(token)
	fmt.Println(payload)
	fmt.Println("TEST")
	fmt.Println(maker.VerifyToken(token))

	c.String(http.StatusOK, "Prijava uspešna: %t, nov token: %s", pravilno, token)
}

func (a *Controller) InsertUser(c *gin.Context) {

	var user DataStructures.User
	err := c.BindJSON(&user)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	err = a.c.InsertUser(c.Request.Context(), user)

	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	c.String(http.StatusOK, "registracija novega userja")
}

func (a *Controller) GetUserByName(c *gin.Context) {

	userByName := c.Param("username")

	user, err := a.c.GetUserByName(c.Request.Context(), userByName)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, user)
}

func CreateToken(userid primitive.ObjectID) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "fallback") //this should be in an env file
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

var ErrExpiredToken = errors.New("token has expired")
