// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"

	medium "github.com/medium/medium-sdk-go"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var listCommand = cli.Command{
	Name:   "list",
	Usage:  "lists to your Medium stories",
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
	i := 1
	for _, post := range detail.Payload.References.Post {
		u := usr.URL + "/" + post.UniqueSlug
		buf.WriteString(fmt.Sprintf("(%d) %s: %s\n", i, post.Title, u))
		i++
	}
	fmt.Printf(buf.String())

	return nil
}
