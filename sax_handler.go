package saxparser

/*
SaxHandler is an interface for the callback functions for the SAX parser.
It defines a function for each XML element type. By implementing the
interface, an XML document can be processed as a series of events. If the
implementing struct is given state, it can act as a state machine that is
able determine the current position is a well-formed XML and act accordingly.
*/
type SaxHandler interface {
	/*
	Characters is a handler function for character data -- i.e. the characters
	between and opening and a closing tag. Please note, that any whitespace between
	two tags is considered to be character data so a lot of these callback will be made
	for formatting spaces. This is relatively easy to filter out though.
	*/
	Characters(chars []byte)

	/*
	Comment is the callback function for comments (the starting and ending symbols are
	not included).
	*/
	Comment(chars []byte)

	/*
	Directive is the callback function for XML directives: <!directive ...>.
	*/
	Directive(chars []byte)

	/*
	EndElement is the callback function for end elements.
	*/
	EndElement(name string)

	/*
	ProcessingInstruction is the callback function for procInsts, i.e.
	elements in the form of: <?...>
	*/
	ProcessingInstruction(target string, instruction []byte)

	/*
	StartElement is the callback function for start elements. The name of
	the element and a map of attributes is given.
	*/
	StartElement(name string, attributes map[string]string)
}
