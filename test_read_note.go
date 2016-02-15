package main

import (
	"./common"
	"strings"
)

func main() {
	reader := strings.NewReader("1 2 3 5.5 2.5 (2.5 2.5) 1.5")

	common.ReadNotes(reader)

}
