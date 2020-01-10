package main

import (
	"encoding/xml"
	"path/filepath"
	"strings"
)

type ContentTypesOverride struct {
	PartName    string `xml:"PartName,attr"`
	ContentType string `xml:"ContentType,attr"`
}

type ContentTypesDefault struct {
	Extension   string `xml:"Extension,attr"`
	ContentType string `xml:"ContentType,attr"`
}

type ContentTypes struct {
	Xlsx     *Xlsx
	XMLName  xml.Name `xml:"Types"`
	Override []*ContentTypesOverride
	Default  []*ContentTypesDefault
}

func (ct *ContentTypes) GetPartNameByType(contentType string) (partName []string, ok bool) {
	for _, ov := range ct.Override {
		if ov.ContentType == contentType {
			partName = append(partName, strings.TrimLeft(ov.PartName, "/"))
		}
	}
	if len(partName) == 0 {
		for _, df := range ct.Default {
			if df.ContentType == contentType {
				for _, f := range ct.Xlsx.Files {
					ext := strings.TrimLeft(filepath.Ext(f.Name), ".")
					if ext == df.Extension {
						partName = append(partName, strings.TrimLeft(f.Name, "/"))
					}
				}
			}
		}
	}
	return partName, len(partName) > 0
}

func NewContentType(xlsx *Xlsx) (ct *ContentTypes, err error) {

	ct = &ContentTypes{}
	ct.Xlsx = xlsx

	var fName string

	for k, f := range xlsx.Files {
		if strings.ToLower(f.Name) == strings.ToLower(ContentTypeFile) {
			fName = k
			break
		}
	}

	contentTypeFile, err := xlsx.Files[fName].Open()

	if err != nil {
		return nil, err
	}

	defer contentTypeFile.Close()

	xmlDecoder := xml.NewDecoder(contentTypeFile)

	err = xmlDecoder.Decode(ct)

	if err != nil {
		return nil, err
	}

	return ct, err
}
