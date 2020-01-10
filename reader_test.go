package main

import (
	"fmt"
	"testing"
)

func Test_reader(t *testing.T) {

	xlsx, err := OpenFile("./test/test_1.xlsx")

	if err != nil {
		t.Fatal(err)
	}

	for _, ws := range xlsx.Worksheets {
		for _, row := range ws.SheetData {

			for _, cell := range row.GetCols() {
				fmt.Printf("%s: %s | ", cell.R, cell.GetString())
			}
			fmt.Println()
			fmt.Println("------------------------------------------------")
		}
	}

}
