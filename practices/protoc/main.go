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
	"strings"
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
	newDocument := &pb.User{Id: pb.NewObjectId(primitive.NewObjectID()), Name: "insertion"}
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
	fmt.Println(result)

	cancel()
}

func ConvertStructToBSON(data interface{}) bson.M {
	const numDefaultProp = 3
	bsonData := bson.M{}
	bytes, _ := bson.Marshal(data)
	_ = bson.Unmarshal(bytes, &bsonData)

	val := reflect.ValueOf(data)
	if reflect.Value.IsNil(val) {
		return bsonData
	}
	typeOfS := reflect.Indirect(val).Type()

	result := bson.M{}
	for i := numDefaultProp; i < len(bsonData)+numDefaultProp; i++ {
		nameField := strings.ToLower(typeOfS.Field(i).Name)
		valueField := reflect.Indirect(val).Field(i)
		typeField := fmt.Sprintf("%T", valueField.Interface())
		if typeField == "string" &&
			strings.Contains(nameField, "id") &&
			primitive.IsValidObjectID(valueField.String()) {
			if nameField == "id" {
				result["_id"], _ = primitive.ObjectIDFromHex(valueField.String())
			} else {
				result[nameField], _ = primitive.ObjectIDFromHex(valueField.String())
			}
		} else if typeField == "[]string" &&
			strings.Contains(nameField, "id") &&
			containsIsValidObjectID(valueField) {
			ids := valueField.Slice(0, valueField.Len()).Interface().([]string)
			bsonA := bson.A{}
			for _, id := range ids {
				objectId, _ := primitive.ObjectIDFromHex(id)
				bsonA = append(bsonA, objectId)
			}
			result[nameField] = bsonA
		} else if strings.Contains(typeField, "*pb") {
			result[nameField] = ConvertStructToBSON(valueField.Interface())
		} else {
			if typeField == "*timestamppb.Timestamp" {
				result[nameField] = valueField.Interface().(*timestamppb.Timestamp).AsTime()
			} else {
				result[nameField] = valueField.Interface()
			}
		}
	}
	return result
}

func containsIsValidObjectID(valueField reflect.Value) bool {
	ids := valueField.Slice(0, valueField.Len()).Interface().([]string)
	for _, id := range ids {
		if !primitive.IsValidObjectID(id) {
			return false
		}
	}
	return true
}
