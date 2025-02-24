package routers

import (
	"BlockHouse_Test_AbishekLwgaun/handler"
	"github.com/gin-gonic/gin"
)

// setting up routes
func Create_routes() *gin.Engine {

	// gin engine instance with default middlewares
	r := gin.Default()

	//routes for handling requests
	r.GET("/orders", handler.GetOrder)
	r.POST("/orders", handler.PostOrder)

	// return the gin router instance
	return r
}
