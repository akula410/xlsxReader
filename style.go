package main

import (
	"encoding/xml"
	"errors"
)

type Style struct {
	XMLName xml.Name  `xml:"styleSheet"`
	NumFmts *NumFmts `xml:"numFmts"`
}

type NumFmts struct {
	Count   int       `xml:"count,attr"`
	NumFmt []*NumFmt `xml:"numFmt"`
}

type NumFmt struct {
	NumFmtId   int    `xml:"numFmtId,attr"`
	FormatCode string `xml:"formatCode,attr"`
}

func NewStyle(xlsx *Xlsx) (st *Style, err error) {
	st = &Style{}

	if partName, ok := xlsx.ContentTypes.GetPartNameByType(StyleContentType); ok {
		f, err := xlsx.Files[partName[0]].Open()

		if err != nil {
			return nil, err
		}

		defer f.Close()

		d := xml.NewDecoder(f)

		err = d.Decode(st)

		if err != nil {
			return nil, err
		}

		return st, err
	}

	return nil, errors.New("style not found")
}
