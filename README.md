# The saxparser module

A SAX parser in Go.

It parses an XML without loading the entire contents into memory. Instead it
uses callback functions to notify the user of the different elements as it goes
along the file. It is a very thin wrapper over `encoding/xml`.

Tests are provided with full code coverage (issue: `go test` and `go test
-cover`).

Documentation is also provided, it can be viewed using the `godoc` tool.
