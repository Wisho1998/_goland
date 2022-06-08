package codecs

import (
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testProtocBsonCoders/pb"
)

var (
	ObjectIDRegistry = reflect.TypeOf(&pb.ObjectId{})
	ObjectIDCodecRef = &objectIDCodec{}

	ObjectIDMongoType = reflect.TypeOf(primitive.ObjectID{})
)

// objectIDCodec is codec for Protobuf ObjectId
type objectIDCodec struct {
}

func (e *objectIDCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(*pb.ObjectId)
	id, err := primitive.ObjectIDFromHex(v.Value)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(ObjectIDMongoType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(id))
}

func (e *objectIDCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(ObjectIDMongoType)
	if err != nil {
		return err
	}
	var id primitive.ObjectID
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&id).Elem()); err != nil {
		return err
	}
	oid := pb.NewObjectId(id)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(oid))
	return nil
}
