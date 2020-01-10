package main

import (
	"archive/zip"
	"errors"
	"strings"
)

const (
	WorkbookRelFile         = "xl/_rels/workbook.xml.rels"
	ContentTypeFile         = "[Content_Types].xml"
	RelationshipContentType = "application/vnd.openxmlformats-package.relationships+xml"
	WorkbookContentType     = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"
	StyleContentType        = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"
	SSTContentType          = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
)

var rels []string

type Xlsx struct {
	Files        map[string]*zip.File
	ContentTypes *ContentTypes
	Rels         map[string]*Relationships
	Workbook     *Workbook
	Worksheets   map[string]*Worksheet
	SST          *SST
	Style        *Style
}

func (xlsx *Xlsx) readRootRels() (err error) {

	xlsx.Rels = make(map[string]*Relationships)

	xlsx.ContentTypes, err = NewContentType(xlsx)

	if err != nil {
		return err
	}

	if partNames, ok := xlsx.ContentTypes.GetPartNameByType(RelationshipContentType); ok {
		for _, pN := range partNames {
			xlsx.Rels[pN] = new(Relationships)
			xlsx.Rels[pN], err = NewRelationships(pN, xlsx)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("wrong xlsx file")
	}

	xlsx.Workbook, err = NewWoorkbook(xlsx)

	if err != nil {
		return err
	}

	workbookRel := xlsx.Rels[WorkbookRelFile]
	xlsx.Worksheets = map[string]*Worksheet{}

	for k, sheet := range xlsx.Workbook.Sheets {
		if rl, ok := workbookRel.GetByID(sheet.RelID); ok {
			if ws, err := NewWorksheet("xl/"+strings.TrimLeft(strings.Replace(rl.Target, "xl/", "", 1), "/"), &xlsx.Workbook.Sheets[k], xlsx); err == nil {
				xlsx.Worksheets[sheet.RelID] = ws
			} else {
				return err
			}
		}
	}

	if len(xlsx.Worksheets) == 0 {
		err = errors.New("sheets not found")
	}

	if err != nil {
		return err
	}

	xlsx.SST, err = NewSST(xlsx)

	if err != nil {
		return err
	}

	xlsx.Style, err = NewStyle(xlsx)

	if err != nil {
		return err
	}

	return err

}
