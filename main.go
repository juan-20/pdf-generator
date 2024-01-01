package main

import (
	"fmt"
	"os"
	htmlparser "pdf-go/htmlParser"
	pdfparser "pdf-go/pdfParser"
)

type Data struct {
	Name string
}

func main() {
	h := htmlparser.New("temp")
	wk := pdfparser.New("temp")

	dataHTML := Data{
		Name: "Teste",
	}

	htmlGenerated ,err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		panic(err)
	}

	defer os.Remove(htmlGenerated)
	fmt.Println("html gerado" + htmlGenerated)

	
	PDFGenerated ,err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err);
		return
	}
	fmt.Println("Pdf gerado",PDFGenerated);
}