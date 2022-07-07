package codecs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
)

var timeStampType = reflect.TypeOf(&timestamppb.Timestamp{})
var TimeCodecRef = &TimeCodec{}

type TimeCodec struct{}

// EncodeValue encodes Protobuf Timestamp value to BSON value
func (e *TimeCodec) EncodeValue(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != timeStampType {
		return bsoncodec.ValueEncoderError{Name: "ObjectIDEncodeValue", Types: []reflect.Type{timeStampType}, Received: val}
	}

	if val.IsNil() {
		return vw.WriteNull()
	}

	v := val.Interface().(*timestamp.Timestamp)
	t := v.AsTime()

	return vw.WriteDateTime(t.UnixMilli())
}

// DecodeValue decodes BSON value to Timestamp value
func (e *TimeCodec) DecodeValue(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if vr.Type() == bson.TypeDateTime {
		readDatetime, err := vr.ReadDateTime()
		if err != nil {
			return err
		}
		ts := timestamppb.New(time.UnixMilli(readDatetime))
		val.Set(reflect.ValueOf(ts))
	} else {
		err := vr.ReadNull()
		if err != nil {
			return err
		}
		var ts *timestamppb.Timestamp = nil
		val.Set(reflect.ValueOf(ts))
	}
	return nil
}
