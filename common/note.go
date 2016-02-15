package common

import (
	"fmt"
)

//                  middle
//                  +---+
//               +--+   +---+
//               |  |   |   +--+
//               |  |   |   |  |
//    thumb      |  |   |   |  |
//        +--+   |  |   |   |  |
//       ++  |   |  |       |  |
//       ++  ++-/              |
//        +                    |
//         \                 |
//           \ -----+   +------+
//                  |   |
//                  |   |
//                  wrist

type Range uint64

//suppose that wirst has same position with middle finger
//this is GlobalNoteIndex range where you are allowed to
//point your wirst to

const NOTE_NUMBER = 7

var NoteName = [NOTE_NUMBER]string{"Do", "Re", "Mi", "Fa", "So", "La", "Si"}
var UpDownName = [3]string{"-1", "", "+1"}

const MAX_GROUP = 7
const MAX_GLOBAL_INDEX = MAX_GROUP * len(NoteName)
const TWO_FINGER = 2
const BIT1 uint64 = 1
const MIN_WRIST_INDEX int = TWO_FINGER
const MAX_WRIST_INDEX int = MAX_GLOBAL_INDEX - TWO_FINGER

type Note struct {
	GroupIndex uint8
	LocalIndex uint8
	UpDown     int8
}

func (n Note) GlobalNoteIndex() uint8 {
	return n.GroupIndex*NOTE_NUMBER + n.LocalIndex
}

func String(n Note) string {
	return fmt.Sprintf("%s%s in Group %d,", NoteName[n.LocalIndex], UpDownName[n.UpDown], n.GroupIndex)
}

func ToRange(n Note) Range {

	var centerIndex int = int(n.GlobalNoteIndex())

	minIndex := MIN_WRIST_INDEX
	maxIndex := MAX_WRIST_INDEX

	if centerIndex-TWO_FINGER > minIndex {
		minIndex = centerIndex - TWO_FINGER
	}

	if centerIndex+TWO_FINGER < maxIndex {
		maxIndex = centerIndex + TWO_FINGER
	}

	var value uint64 = 0
	for i := minIndex; i < maxIndex; i++ {
		value &= (BIT1 << uint(i))
	}

	return Range(value)
}

func HasCommon(lhs, rhs Range) bool {
	return (lhs & rhs) != 0
}

func Merge(lhs, rhs Range) Range {
	return Range(lhs & rhs)
}
