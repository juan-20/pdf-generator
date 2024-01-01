package htmlparser

// Interface
type HTMLParserInterface interface {
	// Parse the HTML file
	Create(template string, data interface{}) (string, error);
}
