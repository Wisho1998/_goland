package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	//newDocument := &pb.User{
	//	Id:    primitive.NewObjectID().Hex(),
	//	Name:  "insertion",
	//	Start: timestamppb.Now(),
	//	End:   timestamppb.Now(),
	//	Nested: &pb.Nested{
	//		Id:   primitive.NewObjectID().Hex(),
	//		Name: "yaseque",
	//	},
	//	NestedId: primitive.NewObjectID().Hex(),
	//	NestedIds: []string{
	//		primitive.NewObjectID().Hex(),
	//		primitive.NewObjectID().Hex(),
	//		primitive.NewObjectID().Hex(),
	//		primitive.NewObjectID().Hex(),
	//	},
	//}
	//insertOneResult, err := collection.InsertOne(ctx, newDocument)
	//if err != nil {
	//	fmt.Println("err insert one", err)
	//	return
	//}
	//fmt.Println("\ndone", insertOneResult.InsertedID)

	//RETRIEVE
	//FIND BY ID
	oid, _ := primitive.ObjectIDFromHex("62c5e6dcba1dd77414a614bc")
	filter := bson.M{"_id": oid}

	findOpt := options.FindOneAndUpdate()
	findOpt.SetReturnDocument(1)

	var result *pb.User
	findResult := collection.FindOneAndUpdate(ctx, filter, bson.D{{"$set", bson.M{"name": "nuevo3"}}}, findOpt).Decode(&result)

	fmt.Println(findResult)
	fmt.Println(result)

	//FIND BY DATE
	//layout := "2006-01-02 15:04:05.000Z"
	//t, _ := time.Parse(layout, "2022-06-17 13:28:05.054Z")
	//fmt.Println(timestamppb.New(t))
	//filter := bson.M{"date": timestamppb.New(t)}
	//var result *pb.User
	//findErr := collection.FindOne(ctx, filter).Decode(&result)
	//
	//if findErr != nil {
	//	log.Fatal(findErr)
	//}
	//fmt.Println(result)

	//UPDATE
	//oid, _ := primitive.ObjectIDFromHex("62ac7f986ad7c68fbf5f5911")
	//filter := bson.M{"_id": oid}
	//modified, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"date": timestamppb.Now()}})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(modified)
	//
	//modified2, err2 := collection.UpdateOne(context.Background(), filter, bson.M{"$push": bson.M{"nestedid": primitive.NewObjectID().Hex()}})
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//fmt.Println(modified2)

	// AGGREGATION
	//pipeAggregation := []bson.M{
	//	{"$match": bson.M{"_id": "62ac7f986ad7c68fbf5f5911"}},
	//	{"$lookup": bson.M{
	//		"from":         "nested",
	//		"localField":   "nestedid",
	//		"foreignField": "_id",
	//		"as":           "nestedid",
	//	}},
	//	{"$unwind": "$nestedid"},
	//	{"$replaceRoot": bson.M{"newRoot": "$nestedid"}},
	//}
	//
	//// Get result and perform DB operations
	//cursor, errAggregate := collection.Aggregate(context.Background(), pipeAggregation)
	//if errAggregate != nil {
	//	fmt.Println("err aggregate ", errAggregate)
	//}
	//if cursorErr := cursor.Err(); cursorErr != nil {
	//	fmt.Println("cursor.Err() aggregate ", cursor.Err())
	//}
	//
	//defer cursor.Close(context.Background())
	//fmt.Println("cursor: ", cursor)
	//for cursor.Next(context.Background()) {
	//	var currentUser *pb.Nested
	//	decodeErr := cursor.Decode(&currentUser)
	//	if decodeErr != nil {
	//		fmt.Println(decodeErr)
	//		return
	//	}
	//	fmt.Println(currentUser)
	//}
	cancel()
}
