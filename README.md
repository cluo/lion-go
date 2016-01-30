[![CircleCI](https://circleci.com/gh/peter-edge/lion-go/tree/master.png)](https://circleci.com/gh/peter-edge/lion-go/tree/master)
[![Go Report Card](http://goreportcard.com/badge/peter-edge/lion-go)](http://goreportcard.com/report/peter-edge/lion-go)
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/go.pedge.io/lion)
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/peter-edge/lion-go/blob/master/LICENSE)

```shell
go get go.pedge.io/lion
```

Documentation coming soon, have not got around to it yet.

All public types are in [lion.go](lion.go) and [lion_level.go](lion_level.go).

In sub packages, are public types are in `nameofpackage.go`, and `nameofpackage.pb.go` in the case of `protolion`. Some of them need to be renamed,
ie [syslog/syslog.go](syslog/syslog.go) should be `syslog/sysloglion.go`, this is a holdover from the previous iteration of lion, called protolog.
