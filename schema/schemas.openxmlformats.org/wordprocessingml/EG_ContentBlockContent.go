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
	"fmt"
	"log"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/math"
)

type EG_ContentBlockContent struct {
	// Block-Level Custom XML Element
	CustomXml *CT_CustomXmlBlock
	// Block-Level Structured Document Tag
	Sdt *CT_SdtBlock
	// Paragraph
	P []*CT_P
	// Table
	Tbl             []*CT_Tbl
	EG_RunLevelElts []*EG_RunLevelElts
}

func NewEG_ContentBlockContent() *EG_ContentBlockContent {
	ret := &EG_ContentBlockContent{}
	return ret
}

func (m *EG_ContentBlockContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CustomXml != nil {
		secustomXml := xml.StartElement{Name: xml.Name{Local: "w:customXml"}}
		e.EncodeElement(m.CustomXml, secustomXml)
	}
	if m.Sdt != nil {
		sesdt := xml.StartElement{Name: xml.Name{Local: "w:sdt"}}
		e.EncodeElement(m.Sdt, sesdt)
	}
	if m.P != nil {
		sep := xml.StartElement{Name: xml.Name{Local: "w:p"}}
		for _, c := range m.P {
			e.EncodeElement(c, sep)
		}
	}
	if m.Tbl != nil {
		setbl := xml.StartElement{Name: xml.Name{Local: "w:tbl"}}
		for _, c := range m.Tbl {
			e.EncodeElement(c, setbl)
		}
	}
	if m.EG_RunLevelElts != nil {
		for _, c := range m.EG_RunLevelElts {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	return nil
}

func (m *EG_ContentBlockContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ContentBlockContent:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXml"}:
				m.CustomXml = NewCT_CustomXmlBlock()
				if err := d.DecodeElement(m.CustomXml, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdt"}:
				m.Sdt = NewCT_SdtBlock()
				if err := d.DecodeElement(m.Sdt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "p"}:
				tmp := NewCT_P()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.P = append(m.P, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tbl"}:
				tmp := NewCT_Tbl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tbl = append(m.Tbl, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "proofErr"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.ProofErr = NewCT_ProofErr()
				if err := d.DecodeElement(tmprunlevelelts.ProofErr, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermStart = NewCT_PermStart()
				if err := d.DecodeElement(tmprunlevelelts.PermStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "permEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.PermEnd = NewCT_Perm()
				if err := d.DecodeElement(tmprunlevelelts.PermEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ins"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Ins = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Ins, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "del"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.Del = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.Del, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFrom"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveFrom = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveFrom, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveTo"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprunlevelelts.MoveTo = NewCT_RunTrackChange()
				if err := d.DecodeElement(tmprunlevelelts.MoveTo, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkStart = NewCT_Bookmark()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bookmarkEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.BookmarkEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.BookmarkEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFromRangeEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveFromRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeStart = NewCT_MoveBookmark()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveToRangeEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.MoveToRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.MoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeStart = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "commentRangeEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CommentRangeEnd = NewCT_MarkupRange()
				if err := d.DecodeElement(tmprangemarkupelements.CommentRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlInsRangeEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlInsRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlInsRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlDelRangeEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlDelRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlDelRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveFromRangeEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveFromRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveFromRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeStart"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeStart = NewCT_TrackChange()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeStart, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "customXmlMoveToRangeEnd"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmprangemarkupelements := NewEG_RangeMarkupElements()
				tmprangemarkupelements.CustomXmlMoveToRangeEnd = NewCT_Markup()
				if err := d.DecodeElement(tmprangemarkupelements.CustomXmlMoveToRangeEnd, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_RangeMarkupElements = append(tmprunlevelelts.EG_RangeMarkupElements, tmprangemarkupelements)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMathPara"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMathPara = math.NewOMathPara()
				if err := d.DecodeElement(tmpmathcontent.OMathPara, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "oMath"}:
				tmprunlevelelts := NewEG_RunLevelElts()
				tmpmathcontent := NewEG_MathContent()
				tmpmathcontent.OMath = math.NewOMath()
				if err := d.DecodeElement(tmpmathcontent.OMath, &el); err != nil {
					return err
				}
				m.EG_RunLevelElts = append(m.EG_RunLevelElts, tmprunlevelelts)
				tmprunlevelelts.EG_MathContent = append(tmprunlevelelts.EG_MathContent, tmpmathcontent)
			default:
				log.Printf("skipping unsupported element on EG_ContentBlockContent %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ContentBlockContent
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ContentBlockContent and its children
func (m *EG_ContentBlockContent) Validate() error {
	return m.ValidateWithPath("EG_ContentBlockContent")
}

// ValidateWithPath validates the EG_ContentBlockContent and its children, prefixing error messages with path
func (m *EG_ContentBlockContent) ValidateWithPath(path string) error {
	if m.CustomXml != nil {
		if err := m.CustomXml.ValidateWithPath(path + "/CustomXml"); err != nil {
			return err
		}
	}
	if m.Sdt != nil {
		if err := m.Sdt.ValidateWithPath(path + "/Sdt"); err != nil {
			return err
		}
	}
	for i, v := range m.P {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/P[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Tbl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tbl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.EG_RunLevelElts {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_RunLevelElts[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
