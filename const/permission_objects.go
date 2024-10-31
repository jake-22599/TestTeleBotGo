package constant

// Object - Custom type to hold value for commands
type Object int

// Declare related constants for each object type
const (
	Help      Object = iota + 1 // EnumIndex = 1
	Git_fetch                   // EnumIndex = 2
	Unity                       // EnumIndex = 3
	Upload                      // EnumIndex = 4
	// sync the last index with the function Max so there
)

var ObjectContants = []Object{Help, Git_fetch, Unity, Upload}

// String - Creating common behavior - give the type a String function
func (w Object) String() string {
	return [...]string{"help", "git_fetch", "unity", "upload"}[w-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (w Object) EnumIndex() int {
	return int(w)
}
