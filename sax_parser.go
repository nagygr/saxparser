package saxparser

import (
	"encoding/xml"
	"errors"
	"io"
)

type SaxParser struct {
	decoder *xml.Decoder
	handler SaxHandler
}

func NewParser(reader io.Reader, handler SaxHandler) *SaxParser {
	return &SaxParser{
		decoder: xml.NewDecoder(reader),
		handler: handler,
	}
}

func (p *SaxParser) Parse() error {
	for token, err := p.decoder.Token(); err != io.EOF; token, err = p.decoder.Token() {
		if err != nil {
			return err
		}

		switch t := token.(type) {
			case xml.CharData:
				p.handler.Characters(t)
			case xml.StartElement:
				name := t.Name.Local
				attrs := make(map[string]string, len(t.Attr))
				for _, v := range t.Attr {
					attrs[v.Name.Local] = v.Value
				}
				p.handler.StartElement(name, attrs)
			case xml.EndElement:
				p.handler.EndElement(t.Name.Local)
			case xml.Comment:
				p.handler.Comment(t)
			case xml.ProcInst:
				p.handler.ProcessingInstruction(t.Target, t.Inst)
			case xml.Directive:
				p.handler.Directive(t)
			default:
				return errors.New("Unknown XML tag")
		}
	}

	return nil
}
