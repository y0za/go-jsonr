package jsonr

// ParentTypeNum represents json property type
type ParentTypeNum int

// ParentTypeNum
const (
	ParentTypeObject ParentTypeNum = iota
	ParentTypeArray
)

// Member records json member info
type Member struct {
	Key        string        // name of its property
	Start      int           // start position
	End        int           // end position
	Parent     int           // the index number of parent in slice and -1 means root Member.
	ParentType ParentTypeNum // parent type (Array or Object)
	Index      int           // for representing the index of Array
}

// ParseJSON parses json to Member slice
func ParseJSON(data []byte) ([]Member, error) {
	var members []Member
	// TODO Implementation
	return members, nil
}

// ResolveJSON seeks json path to the position
func ResolveJSON(members []Member, pos int) ([]Member, error) {
	var nodes []Member
	// TODO Implementation
	return nodes, nil
}
