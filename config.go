// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"strings"

	"github.com/go-yaml/yaml"
	medium "github.com/medium/medium-sdk-go"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var configCommand = cli.Command{
	Name:      "config",
	Usage:     "set gomedium config",
	ArgsUsage: "key=value",
	Before:    initConfig,
	Action:    runConfig,
}

var (
	configKey   string
	configValue string
)

func initConfig(ctx *cli.Context) error {
	if err := checkArgs(ctx, 1, exactArgs, "key=value"); err != nil {
		return err
	}

	configs := strings.SplitN(ctx.Args().First(), "=", 2)
	if len(configs) <= 1 {
		return errors.New("invalid key=value argument")
	}
	configKey = configs[0]
	configValue = configs[1]

	return nil
}

func runConfig(ctx *cli.Context) error {
	fc, err := openConfig()
	if err != nil {
		return err
	}
	defer fc.Close()

	config := new(configSyntax)
	if stat, err := fc.Stat(); err == nil && stat.Size() > 0 {
		data, err := ioutil.ReadAll(fc)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(data, config); err != nil {
			return err
		}
	}

	switch configKey {
	case "license":
		switch medium.License(configValue) {
		case medium.LicenseAllRightsReserved:
		case medium.LicenseCC40By:
		case medium.LicenseCC40ByNC:
		case medium.LicenseCC40ByNCND:
		case medium.LicenseCC40ByNCSA:
		case medium.LicenseCC40ByND:
		case medium.LicenseCC40BySA:
		case medium.LicenseCC40Zero:
		case medium.LicensePublicDomain:
			// nothing to do
		default:
			return errors.Errorf("unknown license: %s. available: [all-rights-reserved, cc-40-by, cc-40-by-sa, cc-40-by-nd, cc-40-by-nc, cc-40-by-nc-nd, cc-40-by-nc-sa, cc-40-zero, public-domain]", configValue)
		}
		config.License = medium.License(configValue)
	default:
		return errors.Errorf("unknown config key: %s", configKey)
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(configFile, data, 0600); err != nil {
		return err
	}

	return nil
}
