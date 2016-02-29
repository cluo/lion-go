/*
Package thriftlion defines the Thrift functionality for lion.
*/
package thriftlion // import "go.pedge.io/lion/thrift"

import "go.pedge.io/lion"

var (
	// Encoding is the name of the encoding.
	Encoding = "thrift"
)

func init() {
	if err := lion.RegisterEncoderDecoder(Encoding, newEncoderDecoder()); err != nil {
		panic(err.Error())
	}
}
