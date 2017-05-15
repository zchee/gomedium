// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"net/url"
	"path"

	medium "github.com/medium/medium-sdk-go"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var listCommand = cli.Command{
	Name:   "list",
	Usage:  "lists to your posted Medium stories",
	Action: runList,
}

func runList(ctx *cli.Context) error {
	token, err := readToken()
	if err != nil {
		return err
	}
	m := medium.NewClientWithAccessToken(token)
	usr, err := m.GetUser("")
	if err != nil {
		return errors.Wrap(err, "could not get medium user information")
	}

	detail, err := GetUserDetail(usr)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	for i, post := range detail.PostReferences() {
		u, err := url.Parse(usr.URL)
		if err != nil {
			return errors.Wrapf(err, "failed to parse get url: '%s'", usr.URL)
		}
		u.Path = path.Join(u.Path, post.UniqueSlug)
		buf.WriteString(fmt.Sprintf("(%d) %s: %s\n", i, post.Title, u))
	}
	fmt.Printf(buf.String())

	return nil
}
