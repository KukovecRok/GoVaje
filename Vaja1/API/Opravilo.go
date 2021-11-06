package API

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todorokvaja1/DataStructures"
)

func (a *Controller) GetOpravilo(c *gin.Context) {
	//Klic vseh opravil
	opravilo, err := a.c.GetOpravilo(c.Request.Context())
	if err != nil {
		// Error - status 400: Bad Request
		fmt.Print("TUKAJ")
		c.String(http.StatusInternalServerError, err.Error())
	}

	// Vse ok - Status 200
	c.JSON(http.StatusOK, opravilo)
}
func (a *Controller) GetOpraviloById(c *gin.Context) {

	//Dobimo user_id tipa string iz naslova in ga pretvorimo v int
	opravilo, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	//Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	OpraviloId, err := a.c.GetOpraviloById(c.Request.Context(), opravilo)
	if err != nil {
		//Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, OpraviloId)
}

func (a *Controller) InsertOpravilo(c *gin.Context) {

	var opravilo DataStructures.Opravilo

	err := c.BindJSON(&opravilo)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.InsertOpravilo(c.Request.Context(), opravilo)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "vstavljam novo opravilo")
}

func (a *Controller) RemoveOpravilo(c *gin.Context) {

	opravilo, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = a.c.RemoveOpravilo(c.Request.Context(), opravilo)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "odstranjujem opravilo")
}

func (a *Controller) UpdateOpravilo(c *gin.Context) {

	opraviloID, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var opraviloUpdate DataStructures.Opravilo

	err = a.c.UpdateOpravilo(c.Request.Context(), opraviloID, opraviloUpdate)
	if err != nil {
		//Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "popravljam novo opravilo")

}
