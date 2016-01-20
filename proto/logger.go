package protolion

import (
	"github.com/gogo/protobuf/proto"
	"go.pedge.io/lion"
)

type logger struct {
	lion.Logger
}

func newLogger(delegate lion.Logger) *logger {
	return &logger{delegate}
}

func (l *logger) WithProtoContext(context proto.Message) Logger {
	return newLogger(l.WithEntryMessageContext(newEntryMessage(context)))
}

func (l *logger) ProtoDebug(event proto.Message) {
	l.LogEntryMessage(lion.LevelDebug, newEntryMessage(event))
}

func (l *logger) ProtoInfo(event proto.Message) {
	l.LogEntryMessage(lion.LevelInfo, newEntryMessage(event))
}

func (l *logger) ProtoWarn(event proto.Message) {
	l.LogEntryMessage(lion.LevelWarn, newEntryMessage(event))
}

func (l *logger) ProtoError(event proto.Message) {
	l.LogEntryMessage(lion.LevelError, newEntryMessage(event))
}

func (l *logger) ProtoFatal(event proto.Message) {
	l.LogEntryMessage(lion.LevelFatal, newEntryMessage(event))
}

func (l *logger) ProtoPanic(event proto.Message) {
	l.LogEntryMessage(lion.LevelPanic, newEntryMessage(event))
}

func (l *logger) ProtoPrint(event proto.Message) {
	l.LogEntryMessage(lion.LevelNone, newEntryMessage(event))
}

func newEntryMessage(message proto.Message) *lion.EntryMessage {
	return &lion.EntryMessage{
		Encoding: Encoding,
		Value:    message,
	}
}
