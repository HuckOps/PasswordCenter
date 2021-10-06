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
	"log"
	"net/http"
)

func MachineList(c *gin.Context) {
	username := c.MustGet("username").(string)
	userTmp := user.User{}
	if err := db.DB.Collection("user").FindOne(context.TODO(), bson.M{"user": username}).Decode(&userTmp); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 101})
		return
	}
	tagName := make(map[primitive.ObjectID]string)
	re := []gin.H{}
	for _, tagTmp := range userTmp.Tag {
		fmt.Println(tagTmp)
		if tagTmp == tag.SATarget {
			result, _ := db.DB.Collection("machine").Find(context.TODO(), bson.M{})
			for result.Next(context.TODO()) {
				machine := Machine{}
				result.Decode(&machine)
				_, err := tagName[machine.Tag]
				if err == false {
					tagStructTmp := tag.Taget{}
					db.DB.Collection("tag").FindOne(context.TODO(), bson.M{"_id": machine.Tag}).Decode(&tagStructTmp)
					tagName[machine.Tag] = tagStructTmp.Tag
					log.Println("置缓存")
				}
				re = append(re, gin.H{
					"hostname": machine.Hostname,
					"IP":       machine.IP,
					"tag":      tagName[machine.Tag],
				})
			}
			c.JSON(http.StatusOK, gin.H{"code": 100, "machine": re})
			return
		}
	}
	result, _ := db.DB.Collection("machine").Find(context.TODO(), bson.M{"tag": bson.M{"$in": userTmp.Tag}})
	for result.Next(context.TODO()) {
		machine := Machine{}
		result.Decode(&machine)
		_, err := tagName[machine.Tag]
		if err == false {
			tagStructTmp := tag.Taget{}
			db.DB.Collection("tag").FindOne(context.TODO(), bson.M{"_id": machine.Tag}).Decode(&tagStructTmp)
			tagName[machine.Tag] = tagStructTmp.Tag
			log.Println("置缓存")
		}
		re = append(re, gin.H{
			"hostname": machine.Hostname,
			"IP":       machine.IP,
			"tag":      tagName[machine.Tag],
		})
	}
	c.JSON(http.StatusOK, gin.H{"code": 100, "machine": re})
}
