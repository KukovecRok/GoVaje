package API

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"todorokvaja1/DataStructures"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
