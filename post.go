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
			Usage: "article status",
		},
		cli.StringSliceFlag{
			Name:  "tags",
			Usage: "article tags",
		},
	},
	Before: initPost,
	Action: runPost,
}

var createOption medium.CreatePostOptions

var (
	postFilename string
	postTitle    string
	postStatus   string
	postTags     []string
)

func initPost(ctx *cli.Context) error {
	postFilename = ctx.Args().First()
	postTitle = ctx.String("title")
	postStatus = ctx.String("status")
	postTags = ctx.StringSlice("tags")

	return nil
}

func runPost(ctx *cli.Context) error {
	if err := checkArgs(ctx, 1, exactArgs, "markdown file"); err != nil {
		return err
	}
	_, err := os.Stat(postFilename)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.Errorf("not fonud %s markdown file", postFilename)
		}
		return errors.Wrapf(err, "could not stat %s", postFilename)
	}

	// TODO(zchee): support CanonicalURL
	// TODO(zchee): support License config
	// Wait for medium-sdk-go exported several internal types.
	// https://github.com/Medium/medium-sdk-go/pull/17
	createOption = medium.CreatePostOptions{
		ContentFormat: medium.ContentFormatMarkdown,
		Tags:          postTags,
	}

	m := medium.NewClientWithAccessToken(os.Getenv("MEDIUM_SECRET_ACCESS_KEY"))
	usr, err := m.GetUser("")
	if err != nil {
		return errors.Wrap(err, "could not get medium user information")
	}
	createOption.UserID = usr.ID

	if postTitle == "" {
		return errors.New("title is empty")
	}
	createOption.Title = postTitle

	switch postStatus {
	case "draft":
		createOption.PublishStatus = medium.PublishStatusDraft
	case "unlisted":
		createOption.PublishStatus = medium.PublishStatusUnlisted
	case "public":
		createOption.PublishStatus = medium.PublishStatusPublic
	default:
		return errors.New("unknown post status")
	}

	buf, err := ioutil.ReadFile(postFilename)
	if err != nil {
		return err
	}
	createOption.Content = string(buf)

	_, err = m.CreatePost(createOption)
	if err != nil {
		return err
	}

	return nil
}
