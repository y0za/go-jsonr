package jsonr

// ParentTypeNum represents json property type
type ParentTypeNum int

// ParentTypeNum
const (
	ParentTypeObject ParentTypeNum = iota
	ParentTypeArray
)

// Token records json token info
type Token struct {
	Key        string        // name of its property
	Start      int           // start position
	End        int           // end position
	Parent     int           // the index number of parent in slice and -1 means root Token.
	ParentType ParentTypeNum // parent type (Array or Object)
	Index      int           // for representing the index of Array
}

// ParseJSON parses json to Token slice
func ParseJSON(data []byte) ([]Token, error) {
	var tokens []Token
	// TODO Implementation
	return tokens, nil
}

// ResolveJSON seeks json path to the position
func ResolveJSON(tokens []Token, pos int) ([]Token, error) {
	var nodes []Token
	// TODO Implementation
	return nodes, nil
}
