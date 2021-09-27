package user

import (
	"PasswordCenter/db"
	"PasswordCenter/tag"
	"PasswordCenter/token"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type User struct {
	User string `json:"user" bson:"user"`
	Password string `json:"password" bson:"password"`
	Tag []primitive.ObjectID `json:"tag" bson:"tag"`
}

func Login(c *gin.Context){
	user := User{}
	if err := c.ShouldBind(&user) ; err != nil{
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}
	userTmp := User{}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": user.User, "password": user.Password}).Decode(&userTmp); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101})
		return
	}
	fmt.Println(userTmp.Tag)
	var tagList []string
	var tagTmp tag.Taget
	for _, t := range userTmp.Tag {
		fmt.Println(t)
		if err := db.DB.Collection("tag").FindOne(context.TODO(), bson.M{"_id": t}).Decode(&tagTmp); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": 102})
			return
		}
		tagList = append(tagList, tagTmp.Tag)
	}
	tokenString, _ := token.GenToken(user.User)
	c.JSON(http.StatusOK, gin.H{"token":tokenString, "tag": tagList})
}

