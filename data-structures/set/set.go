package set

// StringSet maintains a deduped list of strings added to it
type StringSet map[string]bool

// NewStringSet initializes a set with the values form the input string slice
func NewStringSet(stringSlice []string) StringSet {
	set := StringSet{}
	for _, v := range stringSlice {
		set[v] = true
	}
	return set
}

// Has returns true if the value if found in the underlying set
func (s StringSet) Has(val string) bool {
	_, ok := s[val]
	return ok
}

// Add a value to the set
func (s StringSet) Add(val string) {
	s[val] = true
}

// Remove a value from the set
func (s StringSet) Remove(val string) {
	delete(s, val)
}

// Keys returns a slice of all keys in the set
func (s StringSet) Keys() []string {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

// IntSet maintains a deduped list of strings added to it
type IntSet map[int]bool

// NewIntSet initializes a set with the values form the input int slice
func NewIntSet(intSlice []int) IntSet {
	set := IntSet{}
	for _, v := range intSlice {
		set[v] = true
	}
	return set
}

// Has returns true if the value if found in the underlying set
func (s IntSet) Has(val int) bool {
	_, ok := s[val]
	return ok
}

// Add a value to the set
func (s IntSet) Add(val int) {
	s[val] = true
}

// Remove a value from the set
func (s IntSet) Remove(val int) {
	delete(s, val)
}

// Keys returns a slice of all keys in the set
func (s IntSet) Keys() []int {
	var keys []int
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
