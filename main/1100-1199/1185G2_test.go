// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1185/G2
// https://codeforces.com/problemset/status/1185/problem/G2
func Test_cf1185G2(t *testing.T) {
	testCases := [][2]string{
		{
			`3 3
1 1
1 2
1 3`,
			`6`,
		},
		{
			`3 3
1 1
1 1
1 3`,
			`2`,
		},
		{
			`4 10
5 3
2 1
3 2
5 1`,
			`10`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1185G2)
}
