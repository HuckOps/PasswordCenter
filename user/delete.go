package user

import (
	"PasswordCenter/db"
	"PasswordCenter/tag"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type DeleteUserStruct struct {
	User string `json:"user"`
}

func DeleteUser(c *gin.Context) {
	deleteuser := DeleteUserStruct{}
	if err := c.ShouldBind(&deleteuser); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}
	user := User{}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": c.MustGet("username").(string), "tag": tag.SATarget}).Decode(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101})
		return
	}
	_, err := db.DB.Collection("user").DeleteOne(context.TODO(), bson.M{"user": deleteuser.User})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 102})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 100})
}
