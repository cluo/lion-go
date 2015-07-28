/*
Package protolog defines the main protolog functionality.
*/
package protolog // import "go.pedge.io/protolog"

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/peter-edge/go-google-protobuf"
)

var (
	globalLogger = NewStandardLogger(NewFileFlusher(os.Stderr))
	globalLock   = &sync.Mutex{}
)

// GlobalLogger returns the global Logger instance.
func GlobalLogger() Logger {
	return globalLogger
}

// SetLogger sets the global Logger instance.
func SetLogger(logger Logger) {
	globalLock.Lock()
	defer globalLock.Unlock()
	globalLogger = logger
}

// SetLevel sets the global Logger to to be at the given Level.
func SetLevel(level Level) {
	globalLock.Lock()
	defer globalLock.Unlock()
	globalLogger = globalLogger.AtLevel(level)
}

// Message is a proto.Message that also has the ProtologName() method to get the name of the message.
type Message interface {
	proto.Message
	ProtologName() string
}

// Flusher is an object that can be flushed to a persistent store.
type Flusher interface {
	Flush() error
}

// WriteFlusher is an io.Writer that can be flushed.
type WriteFlusher interface {
	io.Writer
	Flusher
}

// Fields are a representation of key/value fields without using protocol buffers.
type Fields map[string]interface{}

