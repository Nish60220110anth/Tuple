// INUmerics Interface
// Interface INumerics act a type constraint over
// type parameter for the size of the collection
// and it can accept only the type unsigned values
// and also it's core type

// Restriction :
// Unsigned numeric's are used to restrict the values
// and also for better readability

package src

type INUmerics interface {
	~uint8 | ~uint16 | ~uint32
}
