package saxparser

import (
	"testing"
	"os"
	"strings"
)

type SaxHandlerTest struct {
	t *testing.T
	inApple bool
}

func (h *SaxHandlerTest) Characters(chars []byte) {
	text := string(chars)

	if strings.TrimSpace(text) == "" {
		return
	}

	expected := "sieversii"

	if text != expected {
		h.t.Errorf("Characters: the characters (%s) are not equal to what was expected (%s)", text, expected)
	}
}

func (h *SaxHandlerTest) Comment(chars []byte) {
	text := string(chars)
	expected := " comment "

	if text != expected {
		h.t.Errorf("Comment: the characters (%s) are not equal to what was expected (%s)", text, expected)
	}
}

func (h *SaxHandlerTest) Directive(chars []byte) {
	text := string(chars)
	expected := "DOCTYPE apple SYSTEM \"example.dtd\""

	if text != expected {
		h.t.Errorf("Directive: the characters (%s) are not equal to what was expected (%s)", text, expected)
	}
}

func (h *SaxHandlerTest) EndElement(name string) {
	if h.inApple {
		h.inApple = false
		if name != "attrs" {
			h.t.Errorf("End element: name (%s) is not what was expected (attrs)", name)
		}
	} else if name != "apple" {
		h.t.Errorf("End element: name (%s) is not what was expected (apple)", name)
	}
}

func (h *SaxHandlerTest) ProcessingInstruction(target string, instruction []byte) {
	if !(target == "xml" && string(instruction) == "version='1.0' encoding='UTF-8'") {
		h.t.Errorf("ProcessingInstruction not what was expected")
	}
}

func (h *SaxHandlerTest) StartElement(name string, attributes map[string]string) {
	if name == "apple" {
		h.inApple = true
	} else if name == "attrs" {
		if value, present := attributes["country"]; !present || value != "China" {
			h.t.Errorf("Missing attribute: country='China'")
		}
		if value, present := attributes["region"]; !present || value != "Tien Shan" {
			h.t.Errorf("Missing attribute: region='Tien Shan'")
		}
	} else {
		h.t.Errorf("Unexpected element: %s", name)
	}
}

type ReaderTest struct {}

func (r* ReaderTest) Read(buf []byte) (n int, err error) {
	return 0, nil
}

func TestNewParser(t *testing.T) {
	parser := NewParser(&ReaderTest{}, &SaxHandlerTest{t:t})

	if parser.decoder == nil {
		t.Errorf("SaxParser.decoder is nil")
	}

	if _, ok := parser.handler.(SaxHandler); !ok {
		t.Errorf("SaxParser.handler is not a SaxParser")
	}
}

func TestExampleXml(t *testing.T) {
	file, err := os.Open("test_data/example.xml")

	if err != nil {
		t.Errorf("Error opening example.xml: %s", err.Error())
	}

	parser := NewParser(file, &SaxHandlerTest{t:t})
	parser.parse()
}
