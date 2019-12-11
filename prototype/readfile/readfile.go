// Package readfile is a node that accepts a filename and sends a buffer
// back to the caller
package readfile

import (
	"io/ioutil"
)

type Bytes []byte

type Error string

func ToError(err error) Error {
	return Error(err.Error())
}

type Buffer struct {
	Bytes
	Error
}

type Node struct {
	Quit     chan struct{}
	Filename <-chan string
	Buffer   chan<- Buffer
}

func New() (n *Node) {
	n = &Node{
		Quit:     make(chan struct{}, 1),
		Filename: make(<-chan string, 1),
		Buffer:   make(chan<- Buffer, 1),
	}
	go func() {
	out:
		for {
			select {
			case filename := <-n.Filename:
				// caller is requesting to read a file
				b, e := ioutil.ReadFile(filename)
				n.Buffer <- Buffer{b, ToError(e)}
			case <-n.Quit:
				break out
			}
		}
	}()
	return
}
