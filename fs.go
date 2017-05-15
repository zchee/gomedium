// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
	medium "github.com/medium/medium-sdk-go"
	"github.com/pkg/errors"
	"github.com/zchee/go-xdgbasedir"
)

var (
	// configDir directory of gomedium config basedy XDG Base Directory spec.
	configDir = filepath.Join(xdgbasedir.ConfigHome(), "gomedium")
	// configFile file of gomedium configs.
	configFile = filepath.Join(configDir, "config.yml")
)

type configSyntax struct {
	License medium.License
}

func openConfig() (*os.File, error) {
	if _, err := os.Stat(configFile); err != nil {
		if err := os.MkdirAll(configDir, 0700); err != nil {
			return nil, err
		}
		return os.Create(configFile)
	}
	return os.Open(configFile)
}

func readConfig(key string) (string, error) {
	f, err := openConfig()
	if err != nil {
		return "", err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	config := new(configSyntax)
	if err := yaml.Unmarshal(data, config); err != nil {
		return "", err
	}

	switch key {
	case "license":
		return string(config.License), nil
	default:
		return "", errors.New("unknown config key")
	}
}

var (
	// tokenFile file of gomedium authenticates token.
	tokenFile = filepath.Join(configDir, "token.yml")
)

type tokenSyntax struct {
	Token string
}

func readToken() (string, error) {
	out, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return "", errors.New(`Unable to locate token. You can configure token by running "gomedium login".`)
	}
	token := new(tokenSyntax)
	if err := yaml.Unmarshal(out, token); err != nil {
		return "", err
	}
	return token.Token, nil
}
