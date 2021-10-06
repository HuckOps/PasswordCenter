package routers

import (
	"PasswordCenter/machine"
	"PasswordCenter/tagAdmin"
	"PasswordCenter/token"
	"PasswordCenter/user"
	"github.com/gin-gonic/gin"
)

func PasswordCenterWebAPI(e *gin.Engine) {
	PasswordCenterWebAPIGroup := e.Group("/")
	{
		PasswordCenterWebAPIGroup.POST("login", user.Login)
		PasswordCenterWebAPIGroup.POST("addtag", token.JWTAuthMiddleware(), tagAdmin.AddTag)
		PasswordCenterWebAPIGroup.POST("addmachine", token.JWTAuthMiddleware(), machine.AddMachine)
		PasswordCenterWebAPIGroup.POST("deletemachine", token.JWTAuthMiddleware(), machine.DeleteMachine)
		PasswordCenterWebAPIGroup.GET("machine", machine.InPasswordCenter)
		PasswordCenterWebAPIGroup.POST("pushpassword", machine.PushPassword)
		PasswordCenterWebAPIGroup.POST("adduser", token.JWTAuthMiddleware(), user.AddUser)
		PasswordCenterWebAPIGroup.POST("deleteuser", token.JWTAuthMiddleware(), user.DeleteUser)
		PasswordCenterWebAPIGroup.GET("machinelist", token.JWTAuthMiddleware(), machine.MachineList)
		PasswordCenterWebAPIGroup.POST("password", token.JWTAuthMiddleware(), machine.GetPassword)
	}
}
