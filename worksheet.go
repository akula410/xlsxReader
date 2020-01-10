package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type Worksheet struct {
	Workbook  *Workbook
	Info      *WbSheet
	XMLName   xml.Name `xml:"worksheet"`
	SheetData []Row    `xml:"sheetData>row"`
	CellMap   map[int]map[int]*Col
}

type Row struct {
	Worksheet *Worksheet
	RowID     int    `xml:"r,attr"`
	Cols      []*Col `xml:"c"`
}

type CellFunction struct {
	Aca   bool   `xml:"aca,attr"`
	Value string `xml:",chardata"`
}

type CellVale struct {
	Value string `xml:",chardata"`
}

type Col struct {
	Row *Row
	R   string       `xml:"r,attr"`
	S   int          `xml:"s,attr"`
	T   string       `xml:"t,attr"`
	V   CellVale     `xml:"v"`
	F   CellFunction `xml:"f"`
}

func (c *Col) GetString() string {
 if c.T == "s" {
		idx, _ := strconv.Atoi(c.V.Value)

		if idx>len(c.Row.Worksheet.Workbook.Xlsx.SST.SSTitems){
			fmt.Println("out of range SST")
		}

		return c.Row.Worksheet.Workbook.Xlsx.SST.SSTitems[idx].GetString()
	}

	return c.V.Value

}

func (rw *Row) GetCols() (cols []*Col) {
	var maxColKey int

	for _, c := range rw.Cols {
		colStr := c.R[:len(c.R)-len(strconv.Itoa(rw.RowID)) ]
		if colNum, err := ColumnStrIdxToNumIdx(colStr); err == nil {
			if colNum > maxColKey {
				maxColKey = colNum
			}
		}
	}
	rCols := rw.Worksheet.CellMap[rw.RowID]
	for i := 1; i <= maxColKey; i++ {

		if c, ok := rCols[i]; ok {
			cols = append(cols, c)
		} else {
			cols = append(cols, &Col{
				Row: rw,
				R:   fmt.Sprintf("%s%d",ColumnNumIdxToStrIdx(i),rw.RowID),
				S:   0,
				T:   "",
				V:   CellVale{},
				F:   CellFunction{},
			})
		}

	}

	return cols
}

func NewWorksheet(partName string, wbSheet *WbSheet, xlsx *Xlsx) (ws *Worksheet, err error) {
	ws = &Worksheet{}
	ws.Workbook = xlsx.Workbook
	ws.Info = wbSheet

	f, err := xlsx.Files[partName].Open()

	if err != nil {
		return nil, err
	}

	defer f.Close()

	d := xml.NewDecoder(f)

	err = d.Decode(ws)

	if err != nil {
		return nil, err
	}

	ws.CellMap = map[int]map[int]*Col{}

	for k, row := range ws.SheetData {
		ws.SheetData[k].Worksheet = ws
		ws.CellMap[row.RowID] = map[int]*Col{}

		for r, cell := range row.Cols {
			row.Cols[r].Row = &ws.SheetData[k]
			colStr := cell.R[:len(cell.R)-len(strconv.Itoa(row.RowID)) ]
			if colNum, err := ColumnStrIdxToNumIdx(colStr); err == nil {
				cell.Row.Worksheet = ws
				ws.CellMap[row.RowID][colNum] = cell
			} else {
				return nil, err
			}
		}
	}

	return ws, err
}
