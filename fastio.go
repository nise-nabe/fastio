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
	"strconv"
)

type InOut struct {
	*bufio.Reader
	*bufio.Writer
}

func NewInOut(r io.Reader, w io.Writer) *InOut {
	return &InOut{bufio.NewReader(r), bufio.NewWriter(w)}
}

// Get Next Integer
func (s *InOut) Next() (r int) {
	b, _ := s.ReadByte()
	p := 1
	for (b < '0' || '9' < b) && b != '-' {
		b, _ = s.ReadByte()
	}
	if b == '-' {
		p = -1
		b, _ = s.ReadByte()
	}
	for '0' <= b && b <= '9' {
		r = 10*r + int(b-'0')
		b, _ = s.ReadByte()
	}
	return r * p
}

// Get Next Line String
func (s *InOut) NextLine() (r string) {
	b, _ := s.ReadByte()
	for b == '\n' {
		b, _ = s.ReadByte()
	}
	buf := make([]byte, 0)
	for ; b != '\n'; b, _ = s.ReadByte() {
		buf = append(buf, b)
	}
	return string(buf)
}

// Get Next String using delimiter whitespace
func (s *InOut) NextStr() (r string) {
	b, _ := s.ReadByte()
	for b == '\n' || b == ' ' {
		b, _ = s.ReadByte()
	}
	buf := make([]byte, 0)
	for ; b != '\n' && b != ' '; b, _ = s.ReadByte() {
		buf = append(buf, b)
	}
	return string(buf)
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
