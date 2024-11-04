package stringutil

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// StringSet defines set that contains some string
type StringSet map[string]struct{}

// ToSlice convert StringSet to slice
func (set StringSet) ToSlice() []string {
	keys := make([]string, 0, len(set))
	for k := range set {
		keys = append(keys, k)
	}
	return keys
}

// Size returns the length of a set
func (set StringSet) Size() int {
	return len(set)
}

// IsEmpty returns weather the set is empty
func (set StringSet) IsEmpty() bool {
	return len(set) == 0
}

// Add adds a string to set
func (set StringSet) Add(s string) {
	set[s] = struct{}{}
}

// Remove removes a string from set
func (set StringSet) Remove(s string) {
	delete(set, s)
}

// Contains returns weather a string is in the set
func (set StringSet) Contains(s string) bool {
	_, ok := set[s]
	return ok
}

// Equals returns weather 2 sets' elements are equal
func (set1 StringSet) Equals(set2 StringSet) bool {
	if len(set1) != len(set2) {
		return false
	}

	for k := range set1 {
		if _, ok := set2[k]; !ok {
			return false
		}
	}

	return true
}

// Intersection returns the intersection part of 2 sets
func (set1 StringSet) Intersection(set2 StringSet) StringSet {
	nset := NewStringSet()
	for k := range set1 {
		if _, ok := set2[k]; ok {
			nset.Add(k)
		}
	}

	return nset
}

// Difference returns the different part of 2 sets
func (set1 StringSet) Difference(set2 StringSet) StringSet {
	nset := NewStringSet()
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			nset.Add(k)
		}
	}

	return nset
}

// Union returns the set contains set1's and set2's elements
func (set1 StringSet) Union(set2 StringSet) StringSet {
	nset := NewStringSet()
	for k := range set1 {
		nset.Add(k)
	}

	for k := range set2 {
		nset.Add(k)
	}

	return nset
}

// MarshalJSON convert set to json
func (set StringSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.ToSlice())
}

// UnmarshalJSON convert a json format bytes to stringset structure
func (set *StringSet) UnmarshalJSON(data []byte) error {
	sl := []interface{}{}
	var err error
	if err = json.Unmarshal(data, &sl); err == nil {
		*set = make(StringSet)
		for _, s := range sl {
			set.Add(fmt.Sprintf("%v", s))
		}
	} else {
		var s interface{}
		if err = json.Unmarshal(data, &s); err == nil {
			*set = make(StringSet)
			set.Add(fmt.Sprintf("%v", s))
		}
	}

	return err
}

// String converts stringset to string type
// format like `[elem1 elem2 elem3]`
func (set StringSet) String() string {
	return fmt.Sprintf("%s", set.ToSlice())
}

// NewStringSet Initialize a stringset
func NewStringSet() StringSet {
	return make(StringSet)
}

// CreateStringSet creates a stringset contains given strings
func CreateStringSet(sl ...string) StringSet {
	set := make(StringSet)
	for _, k := range sl {
		set.Add(k)
	}
	return set
}

// CopyStringSet deep copy a set and return the copied one
func CopyStringSet(set StringSet) StringSet {
	nset := NewStringSet()
	for k, v := range set {
		nset[k] = v
	}
	return nset
}

// Scan implements the Scanner interface of gorm
func (set *StringSet) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return set.scanBytes(src)
	case string:
		return set.scanBytes([]byte(src))
	case nil:
		set = nil
		return nil
	}

	return fmt.Errorf("util: cannot convert %T to StringSet", src)
}

// scanBytes scan bytes to set
func (set *StringSet) scanBytes(src []byte) error {
	return set.UnmarshalJSON(src)
}

// Value implements the Valuer interface of gorm
func (set *StringSet) Value() (driver.Value, error) {
	return set.String(), nil
}
