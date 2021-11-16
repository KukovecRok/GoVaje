package GinMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Router struct{}

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
