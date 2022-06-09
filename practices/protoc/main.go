package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"testProtocBsonCoders/codecs"
	"testProtocBsonCoders/pb"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	const uri = "mongodb://super:passw0rd@localhost:27017/dbtest"
	client, connErr := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetRegistry(codecs.Register(bson.NewRegistryBuilder()).Build()))

	if connErr != nil {
		log.Fatal(connErr)
	}

	collection := client.Database("dbtest").Collection("general")

	//INSERT
	newDocument := &pb.User{
		Id:   primitive.NewObjectID().Hex(),
		Date: timestamppb.Now(),
		Name: "insertion",
		NestedId: []string{
			primitive.NewObjectID().Hex(),
			primitive.NewObjectID().Hex(),
			primitive.NewObjectID().Hex(),
			primitive.NewObjectID().Hex(),
		},
	}

	insertOneResult, err := collection.InsertOne(ctx, newDocument)
	if err != nil {
		fmt.Println("err insert one")
		return
	}
	fmt.Println("done", insertOneResult.InsertedID)

	//RETRIEVE
	filter := bson.M{"_id": insertOneResult.InsertedID.(primitive.ObjectID)}
	var result *pb.User
	findErr := collection.FindOne(ctx, filter).Decode(&result)

	if findErr != nil {
		log.Fatal(findErr)
	}
	fmt.Println(result)
	cancel()
}
