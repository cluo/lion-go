/*
Package lion_logrus defines functionality for integration with Logrus.
*/
package lion_logrus // import "go.pedge.io/lion/logrus"

import (
	"io"

	"github.com/Sirupsen/logrus"

	"go.pedge.io/lion"
)

// PusherOptions defines options for constructing a new Logrus lion.Pusher.
type PusherOptions struct {
	Out             io.Writer
	Hooks           []logrus.Hook
	Formatter       logrus.Formatter
	DisableContexts bool
}

// NewPusher creates a new lion.Pusher that logs using Logrus.
func NewPusher(options PusherOptions) lion.Pusher {
	return newPusher(options)
}

// RedirectStandardLogger redirects the logrus StandardLogger to lion's global Logger instance.
func RedirectStandardLogger() {
	lion.AddGlobalHook(
		func(logger lion.Logger) {
			logrus.SetOutput(logger.Writer())
		},
	)
}
