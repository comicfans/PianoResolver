package common

import (
	"fmt"
)

type Range struct {
	Value uint64
	Note  *Note
}

const NoteName []string = []string{"Do", "Re", "Mi", "Fa", "So", "La", "Si"}
const UpDownName []string = []string{"-1", "", "+1"}

type Note struct {
	Group  uint8
	Index  uint8
	UpDown int8
}

func String(n Note) string {
	fmt.Sprintf("%s%s in Group %d,", NoteName[n.Index], UpDownName[n.UpDown], n.Group)
}

func ToRange()

func init() {

}
