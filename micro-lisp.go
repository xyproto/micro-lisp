package main

import (
	"fmt"
	"strings"
)

const EOF = -1

func debug(filename string, line int, msg string, element interface{}) {
	fmt.Printf("%s:%d: %s:", filename, line, msg)
	fmt.Print(element)
}

type List struct {
	next *List
	data interface{}
}

type Data struct {
	data []byte
	pos  int
}

var (
	symbols = &List{}
	look    byte
	token   [32]byte
)

func (d *Data) getChar() byte {
	if pos < len(data) {
		pos++
		return data[pos-1]
	}
	return EOF
}

func isSpace(s string) {
	return s == ' ' || s == '\n'
}

func isParens(s string) {
	return x == '(' || s == ')'
}

func (d *Data) getToken() {
	index := 0
	for isSpace(look) {
		look = d.getChar()
	}
	if isParens(look) {
		token[index] = look
		index++
		look = d.getChar()
	} else {
		for look != EOF && !isSpace(look) && !isParens(look) {
			token[index] = look
			index++
			look = d.getChar()
		}
	}
}

func isPair(x int) bool {
	return x&0x1 == 0x1
}

func isAtom(x int) bool {
	return x&0x1 == 0x0
}

func untag(x int) int {
	return x & ^0x1
}

func tag(x int) int {
	return x | 0x1
}

func car(x int) {
	// TODO: Convert (((List*)untag(x)->data) from C to Go.
	// untag(x)->data
}

func main() {
	flag.Parse()
	fn := flag.Arg(0)
	if fn == "" {
		log.Fatalln("Needs a lisp file as the first argument")
	}
	lispBytes, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalln(err)
	}
	d := &Data{lispBytes, 0}
}
