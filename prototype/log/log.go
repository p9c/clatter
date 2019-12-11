// Package log is
package log

import (
	"fmt"
	"io"
	"time"
)

type Level byte

const (
	Off Level = iota
	Fatal
	Error
	Info
	Debug
	Trace
)

var Labels = []string{
	"Off",
	"Fatal",
	"Error",
	"Info",
	"Debug",
	"Trace",
}

type Message struct {
	time.Time
	Text string
}

type Node struct {
	io.Writer
	Quit    chan struct{}
	Message chan Message
}

func New(l Level, w io.Writer) (n *Node) {
	n.Writer = w
	go func() {
	out:
		for {
			select {
			case msg := <-n.Message:
				fmt.Fprintf(w, "%s %s %v", Labels[l], msg.Text, msg.Time)
			case <-n.Quit:
				break out
			}
		}
	}()
	return
}
