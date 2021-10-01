package user

import (
	"PasswordCenter/db"
	"PasswordCenter/tag"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type NewUser struct {
	User     string   `json:"user" bson:"user"`
	Password string   `json:"password" bson:"password"`
	Tag      []string `json:"tag" bson:"tag"`
	TagID    []primitive.ObjectID
}

func AddUser(c *gin.Context) {
	newuser := NewUser{}
	if err := c.ShouldBind(&newuser); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}

	user := User{}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": newuser.User}).Decode(&user); err == nil {
		c.JSON(http.StatusOK, gin.H{"code": 104})
		return
	}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": c.MustGet("username").(string), "tag": tag.SATarget}).Decode(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101})
		return
	}

	result, _ := db.DB.Collection("tag").Find(context.TODO(), bson.M{"tag": bson.M{"$in": newuser.Tag}})
	for result.Next(context.TODO()) {
		tagTmp := tag.Taget{}
		if err := result.Decode(&tagTmp); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 102})
			return
		}
		newuser.TagID = append(newuser.TagID, tagTmp.ID)
	}
	_, err := db.DB.Collection("user").InsertOne(context.TODO(), bson.M{"user": newuser.User, "password": newuser.Password, "tag": newuser.TagID})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 103})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 100})
}
