package routes

import (
	"log-processor/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/traces", controllers.GetTraces)
}
