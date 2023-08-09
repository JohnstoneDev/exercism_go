package stringset

import (
	"log"
	"fmt"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]string

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	set := make(Set, len(l))

	for _, elem := range l {
		set.Add(elem)
	}

	return set
}

func (s Set) String() string {
	ret := "{"
	for _, k := range s {
		if ret != "{" {
			ret += ", "
		}
		ret += fmt.Sprintf("%q", k)
	}
	ret += "}"
	return ret
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, found := s[elem]
	return found
}

func (s Set) Add(elem string) {
	if s.Has(elem) {
		log.Printf("%s already exists", elem)
	}

	s[elem] = elem
}

func Equal(s1, s2 Set) bool {
	equality := len(s1) == len(s2)

	if !equality {
		return false
	}

	return Subset(s1,s2)
}

func Subset(s1, s2 Set) bool {
	for _, elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}
	return true
}


func Difference(s1, s2 Set) Set {
	set := New()

	for _, elem := range s1 {
		if !s2.Has(elem) {
			set.Add(elem)
		}
	}

	return set
}

func Disjoint(s1, s2 Set) bool {
	return Intersection(s1,s2).IsEmpty()
}


func Intersection(s1, s2 Set) Set {
	set := make(Set, len(s1))

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	for _, elem := range s1 {
		if s2.Has(elem) {
			set.Add(elem)
		}
	}

	return set
}


// Creates a Set with elements from s1 & s2
func Union(s1, s2 Set) Set {
	set := make(Set, len(s1) + len(s2))

	for key := range s1 {
		set.Add(key)
	}

	for key := range s2 {
		set.Add(key)
	}

	return set
}
