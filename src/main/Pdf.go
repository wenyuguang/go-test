package main

import "github.com/jung-kurt/gofpdf"

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 7, "hello word", "0", 0, "CM", false, 0, "")
	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"D://1.JPG",
		0, 0,
		0, 0,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)

	pdf.OutputFileAndClose("d://test//test.pdf")
}
