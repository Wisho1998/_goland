package codecs

import (
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

// Register registers Google protocol buffers types codecs
func Register(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	rb.RegisterTypeEncoder(tOID, bsoncodec.ValueEncoderFunc(ObjectIDEncodeValue))
	rb.RegisterTypeDecoder(tOID, bsoncodec.ValueDecoderFunc(ObjectIDDecodeValue))

	rb.RegisterCodec(TimestampRegistry, TimestampCodecRef)

	return rb
}
