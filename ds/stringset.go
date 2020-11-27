package ds

import "sync"

// New StringSet.
func New() *StringSet {
	s := StringSet{
		d: make(map[string]struct{}),
	}
	return &s
}

var special = struct{}{}

// StringSet is our struct that acts as a set data structure
// with string as members.
type StringSet struct {
	l sync.RWMutex
	d map[string]struct{}
}

// Add method to add a member to the StringSet.
func (s *StringSet) Add(member string) {
	s.l.Lock()
	defer s.l.Unlock()

	s.d[member] = special
}

// Remove method to remove a member from the StringSet.
func (s *StringSet) Remove(member string) {
	s.l.Lock()
	defer s.l.Unlock()

	delete(s.d, member)
}

// IsMember method to check if a member is present in the StringSet.
func (s *StringSet) IsMember(member string) bool {
	s.l.RLock()
	defer s.l.RUnlock()

	_, found := s.d[member]
	return found
}

// Members method to retrieve all members of the StringSet.
func (s *StringSet) Members() []string {
	s.l.RLock()
	defer s.l.RUnlock()

	keys := make([]string, 0)
	for k := range s.d {
		keys = append(keys, k)
	}
	return keys
}

// Size method to get the cardinality of the StringSet.
func (s *StringSet) Size() int {
	s.l.RLock()
	defer s.l.RUnlock()

	return len(s.d)
}

// Clear method to remove all members from the StringSet.
func (s *StringSet) Clear() {
	s.l.Lock()
	defer s.l.Unlock()

	s.d = make(map[string]struct{})
}
