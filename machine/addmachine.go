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

type Machine struct {
	Hostname   string             `json:"hostname" bson:"hostname"`
	IP         string             `json:"ip" bson:"ip"`
	Tag        primitive.ObjectID `bson:"tag"`
	Target     string             `json:"target"`
	Ciphertext string             `bson:"ciphertext"`
	IV         string             `bson:"iv"`
	Key        string             `bson:"key"`
}

func AddMachine(c *gin.Context) {
	machine := Machine{}
	if err := c.ShouldBind(&machine); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}
	tagTmp := tag.Taget{}
	if err := db.DB.Collection("tag").FindOne(context.TODO(), bson.M{"tag": machine.Target}).Decode(&tagTmp); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101})
		return
	}
	tagSearchList := []primitive.ObjectID{tagTmp.ID, tag.SATarget}
	userTmp := user.User{}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": c.MustGet("username").(string), "tag": bson.M{"$in": tagSearchList}}).Decode(&userTmp); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 102})
		return
	}
	machineTmp := Machine{}
	err := db.DB.Collection("machine").FindOne(context.TODO(), bson.M{"$or": []bson.M{bson.M{"hostname": machine.Hostname}, bson.M{"ip": machine.IP}}}).Decode(&machineTmp)
	if err != nil {
		db.DB.Collection("machine").InsertOne(context.TODO(), bson.M{"hostname": machine.Hostname, "ip": machine.IP, "tag": tagTmp.ID})
		c.JSON(http.StatusOK, gin.H{"code": 100})
		return
	}
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{"code": 103})
	return

}
