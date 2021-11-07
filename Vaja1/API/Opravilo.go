package API

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"todorokvaja1/DataStructures"
)

func (a *Controller) GetOpravilo(c *gin.Context) {
	//[]DataStructure Opravilo - Array vseh opravil
	opravilo, err := a.c.GetOpravilo(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	// Vse ok - Status 200
	c.JSON(http.StatusOK, opravilo)
}
func (a *Controller) GetOpraviloById(c *gin.Context) {

	opravilo, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	OpraviloId, err := a.c.GetOpraviloById(c.Request.Context(), opravilo)
	if err != nil {
		//Vrne error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	//JSON serializacija
	c.JSON(http.StatusOK, OpraviloId)
}

func (a *Controller) InsertOpravilo(c *gin.Context) {

	var opravilo DataStructures.Opravilo

	err := c.BindJSON(&opravilo)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = a.c.InsertOpravilo(c.Request.Context(), opravilo)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "vstavljam novo opravilo")
}

func (a *Controller) RemoveOpravilo(c *gin.Context) {

	opravilo, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = a.c.RemoveOpravilo(c.Request.Context(), opravilo)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.String(http.StatusOK, "odstranjujem opravilo")
}

func (a *Controller) UpdateOpravilo(c *gin.Context) {

	opraviloID, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var opraviloUpdate DataStructures.Opravilo
	err = c.BindJSON(&opraviloUpdate)
	if err != nil {
		//Vrne error 409 - Conflict
		c.String(http.StatusConflict, err.Error())
		return
	}

	err = a.c.UpdateOpravilo(c.Request.Context(), opraviloID, opraviloUpdate)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "popravljam novo opravilo")
}
