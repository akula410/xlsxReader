package main

import (
	"errors"
	"math"
)

var cols = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"D": 4,
	"E": 5,
	"F": 6,
	"G": 7,
	"H": 8,
	"I": 9,
	"J": 10,
	"K": 11,
	"L": 12,
	"M": 13,
	"N": 14,
	"O": 15,
	"P": 16,
	"Q": 17,
	"R": 18,
	"S": 19,
	"T": 20,
	"U": 21,
	"V": 22,
	"W": 23,
	"X": 24,
	"Y": 25,
	"Z": 26,
}
var idx = map[int]string{
	1:  "A",
	2:  "B",
	3:  "C",
	4:  "D",
	5:  "E",
	6:  "F",
	7:  "G",
	8:  "H",
	9:  "I",
	10: "J",
	11: "K",
	12: "L",
	13: "M",
	14: "N",
	15: "O",
	16: "P",
	17: "Q",
	18: "R",
	19: "S",
	20: "T",
	21: "U",
	22: "V",
	23: "W",
	24: "X",
	25: "Y",
	26: "Z",
}

func ColumnStrIdxToNumIdx(strIndex string) (num int, err error) {

	runes := []rune(strIndex)

	for k, v := range runes {
		if n, ok := cols[string(v)]; ok {
			s := int(math.Pow(float64(len(cols)), float64(len(runes)-(k+1))))
			num = num + (n * s)
		} else {
			return num, errors.New("wrong column index")
		}
	}

	return num, nil
}

func ColumnNumIdxToStrIdx(numIndex int) (str string) {
	var mem []int
	n := numIndex
	r := len(idx)

	for {
		mem = append(mem, n%r)
		n = n / r
		if n < r {
			mem = append(mem, n)
			break
		}
	}

	for i := len(mem)-1; i >= 0; i-- {
		str = str + idx[mem[i]]
	}

	return str
}
