package protolion // import "go.pedge.io/lion/proto"

import "sync"

var (
	// Encoding is the name of the encoding.
	Encoding = "proto"

	// DelimitedMarshaller is a Marshaller that uses the protocol buffers write delimited scheme.
	DelimitedMarshaller = &delimitedMarshaller{}
	// DelimitedUnmarshaller is an Unmarshaller that uses the protocol buffers write delimited scheme.
	DelimitedUnmarshaller = &delimitedUnmarshaller{}

	globalPrimaryPackage     = "golang"
	globalSecondaryPackage   = "gogo"
	globalOnlyPrimaryPackage = true
	globalLock               = &sync.Mutex{}
)

// GolangFirst says to check both golang and gogo for message names and types, but golang first.
func GolangFirst() {
	globalLock.Lock()
	defer globalLock.Unlock()
	globalPrimaryPackage = "golang"
	globalSecondaryPackage = "gogo"
	globalOnlyPrimaryPackage = false
}

// GolangOnly says to check only golang for message names and types, but not gogo.
func GolangOnly() {
	globalLock.Lock()
	defer globalLock.Unlock()
	globalPrimaryPackage = "golang"
	globalSecondaryPackage = "gogo"
	globalOnlyPrimaryPackage = true
}

// GogoFirst says to check both gogo and golang for message names and types, but gogo first.
func GogoFirst() {
	globalLock.Lock()
	defer globalLock.Unlock()
	globalPrimaryPackage = "gogo"
	globalSecondaryPackage = "golang"
	globalOnlyPrimaryPackage = false
}

// GogoOnly says to check only gogo for message names and types, but not golang.
func GogoOnly() {
	globalLock.Lock()
	defer globalLock.Unlock()
	globalPrimaryPackage = "gogo"
	globalSecondaryPackage = "golang"
	globalOnlyPrimaryPackage = true
}
