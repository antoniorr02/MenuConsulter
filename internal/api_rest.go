package internal

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Router() {
	if router == nil {
		router.GET("/Plato/{Nombre}", getPlato)
		router.PUT("/Plato/{Nombre}", putPlato)
		router.DELETE("/Plato/{Nombre}", deletePlato)
		router.GET("/Comedor/{Nombre}", getComedor)
		router.PUT("/Comedor/{Nombre}", putComedor)
		router.DELETE("/Comedor/{Nombre}", deleteComedor)
		router.GET("/Menu/{Fecha}", getMenu)
		router.PUT("/Menu/{Fecha}", putMenu)
		router.DELETE("/Menu/{Fecha}", deleteMenu)
	}
}

func getPlato(c *gin.Context) {

}

func putPlato(c *gin.Context) {

}

func deletePlato(c *gin.Context) {

}

func getComedor(c *gin.Context) {

}

func putComedor(c *gin.Context) {

}

func deleteComedor(c *gin.Context) {

}

func getMenu(c *gin.Context) {

}

func putMenu(c *gin.Context) {

}

func deleteMenu(c *gin.Context) {

}