// Logger is the main logging interface. All methods are also replicated
// on the package and attached to a global Logger.
type Logger interface {
	Flusher

	AtLevel(level Level) Logger

	WithContext(context Message) Logger
	Debug(event Message)
	Info(event Message)
	Warn(event Message)
	Error(event Message)
	Fatal(event Message)
	Panic(event Message)
	Print(event Message)

	DebugWriter() io.Writer
	InfoWriter() io.Writer
	WarnWriter() io.Writer
	ErrorWriter() io.Writer
	Writer() io.Writer

	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

// Pusher is the interface used to push Entry objects to a persistent store.
type Pusher interface {
	Flusher
	Push(entry *Entry) error
}

// IDAllocator allocates unique IDs for Entry objects. The default
// behavior is to allocate a new UUID for the process, then add an
// incremented integer to the end.
type IDAllocator interface {
	Allocate() string
}

// Timer returns the current time. The default behavior is to
// call time.Now().UTC().
type Timer interface {
	Now() time.Time
}

// ErrorHandler handles errors when logging. The default behavior
// is to panic.
type ErrorHandler interface {
	Handle(err error)
}

// LoggerOptions defines options for the Logger constructor.
type LoggerOptions struct {
	IDAllocator  IDAllocator
	Timer        Timer
	ErrorHandler ErrorHandler
	// MarshalFunc is the function used to serialize protobuf Messages.
	// The default behavior is to call proto.Marshal.
	MarshalFunc func(proto.Message) ([]byte, error)
}

// NewLogger constructs a new Logger using the given Pusher.
func NewLogger(pusher Pusher, options LoggerOptions) Logger {
	return newLogger(pusher, options)
}

// NewStandardLogger constructs a new Logger that logs using a text Marshaller.
func NewStandardLogger(writeFlusher WriteFlusher) Logger {
	return NewLogger(
		NewWritePusher(
			writeFlusher,
			WritePusherOptions{
				Marshaller: NewTextMarshaller(MarshallerOptions{}),
			},
		),
		LoggerOptions{},
	).AtLevel(
		Level_LEVEL_INFO,
	)
}

// Marshaller marshals Entry objects to be written.
type Marshaller interface {
	Marshal(entry *Entry) ([]byte, error)
}

// Encoder encodes a marshalled Entry. This is used for adding metadata
// if the marshalled Entry is being sent over the wire.
type Encoder interface {
	Encode(p []byte) ([]byte, error)
}

// WritePusherOptions defines options for constructing a new write Pusher.
type WritePusherOptions struct {
	Marshaller Marshaller
	Encoder    Encoder
}

// NewWritePusher constructs a new Pusher that writes to the given WriteFlusher.
func NewWritePusher(writeFlusher WriteFlusher, options WritePusherOptions) Pusher {
	return newWritePusher(writeFlusher, options)
}

// Puller pulls Entry objects from a persistent store.
type Puller interface {
	Pull() (*Entry, error)
}

// Unmarshaller unmarshalls a marshalled Entry object.
type Unmarshaller interface {
	Unmarshal(p []byte, entry *Entry) error
}

// Decoder decodes Entry objects from an io.Reader. At the end
// of a stream, Decoder will return io.EOF.
type Decoder interface {
	Decode(reader io.Reader) ([]byte, error)
}

// ReadPullerOptions defines options for a read Puller.
type ReadPullerOptions struct {
	Unmarshaller  Unmarshaller
	UnmarshalFunc func([]byte, proto.Message) error
}

// NewReadPuller constructs a new Puller that reads from the given Reader
// and decodes using the given Decoder.
func NewReadPuller(reader io.Reader, decoder Decoder, options ReadPullerOptions) Puller {
	return newReadPuller(reader, decoder, options)
}

// MarshallerOptions provides options for creating Marshallers.
type MarshallerOptions struct {
	// EnableID will add the printing of Entry IDs.
	EnableID bool
	// DisableTimestamp will suppress the printing of Entry Timestamps.
	DisableTimestamp bool
	// DisableLevel will suppress the printing of Entry Levels.
	DisableLevel bool
	// DisableContexts will suppress the printing of Entry contexts.
	DisableContexts bool
	// UnmarshalFunc is the function used to unmarshal previously
	// marshalled protobuf Messages. The default behavior is to use
	// proto.Unmarshal
	UnmarshalFunc func([]byte, proto.Message) error
}

// NewTextMarshaller constructs a new Marshaller that produces human-readable
// marshalled Entry objects. This Marshaller is current inefficient.
func NewTextMarshaller(options MarshallerOptions) Marshaller {
	return newTextMarshaller(options)
}

// NewWireEncoder constructs an Encoder that produces values which can
// be sent over the wire.
func NewWireEncoder() Encoder {
	return wireEncoderInstance
}

// NewWireDecoder constructs a Decoder that can decode values encoded
// with a wire Encoder.
func NewWireDecoder() Decoder {
	return wireDecoderInstance
}

// NewWriterFlusher wraps an io.Writer into a WriteFlusher.
// Flush() is a no-op on the returned WriteFlusher.
func NewWriterFlusher(writer io.Writer) WriteFlusher {
	return newWriterFlusher(writer)
}

// FileFlusher wraps an *os.File into a Flusher.
// Flush() will call Sync().
type FileFlusher struct {
	*os.File
}

// NewFileFlusher constructs a new FileFlusher for the given *os.File.
func NewFileFlusher(file *os.File) *FileFlusher {
	return &FileFlusher{file}
}

// Flush calls Sync.
func (f *FileFlusher) Flush() error {
	return f.Sync()
}

// UnmarshalledContexts returns the context Messages marshalled on an Entry object.
func (m *Entry) UnmarshalledContexts(unmarshalFunc func([]byte, proto.Message) error) ([]Message, error) {
	return entryMessagesToMessages(m.Context, unmarshalFunc)
}

// UnmarshalledEvent returns the event Message marshalled on an Entry object.
func (m *Entry) UnmarshalledEvent(unmarshalFunc func([]byte, proto.Message) error) (Message, error) {
	return entryMessageToMessage(m.Event, unmarshalFunc)
}

// TimeToTimestamp is a utility function that converts a time.Time into
// a *google_protobuf.Timestamp.
func TimeToTimestamp(t time.Time) *google_protobuf.Timestamp {
	unixNano := t.UnixNano()
	return &google_protobuf.Timestamp{
		Seconds: unixNano / int64(time.Second),
		Nanos:   int32(unixNano % int64(time.Second)),
	}
}

// TimestampToTime is a utility function that converts a *google_protobuf.Timestamp
// into a time.Time.
func TimestampToTime(timestamp *google_protobuf.Timestamp) time.Time {
	return time.Unix(timestamp.Seconds, int64(timestamp.Nanos)).UTC()
}

// Flush calls Flush on the global Logger.
func Flush() error {
	return globalLogger.Flush()
}

// AtLevel calls AtLevel on the global Logger.
func AtLevel(level Level) Logger {
	return globalLogger.AtLevel(level)
}

// WithContext calls WithContext on the global Logger.
func WithContext(context Message) Logger {
	return globalLogger.WithContext(context)
}

// Debug calls Debug on the global Logger.
func Debug(Event Message) {
	globalLogger.Debug(Event)
}

// Info calls Info on the global Logger.
func Info(Event Message) {
	globalLogger.Info(Event)
}

// Warn calls Warn on the global Logger.
func Warn(Event Message) {
	globalLogger.Warn(Event)
}

// Error calls Error on the global Logger.
func Error(Event Message) {
	globalLogger.Error(Event)
}

// Fatal calls Fatal on the global Logger.
func Fatal(Event Message) {
	globalLogger.Fatal(Event)
}

// Panic calls Panic on the global Logger.
func Panic(Event Message) {
	globalLogger.Panic(Event)
}

// Print calls Print on the global Logger.
func Print(Event Message) {
	globalLogger.Print(Event)
}

// DebugWriter calls DebugWriter on the global Logger.
func DebugWriter() io.Writer {
	return globalLogger.DebugWriter()
}

// InfoWriter calls InfoWriter on the global Logger.
func InfoWriter() io.Writer {
	return globalLogger.InfoWriter()
}

// WarnWriter calls WarnWriter on the global Logger.
func WarnWriter() io.Writer {
	return globalLogger.WarnWriter()
}

// ErrorWriter calls ErrorWriter on the global Logger.
func ErrorWriter() io.Writer {
	return globalLogger.ErrorWriter()
}

// Writer calls Writer on the global Logger.
func Writer() io.Writer {
	return globalLogger.Writer()
}

// WithField calls WithField on the global Logger.
func WithField(key string, value interface{}) Logger {
	return globalLogger.WithField(key, value)
}

// WithFields calls WithFields on the global Logger.
func WithFields(fields Fields) Logger {
	return globalLogger.WithFields(fields)
}

// Debugf calls Debugf on the global Logger.
func Debugf(format string, args ...interface{}) {
	globalLogger.Debugf(format, args...)
}

// Debugln calls Debugln on the global Logger.
func Debugln(args ...interface{}) {
	globalLogger.Debugln(args...)
}

// Infof calls Infof on the global Logger.
func Infof(format string, args ...interface{}) {
	globalLogger.Infof(format, args...)
}

// Infoln calls Infoln on the global Logger.
func Infoln(args ...interface{}) {
	globalLogger.Infoln(args...)
}

// Warnf calls Warnf on the global Logger.
func Warnf(format string, args ...interface{}) {
	globalLogger.Warnf(format, args...)
}

// Warnln calls Warnln on the global Logger.
func Warnln(args ...interface{}) {
	globalLogger.Warnln(args...)
}

// Errorf calls Errorf on the global Logger.
func Errorf(format string, args ...interface{}) {
	globalLogger.Errorf(format, args...)
}

// Errorln calls Errorln on the global Logger.
func Errorln(args ...interface{}) {
	globalLogger.Errorln(args...)
}

// Fatalf calls Fatalf on the global Logger.
func Fatalf(format string, args ...interface{}) {
	globalLogger.Fatalf(format, args...)
}

// Fatalln calls Fatalln on the global Logger.
func Fatalln(args ...interface{}) {
	globalLogger.Fatalln(args...)
}

// Panicf calls Panicf on the global Logger.
func Panicf(format string, args ...interface{}) {
	globalLogger.Panicf(format, args...)
}

// Panicln calls Panicln on the global Logger.
func Panicln(args ...interface{}) {
	globalLogger.Panicln(args...)
}

// Printf calls Printf on the global Logger.
func Printf(format string, args ...interface{}) {
	globalLogger.Printf(format, args...)
}

// Println calls Println on the global Logger.
func Println(args ...interface{}) {
	globalLogger.Println(args...)
}
