package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *Controller) GetOpravilo(c *gin.Context) {
	//Klic vseh opravil
	opravilo, err := a.c.GetOpravilo(c.Request.Context())
	if err != nil {
		// Error - status 400: Bad Request
		c.String(http.StatusBadRequest, err.Error())
	}
	// Vse ok - Status 200
	c.JSON(http.StatusOK, opravilo)
}
