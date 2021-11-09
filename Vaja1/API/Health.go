package API

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (a *Controller) GetHealth(c *gin.Context) {

	health, err := a.c.GetHealth(c.Request.Context())
	if err != nil {
		// /health - Naredi ping na bazo in vrne status 200 OK če je vse ok v nasprotnem primeru pa status 500 Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
	}
	// /health - Naredi ping na bazo in vrne status 200 OK če je vse ok v nasprotnem primeru pa status 500 Internal server error
	c.JSON(http.StatusOK, health)
}
