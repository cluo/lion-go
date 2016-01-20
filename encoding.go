package lion

// *** this file is for everything to do with encoding and decoding ***

import "fmt"

var (
	encodingToEncoderDecoderPair = make(map[string]*encoderDecoderPair, 0)
)

// *** EncodedPusherToPusher ***

type encodedPusherToPusherWrapper struct {
	EncodedPusher
}

func newEncodedPusherToPusherWrapper(encodedPusher EncodedPusher) *encodedPusherToPusherWrapper {
	return &encodedPusherToPusherWrapper{encodedPusher}
}

func (e *encodedPusherToPusherWrapper) Push(entry *Entry) error {
	encodedEntry, err := entry.Encode()
	if err != nil {
		return err
	}
	return e.EncodedPusher.Push(encodedEntry)
}

// *** registration ***

type encoderDecoderPair struct {
	encoder Encoder
	decoder Decoder
}

// TODO(pedge): rw lock?
func registerEncoderDecoder(encoding string, encoderDecoder EncoderDecoder) error {
	if err := checkNoRegisteredEncoding(encoding); err != nil {
		return err
	}
	encodingToEncoderDecoderPair[encoding] = &encoderDecoderPair{
		encoder: encoderDecoder,
		decoder: encoderDecoder,
	}
	return nil
}

func getEncoder(encoding string) (Encoder, error) {
	encoderDecoderPair, err := getEncoderDecoderPair(encoding)
	if err != nil {
		return nil, err
	}
	if encoderDecoderPair.encoder == nil {
		return nil, fmt.Errorf("lion: encoding %s has no encoder", encoding)
	}
	return encoderDecoderPair.encoder, nil
}

func getDecoder(encoding string) (Decoder, error) {
	encoderDecoderPair, err := getEncoderDecoderPair(encoding)
	if err != nil {
		return nil, err
	}
	if encoderDecoderPair.decoder == nil {
		return nil, fmt.Errorf("lion: encoding %s has no decoder", encoding)
	}
	return encoderDecoderPair.decoder, nil
}

func getEncoderDecoderPair(encoding string) (*encoderDecoderPair, error) {
	encoderDecoderPair, ok := encodingToEncoderDecoderPair[encoding]
	if !ok {
		return nil, fmt.Errorf("lion: encoding %s not registered", encoding)
	}
	return encoderDecoderPair, nil
}

func checkRegisteredEncoding(encoding string) error {
	if _, ok := encodingToEncoderDecoderPair[encoding]; !ok {
		return fmt.Errorf("lion: encoding %s not registered", encoding)
	}
	return nil
}

func checkNoRegisteredEncoding(encoding string) error {
	if _, ok := encodingToEncoderDecoderPair[encoding]; ok {
		return fmt.Errorf("lion: encoding %s already registered", encoding)
	}
	return nil
}

/// *** util ***

func encodeEntry(entry *Entry) (*EncodedEntry, error) {
	encodedContexts, err := encodeEntryMessages(entry.Contexts)
	if err != nil {
		return nil, err
	}
	encodedEvent, err := encodeEntryMessage(entry.Event)
	if err != nil {
		return nil, err
	}
	return &EncodedEntry{
		ID:           entry.ID,
		Level:        entry.Level,
		Time:         entry.Time,
		Contexts:     encodedContexts,
		Fields:       entry.Fields,
		Event:        encodedEvent,
		Message:      entry.Message,
		WriterOutput: entry.WriterOutput,
	}, nil
}

func encodeEntryMessages(entryMessages []*EntryMessage) ([]*EncodedEntryMessage, error) {
	encodedEntryMessages := make([]*EncodedEntryMessage, len(entryMessages))
	for i, entryMessage := range entryMessages {
		encodedEntryMessage, err := encodeEntryMessage(entryMessage)
		if err != nil {
			return nil, err
		}
		encodedEntryMessages[i] = encodedEntryMessage
	}
	return encodedEntryMessages, nil
}

func encodeEntryMessage(entryMessage *EntryMessage) (*EncodedEntryMessage, error) {
	encoder, err := getEncoder(entryMessage.Encoding)
	if err != nil {
		return nil, err
	}
	return encoder.Encode(entryMessage)
}

func decodeEncodedEntry(encodedEntry *EncodedEntry) (*Entry, error) {
	contexts, err := decodeEncodedEntryMessages(encodedEntry.Contexts)
	if err != nil {
		return nil, err
	}
	event, err := decodeEncodedEntryMessage(encodedEntry.Event)
	if err != nil {
		return nil, err
	}
	return &Entry{
		ID:           encodedEntry.ID,
		Level:        encodedEntry.Level,
		Time:         encodedEntry.Time,
		Contexts:     contexts,
		Fields:       encodedEntry.Fields,
		Event:        event,
		Message:      encodedEntry.Message,
		WriterOutput: encodedEntry.WriterOutput,
	}, nil
}

func decodeEncodedEntryMessages(encodedEntryMessages []*EncodedEntryMessage) ([]*EntryMessage, error) {
	entryMessages := make([]*EntryMessage, len(encodedEntryMessages))
	for i, encodedEntryMessage := range encodedEntryMessages {
		entryMessage, err := decodeEncodedEntryMessage(encodedEntryMessage)
		if err != nil {
			return nil, err
		}
		entryMessages[i] = entryMessage
	}
	return entryMessages, nil
}

func decodeEncodedEntryMessage(encodedEntryMessage *EncodedEntryMessage) (*EntryMessage, error) {
	decoder, err := getDecoder(encodedEntryMessage.Encoding)
	if err != nil {
		return nil, err
	}
	return decoder.Decode(encodedEntryMessage)
}

func entryMessageName(entryMessage *EntryMessage) (string, error) {
	encoder, err := getEncoder(entryMessage.Encoding)
	if err != nil {
		return "", err
	}
	return encoder.Name(entryMessage)
}
