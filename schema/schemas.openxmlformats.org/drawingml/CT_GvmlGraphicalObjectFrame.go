// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_GvmlGraphicalObjectFrame struct {
	NvGraphicFramePr *CT_GvmlGraphicFrameNonVisual
	Graphic          *Graphic
	Xfrm             *CT_Transform2D
	ExtLst           *CT_OfficeArtExtensionList
}

func NewCT_GvmlGraphicalObjectFrame() *CT_GvmlGraphicalObjectFrame {
	ret := &CT_GvmlGraphicalObjectFrame{}
	ret.NvGraphicFramePr = NewCT_GvmlGraphicFrameNonVisual()
	ret.Graphic = NewGraphic()
	ret.Xfrm = NewCT_Transform2D()
	return ret
}

func (m *CT_GvmlGraphicalObjectFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	senvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "a:nvGraphicFramePr"}}
	e.EncodeElement(m.NvGraphicFramePr, senvGraphicFramePr)
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	sexfrm := xml.StartElement{Name: xml.Name{Local: "a:xfrm"}}
	e.EncodeElement(m.Xfrm, sexfrm)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlGraphicalObjectFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvGraphicFramePr = NewCT_GvmlGraphicFrameNonVisual()
	m.Graphic = NewGraphic()
	m.Xfrm = NewCT_Transform2D()
lCT_GvmlGraphicalObjectFrame:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "nvGraphicFramePr"}:
				if err := d.DecodeElement(m.NvGraphicFramePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphic"}:
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "xfrm"}:
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_GvmlGraphicalObjectFrame %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlGraphicalObjectFrame
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlGraphicalObjectFrame and its children
func (m *CT_GvmlGraphicalObjectFrame) Validate() error {
	return m.ValidateWithPath("CT_GvmlGraphicalObjectFrame")
}

// ValidateWithPath validates the CT_GvmlGraphicalObjectFrame and its children, prefixing error messages with path
func (m *CT_GvmlGraphicalObjectFrame) ValidateWithPath(path string) error {
	if err := m.NvGraphicFramePr.ValidateWithPath(path + "/NvGraphicFramePr"); err != nil {
		return err
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
