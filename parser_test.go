package jsonr

import (
	"fmt"
	"testing"
)

func TestParseJSON(t *testing.T) {
	tests := []struct {
		json     []byte
		err      error
		expected []Token
	}{
		{
			json: []byte(`{"a":1}`),
			err:  nil,
			expected: []Token{
				Token{Key: "a", Start: 6, End: 6, Parent: -1, ParentType: ParentTypeObject},
			},
		},
		{
			json: []byte(`{"a":1,"b":"bbb"}`),
			err:  nil,
			expected: []Token{
				Token{Key: "a", Start: 1, End: 5, Parent: -1, ParentType: ParentTypeObject},
				Token{Key: "b", Start: 7, End: 15, Parent: -1, ParentType: ParentTypeObject},
			},
		},
	}

	for _, tc := range tests {
		tokens, err := ParseJSON(tc.json)
		if err != tc.err {
			t.Errorf("Expected error %v, actual %v", tc.err, err)
			continue
		}
		if len(tokens) != len(tc.expected) {
			t.Errorf("Expected tokens length %d, actual %d", len(tc.expected), len(tokens))
			continue
		}
		for i, token := range tokens {
			eq, msg := equalToken(token, tc.expected[i])
			if !eq {
				t.Errorf("Different %dth token from expected: %s", i, msg)
			}
		}
	}
}

// equalToken compare two tokens
func equalToken(t1, t2 Token) (bool, string) {
	if t1.Key != t2.Key {
		return false, fmt.Sprintf("Key first: '%s', second: '%s'", t1.Key, t2.Key)
	} else if t1.Start != t2.Start {
		return false, fmt.Sprintf("Start first: %d, second: %d", t1.Start, t2.Start)
	} else if t1.End != t2.End {
		return false, fmt.Sprintf("End first: %d, second: %d", t1.End, t2.End)
	} else if t1.Parent != t2.Parent {
		return false, fmt.Sprintf("Parent first: %d, second: %d", t1.Parent, t2.Parent)
	} else if t1.ParentType != t2.ParentType {
		return false, fmt.Sprintf("ParentType first: %d, second: %d", t1.ParentType, t2.ParentType)
	} else if t1.Index != t2.Index {
		return false, fmt.Sprintf("Index first: %d, second: %d", t1.Index, t2.Index)
	}
	return true, ""
}
