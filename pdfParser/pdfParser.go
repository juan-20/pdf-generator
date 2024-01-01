package pdfparser

type PDFParserInterface interface {
	Create(file string) (string, error)
}