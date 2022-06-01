package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"reflect"
	"time"
)

var tDateTime = reflect.TypeOf(primitive.DateTime(0))

func dateTimeEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tDateTime {
		return bsoncodec.ValueEncoderError{Name: "DateTimeEncodeValue", Types: []reflect.Type{tDateTime}, Received: val}
	}

	into := val.Interface().(string)
	return vw.WriteString(into)
}

func createCustomRegistry() *bsoncodec.RegistryBuilder {
	var primitiveCodecs bson.PrimitiveCodecs
	rb := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	rb.RegisterTypeEncoder(tDateTime, bsoncodec.ValueEncoderFunc(dateTimeEncodeValue))
	primitiveCodecs.RegisterPrimitiveCodecs(rb)
	return rb
}

func main() {
	var result struct {
		Id   string `bson:"_id"`
		Name string
		Date timestamppb.Timestamp
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, connErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://super:passw0rd@localhost:27017/dbtest"))

	if connErr != nil {
		log.Fatal(connErr)
	}

	collection := client.Database("dbtest").Collection("general")

	filter := bson.M{"name": "General"}
	findErr := collection.FindOne(ctx, filter).Decode(&result)
	var customRegistry = createCustomRegistry().Build()
	customData, _ := bson.MarshalExtJSONWithRegistry(customRegistry, result, false, false)

	if findErr != nil {
		log.Fatal(findErr)
	}

	fmt.Println(string(customData))

	//data, writeErr := bson.MarshalExtJSON(result, false, false)
	//if writeErr != nil {
	//	log.Fatal(writeErr)
	//}
	//
	//fmt.Println(string(data))
	cancel()
}
