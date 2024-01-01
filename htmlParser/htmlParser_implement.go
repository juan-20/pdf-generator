package htmlparser

import (
	"html/template"
	"os"

	"github.com/google/uuid"
)

type HTMLParserStruct struct {
	rootPath string
}

// creating the struct for the interface
func New(rootPath string) HTMLParserInterface {
	return &HTMLParserStruct{rootPath: rootPath}
}

// if the pointer exist the function will be called
func (h *HTMLParserStruct) Create(templateName string, data interface{}) (string, error) {
	// get html template
	templateViewer, err := template.ParseFiles(templateName)
	if err != nil {
		return "", err
	}
	fileName := h.rootPath + "/" + uuid.New().String() + ".html"
	
	// execute the template
	fileWriter ,err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	err = templateViewer.Execute(fileWriter, data)
	if err != nil {
		return "", err
	}

	return fileName, nil
}