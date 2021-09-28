package tag

import (
	"PasswordCenter/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
)

var SATarget primitive.ObjectID

func GetSATarget() {
	satag := Taget{}
	if err := db.DB.Collection("tag").FindOne(context.TODO(), bson.M{"tag": "owt.sa"}).Decode(&satag); err != nil {
		log.Println("不存在SA Tag")
		os.Exit(1)
	}
	SATarget = satag.ID
	fmt.Println(satag)
}
