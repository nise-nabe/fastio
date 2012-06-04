fastio
======

This library use for programming contest.

## Installation

    go get github.com/nise-nabe/fastio

## Usage Exapmle

    s := fastio.NewInOut(os.Stdin, os.Stdout)
    n := s.Next()
    s.Println("Hello, World! ", n)
    s.Flush()
