## xlsxReader

Библиотека для чтения xlsx файлов / Golang read xlsx library

## Установка / Installation

`$ go get github.com/akula410/xlsxReader`

## Использование / Using

```go
xlsx, err := OpenFile("file.xlsx")

if err != nil {
    panic(err)
}

for _, ws := range xlsx.Worksheets {
	for _, row := range ws.SheetData {
		for _, cell := range row.GetCols() {
			fmt.Println(cell.GetString())
		}			
	}
}
```
 