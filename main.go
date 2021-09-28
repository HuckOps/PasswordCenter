package main

import (
	"PasswordCenter/db"
	"PasswordCenter/routers"
	"PasswordCenter/tag"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitMongoDB()
	tag.GetSATarget()
	r := gin.Default()
	routers.PasswordCenterWebAPI(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
