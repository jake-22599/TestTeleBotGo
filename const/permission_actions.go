package constant

// Action - Custom type to hold value of the action permissible
type Action int

// Declare related constants for each value starting with index 1
const (
	Read    Action = iota + 1 // EnumIndex = 1
	Write                     // EnumIndex = 2
	Execute                   // EnumIndex = 3
)

var ActionContants = []Action{Read, Write, Execute}

// String - Creating common behavior - give the type a String function
func (d Action) String() string {
	return [...]string{"read", "write", "execute"}[d-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (d Action) EnumIndex() int {
	return int(d)
}
