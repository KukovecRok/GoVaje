package GinMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (r *Router) basicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Ta del se izvede pred nadaljnimi stvarmi
		fmt.Println("Before")

		//Z tem nadaljujemo zahtevo
		c.Next()

		//Ta del se izvede za nadaljnimi stvarmi
		fmt.Println("After")

	}
}

func (r *Router) checkPermissions(permissionNumber int) gin.HandlerFunc {
	return func(c *gin.Context) {

		/*
			Izvedemo kodo za preverjanje pravic in spustimo zahtevo naprej ali jo zavrnemo
		*/

	}
}
