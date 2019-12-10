// Package clatter is a code generator that simplifies the writing of
// multi-threaded applications that purely share by communicating, encapsulating
// variables inside a package and sharing their results to other nodes in the
// processing graph where eventually they produce an output. The initial
// implementation is just a code generator but it is planned to eventually make
// transparent transport plugs, copy on write immutable functional variables
// and several other features found in the Roadmap in readme.md
package clatter
