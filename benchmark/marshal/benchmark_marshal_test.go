package lion_benchmark_marshal

import (
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"go.pedge.io/lion"
	"go.pedge.io/lion/testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkDelimitedMarshaller(b *testing.B) {
	benchmarkMarshaller(b, lion.DelimitedMarshaller)
}

func BenchmarkDefaultTextMarshaller(b *testing.B) {
	benchmarkMarshaller(b, lion.NewTextMarshaller())
}

func benchmarkMarshaller(b *testing.B, marshaller lion.Marshaller) {
	b.StopTimer()
	entry := getBenchEntry()
	_, err := marshaller.Marshal(entry)
	require.NoError(b, err)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = marshaller.Marshal(entry)
	}
}

func getBenchEntry() *lion.Entry {
	foo := &lion_testing.Foo{
		StringField: "one",
		Int32Field:  2,
	}
	bar := &lion_testing.Bar{
		StringField: "one",
		Int32Field:  2,
	}
	baz := &lion_testing.Baz{
		Bat: &lion_testing.Baz_Bat{
			Ban: &lion_testing.Baz_Bat_Ban{
				StringField: "one",
				Int32Field:  2,
			},
		},
	}
	entry := &lion.Entry{
		ID:    "123",
		Level: lion.LevelInfo,
		Time:  time.Now().UTC(),
		Contexts: []proto.Message{
			foo,
			bar,
		},
		Event: baz,
	}
	return entry
}
