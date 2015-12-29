package lion_glog

import (
	"github.com/golang/glog"
	"go.pedge.io/lion"
)

var (
	levelToLogFunc = map[lion.Level]func(...interface{}){
		lion.LevelNone:  glog.Infoln,
		lion.LevelDebug: glog.Infoln,
		lion.LevelInfo:  glog.Infoln,
		lion.LevelWarn:  glog.Warningln,
		lion.LevelError: glog.Errorln,
		lion.LevelFatal: glog.Errorln,
		lion.LevelPanic: glog.Errorln,
	}
)

type pusher struct {
	marshaller lion.Marshaller
}

func newPusher(options ...PusherOption) *pusher {
	pusher := &pusher{DefaultTextMarshaller}
	for _, option := range options {
		option(pusher)
	}
	return pusher
}

func (p *pusher) Flush() error {
	glog.Flush()
	return nil
}

func (p *pusher) Push(entry *lion.Entry) error {
	data, err := p.marshaller.Marshal(entry)
	if err != nil {
		return err
	}
	levelToLogFunc[entry.Level](string(data))
	return nil
}
