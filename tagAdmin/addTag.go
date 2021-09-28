package tagAdmin

import (
	"PasswordCenter/db"
	"PasswordCenter/tag"
	"PasswordCenter/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"net/http"
)

type AddTagStruct struct {
	Tag string `json:"tag"`
}

func AddTag(c *gin.Context) {
	var addtag AddTagStruct
	if err := c.ShouldBind(&addtag); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}
	fmt.Println(c.MustGet("username").(string))
	user := user.User{}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": c.MustGet("username").(string), "tag": tag.SATarget}).Decode(&user); err == nil {
		var tagDecode tag.Taget
		if err := db.DB.Collection("tag").FindOne(context.TODO(), bson.M{"tag": addtag.Tag}).Decode(&tagDecode); err != nil {
			db.DB.Collection("tag").InsertOne(context.TODO(), bson.M{"tag": addtag.Tag})
			c.JSON(http.StatusOK, gin.H{"code": 100})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 103})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"code": 1})
}
