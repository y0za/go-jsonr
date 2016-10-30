package jsonr

import (
	"fmt"
	"testing"
)

func TestParseJSON(t *testing.T) {
	tests := []struct {
		json     []byte
		err      error
		expected []Member
	}{
		{
			json: []byte(`{"a":1}`),
			err:  nil,
			expected: []Member{
				Member{Key: "a", Start: 6, End: 6, Parent: -1, ParentType: ParentTypeObject},
			},
		},
		{
			json: []byte(`{"a":1,"b":"bbb"}`),
			err:  nil,
			expected: []Member{
				Member{Key: "a", Start: 1, End: 5, Parent: -1, ParentType: ParentTypeObject},
				Member{Key: "b", Start: 7, End: 15, Parent: -1, ParentType: ParentTypeObject},
			},
		},
	}

	for _, tc := range tests {
		members, err := ParseJSON(tc.json)
		if err != tc.err {
			t.Errorf("Expected error %v, actual %v", tc.err, err)
			continue
		}
		if len(members) != len(tc.expected) {
			t.Errorf("Expected members length %d, actual %d", len(tc.expected), len(members))
			continue
		}
		for i, member := range members {
			eq, msg := equalMember(member, tc.expected[i])
			if !eq {
				t.Errorf("Different %dth member from expected: %s", i, msg)
			}
		}
	}
}

// equalMember compare two members
func equalMember(m1, m2 Member) (bool, string) {
	if m1.Key != m2.Key {
		return false, fmt.Sprintf("Key first: '%s', second: '%s'", m1.Key, m2.Key)
	} else if m1.Start != m2.Start {
		return false, fmt.Sprintf("Start first: %d, second: %d", m1.Start, m2.Start)
	} else if m1.End != m2.End {
		return false, fmt.Sprintf("End first: %d, second: %d", m1.End, m2.End)
	} else if m1.Parent != m2.Parent {
		return false, fmt.Sprintf("Parent first: %d, second: %d", m1.Parent, m2.Parent)
	} else if m1.ParentType != m2.ParentType {
		return false, fmt.Sprintf("ParentType first: %d, second: %d", m1.ParentType, m2.ParentType)
	} else if m1.Index != m2.Index {
		return false, fmt.Sprintf("Index first: %d, second: %d", m1.Index, m2.Index)
	}
	return true, ""
}
