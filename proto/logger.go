package protolion

import (
	"go.pedge.io/lion"

	"github.com/golang/protobuf/proto"
)

type logger struct {
	lion.Logger
	l lion.Level
}

func newLogger(delegate lion.Logger) *logger {
	return &logger{delegate, delegate.Level()}
}

func (l *logger) AtLevel(level lion.Level) Logger {
	return newLogger(l.Logger.AtLevel(level))
}

func (l *logger) WithField(key string, value interface{}) Logger {
	return newLogger(l.Logger.WithField(key, value))
}

func (l *logger) WithFields(fields map[string]interface{}) Logger {
	return newLogger(l.Logger.WithFields(fields))
}

func (l *logger) WithKeyValues(keyValues ...interface{}) Logger {
	return newLogger(l.Logger.WithKeyValues(keyValues...))
}

func (l *logger) WithContext(context proto.Message) Logger {
	return newLogger(l.WithEntryMessageContext(newEntryMessage(context)))
}

func (l *logger) Debug(event proto.Message) {
	if lion.LevelDebug < l.l {
		return
	}
	l.LogEntryMessage(lion.LevelDebug, newEntryMessage(event))
}

func (l *logger) Info(event proto.Message) {
	if lion.LevelInfo < l.l {
		return
	}
	l.LogEntryMessage(lion.LevelInfo, newEntryMessage(event))
}

func (l *logger) Warn(event proto.Message) {
	if lion.LevelWarn < l.l {
		return
	}
	l.LogEntryMessage(lion.LevelWarn, newEntryMessage(event))
}

func (l *logger) Error(event proto.Message) {
	if lion.LevelError < l.l {
		return
	}
	l.LogEntryMessage(lion.LevelError, newEntryMessage(event))
}

func (l *logger) Fatal(event proto.Message) {
	if lion.LevelFatal < l.l {
		return
	}
	l.LogEntryMessage(lion.LevelFatal, newEntryMessage(event))
}

func (l *logger) Panic(event proto.Message) {
	if lion.LevelPanic < l.l {
		return
	}
	l.LogEntryMessage(lion.LevelPanic, newEntryMessage(event))
}

func (l *logger) Print(event proto.Message) {
	l.LogEntryMessage(lion.LevelNone, newEntryMessage(event))
}

func (l *logger) LionLogger() lion.Logger {
	return l.Logger
}

func newEntryMessage(message proto.Message) *lion.EntryMessage {
	return &lion.EntryMessage{
		Encoding: Encoding,
		Value:    message,
	}
}
