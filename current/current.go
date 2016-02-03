package currentlion // import "go.pedge.io/lion/current"

import (
	"log/syslog"

	"go.pedge.io/lion"
	"go.pedge.io/lion/syslog"
)

const (
	syslogNetwork = "tcp"
)

// NewPusher returns a new Pusher for current.
func NewPusher(
	appName string,
	syslogAddress string,
	token string,
) (lion.Pusher, error) {
	writer, err := syslog.Dial(
		syslogNetwork,
		syslogAddress,
		syslog.LOG_INFO,
		appName,
	)
	if err != nil {
		return nil, err
	}
	return sysloglion.NewPusher(
		writer,
		sysloglion.PusherWithMarshaller(
			newMarshaller(token),
		),
	), nil
}
