// Package clatter is an implementation of a CSP based functional concurrent
// internal message passing system that can be used to easily build finite state
// automata without the use of any new syntax or grammar such as used in `lex`
// and 'bison', as well as concurrent scheduling, memory management, GUI
// implementations and all kinds of data processing systems, and with the
// creation of patterns and generators, enabling a pointer-free language, with
// the main downside being in the loss of prepositional infix operator syntax
// that can fall back to Go's native infix operators on primitive simple types
// within function scope but still allow the underlying implementation to pass
// parameters without copying them to the stack while allowing receiving
// functions to mutate the values without mutating the value for the caller, by
// using methods.
package clatter

// Source is an object which abstracts a wrapper around a handler that receives
// inputs usually from input devices but also such as tickers, which are based
// on timer interrupts
type Source interface {
	Append(Sink)
	Prepend(Sink)
	Delete()
	GetSinks() []Sink
}
// this is the concrete implementation part of a Source. The sinks are like a
// subscriber list, and are iterated in order when a message is generated
type source struct{
	sinks []Sink
}

// Sink is an object which abstracts functions that process input
type Sink interface {

}
// this is the core concrete implementation which provides an array of Sources
// that the Sink generates
type sink struct {
	sources []Source
}

// Message is a packet of information generated in a Source and sent to Sinks
// The type of data is abstracted into []byte and copy-free decoded into its
// type using unsafe pointer type casts
type Message interface {
	Load() []byte
	Store([]byte)
}
