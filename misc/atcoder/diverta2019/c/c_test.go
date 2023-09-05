// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"strings"
	"testing"
)

// 题目：https://atcoder.jp/contests/diverta2019/tasks/diverta2019_c
// 提交：https://atcoder.jp/contests/diverta2019/submit?taskScreenName=diverta2019_c
// 对拍：https://atcoder.jp/contests/diverta2019/submissions?f.LanguageName=Go&f.Status=AC&f.Task=diverta2019_c&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`3
ABCA
XBAZ
BAD`,
			`2`,
		},
		{
			`9
BEWPVCRWH
ZZNQYIJX
BAVREA
PA
HJMYITEOX
BCJHMRMNK
BP
QVFABZ
PRGKSPUNA`,
			`4`,
		},
		{
			`7
RABYBBE
JOZ
BMHQUVA
BPA
ISU
MCMABAOBHZ
SZMEHMA`,
			`4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}

func TestCompare(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 5)
		rg.NewLine()
		for i := 0; i < n; i++ {
			rg.Str(2, 5, 'A', 'C')
			rg.NewLine()
		}
		return rg.String()
	}

	permutations := func(n, r int, do func(ids []int) (Break bool)) {
		ids := make([]int, n)
		for i := range ids {
			ids[i] = i
		}
		if do(ids[:r]) {
			return
		}
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					tmp := ids[i]
					copy(ids[i:], ids[i+1:])
					ids[n-1] = tmp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					ids[i], ids[n-j] = ids[n-j], ids[i]
					if do(ids[:r]) {
						return
					}
					break
				}
			}
			if i == -1 {
				return
			}
		}
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n int
		Fscan(in, &n)
		a := make([]string, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		permutations(n, n, func(ids []int) (Break bool) {
			s := ""
			for _, i := range ids {
				s += a[i]
			}
			c := strings.Count(s, "AB")
			if c > ans {
				ans = c
			}
			return
		})
		Fprintln(out, ans)
	}
	
	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, run)
}