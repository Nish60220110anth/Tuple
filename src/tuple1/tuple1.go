// Package tuple1 supports the tuple using the underlying struct .
// At creation of tuple , types need to provided but it can also be
// inferred automatically by compiler but on further  type-constrained
// changes in the underlying struct will be based on the previously
// given value , so it makes easy to use values of same core types
// to assigned without explicit conversions

package tuple1

import (
	"fmt"
	"reflect"
	"tuple/src"
)

const tupleSize = 1

// NewTuple[types](...)  Tuple.At(index)
type Tuple1[T1 any] struct {
	V1 T1
}

func (tuple *Tuple1[T1]) At(index uint) any {
	if index > tupleSize {
		src.OverSizeHandler(tupleSize, int(index))
	}
	return tuple.V1
}

func (tuple *Tuple1[T1]) Size() uint {
	return tupleSize
}

func (tuple *Tuple1[T1]) String() string {
	return fmt.Sprintf("{T1 : %v}", tuple.V1)
}

func (tuple *Tuple1[T1]) Set(index uint, pVal T1) {
	if reflect.TypeOf((*tuple).V1) != reflect.TypeOf(pVal) {
		src.TypeErrorMatchHandler(reflect.TypeOf((*tuple).V1), reflect.TypeOf(pVal))
	}
	(*tuple).V1 = pVal
}

func NewTuple1[T1 any](pV1 T1) *src.Collection {
	var tuple src.Collection = &Tuple1[any]{V1: pV1}
	return &tuple
}
