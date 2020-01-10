package main

import (
	"encoding/xml"
	"strings"
)

type SST struct {
	XMLName     xml.Name   `xml:"sst"`
	Count       int        `xml:"count,attr"`
	UniqueCount int        `xml:"uniqueCount,attr"`
	SSTitems    []*SSTitem `xml:"si"`
}

type SSTitemVal struct {
	Value string `xml:",chardata"`
	Space string `xml:"space,attr"`
}
type SSTitem struct {
	T SSTitemVal   `xml:"t"`
	R []SSTitemVal `xml:"r>t"`
}

func (i *SSTitem) GetString() string {
	if len(i.T.Value) > 0 {
		return i.T.Value
	} else if len(i.R) > 0 {
		var s string
		for _, v := range i.R {
			s = s + v.Value
		}
		return s
	}
	return ""
}

func NewSST(xlsx *Xlsx) (sst *SST, err error) {
	sst = &SST{}

	var fName string

	if partName, ok := xlsx.ContentTypes.GetPartNameByType(SSTContentType); ok {

		for k, f := range xlsx.Files {
			if strings.ToLower(f.Name) == strings.ToLower(partName[0]) {
				fName = k
				break
			}
		}

		f, err := xlsx.Files[fName].Open()

		if err != nil {
			return nil, err
		}

		defer f.Close()

		d := xml.NewDecoder(f)

		err = d.Decode(sst)

		if err != nil {
			return nil, err
		}

		return sst, err
	}

	return sst, err
}
