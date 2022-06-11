//tuple1 package
package src

//Collection1 interface provides method for
// Array , Tuple , List ...
type Collection interface {
	At(index uint) any
	Size() uint
	String() string
	Set(index uint, pVal any)
}
