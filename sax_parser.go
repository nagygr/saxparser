/*
Package saxparser is a very thin wrapper over encoding/xml providing a SAX-like
API with callback functions corresponding to all the XML element types.
*/
package saxparser

import (
	"encoding/xml"
	"io"
)

/*
SaxParser is the struct encompassing the XML parser. It holds a pointer to an
xml.Decoder and the SaxHandler implementation.
*/
type SaxParser struct {
	decoder *xml.Decoder
	handler SaxHandler
}

/*
NewParser creates a new SaxParser. It expects an io.Reader for the XML document
and a SaxHandler implementation.
*/
func NewParser(reader io.Reader, handler SaxHandler) *SaxParser {
	return &SaxParser{
		decoder: xml.NewDecoder(reader),
		handler: handler,
	}
}

/*
Parse performs the actual parsing of the XML document. It processes the XML
file, calls the callback functions of the provided SaxHandler implementation and
returns an error if anything goes wrong -- it simply forwards the errors of
encoding/xml.Decoder.Token().
*/
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
			/*
			 * There's no default branch as all the token types have been
			 * handled.
			 */
		}
	}

	return nil
}
