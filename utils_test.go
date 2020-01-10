package main

import (
	"errors"
	"fmt"
	"testing"
)

func Test_utils(t *testing.T)  {

	var testLtrs = map[string]int{
		"A":1,
		"BA":53,
		"AAA":703,
		"AMJ":1024,
	}

	for k, v := range testLtrs {
		num,err := ColumnStrIdxToNumIdx(k)
		idx := ColumnNumIdxToStrIdx(num)

		if err!=nil {
			t.Error(err)
		}

		if num !=v {
			t.Error(errors.New("fail converter test"))
		}

		if idx !=k {
			fmt.Println(idx)
			fmt.Println(k)
			t.Error(errors.New("fail converter test"))
		}
	}



}
