// +build cgo

package zstd

import (
	"github.com/DataDog/zstd"
	"github.com/segmentio/kafka-go"
)

func init() {
	kafka.RegisterCompressionCodec(func() kafka.CompressionCodec {
		return NewCompressionCodec()
	})
}

type CompressionCodec struct {
	// CompressionLevel is the level of compression to use on messages.
	CompressionLevel int
}

const Code = 4

func NewCompressionCodec() CompressionCodec {
	return CompressionCodec{}
}

// Code implements the kafka.CompressionCodec interface.
func (c CompressionCodec) Code() int8 {
	return Code
}

// Encode implements the kafka.CompressionCodec interface.
func (c CompressionCodec) Encode(src []byte) ([]byte, error) {
	return zstd.CompressLevel(nil, src, c.CompressionLevel)
}

// Decode implements the kafka.CompressionCodec interface.
func (c CompressionCodec) Decode(src []byte) ([]byte, error) {
	return zstd.Decompress(nil, src)
}
