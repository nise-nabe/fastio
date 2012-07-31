// Copyright 2012 nise_nabe. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This library use for programming contest like a SPOJ or GCJ.
//
// Example:
// 	js := fastio.NewInOut(os.Stdin, os.Stdout)
// 	jn := s.Next()
// 	js.Println("Hello, World! ", n)
// 	js.Flush()

package fastio

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
)

type InOut struct {
	in []byte
	*bufio.Writer
}

func NewInOut(r io.Reader, w io.Writer) *InOut {
	in, _ := ioutil.ReadAll(r)
	return &InOut{in, bufio.NewWriter(w)}
}

// Get Next Integer
func (s *InOut) Next() (r int) {
	buf := s.in
	p := 1
	for (buf[0] < '0' || '9' < buf[0]) && buf[0] != '-' {
		buf = buf[1:]
	}
	if buf[0] == '-' {
		p = -1
		buf = buf[1:]

	}
	for '0' <= buf[0] && buf[0] <= '9' {
		r = 10*r + int(buf[0]-'0')
		buf = buf[1:]
	}
	r *= p
	s.in = buf
	return
}

// Get Next Line String
func (s *InOut) NextLine() (r string) {
	buf := s.in
	for buf[0] == '\n' {
		buf = buf[1:]
	}
	p := 0
	for buf[p] != '\n' {
		p++
	}
	r = string(buf[0:p])
	s.in = buf[p:]
	return
}

// Get Next String using delimiter whitespace
func (s *InOut) NextStr() (r string) {
	buf := s.in
	for buf[0] == '\n' || buf[0] == ' ' {
		buf = buf[1:]
	}
	p := 0
	for buf[p] != '\n' && buf[p] != ' ' {
		p++
	}
	r = string(buf[0:p])
	s.in = buf[p:]
	return
}

// Print Strings using the suitable way to change type
func (s *InOut) Print(os ...interface{}) {
	for _, o := range os {
		switch o.(type) {
		case byte:
			s.WriteByte(o.(byte))
		case string:
			s.WriteString(o.(string))
		case int:
			s.WriteString(strconv.Itoa(o.(int)))
		case int64:
			s.WriteString(strconv.FormatInt(o.(int64), 10))
		default:
			s.WriteString(fmt.Sprint(o))
		}
	}
}

// Print Strings using the suitable way to change type with new line
func (s *InOut) Println(os ...interface{}) {
	for _, o := range os {
		s.Print(o)
	}
	s.Print("\n")
}

// Print immediately with new line (for debug)
func (s *InOut) PrintlnNow(o interface{}) {
	fmt.Println(o)
}
