package routes

import (
	ctrl "github.com/mendezandrewm/monGo-UserDB/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(req *gin.Engine) {
	req.GET("/", ctrl.TestAPI())
	req.POST("/signup", ctrl.SignUp())
	req.POST("/login", ctrl.Login())
}
