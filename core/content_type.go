package core

import (
	"encoding/xml"
	"strings"
)

type ContentTypesOverride struct {
	PartName    string `xml:"PartName,attr"`
	ContentType string `xml:"ContentType,attr"`
}
type ContentTypes struct {
	XMLName  xml.Name `xml:"Types"`
	Override []*ContentTypesOverride
}

func (ct *ContentTypes) GetPartNameByType(contentType string) (partName []string, ok bool){
	for _, ov := range ct.Override {
		if ov.ContentType == contentType {
			partName = append(partName, strings.TrimLeft(ov.PartName, "/"))
		}
	}
	return partName, len(partName)>0
}

func NewContentType(xlsx *Xlsx) (ct *ContentTypes, err error) {

	ct = &ContentTypes{}

	contentTypeFile, err := xlsx.Files[ContentTypeFile].Open()

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
