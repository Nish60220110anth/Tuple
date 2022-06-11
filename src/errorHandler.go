package src

import (
	"fmt"
	"reflect"
)

const overSizeErrorStr = "index value is more than required value"
const typeErrrorStr = "core type of operands is not matched "

// overSizeHandler is handler for the illegal access of tuple collection
// over it's required value
func OverSizeHandler(rmaxInt, gInt int) {
	panic(fmt.Sprintf("%s\nRequired(at max inclusive) : %d\nGot : %d", overSizeErrorStr, rmaxInt, gInt))
}

func TypeErrorMatchHandler(rType, gType reflect.Type) {
	panic(fmt.Sprintf("%s\nRequired Core Type : %s\nGot Core Type : %s", typeErrrorStr, rType, gType))
}
