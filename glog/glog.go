/*
Package lion_glog defines functionality for integration with glog.
*/
package lion_glog // import "go.pedge.io/lion/glog"

import (
	"flag"

	"go.pedge.io/lion"
)

var (
	// DefaultTextMarshaller is the default text Marshaller for glog.
	DefaultTextMarshaller = lion.NewTextMarshaller(
		lion.TextMarshallerDisableTime(),
		lion.TextMarshallerDisableLevel(),
	)
)

// PusherOption is an option for constructing a new Pusher.
type PusherOption func(*pusher)

// PusherWithMarshaller uses the Marshaller for the Pusher.
//
// By default, DefaultTextMarshaller is used.
func PusherWithMarshaller(marshaller lion.Marshaller) PusherOption {
	return func(pusher *pusher) {
		pusher.marshaller = marshaller
	}
}

// NewPusher constructs a new Pusher that pushes to glog.
//
// Note that glog is only global, so two glog Pushers push to the same source.
// If using glog, it is recommended register one glog Pusher as the global lion.Logger.
func NewPusher(options ...PusherOption) lion.Pusher {
	return newPusher(options...)
}

// LogToStderr sets the -logtostderr flag.
func LogToStderr() error {
	return flag.Set("logtostderr", "true")
}

// AlsoLogToStderr sets the -alsologtostderr flag.
func AlsoLogToStderr() error {
	return flag.Set("alsologtostderr", "true")
}
