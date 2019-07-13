package go_commons

import (
	"errors"
	"github.com/golang/protobuf/proto"
)

var (
	ErrorNilMessage = errors.New("Nil messsage")
)

func EncodeProto(msg proto.Message) ([]byte, error) {
	if msg == nil {
		return nil, ErrorNilMessage
	}
	var b []byte
	b, err := proto.Marshal(msg)
	return b, err
}

func DecodeProto(b []byte, output proto.Message) error {
	if b == nil || output == nil {
		return ErrorNilMessage
	}
	err := proto.Unmarshal(b, output)
	return err
}
