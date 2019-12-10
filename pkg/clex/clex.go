// Package clex is a lexical analyser that takes a buffer containing a clatter
// node declaration and spits out a slice of structs containing the symbol and
// its type, including the block inside the node, which is wrapped in a Node()
// function in a subfolder of the clatter source code file.
// Just as in Go, main is a special function name, and the nodes are in fact
// implemented as init because they run whether they work or not.
package clex

type Node struct {}

func NewNode(filename string) (out *Node) {
	return
}

// func (n *Node)
