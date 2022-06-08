package codecs

import "go.mongodb.org/mongo-driver/bson/bsoncodec"

// Register registers Google protocol buffers types codecs
func Register(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	return rb.RegisterCodec(ObjectIDRegistry, ObjectIDCodecRef).
		RegisterCodec(TimestampRegistry, TimestampCodecRef)
}
