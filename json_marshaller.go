package lion

type jsonMarshaller struct{}

func newJSONMarshaller() *jsonMarshaller {
	return &jsonMarshaller{}
}

func (t *jsonMarshaller) Marshal(entry *Entry) ([]byte, error) {
	return jsonMarshalEntry(entry)
}

func jsonMarshalEntry(entry *Entry) ([]byte, error) {
	return nil, nil
}

//func jsonMarshalMessage(buffer *bytes.Buffer, entry *Entry) ([]byte, error) {
//}
