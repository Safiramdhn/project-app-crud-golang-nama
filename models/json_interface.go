package models

type JsonCodec interface {
	JsonEncode()
	JsonDecode()
}
