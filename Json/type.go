package Json

type Type int

const (
	Object Type = iota
	Array
	String
	Boolean
	Number
	Null
)

type Value interface {
	Type() Type
	Get() any
}

type Pair struct {
	Key string
	Value
}
