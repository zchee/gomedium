// Copyright 2017 The gomedium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	medium "github.com/medium/medium-sdk-go"
	"github.com/pkg/errors"
	"github.com/zchee/gomedium/internal/api"
)

const (
	siteEndpoint = "https://medium.com/"
)

// GetUserDetail gets more detailed user information than the official package.
func GetUserDetail(usr *medium.User) (*api.User, error) {
	u, err := url.Parse(siteEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse get url: '%s'", siteEndpoint)
	}
	u.Path = "@" + usr.Username
	u.RawQuery = "format=json"

	ctx := context.Background()
	data, err := request(ctx, u.String())
	if err != nil {
		return nil, err
	}

	usrDetail := new(api.User)
	if err := json.Unmarshal(data, usrDetail); err != nil {
		return nil, err
	}

	return usrDetail, nil
}

func request(pctx context.Context, urlStr string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, errors.New("failed to create new Get request")
	}

	ctx, cancel := context.WithCancel(pctx)
	defer cancel()

	client := http.DefaultClient
	req = req.WithContext(ctx)
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request to medium")
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// trim the first extra content that prevents JSON hacking (])}while(1);</x>)
	// https://medium.com/statuscode/building-a-basic-web-service-to-display-your-medium-blog-posts-on-your-website-using-aws-api-48597b1771c5
	data = bytes.TrimPrefix(data, []byte("])}while(1);</x>"))

	return data, nil
}
