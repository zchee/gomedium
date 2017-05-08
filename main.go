// Copyright 2017 The gmedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli"
)

func init() {
	log.SetPrefix("gmedium: ")
}

func main() {
	app := cli.NewApp()
	app.Name = "gmedium"
	app.Usage = "A command line tool for Medium stories."
	app.Version = fmt.Sprintf("%s (%s)", version, gitCommit)
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{}
	app.ErrWriter = &fatalWriter{cli.ErrWriter}
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

type fatalWriter struct {
	cliErrWriter io.Writer
}

func (f *fatalWriter) Write(b []byte) (n int, err error) {
	return f.cliErrWriter.Write(b)
}
