/*
Package logrus defines functionality for integration with Logrus.
*/
package logrus // import "go.pedge.io/protolog/logrus"

import (
	"github.com/Sirupsen/logrus"

	"go.pedge.io/protolog"
)

var (
	// DefaultPusher is the default logrus Pusher.
	DefaultPusher = NewPusher(PusherOptions{})
)

// PusherOptions defines options for constructing a new Logrus protolog.Pusher.
type PusherOptions struct {
	Out             protolog.WriteFlusher
	Hooks           []logrus.Hook
	Formatter       logrus.Formatter
	DisableContexts bool
	JSONMarshaller  protolog.JSONMarshaller
}

// NewPusher creates a new protolog.Pusher that logs using Logrus.
func NewPusher(options PusherOptions) protolog.Pusher {
	return newPusher(options)
}

// RedirectStandardLogger redirects the logrus StandardLogger to protolog's global Logger instance.
func RedirectStandardLogger() {
	protolog.AddGlobalHook(
		func(logger protolog.Logger) {
			logrus.SetOutput(logger.Writer())
		},
	)
}
