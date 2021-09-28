package machine

import (
	"PasswordCenter/db"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func InPasswordCenter(c *gin.Context) {
	machine := Machine{}
	if err := db.DB.Collection("machine").FindOne(context.TODO(), bson.M{"ip": exnet.RemoteIP(c.Request)}).Decode(&machine); err != nil {
		fmt.Println(exnet.RemoteIP(c.Request))
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": -1})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 100})
}
