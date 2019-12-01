package core

import (
	"encoding/xml"
	"errors"
)

type Workbook struct {
	XMLName     xml.Name    `xml:"workbook"`
	FileVersion FileVersion `xml:"fileVersion"`
	WorkbookPr  WorkbookPr  `xml:"workbookPr"`
	Sheets      []WbSheet   `xml:"sheets>sheet"`
}
type FileVersion struct {
	AppName string `xml:"appName,attr"`
}

type WorkbookPr struct {
	BackupFile  bool   `xml:"backupFile,attr"`
	ShowObjects string `xml:"showObjects,attr"`
	Date1904    bool   `xml:"date1904,attr"`
}

type WbSheet struct {
	Name    string `xml:"name,attr"`
	SheetID string `xml:"sheetId,attr"`
	State   string `xml:"state,attr"`
	RelID   string `xml:"id,attr"`
}

func NewWoorkbook(xlsx *Xlsx) (wb *Workbook, err error) {

	wb = &Workbook{}

	if partName, ok := xlsx.ContentTypes.GetPartNameByType(WorkbookContentType); ok {
		f, err := xlsx.Files[partName[0]].Open()

		if err != nil {
			return nil, err
		}

		defer f.Close()

		d := xml.NewDecoder(f)

		err = d.Decode(wb)

		if err != nil {
			return nil, err
		}

		return wb, err
	}

	return nil, errors.New("Woorkbook not found")

}
