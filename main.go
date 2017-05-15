// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func init() {
	log.SetPrefix("gomedium: ")
	log.SetFlags(0)
}

func main() {
	app := cli.NewApp()
	app.Name = "gomedium"
	app.Usage = "A command line tool for Medium stories."
	app.Version = fmt.Sprintf("%s (%s)", version, gitCommit)
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		configCommand,
		listCommand,
		loginCommand,
		postCommand,
	}
	app.ErrWriter = &fatalWriter{cli.ErrWriter}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type fatalWriter struct {
	cliErrWriter io.Writer
}

func (f *fatalWriter) Write(b []byte) (n int, err error) {
	return f.cliErrWriter.Write(b)
}

const (
	exactArgs = iota
	minArgs
	maxArgs
)

func checkArgs(ctx *cli.Context, expected, checkType int, args ...string) error {
	cmdName := ctx.Command.FullName()
	var err error
	switch checkType {
	case exactArgs:
		if ctx.NArg() != expected {
			err = fmt.Errorf("%q command requires exactly <%s> %d argument(s)", cmdName, strings.Join(args, " "), expected)
		}
	case minArgs:
		if ctx.NArg() < expected {
			err = fmt.Errorf("%q command requires a minimum of <%s> %d argument(s)", cmdName, strings.Join(args, " "), expected)
		}
	case maxArgs:
		if ctx.NArg() > expected {
			err = fmt.Errorf("%q command requires a maximum of <%s> %d argument(s)", cmdName, strings.Join(args, " "), expected)
		}
	}

	if err != nil {
		return err
	}
	return nil
}
