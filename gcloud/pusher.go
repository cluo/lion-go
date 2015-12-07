package gcloud

import (
	"net/http"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"go.pedge.io/google-protobuf"
	"go.pedge.io/protolog"
	"google.golang.org/api/logging/v1beta3"
)

const customServiceName = "compute.googleapis.com"

var (
	marshaler = &jsonpb.Marshaler{}

	// https://cloud.google.com/logging/docs/api/ref/rest/v1beta3/projects.logs.entries/write#LogSeverity
	severityName = map[protolog.Level]string{
		protolog.Level_LEVEL_NONE:  "DEFAULT",
		protolog.Level_LEVEL_DEBUG: "DEBUG",
		protolog.Level_LEVEL_INFO:  "INFO",
		protolog.Level_LEVEL_WARN:  "WARNING",
		protolog.Level_LEVEL_ERROR: "ERROR",
		protolog.Level_LEVEL_FATAL: "ERROR",
		protolog.Level_LEVEL_PANIC: "ALERT",
	}
)

type pusher struct {
	service   *logging.ProjectsLogsEntriesService
	projectId string
	logName   string
}

func newPusher(client *http.Client, projectId string, logName string) *pusher {
	service, err := logging.New(client)
	if err != nil {
		panic(err)
	}
	return &pusher{service.Projects.Logs.Entries, projectId, logName}
}

func (p *pusher) Push(entry *protolog.Entry) error {
	logEntry, err := p.newLogEntry(entry)
	if err != nil {
		return err
	}
	request := p.service.Write(
		p.projectId,
		p.logName,
		&logging.WriteLogEntriesRequest{
			Entries: []*logging.LogEntry{logEntry},
		},
	)
	_, err = request.Do()
	return err
}

func (p *pusher) Flush() error {
	return nil
}

func (p *pusher) newLogEntry(entry *protolog.Entry) (*logging.LogEntry, error) {
	payload, err := p.marshalEntry(entry)
	if err != nil {
		return nil, err
	}

	return &logging.LogEntry{
		InsertId:    entry.Id,
		TextPayload: payload,
		Metadata: &logging.LogEntryMetadata{
			ServiceName: customServiceName,
			Severity:    severityName[entry.Level],
			Timestamp:   newTimestamp(entry.Timestamp),
		},
	}, nil
}

func newTimestamp(timestamp *google_protobuf.Timestamp) string {
	return time.Unix(
		timestamp.Seconds,
		int64(timestamp.Nanos),
	).Format(time.RFC3339)
}

func (p *pusher) marshalEntry(entry *protolog.Entry) (string, error) {
	return marshaler.MarshalToString(entry)
}
