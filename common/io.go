package common

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

const CENTER_C_GROUP_INDEX = 3
const CENTER_C_DO_LOCAL_INDEX = 1

func Parse5(s string, sign int) (ret Note, err error) {
	/*

					  ------- +2
							  +1.5
		              ------- +1
		                      +0.5
				------------------  5
									4.5
				------------------  4
									3.5
				------------------  3
									2.5
				------------------  2
									1.5
				------------------  1
				                   -0.5
				       ------- -1
					   		   -1.5
					   ------- -2


	*/

	f, err := strconv.ParseFloat(s, 10)

	if sign == 0 {
		//center C

		ret.GroupIndex = CENTER_C_GROUP_INDEX
		ret.LocalIndex = (f - 1) * 2
	}

	return ret, nil
}

func ReadNotes(src io.Reader) (ret []interface{}, err error) {

	var s scanner.Scanner
	s.Init(src)

	isGroup := false

	curSign := 0

	var lastGroup []Note

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()

		v := s.TokenText()

		fmt.Println(v)

		switch v {
		case "(":
			isGroup = true
			continue
		case ")":
			isGroup = false
			continue
		case "+":
			curSing = 1
			continue
		case "-":
			curSign = -1
			continue
		}

		if !isGroup {

			note, err := Parse5(v, curSign)

			if err != nil {
				return ret, err
			}

			lastGroup = make([]Note, 1)
			lastGroup[0] = note
		}

	}
	return
}
