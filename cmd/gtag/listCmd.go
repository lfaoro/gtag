// Copyright (c) 2019 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

type tagInfo struct {
	tag    string
	commit string
	tagger string
	time   string
}

var listCmd = cli.Command{
	Name:    "list",
	Aliases: []string{"l", "lis"},
	Usage:   "list all git tags",
	Flags:   []cli.Flag{},
	Action: func(c *cli.Context) error {
		out, err := shellCmd("git --no-pager tag -l")
		if err != nil {
			return errors.Wrap(err, "unable to list tags")
		}
		s := strings.TrimSpace(out)
		if s == "" {
			return errors.New("repo has no tags, create one w/ `gtag zero`")
		}
		tags := strings.Split(s, "\n")

		var tagData [][]string
		for _, tag := range tags {
			data := dataFor(tag)
			tagData = append(tagData, data)
		}

		var tagInfo []tagInfo
		for _, data := range tagData {
			info := genTagInfoFrom(data)
			tagInfo = append(tagInfo, info)
		}

		for _, d := range tagInfo {
			if d.tag == "" {
				continue
			}
			output := fmt.Sprintf("tag %v [%v] by %v on %v",
				d.tag, d.commit, d.tagger, d.time)
			fmt.Println(output)
		}

		return nil
	},
}

func dataFor(tag string) []string {
	cmd := fmt.Sprintf("git cat-file tag %v", tag)
	out, err := shellCmd(cmd)
	if err != nil {
		if debugFlag {
			fmt.Println(err)
		}
		fmt.Printf("tag %s has no data available\n", tag)
		return []string{}
	}

	tagInfo := strings.Split(out, "\n")
	return tagInfo
}

func genTagInfoFrom(data []string) tagInfo {
	var tag = tagInfo{}
	for _, el := range data {
		if strings.HasPrefix(el, "object") {
			ss := strings.Split(el, " ")
			commit := ss[1]
			// fmt.Println("commit: ", ss[1]) // d846485eb58703b448fb26317e9da541c452bd06
			tag.commit = commit[:8]
		}
		if strings.HasPrefix(el, "tagger") {
			ss := strings.Split(el, " ")
			// fmt.Println("name:", ss[1]) // Leonardo
			tag.tagger = ss[1]
			// fmt.Println("time: ", ss[len(ss)-2]) // 1565197008
			i, err := strconv.Atoi(ss[len(ss)-2])
			exitIfError(err)

			tagTime := time.Unix(int64(i), 0)
			// 2006-01-02T15:04:05
			format := "January 02 2006 at 15:04:05"
			t := tagTime.Format(format)
			tag.time = t
		}
		if strings.HasPrefix(el, "tag ") {
			ss := strings.Split(el, " ")
			// fmt.Println(ss[1]) // 0.3.0
			tag.tag = ss[1]
		}
	}
	return tag
}
