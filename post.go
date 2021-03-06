// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"io/ioutil"
	"os"

	medium "github.com/medium/medium-sdk-go"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var postCommand = cli.Command{
	Name:      "post",
	Usage:     "post the article",
	ArgsUsage: "<markdown file>",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "title, t",
			Usage: "article title",
		},
		cli.StringFlag{
			Name:  "status, s",
			Usage: "article status [draft, unlisted, public]",
		},
		cli.StringSliceFlag{
			Name:  "tags",
			Usage: "article tags",
		},
	},
	Before: initPost,
	Action: runPost,
}

var (
	postFilename string
	postTitle    string
	postStatus   string
	postTags     []string
)

func initPost(ctx *cli.Context) error {
	if err := checkArgs(ctx, 1, exactArgs, "markdown file"); err != nil {
		return err
	}

	postFilename = ctx.Args().First()
	postTitle = ctx.String("title")
	postStatus = ctx.String("status")
	postTags = ctx.StringSlice("tags")

	_, err := os.Stat(postFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.Errorf("not fonud %s markdown file", postFilename)
		}
		return errors.Wrapf(err, "could not stat %s", postFilename)
	}

	if postTitle == "" {
		return errors.New("title is empty")
	}

	switch postStatus {
	case "draft", "unlisted", "public":
		// nothing to do
	case "":
		postStatus = "draft" // defalut is "draft"
	default:
		return errors.New("unknown post status")
	}

	return nil
}

func runPost(ctx *cli.Context) error {
	buf, err := ioutil.ReadFile(postFilename)
	if err != nil {
		return err
	}

	token, err := readToken()
	if err != nil {
		return err
	}
	m := medium.NewClientWithAccessToken(string(token))
	usr, err := m.GetUser("")
	if err != nil {
		return errors.Wrap(err, "could not get medium user information")
	}

	// TODO(zchee): support CanonicalURL
	createOption := medium.CreatePostOptions{
		UserID:        usr.ID,
		Title:         postTitle,
		Content:       string(buf),
		ContentFormat: medium.ContentFormatMarkdown,
		Tags:          postTags,
		PublishStatus: medium.PublishStatus(postStatus),
	}
	license, err := readConfig("license")
	if err != nil {
		return err
	}
	createOption.License = medium.License(license)

	_, err = m.CreatePost(createOption)
	if err != nil {
		return err
	}

	return nil
}
