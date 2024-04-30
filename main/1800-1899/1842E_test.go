// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1842/E
// https://codeforces.com/problemset/status/1842/problem/E
func Test_cf1842E(t *testing.T) {
	testCases := [][2]string{
		{
			`4 6 1
1 2 1
2 1 1
1 1 1
3 2 6`,
			`4`,
		},
		{
			`6 7 1
4 2 1
3 3 1
5 1 4
3 2 5
4 1 1
0 6 4`,
			`4`,
		},
		{
			`10 4 100
0 0 1
0 1 1
0 2 50
0 3 200
1 0 1
1 1 1
1 2 1
2 0 200
2 1 200
3 0 200`,
			`355`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1842E)
}