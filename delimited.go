package lion

import (
	"bytes"
	"fmt"
	"io"
	"reflect"

	"go.pedge.io/lion/pb"
	"go.pedge.io/proto/time"

	"github.com/golang/protobuf/proto"
	"github.com/matttproud/golang_protobuf_extensions/pbutil"
)

var (
	levelToPB = map[Level]lionpb.Level{
		LevelNone:  lionpb.Level_LEVEL_NONE,
		LevelDebug: lionpb.Level_LEVEL_DEBUG,
		LevelInfo:  lionpb.Level_LEVEL_INFO,
		LevelWarn:  lionpb.Level_LEVEL_WARN,
		LevelError: lionpb.Level_LEVEL_ERROR,
		LevelFatal: lionpb.Level_LEVEL_FATAL,
		LevelPanic: lionpb.Level_LEVEL_PANIC,
	}
	pbToLevel = map[lionpb.Level]Level{
		lionpb.Level_LEVEL_NONE:  LevelNone,
		lionpb.Level_LEVEL_DEBUG: LevelDebug,
		lionpb.Level_LEVEL_INFO:  LevelInfo,
		lionpb.Level_LEVEL_WARN:  LevelWarn,
		lionpb.Level_LEVEL_ERROR: LevelError,
		lionpb.Level_LEVEL_FATAL: LevelFatal,
		lionpb.Level_LEVEL_PANIC: LevelPanic,
	}
)

type delimitedMarshaller struct{}

func (m *delimitedMarshaller) Marshal(entry *Entry) ([]byte, error) {
	pbEntry, err := entryToPBEntry(entry)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(nil)
	if _, err := pbutil.WriteDelimited(buffer, pbEntry); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

type delimitedUnmarshaller struct{}

func (u *delimitedUnmarshaller) Unmarshal(reader io.Reader, entry *Entry) error {
	pbEntry := &lionpb.Entry{}
	if _, err := pbutil.ReadDelimited(reader, pbEntry); err != nil {
		return err
	}
	iEntry, err := pbEntryToEntry(pbEntry)
	if err != nil {
		return err
	}
	*entry = *iEntry
	return nil
}

func entryToPBEntry(entry *Entry) (*lionpb.Entry, error) {
	contexts, err := messagesToEntryMessages(entry.Contexts)
	if err != nil {
		return nil, err
	}
	event, err := messageToEntryMessage(entry.Event)
	if err != nil {
		return nil, err
	}
	pbLevel, ok := levelToPB[entry.Level]
	if !ok {
		return nil, fmt.Errorf("lion: unknown level: %v", entry.Level)
	}
	return &lionpb.Entry{
		Id:           entry.ID,
		Level:        pbLevel,
		Timestamp:    prototime.TimeToTimestamp(entry.Time),
		Context:      contexts,
		Fields:       entry.Fields,
		Event:        event,
		Message:      entry.Message,
		WriterOutput: entry.WriterOutput,
	}, nil
}

func pbEntryToEntry(pbEntry *lionpb.Entry) (*Entry, error) {
	contexts, err := entryMessagesToMessages(pbEntry.Context)
	if err != nil {
		return nil, err
	}
	event, err := entryMessageToMessage(pbEntry.Event)
	if err != nil {
		return nil, err
	}
	level, ok := pbToLevel[pbEntry.Level]
	if !ok {
		return nil, fmt.Errorf("lion: unknown level: %v", pbEntry.Level)
	}
	return &Entry{
		ID:           pbEntry.Id,
		Level:        level,
		Time:         prototime.TimestampToTime(pbEntry.Timestamp),
		Contexts:     contexts,
		Fields:       pbEntry.Fields,
		Event:        event,
		Message:      pbEntry.Message,
		WriterOutput: pbEntry.WriterOutput,
	}, nil
}

// NOTE: the jsoonpb.Marshaler was EPICALLY SLOW in benchmarks
// When using the stdlib json.Marshal function instead for the text Marshaller,
// a speedup of 6X was observed!

func messageToEntryMessage(message proto.Message) (*lionpb.Entry_Message, error) {
	if message == nil {
		return nil, nil
	}
	value, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return &lionpb.Entry_Message{
		Name:  messageName(message),
		Value: value,
	}, nil
}

func entryMessageToMessage(entryMessage *lionpb.Entry_Message) (proto.Message, error) {
	if entryMessage == nil {
		return nil, nil
	}
	message, err := newMessage(entryMessage.Name)
	if err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(entryMessage.Value, message); err != nil {
		return nil, err
	}
	return message, nil
}

func messagesToEntryMessages(messages []proto.Message) ([]*lionpb.Entry_Message, error) {
	if messages == nil {
		return nil, nil
	}
	entryMessages := make([]*lionpb.Entry_Message, len(messages))
	for i, message := range messages {
		entryMessage, err := messageToEntryMessage(message)
		if err != nil {
			return nil, err
		}
		entryMessages[i] = entryMessage
	}
	return entryMessages, nil
}

func entryMessagesToMessages(entryMessages []*lionpb.Entry_Message) ([]proto.Message, error) {
	if entryMessages == nil {
		return nil, nil
	}
	messages := make([]proto.Message, len(entryMessages))
	for i, entryMessage := range entryMessages {
		message, err := entryMessageToMessage(entryMessage)
		if err != nil {
			return nil, err
		}
		messages[i] = message
	}
	return messages, nil
}

func newMessage(name string) (proto.Message, error) {
	reflectType := proto.MessageType(name)
	if reflectType == nil {
		return nil, fmt.Errorf("lion: no Message registered for name: %s", name)
	}

	return reflect.New(reflectType.Elem()).Interface().(proto.Message), nil
}

func messageName(message proto.Message) string {
	if message == nil {
		return ""
	}
	return proto.MessageName(message)
}
