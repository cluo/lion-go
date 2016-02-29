package thriftlion

import (
	"fmt"
	"reflect"
	"sync"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var (
	nameToConstructor = make(map[string]func() thrift.TStruct)
	lock              sync.RWMutex
)

// MustRegister calls Register and panics on error.
func MustRegister(constructor func() thrift.TStruct) {
	if err := Register(constructor); err != nil {
		panic(err.Error())
	}
}

// Register registers the given thrift.TStruct using the generated constructor.
func Register(constructor func() thrift.TStruct) error {
	lock.Lock()
	defer lock.Unlock()

	name := getName(constructor())
	if _, ok := nameToConstructor[name]; ok {
		return fmt.Errorf("thriftlion: duplicate name %s", name)
	}
	nameToConstructor[name] = constructor
	return nil
}

func newTStruct(name string) (thrift.TStruct, error) {
	constructor, err := getConstructor(name)
	if err != nil {
		return nil, err
	}
	return constructor(), nil
}

func getConstructor(name string) (func() thrift.TStruct, error) {
	lock.RLock()
	defer lock.RUnlock()

	constructor, ok := nameToConstructor[name]
	if !ok {
		return nil, fmt.Errorf("thriftlion: unknown name: %s", name)
	}
	return constructor, nil
}

func getName(tStruct thrift.TStruct) string {
	reflectType := reflect.TypeOf(tStruct)
	return fmt.Sprintf("%s.%s", reflectType.PkgPath(), reflectType.Name())
}
