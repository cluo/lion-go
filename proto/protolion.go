package protolion // import "go.pedge.io/lion/proto"

var (
	// DelimitedMarshaller is a Marshaller that uses the protocol buffers write delimited scheme.
	DelimitedMarshaller = &delimitedMarshaller{}
	// DelimitedUnmarshaller is an Unmarshaller that uses the protocol buffers write delimited scheme.
	DelimitedUnmarshaller = &delimitedUnmarshaller{}
)
