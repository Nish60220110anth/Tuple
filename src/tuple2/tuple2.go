package tuple2

import (
	"fmt"
	"reflect"
	"tuple/src"
)

const tupleSize = 2

// NewTuple[types](...)  Tuple.At(index)
type Tuple2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

func (tuple *Tuple2[T1, T2]) At(index uint) any {
	if index > tupleSize {
		src.OverSizeHandler(tupleSize, int(index))
	}

	switch index {
	case 1:
		return tuple.V1
	default:
		return tuple.V2
	}
}

func (tuple *Tuple2[T1, T2]) Size() uint {
	return tupleSize
}

func (tuple *Tuple2[T1, T2]) String() string {
	return fmt.Sprintf("{T1 : %v} {T2 : %v}", tuple.V1, tuple.V2)
}

func (tuple *Tuple2[T1,T2]) Set(index uint, pVal any) {

	switch index{
	case 1:
		if reflect.TypeOf((*tuple).V1)!=reflect.TypeOf(pVal){

		}
		t,ok:=pVal.(reflect.TypeOf((*tuple).V1))
	}
}

func NewTuple2[T1, T2 any](pV1 T1, pV2 T2) *src.Collection {
	var tuple src.Collection = &Tuple2[any, any]{V1: pV1, V2: pV2}
	return &tuple
}
