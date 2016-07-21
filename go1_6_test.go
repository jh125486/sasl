// Copyright 2016 Sam Whited.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

// +build !go1.7

package sasl

import (
	"fmt"
	"testing"
)

// doTests runs all tests without using t.Run. As a result, context may be
// missing, but at least all tests are run.
func doTests(t *testing.T, fn func(t *testing.T, tc saslTest)) {
	for _, g := range testCases {
		for i, tc := range g.cases {
			name := fmt.Sprintf("%s:%d:%s", g.name, i, tc.client.mechanism.Name)
			t.Log("Testing ", name)
			fn(t, tc)
		}
	}
}
