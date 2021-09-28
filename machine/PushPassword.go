package machine

import (
	"PasswordCenter/db"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type Password struct {
	Ciphertext string `json:"ciphertext"`
	Key        string `json:"key"`
	IV         string `json:"iv"`
}

func PushPassword(c *gin.Context) {
	password := Password{}
	if err := c.ShouldBind(&password); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}
	_, err := db.DB.Collection("machine").UpdateOne(context.TODO(), bson.M{"ip": exnet.RemoteIP(c.Request)},
		bson.M{"$set": bson.M{"ciphertext": password.Ciphertext, "key": password.Key, "iv": password.IV}})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101})
	}
	c.JSON(http.StatusOK, gin.H{"code": 100})
}
