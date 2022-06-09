package pb

//// NewObjectId creates proto ObjectId from MongoDB ObjectID
//func NewObjectId(id primitive.ObjectID) *ObjectId {
//	return &ObjectId{Value: id.Hex()}
//}
//
////GetObjectID returns MongoDB object ID from struct
//func (x *User) GetObjectID() (primitive.ObjectID, error) {
//	return primitive.ObjectIDFromHex(x.Id.Value)
//}
//
////GetHexID returns MongoDB hexadecimal ID from struct
//func (x *User) GetHexID() string {
//	if x.Id == nil {
//		return ""
//	}
//	return x.Id.Value
//}
//
////GetObjectID returns MongoDB object ID from field
//func (x *ObjectId) GetObjectID() (primitive.ObjectID, error) {
//	return primitive.ObjectIDFromHex(x.Value)
//}
//
////GetHexID returns MongoDB hexadecimal ID from field
//func (x *ObjectId) GetHexID() string {
//	return x.Value
//}
