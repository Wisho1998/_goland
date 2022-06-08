package codecs

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"time"
)

var (
	TimestampRegistry = reflect.TypeOf(&timestamp.Timestamp{})
	TimestampCodecRef = &timestampCodec{}

	DateMongoType = reflect.TypeOf(time.Time{})
)

// timestampCodec is codec for Protobuf Timestamp
type timestampCodec struct {
}

// EncodeValue encodes Protobuf Timestamp value to BSON value
func (e *timestampCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(*timestamp.Timestamp)
	t := time.Unix(v.Seconds, int64(v.Nanos))
	enc, err := ectx.LookupEncoder(DateMongoType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(t.In(time.UTC)))
}

// DecodeValue decodes BSON value to Timestamp value
func (e *timestampCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(DateMongoType)
	if err != nil {
		return err
	}
	var t time.Time
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	ts := timestamppb.New(t.In(time.UTC))
	val.Set(reflect.ValueOf(ts))
	return nil
}
