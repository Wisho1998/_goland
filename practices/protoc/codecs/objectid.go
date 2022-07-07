package codecs

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

var stringType = reflect.TypeOf("")
var ObjectIDCodecRef = &ObjectIDCodec{}

type ObjectIDCodec struct{}

// EncodeValue encodes Protobuf ObjectId value to BSON value
func (e *ObjectIDCodec) EncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != stringType {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{stringType}, Received: val}
	}
	v := val.Interface().(string)

	if isValidObjectID := primitive.IsValidObjectID(v); isValidObjectID {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return err
		}
		return vw.WriteObjectID(id)
	}
	return vw.WriteString(v)
}

// DecodeValue decodes BSON value to ObjectId value
func (e *ObjectIDCodec) DecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if vr.Type() == bson.TypeObjectID {
		readOID, err := vr.ReadObjectID()
		if err != nil {
			return err
		}
		oid := readOID.Hex()
		val.Set(reflect.ValueOf(oid))
	} else if vr.Type() == bson.TypeNull {
		err := vr.ReadNull()
		if err != nil {
			return err
		}
		val.Set(reflect.ValueOf(""))
	} else {
		readString, err := vr.ReadString()
		if err != nil {
			return err
		}
		val.Set(reflect.ValueOf(readString))
	}
	return nil
}
