package codecs

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

var tOID = reflect.TypeOf("")

func ObjectIDEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tOID {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{tOID}, Received: val}
	}
	s := val.Interface().(string)
	isValidObjectID := primitive.IsValidObjectID(s)
	if isValidObjectID {
		id, err := primitive.ObjectIDFromHex(s)
		if err != nil {
			return err
		}
		return vw.WriteObjectID(id)
	}
	return vw.WriteString(s)
}

func ObjectIDDecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	// this is the function when we read the datetime format
	if vr.Type() == bson.TypeObjectID {
		readOID, err := vr.ReadObjectID()
		if err != nil {
			return err
		}
		oid := readOID.Hex()
		val.Set(reflect.ValueOf(oid))
	} else {
		readString, err := vr.ReadString()
		if err != nil {
			return err
		}
		val.Set(reflect.ValueOf(readString))
	}

	return nil
}
