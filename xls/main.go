package main

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/nelsonsaake/go/do"
	"github.com/nelsonsaake/go/ufs"
	"github.com/nelsonsaake/go/xls"
	"github.com/xuri/excelize/v2"
)

var (
	//go:embed data
	data string

	// xls sheet name for output
	oututSheetName = "sheet1"

	// output file name
	outputfile string = "output.xlsx"
)

func main() {

	// parse content the file `data` into lines
	ls := do.Lines(data)

	// isDataLine: check if this line contains data
	// it's either a dataLine or section header
	isDataLine := func(line string) bool {
		return strings.HasSuffix(line, "%")
	}

	// split a line into it's component data sets
	split := func(line string) (label, f, percentage string) {

		// `line` is expected to look like this: `Not at all 0 0%`
		ls := strings.Split(line, " ")

		// less than 2, means we don't have enough parts to make sense of the data
		// so we don't process
		if len(ls) < 3 {
			return
		}

		// `percentage` should be last entry
		percentage = ls[len(ls)-1]
		percentage = strings.TrimSpace(percentage)

		// `f`, frequency should be the last but one entry
		f = ls[len(ls)-2]
		f = strings.TrimSpace(f)

		// label separator, use space or not
		labelSeparator := " "
		if isDataLine(line) {
			labelSeparator = ""
		}

		// `label`, should be everything before `f`
		label = strings.Join(ls[:len(ls)-2], labelSeparator)
		label = strings.TrimSpace(label)

		return
	}

	// new xls workbook
	doc := xls.New()

	// new xls sheet in the workbook
	sheet := doc.NewSheet(oututSheetName)

	// xls row, where the data would be written
	row := 1

	// cummulativeFrequency frequency
	var (
		cummulativeFrequency int64 = 0
	)

	// process the data from the file, line by line
	for i, line := range ls {

		// add some spacing before each section
		// non dataline indicated the start of a new section
		if !isDataLine(line) && i != 0 {
			err := sheet.Println(row, "")
			do.Die(err)
			row++
		}

		// etract component data from the line
		label, fstring, percentage := split(line)

		// try converting the frequency value from string to int
		fint, _ := strconv.ParseInt(fstring, 10, 64)

		// if it's a data line, use the integer version of frequency
		// else use the string version
		var f any
		if isDataLine(line) {
			f = fint
		} else {
			f = fstring
		}

		// if it's non data line clear the current cummulative frequency it
		// non data line indicates a new section
		// else add the current frequency to the cummulative frequency
		// we don't add the frequency of `Total`, is it's not a raw data
		if !isDataLine(line) {
			cummulativeFrequency = 0
		} else if label != "Total" {
			cummulativeFrequency += fint
		}

		// write the data to xls file
		var err error
		if !isDataLine(line) {
			err = sheet.Println(row, label, f, percentage, "Valid Percentage", "Cummulative Frequency")
		} else {
			err = sheet.Println(row, label, f, percentage, percentage, cummulativeFrequency)
		}
		do.Die(err)

		// move the row cursor down one line
		row++
	}

	// adjust the column width for the columns we've written to
	for i := 1; i <= 6; i++ {

		col, err := excelize.ColumnNumberToName(i)
		do.Die(err)

		err = sheet.SetColWidth(col, 30)
		do.Die(err)
	}

	// delete old output file
	// we don't want any situation of old data, sort of overlapping the new one
	err := ufs.DelFile(outputfile)
	do.Die(err)

	// write new output file
	err = doc.SaveAs(outputfile)
	do.Die(err)
}
