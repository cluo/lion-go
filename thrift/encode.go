package thriftlion

import "go.pedge.io/lion"

type encoderDecoder struct{}

func newEncoderDecoder() *encoderDecoder {
	return &encoderDecoder{}
}

func (e *encoderDecoder) Encode(entryMessage *lion.EntryMessage) (*lion.EncodedEntryMessage, error) {
	return nil, nil
}

func (e *encoderDecoder) Name(entryMessage *lion.EntryMessage) (string, error) {
	return "", nil
}

func (e *encoderDecoder) Decode(encodedEntryMessage *lion.EncodedEntryMessage) (*lion.EntryMessage, error) {
	return nil, nil
}
