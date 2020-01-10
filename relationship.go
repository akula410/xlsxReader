package main

import "encoding/xml"

type Relationship struct {
	ID     string `xml:"Id,attr"`
	Type   string `xml:"Type,attr"`
	Target string `xml:"Target,attr"`
}
type Relationships struct {
	partName     string
	XMLName      xml.Name `xml:"Relationships"`
	Relationship []*Relationship
}

func (r *Relationships) GetByID(ID string) (rl *Relationship, ok bool) {
	for _, rl := range r.Relationship {
		if rl.ID == ID {
			return rl, true
		}
	}
	return nil, false
}

func NewRelationships(partName string, xlsx *Xlsx) (r *Relationships, err error) {

	r = &Relationships{
		partName:     partName,
		XMLName:      xml.Name{},
		Relationship: nil,
	}

	f, err := xlsx.Files[r.partName].Open()

	if err != nil {
		return nil, err
	}

	defer f.Close()

	d := xml.NewDecoder(f)

	err = d.Decode(r)

	if err != nil {
		return nil, err
	}

	return r, err
}
