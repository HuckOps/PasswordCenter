package machine

import (
	"PasswordCenter/db"
	"PasswordCenter/tag"
	"PasswordCenter/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type DeleteMachineStruct struct {
	IP string `json:"ip"`
}

func DeleteMachine(c *gin.Context) {
	deletemachine := DeleteMachineStruct{}
	if err := c.ShouldBind(&deletemachine); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}
	userTmp := user.User{}
	machine := Machine{}
	if err := db.DB.Collection("machine").FindOne(context.TODO(), bson.M{"ip": deletemachine.IP}).Decode(&machine); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101})
		return
	}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": c.MustGet("username").(string), "tag": bson.M{"$in": []primitive.ObjectID{machine.Tag, tag.SATarget}}}).Decode(&userTmp); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": 102})
		return
	}
	_, err := db.DB.Collection("machine").DeleteOne(context.TODO(), bson.M{"ip": deletemachine.IP})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 103})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 100})
}
