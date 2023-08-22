package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, "Hello, world")

	err := pdf.OutputFileAndClose("golang.pdf")
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo PDF: ", err)
	} else {
		fmt.Println("PDF gerado com sucesso")
	}
}