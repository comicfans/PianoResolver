package common

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

func ReadNotes(src io.Reader) []interface {

	var s scanner.Scanner
	s.Init(src)

	ret := make([][]Note, 0)

	isGroup := false

	var lastGroup []Note

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()

		v := s.TokenText()

		if v == "(" {
			isGroup = true
			continue
		}

		if v == ")" {
			isGroup = false
			continue
		}

		if !isGroup {

			lastGroup =make([]Note,1)
			lastGroup[0],err =strconv.Atoi(v)

			if err!=nil {
				return nil
				
			}
		}

	}

	return ret

}
