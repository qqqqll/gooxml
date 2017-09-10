// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_Frame struct {
	// Frame Size
	Sz *CT_String
	// Frame Name
	Name *CT_String
	// Frame or Frameset Title
	Title *CT_String
	// Frame Long Description
	LongDesc *CT_Rel
	// Source File for Frame
	SourceFileName *CT_Rel
	// Left and Right Margin for Frame
	MarW *CT_PixelsMeasure
	// Top and Bottom Margin for Frame
	MarH *CT_PixelsMeasure
	// Scrollbar Display Option
	Scrollbar *CT_FrameScrollbar
	// Frame Cannot Be Resized
	NoResizeAllowed *CT_OnOff
	// Maintain Link to Existing File
	LinkedToFile *CT_OnOff
}

func NewCT_Frame() *CT_Frame {
	ret := &CT_Frame{}
	return ret
}

func (m *CT_Frame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Sz != nil {
		sesz := xml.StartElement{Name: xml.Name{Local: "w:sz"}}
		e.EncodeElement(m.Sz, sesz)
	}
	if m.Name != nil {
		sename := xml.StartElement{Name: xml.Name{Local: "w:name"}}
		e.EncodeElement(m.Name, sename)
	}
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "w:title"}}
		e.EncodeElement(m.Title, setitle)
	}
	if m.LongDesc != nil {
		selongDesc := xml.StartElement{Name: xml.Name{Local: "w:longDesc"}}
		e.EncodeElement(m.LongDesc, selongDesc)
	}
	if m.SourceFileName != nil {
		sesourceFileName := xml.StartElement{Name: xml.Name{Local: "w:sourceFileName"}}
		e.EncodeElement(m.SourceFileName, sesourceFileName)
	}
	if m.MarW != nil {
		semarW := xml.StartElement{Name: xml.Name{Local: "w:marW"}}
		e.EncodeElement(m.MarW, semarW)
	}
	if m.MarH != nil {
		semarH := xml.StartElement{Name: xml.Name{Local: "w:marH"}}
		e.EncodeElement(m.MarH, semarH)
	}
	if m.Scrollbar != nil {
		sescrollbar := xml.StartElement{Name: xml.Name{Local: "w:scrollbar"}}
		e.EncodeElement(m.Scrollbar, sescrollbar)
	}
	if m.NoResizeAllowed != nil {
		senoResizeAllowed := xml.StartElement{Name: xml.Name{Local: "w:noResizeAllowed"}}
		e.EncodeElement(m.NoResizeAllowed, senoResizeAllowed)
	}
	if m.LinkedToFile != nil {
		selinkedToFile := xml.StartElement{Name: xml.Name{Local: "w:linkedToFile"}}
		e.EncodeElement(m.LinkedToFile, selinkedToFile)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Frame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Frame:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sz"}:
				m.Sz = NewCT_String()
				if err := d.DecodeElement(m.Sz, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "name"}:
				m.Name = NewCT_String()
				if err := d.DecodeElement(m.Name, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "title"}:
				m.Title = NewCT_String()
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "longDesc"}:
				m.LongDesc = NewCT_Rel()
				if err := d.DecodeElement(m.LongDesc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sourceFileName"}:
				m.SourceFileName = NewCT_Rel()
				if err := d.DecodeElement(m.SourceFileName, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "marW"}:
				m.MarW = NewCT_PixelsMeasure()
				if err := d.DecodeElement(m.MarW, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "marH"}:
				m.MarH = NewCT_PixelsMeasure()
				if err := d.DecodeElement(m.MarH, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "scrollbar"}:
				m.Scrollbar = NewCT_FrameScrollbar()
				if err := d.DecodeElement(m.Scrollbar, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noResizeAllowed"}:
				m.NoResizeAllowed = NewCT_OnOff()
				if err := d.DecodeElement(m.NoResizeAllowed, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "linkedToFile"}:
				m.LinkedToFile = NewCT_OnOff()
				if err := d.DecodeElement(m.LinkedToFile, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element on CT_Frame %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Frame
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Frame and its children
func (m *CT_Frame) Validate() error {
	return m.ValidateWithPath("CT_Frame")
}

// ValidateWithPath validates the CT_Frame and its children, prefixing error messages with path
func (m *CT_Frame) ValidateWithPath(path string) error {
	if m.Sz != nil {
		if err := m.Sz.ValidateWithPath(path + "/Sz"); err != nil {
			return err
		}
	}
	if m.Name != nil {
		if err := m.Name.ValidateWithPath(path + "/Name"); err != nil {
			return err
		}
	}
	if m.Title != nil {
		if err := m.Title.ValidateWithPath(path + "/Title"); err != nil {
			return err
		}
	}
	if m.LongDesc != nil {
		if err := m.LongDesc.ValidateWithPath(path + "/LongDesc"); err != nil {
			return err
		}
	}
	if m.SourceFileName != nil {
		if err := m.SourceFileName.ValidateWithPath(path + "/SourceFileName"); err != nil {
			return err
		}
	}
	if m.MarW != nil {
		if err := m.MarW.ValidateWithPath(path + "/MarW"); err != nil {
			return err
		}
	}
	if m.MarH != nil {
		if err := m.MarH.ValidateWithPath(path + "/MarH"); err != nil {
			return err
		}
	}
	if m.Scrollbar != nil {
		if err := m.Scrollbar.ValidateWithPath(path + "/Scrollbar"); err != nil {
			return err
		}
	}
	if m.NoResizeAllowed != nil {
		if err := m.NoResizeAllowed.ValidateWithPath(path + "/NoResizeAllowed"); err != nil {
			return err
		}
	}
	if m.LinkedToFile != nil {
		if err := m.LinkedToFile.ValidateWithPath(path + "/LinkedToFile"); err != nil {
			return err
		}
	}
	return nil
}
