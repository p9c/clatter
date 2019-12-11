// Package clex is a lexical analyser that takes a buffer containing a clatter
// node declaration and spits out a slice of structs containing the symbol and
// its type, including the block inside the node, which is wrapped in a Node()
// function in a subfolder of the clatter source code file.
// Just as in Go, main is a special function name, and the nodes are in fact
// implemented as init because they run whether they work or not.
package clex

// Pattern is the regular expression or literal string of a lexeme
type Pattern []byte

// Lexicon is a map of patterns with human readable short identifiers
type Lexicon map[string]Pattern

// Lexeme is a single symbol extracted using a pattern from the lexicon
type Lexeme []byte

// Stream is a string of symbols extracted from the source file in
// order of appearance that match
type Stream []Lexeme

type Node struct {
	Quit    chan struct{}
	Error   chan<- []string
	Buffer  chan<- []byte
	Symbols <-chan Stream
}

func New() (n *Node) {
	// init
	n = &Node{
		Quit: make(chan struct{}, 1),
	}
	go func() {
	out:
		for {
			// select
			select {
			case <-n.Quit:
				break out
			}
		}
	}()
	// quit
	return
}

// func (n *Node)
