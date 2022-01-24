// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"math/bits"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[1,2,3,4]`,
			`6`,
		},
		{
			`[4,2,3,15]`,
			`5`,
		},
		{
			`[1]`,
			`0`,
		},
		{
			`[18,28,2,17,29,30,15,9,12]`,
			`19`,
		},
		{
			`[5,10,1,26,24,21,24,23,11,12,27,4,17,16,2,6,1,1,6,8,13,30,24,20,2,19]`,
			`5368`,
		},
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, numberOfGoodSubsets, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode-cn.com/contest/biweekly-contest-60/problems/the-number-of-good-subsets/

func TestCompareInf(t *testing.T) {
	inputGenerator := func() (a []int) {
		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 5)
		a = rg.IntSlice(n, 1, 3)
		return
	}

	runAC := func(a []int) (ans int) {
		calc := func(sub uint) bool {
			n := 1
			for sub := sub; sub > 0; sub &= sub - 1 {
				p := bits.TrailingZeros(sub)
				v := a[p]
				n *= v
			}
			if n == 1 {
				return false
			}
			for _, p := range primes {
				if n%p == 0 {
					c := 0
					for n%p == 0 {
						n /= p
						c++
					}
					if c > 1 {
						return false
					}
				}
			}
			return true
		}
		for sub := uint(0); sub < 1<<len(a); sub++ {
			res := calc(sub)
			if res {
				ans++
			}
		}
		return
	}

	// test examples first (or make it global)
	//examples := [][]string{
	//	{
	//		`[1,2,3,4]`,
	//		`6`,
	//	},
	//	{
	//		`[4,2,3,15]`,
	//		`5`,
	//	},
	//	{
	//		`[1]`,
	//		`0`,
	//	},
	//	{
	//		`[18,28,2,17,29,30,15,9,12]`,
	//		`19`,
	//	},
	//}
	//if err := testutil.RunLeetCodeFuncWithExamples(t, runAC, examples, 0); err != nil {
	//	t.Fatal(err)
	//}
	//return

	testutil.CompareInf(t, inputGenerator, runAC, numberOfGoodSubsets)
}