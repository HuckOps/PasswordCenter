package machine

import (
	"PasswordCenter/db"
	"PasswordCenter/tag"
	"PasswordCenter/user"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

type GetPasswordRequest struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

func GetPassword(c *gin.Context) {
	getpasswordstruct := GetPasswordRequest{}
	if err := c.ShouldBind(&getpasswordstruct); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}
	username := c.MustGet("username")
	userTmp := user.User{}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": username}).Decode(&userTmp); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101})
		return
	}
	machine := Machine{}
	if err := db.DB.Collection("machine").FindOne(context.TODO(), bson.M{"hostname": getpasswordstruct.Hostname, "ip": getpasswordstruct.IP}).Decode(&machine); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 102})
		return
	}
	for _, tagTmp := range userTmp.Tag {
		if tagTmp == tag.SATarget || tagTmp == machine.Tag {
			c.JSON(http.StatusOK, gin.H{
				"code":       100,
				"ciphertext": machine.Ciphertext,
				"iv":         machine.IV,
				"key":        machine.Key,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 103})
	return
}
