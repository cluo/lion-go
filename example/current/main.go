/*
Package main implements an example integration with Current.

https://current.sh

Set the CURRENT_TOKEN and CURRENT_SYSLOG_ADDRESS environment variables to match the output of current syslog -n <org>.
Only plaintext connections are supported for now.
To output the debug messages, set the LOG_LEVEL environment variable to DEBUG.
Run using make run.
*/
package main

import (
	"fmt"
	"os"

	"go.pedge.io/lion/env"
	"go.pedge.io/lion/proto"
)

func main() {
	if err := do(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
	if err := envlion.Setup(); err != nil {
		return err
	}
	for i := 1; i < 5; i++ {
		protolion.Info(
			&Foo{
				Bar: &Bar{
					One: "one",
				},
				Two:   fmt.Sprintf("two%d", i),
				Three: uint64(i),
			},
		)
		protolion.Debugf("hello%d", i*i*i)
	}
	return nil
}
