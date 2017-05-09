// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"github.com/zchee/go-xdgbasedir"
)

var (
	// configDir directory of gomedium config basedy XDG Base Directory spec.
	configDir = filepath.Join(xdgbasedir.ConfigHome(), "gomedium")
	// configFile file of gomedium configs.
	configFile = filepath.Join(configDir, "config.yml")
)

type configSyntax struct{}

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
	yaml.Unmarshal(out, token)
	return token.Token, nil
}
