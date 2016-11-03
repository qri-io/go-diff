// Copyright (c) 2012 Sergi Mansilla <sergi.mansilla@gmail.com>
// https://github.com/sergi/go-diff
// See the included LICENSE file for license details.
//
// go-diff is a Go implementation of Google's Diff, Match, and Patch library
// Original library is Copyright (c) 2006 Google Inc.
// http://code.google.com/p/google-diff-match-patch/

package diffmatchpatch

import (
	"io/ioutil"
)

func speedtestTexts() (s1 string, s2 string) {
	d1, err := ioutil.ReadFile("../testdata/speedtest1.txt")
	if err != nil {
		panic(err)
	}
	d2, err := ioutil.ReadFile("../testdata/speedtest2.txt")
	if err != nil {
		panic(err)
	}

	return string(d1), string(d2)
}
