// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_XmlPr struct {
	// XML Map Id
	MapIdAttr uint32
	// XPath
	XpathAttr string
	// XML Data Type
	XmlDataTypeAttr string
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_XmlPr() *CT_XmlPr {
	ret := &CT_XmlPr{}
	return ret
}

func (m *CT_XmlPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "mapId"},
		Value: fmt.Sprintf("%v", m.MapIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xpath"},
		Value: fmt.Sprintf("%v", m.XpathAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlDataType"},
		Value: fmt.Sprintf("%v", m.XmlDataTypeAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_XmlPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "mapId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.MapIdAttr = uint32(parsed)
		}
		if attr.Name.Local == "xpath" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.XpathAttr = parsed
		}
		if attr.Name.Local == "xmlDataType" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.XmlDataTypeAttr = parsed
		}
	}
lCT_XmlPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_XmlPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_XmlPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_XmlPr and its children
func (m *CT_XmlPr) Validate() error {
	return m.ValidateWithPath("CT_XmlPr")
}

// ValidateWithPath validates the CT_XmlPr and its children, prefixing error messages with path
func (m *CT_XmlPr) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
