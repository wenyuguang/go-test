package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

	excelFile := excelize.NewFile()

	sheet := excelFile.NewSheet("testSheet")

	excelFile.SetCellValue("testSheet", "A2", "hello world")
	excelFile.SetCellValue("testSheet", "B2", 100)

	excelFile.SetActiveSheet(sheet)

	if err := excelFile.SaveAs("d://test//test.xlsx"); err != nil {
		fmt.Println(err)
	}

}
