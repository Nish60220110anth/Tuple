package tuple3

import (
	"fmt"
	"tuple/src"
)

const tupleSize = 3

// NewTuple[types](...)  Tuple.At(index)
type Tuple3[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

func (tuple *Tuple3[T1, T2, T3]) At(index uint) any {
	if index > tupleSize {
		src.OverSizeHandler(tupleSize, int(index))
	}

	switch index {
	case 1:
		return tuple.V1
	case 2:
		return tuple.V2
	default:
		return tuple.V3
	}
}

func (tuple *Tuple3[T1, T2, T3]) Size() uint {
	return tupleSize
}

func (tuple *Tuple3[T1, T2, T3]) String() string {
	return fmt.Sprintf("{T1 : %v} {T2 : %v} {T3 : %v}", tuple.V1, tuple.V2, tuple.V3)
}

func NewTuple3[T1, T2, T3 any](pV1 T1, pV2 T2, pV3 T3) *src.Collection {
	var tuple src.Collection = &Tuple3[T1, T2, T3]{V1: pV1, V2: pV2, V3: pV3}
	return &tuple
}
