// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	medium "github.com/medium/medium-sdk-go"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var loginCommand = cli.Command{
	Name:  "login",
	Usage: "authenticates against the Medium API and stores the token",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "token, t",
			Usage:  "Medium secret token",
			EnvVar: "MEDIUM_SECRET_ACCESS_KEY",
		},
	},
	Before: initLogin,
	Action: runLogin,
}

var (
	loginToken string
)

func initLogin(ctx *cli.Context) error {
	loginToken = ctx.String("token")
	if loginToken == "" {
		return errors.New("token is empty")
	}
	return nil
}

func runLogin(ctx *cli.Context) error {
	if _, err := os.Stat(tokenFile); err == nil {
		log.Print("already logged in")
		return nil
	}

	m := medium.NewClientWithAccessToken(os.Getenv("MEDIUM_SECRET_ACCESS_KEY"))
	usr, err := m.GetUser("")
	if err != nil {
		return err
	}

	token := tokenSyntax{Token: loginToken}
	data, err := yaml.Marshal(token)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return err
	}
	if err := ioutil.WriteFile(tokenFile, data, 0600); err != nil {
		return err
	}

	log.Printf("Successfully logged in as %s (%s)", usr.Username, usr.Name)

	return nil
}
