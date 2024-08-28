package set

import (
	"fmt"
	"sort"
	"strings"
)

type Set struct {
	elements map[int]bool
}

func NewSet() *Set {
	return &Set{
		elements: make(map[int]bool),
	}
}

func (s *Set) Elements() []int {
	keys := make([]int, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}

func (s Set) Clone() *Set {
	clone := NewSet()
	for key := range s.elements {
		clone.Add(key)
	}
	return clone
}

func (s Set) Add(val int) {
  s.elements[val] = true
}

func (s Set) Remove(val int) {
  delete(s.elements, val)
}

func (s Set) Contains(val int) bool {
	_, ok := s.elements[val]
	return ok
}

func (s Set) String() string {
	var elements []int
	for key := range s.elements {
		elements = append(elements, key)
	}
	return "{" + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(elements)), ", "), "[]") + "}"
}
