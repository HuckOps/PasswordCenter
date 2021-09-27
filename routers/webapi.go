package routers

import (
	"PasswordCenter/machine"
	"PasswordCenter/tagAdmin"
	"PasswordCenter/token"
	"PasswordCenter/user"
	"github.com/gin-gonic/gin"
)

func PasswordCenterWebAPI(e *gin.Engine){
	PasswordCenterWebAPIGroup := e.Group("/")
	{
		PasswordCenterWebAPIGroup.POST("login", user.Login)
		PasswordCenterWebAPIGroup.POST("addtag", token.JWTAuthMiddleware(), tagAdmin.AddTag)
		PasswordCenterWebAPIGroup.POST("addmachine", token.JWTAuthMiddleware(), machine.AddMachine)
	}
}