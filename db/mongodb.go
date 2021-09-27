package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var DB mongo.Database

func InitMongoDB(){
	clientOptions := options.Client().
		ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("数据库连接失败")
		os.Exit(1)
	}
	log.Println("数据库连接成功")
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Println("数据库通信失败")
		os.Exit(1)
	}else {
		log.Println("数据库通信成功")
	}
	DB = *client.Database("password_center")
}
