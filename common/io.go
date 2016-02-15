package common

import (
	"io"
	"regex"
	"strings"
	"text/scanner"
)

func ReadNotes(src io.Reader) []Note {

	var s scanner.Scanner
	s.Init(src)

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()

		v := s.TokenText()

	}
}
