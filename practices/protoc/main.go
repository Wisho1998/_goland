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
	"log"
	"reflect"
	"testProtoc/pb"
	"time"
)

var tOID = reflect.TypeOf(&pb.ObjectId{})

func objectIDEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tOID {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{tOID}, Received: val}
	}
	s := val.Interface().(*pb.ObjectId)
	id, err := primitive.ObjectIDFromHex(s.Value)
	if err != nil {
		return err
	}

	return vw.WriteObjectID(id)
}

func objectIDDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	// this is the function when we read the datetime format
	read, err := vr.ReadObjectID()
	if err != nil {
		return err
	}
	oid := *pb.NewObjectId(read)
	val.Set(reflect.ValueOf(&oid))
	return nil
}

func createCustomRegistry() *bsoncodec.RegistryBuilder {
	var primitiveCodecs bson.PrimitiveCodecs
	rb := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	rb.RegisterTypeEncoder(tOID, bsoncodec.ValueEncoderFunc(objectIDEncodeValue))
	rb.RegisterTypeDecoder(tOID, bsoncodec.ValueDecoderFunc(objectIDDecodeValue))
	primitiveCodecs.RegisterPrimitiveCodecs(rb)
	return rb
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	const uri = "mongodb://super:passw0rd@localhost:27017/dbtest"
	client, connErr := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetRegistry(createCustomRegistry().Build()))

	if connErr != nil {
		log.Fatal(connErr)
	}

	collection := client.Database("dbtest").Collection("general")

	//INSERT
	newDocument := &pb.User{
		Id:   pb.NewObjectId(primitive.NewObjectID()),
		Name: "insertion",
		NestedId: []*pb.ObjectId{
			pb.NewObjectId(primitive.NewObjectID()),
			pb.NewObjectId(primitive.NewObjectID()),
			pb.NewObjectId(primitive.NewObjectID()),
			pb.NewObjectId(primitive.NewObjectID()),
		},
	}
	//bsonConverted := ConvertStructToBSON(newDocument)
	//fmt.Printf("- new document: \n %v \n- bson converted: \n%v \n\n", newDocument, bsonConverted)
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
	fmt.Println(result.NestedId)

	for _, id := range result.GetNestedId() {
		fmt.Println(id.GetHexID())
	}

	cancel()
}
