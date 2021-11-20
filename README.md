# SaxParser

A SAX parser in Go.

It parses an XML without loading the entire contents into memory. Instead it
uses callback functions to notify the user of the different elements as it goes
along the file.
