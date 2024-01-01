package pdfparser

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wk struct {
	rootPath string
}

func New(rootPath string) PDFParserInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile string) (string, error) {
	// get the file and create the pdf
	f, err := os.Open(htmlFile)
	if err != nil {
		return "", err
	}
	
	newPDF, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", err
	}
	
	newPDF.AddPage(wkhtmltopdf.NewPageReader(f))

	if err := newPDF.Create(); err != nil {
		return "", err
	}

	filename := w.rootPath + "/" + uuid.New().String() + ".pdf"

	// write the pdf
	if err := newPDF.WriteFile(filename); err != nil {
		return "", err
	}
	return  filename, nil
}
