// CollModes are used by Collection to decide
// the type of collection you require and
// it is just abstraction of underlying of
// array

package src

type CollMode uint16

const UnderlyingArraySize = 100

const (
	FixedColl    CollMode = 1 << (16 - 1 - iota) // Fixed Size
	NotFixedColl                                 // Non-Fixed Size
	MutableColl                                  // Mutable
	IMutableColl                                 // IMutable
	UniTypeColl                                  // Uniform Type
	NUniTypeColl                                 // Non Uniform type
	ListColl                                     // Immutable fixed size
	ArrayColl                                    // Mutable fixed size
	TupleColl                                    // Mutable fixed size
)
