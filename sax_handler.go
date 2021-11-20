package saxparser

type SaxHandler interface {
	Characters(chars []byte)
	Comment(chars []byte)
	Directive(chars []byte)
	EndElement(name string)
	ProcessingInstruction(target string, instruction []byte)
	StartElement(name string, attributes map[string]string)
}
