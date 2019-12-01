package core

import (
	"archive/zip"
	"errors"
)

const (
	RootRelFile             = "_rels/.rels"
	WorkbookRelFile         = "xl/_rels/workbook.xml.rels"
	ContentTypeFile         = "[Content_Types].xml"
	RelationshipContentType = "application/vnd.openxmlformats-package.relationships+xml"
	WorkbookContentType     = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"
	WorksheetContentType    = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
)

var rels []string

type Xlsx struct {
	Files        map[string]*zip.File
	ContentTypes *ContentTypes
	Rels         map[string]*Relationships
	Workbook     *Workbook
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
		return errors.New("Bad xlsx file")
	}

	xlsx.Workbook, err = NewWoorkbook(xlsx)

	if err != nil {
		return err
	}

	return err

}
