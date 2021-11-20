package saxparser

import (
	"fmt"
	"testing"
)

type SaxHandlerTest struct {}

func (h *SaxHandlerTest) Characters(chars []byte) {
	fmt.Printf("Characters: %s\n", string(chars))
}

func (h *SaxHandlerTest) Comment(chars []byte) {
	fmt.Printf("Comment: %s\n", string(chars))
}

func (h *SaxHandlerTest) Directive(chars []byte) {
	fmt.Printf("Directive: %s\n", string(chars))
}

func (h *SaxHandlerTest) EndElement(name string) {
	fmt.Printf("End element: %s\n", name)
}

func (h *SaxHandlerTest) ProcessingInstruction(target string, instruction []byte) {
	fmt.Printf("Target: %s Instruction: %s\n", target, string(instruction))
}

func (h *SaxHandlerTest) StartElement(name string, attributes map[string]string) {
	fmt.Printf("Start element: %s\n", name)
	for key, value := range attributes {
		fmt.Printf("\tAttribute: %s %s", key, value)
	}
}

type ReaderTest struct {}

func (r* ReaderTest) Read(buf []byte) (n int, err error) {
	return 0, nil
}

func TestNewParser(t *testing.T) {
	parser := NewParser(&ReaderTest{}, &SaxHandlerTest{})

	if parser.decoder == nil {
		t.Errorf("SaxParser.decoder is nil")
	}

	if _, ok := parser.handler.(SaxHandler); !ok {
		t.Errorf("SaxParser.handler is not a SaxParser")
	}
}
