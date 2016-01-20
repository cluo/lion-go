package protolion // import "go.pedge.io/lion/proto"

import (
	"sync"

	"go.pedge.io/lion"

	"github.com/golang/protobuf/proto"
)

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
	globalLogger             = NewLogger(lion.GlobalLogger())
	globalLock               = &sync.Mutex{}
)

func init() {
	if err := lion.RegisterEncoderDecoder(Encoding, newEncoderDecoder()); err != nil {
		panic(err.Error())
	}
	lion.AddGlobalHook(setGlobalLogger)
}

func setGlobalLogger(logger lion.Logger) {
	globalLock.Lock()
	defer globalLock.Unlock()
	globalLogger = NewLogger(logger)
}

// Logger is a lion.Logger that also has proto logging methods.
type Logger interface {
	lion.Logger

	WithProtoContext(context proto.Message) Logger
	WithProtoField(key string, value interface{}) Logger
	WithProtoFields(fields map[string]interface{}) Logger

	ProtoDebug(event proto.Message)
	ProtoInfo(event proto.Message)
	ProtoWarn(event proto.Message)
	ProtoError(event proto.Message)
	ProtoFatal(event proto.Message)
	ProtoPanic(event proto.Message)
	ProtoPrint(event proto.Message)
}

// NewLogger returns a new Logger.
func NewLogger(delegate lion.Logger) Logger {
	return newLogger(delegate)
}

// WithContext calls WithProtoContext on the global Logger.
func WithContext(context proto.Message) Logger {
	return globalLogger.WithProtoContext(context)
}

// Debug calls ProtoDebug on the global Logger.
func Debug(event proto.Message) {
	globalLogger.ProtoDebug(event)
}

// Info calls ProtoInfo on the global Logger.
func Info(event proto.Message) {
	globalLogger.ProtoInfo(event)
}

// Warn calls ProtoWarn on the global Logger.
func Warn(event proto.Message) {
	globalLogger.ProtoWarn(event)
}

// Error calls ProtoError on the global Logger.
func Error(event proto.Message) {
	globalLogger.ProtoError(event)
}

// Fatal calls ProtoFatal on the global Logger.
func Fatal(event proto.Message) {
	globalLogger.ProtoFatal(event)
}

// Panic calls ProtoPanic on the global Logger.
func Panic(event proto.Message) {
	globalLogger.ProtoPanic(event)
}

// Print calls ProtoPrint on the global Logger.
func Print(event proto.Message) {
	globalLogger.ProtoPrint(event)
}

//// GolangFirst says to check both golang and gogo for message names and types, but golang first.
//func GolangFirst() {
//globalLock.Lock()
//defer globalLock.Unlock()
//globalPrimaryPackage = "golang"
//globalSecondaryPackage = "gogo"
//globalOnlyPrimaryPackage = false
//}

//// GolangOnly says to check only golang for message names and types, but not gogo.
//func GolangOnly() {
//globalLock.Lock()
//defer globalLock.Unlock()
//globalPrimaryPackage = "golang"
//globalSecondaryPackage = "gogo"
//globalOnlyPrimaryPackage = true
//}

//// GogoFirst says to check both gogo and golang for message names and types, but gogo first.
//func GogoFirst() {
//globalLock.Lock()
//defer globalLock.Unlock()
//globalPrimaryPackage = "gogo"
//globalSecondaryPackage = "golang"
//globalOnlyPrimaryPackage = false
//}

//// GogoOnly says to check only gogo for message names and types, but not golang.
//func GogoOnly() {
//globalLock.Lock()
//defer globalLock.Unlock()
//globalPrimaryPackage = "gogo"
//globalSecondaryPackage = "golang"
//globalOnlyPrimaryPackage = true
//}
