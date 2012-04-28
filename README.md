fastio
======

This library use in programming contest for nise\_nabe.

## Installation

    go get github.com/nise-nabe/fastio

## Usage Exapmle

   s := fastio.NewInOut(os.Stdin, os.Stdout)
   n := s.next()
   s.println("Hello, World! ", n)
   s.Flush()
