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

var tOID = reflect.TypeOf(primitive.ObjectID{})
var tDateTime = reflect.TypeOf(primitive.DateTime(0))

//func dateTimeEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
//	v := val.Interface().(*timestamp.Timestamp)
//	t, err := ptypes.Timestamp(v)
//	if err != nil {
//		return err
//	}
//
//	return vw.WriteTimestamp(uint32(t.Second()), uint32(t.Nanosecond()))
//}

func dateTimeDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	// this is the function when we read the datetime format
	read, err := vr.ReadDateTime()
	if err != nil {
		return err
	}
	val.SetInt(read)
	return nil
}

func objectIDEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tOID {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{tOID}, Received: val}
	}
	s := val.Interface().(primitive.ObjectID).Hex()
	return vw.WriteString(s)
}

func createCustomRegistry() *bsoncodec.RegistryBuilder {
	var primitiveCodecs bson.PrimitiveCodecs
	rb := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	//rb.RegisterTypeEncoder(tDateTime, bsoncodec.ValueEncoderFunc(dateTimeEncodeValue))
	rb.RegisterTypeDecoder(tDateTime, bsoncodec.ValueDecoderFunc(dateTimeDecodeValue))
	rb.RegisterTypeEncoder(tOID, bsoncodec.ValueEncoderFunc(objectIDEncodeValue))
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

	filter := bson.M{"name": "General"}
	var result *pb.User
	findErr := collection.FindOne(ctx, filter).Decode(&result)

	if findErr != nil {
		log.Fatal(findErr)
	}
	fmt.Println(result)
	cancel()
}
