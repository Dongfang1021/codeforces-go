// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[1,2,3,4]`, `[2]`, 
			`false`,
		},
		{
			`[1,2,3,3]`, `[2]`, 
			`true`,
		},
		{
			`[1,1,2,2]`, `[2,2]`, 
			`true`,
		},
		{
			`[1,1,2,3]`, `[2,2]`, 
			`false`,
		},
		{
			`[1,1,1,1,1]`, `[2,3]`, 
			`true`,
		},
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, canDistribute, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-39/problems/distribute-repeating-integers/