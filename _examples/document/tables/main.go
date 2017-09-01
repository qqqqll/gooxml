// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"baliance.com/gooxml/color"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"

	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

func main() {
	doc := document.New()

	table := doc.AddTable()
	// width of the page
	table.SetWidthPercent(100)
	// with thick borers
	borders := table.Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

	row := table.AddRow()
	run := row.AddCell().AddParagraph().AddRun()
	run.AddText("Name")
	run.SetHighlight(wml.ST_HighlightColorYellow)
	row.AddCell().AddParagraph().AddRun().AddText("John Smith")
	row = table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("Street Address")
	row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")

	doc.AddParagraph()

	table = doc.AddTable()
	// 4 inches wide
	table.SetWidth(4 * measurement.Inch)
	borders = table.Borders()
	// thin borders
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	row = table.AddRow()
	cell := row.AddCell()
	// column span
	//cell.Properties().SetColumnSpan(2)
	run = cell.AddParagraph().AddRun()
	run.AddText("Cells can span multiple columns")

	row = table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText("Street Address")
	row.AddCell().AddParagraph().AddRun().AddText("111 Country Road")

	doc.SaveToFile("tables.docx")

}
