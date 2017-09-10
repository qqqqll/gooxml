// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

func TestWdCT_WordprocessingCanvasChoiceConstructor(t *testing.T) {
	v := wordprocessingml.NewWdCT_WordprocessingCanvasChoice()
	if v == nil {
		t.Errorf("wordprocessingml.NewWdCT_WordprocessingCanvasChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wordprocessingml.WdCT_WordprocessingCanvasChoice should validate: %s", err)
	}
}

func TestWdCT_WordprocessingCanvasChoiceMarshalUnmarshal(t *testing.T) {
	v := wordprocessingml.NewWdCT_WordprocessingCanvasChoice()
	buf, _ := xml.Marshal(v)
	v2 := wordprocessingml.NewWdCT_WordprocessingCanvasChoice()
	xml.Unmarshal(buf, v2)
}
