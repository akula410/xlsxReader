package core

import "testing"

func Test_reader(t *testing.T)  {
		_, err := openFile("../test/test_1.xlsx")

		if err!=nil {
			t.Error(err)
		}
}
