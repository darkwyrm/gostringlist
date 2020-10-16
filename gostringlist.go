// Package gostringlist - an object to make managing lists of strings more convenient. In some
// cases, convenient frontends for functionality defined in other packages are provided
//
// ©2020 Jon Yoder <jsyoder@mailfence.com>
// Released under the MIT License
package gostringlist

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

// StringList - a class for managing and manipulating lists of strings. It stores the items
// internally as a slice, so memory management caveats apply.
type StringList struct {
	Items []string
}

// New creates a new empty list
func New() *StringList {
	var newList StringList
	newList.Items = make([]string, 0)
	return &newList

// Append appends a string to the list
func (list *StringList) Append(str string) {
	list.Items = append(list.Items, str)
}

// Copy creates a duplicate of the existing object
func (list StringList) Copy() StringList {
	var newList StringList
	copy(newList.Items, list.Items)
	return newList
}

// IsEqual returns true if the current object's items match exactly those of the passed StringList.
func (list StringList) IsEqual(list2 StringList) bool {
	if len(list.Items) != len(list2.Items) {
		return false
	}

	for i := range list.Items {
		if list.Items[i] != list2.Items[i] {
			return false
		}
	}

	return true
}

// IsEmpty returns true if the object contains no items.
func (list StringList) IsEmpty() bool {
	return len(list.Items) == 0
}

// IndexOf returns the index of item if the list contains an exact match of the specified string
// or -1 if not found
func (list StringList) IndexOf(str string) int {
	for i, v := range list.Items {
		if str == v {
			return i
		}
	}
	return -1
}

// ToString converts the list to a string conveying its contents
func (list StringList) ToString() string {
	parts := make([]string, (len(list.Items)*2)+1)
	parts[0] = "["
	partsIndex := 1
	for _, v := range list.Items {
		parts[partsIndex] = "\"" + v + "\""
		parts[partsIndex+1] = ","
		partsIndex += 2
	}
	parts[len(parts)-1] = "]"
	return strings.Join(parts, "")
}

// Contains returns true if the list contains an exact match of the specified string
func (list StringList) Contains(str string) bool {
	for _, v := range list.Items {
		if str == v {
			return true
		}
	}
	return false
}

// Insert inserts the specified string into the list at the specified index. Like Remove(), this
// method performs some memory reallocations and copying as needed in order to provide convenience.
// As such, it is expensive, and if you need to do a lot of insertions, a task-specific
// implementation will perform better.
func (list *StringList) Insert(str string, index int) error {
	if index < 0 || index > len(list.Items) {
		return errors.New("index out of range")
	}
	list.Items = append(list.Items, str)
	copy(list.Items[index+1:], list.Items[index:])
	list.Items[index] = str

	return nil
}

// Remove deletes the string from the list. This method removes the item by copying each element
// after the one deleted to the slot before it, which is the method recommended from The Go
// Programming Language. Speed is sacrificed for the sake of convenience. If you intend to do a lot
// of removal, you are better off implementing a task-specific version.
func (list *StringList) Remove(str string) {
	index := list.IndexOf(str)
	if index < 0 {
		return
	}

	copy(list.Items[index:], list.Items[index+1:])
	list.Items = list.Items[:len(list.Items)-1]
}

// RemoveUnordered deletes the string from the list like Remove(), but it rearranges the items for
// the sake of speed. If the order of the list items doesn't matter, this method should be
// preferred over Remove()
func (list *StringList) RemoveUnordered(str string) {
	index := list.IndexOf(str)
	if index < 0 {
		return
	}

	length := len(list.Items)
	list.Items[index] = list.Items[length-1]
	list.Items = list.Items[:length-1]
}

// Sort - sorts the list in ascending alphabetical order
func (list *StringList) Sort() {
	sort.Strings(list.Items)
}

// Join - convenience function to return all items joined by the specified character
func (list StringList) Join(sep string) string {
	return strings.Join(list.Items, sep)
}

// Filter is a generic interface to creating new StringLists from the original, similar to Python's
// list comprehensions. It takes a pointer to a filter function. The filter function is passed
// an index to the current item and the slice of strings to be used as the source. It is expected
// to return a boolean and a string. The boolean value specifies whether the returned string is to
// be added to the filtered list.
func (list StringList) Filter(op func(int, []string) (bool, string)) StringList {
	var newList StringList
	newList.Items = make([]string, 0, len(list.Items))
	for i := range list.Items {
		addItem, out := op(i, list.Items)
		if addItem {
			newList.Items = append(newList.Items, out)
		}
	}

	return newList
}

// MatchFilter returns a new StringList containing all the items in the list which match the
// supplied regular expression
func (list StringList) MatchFilter(pattern string) (StringList, error) {
	var newList StringList
	re, err := regexp.Compile(pattern)
	if err != nil {
		return newList, err
	}

	newList.Items = make([]string, 0, len(list.Items))

	for _, item := range list.Items {
		if re.MatchString(item) {
			newList.Items = append(newList.Items, item)
		}
	}

	return newList, nil
}

// ReplaceAllFilter returns a new StringList containing all the items in the list which match the
// supplied regular expression
func (list StringList) ReplaceAllFilter(pattern string, repl string) (StringList, error) {
	var newList StringList
	newList.Items = make([]string, 0, len(list.Items))

	re, err := regexp.Compile(pattern)
	if err != nil {
		return newList, err
	}
	for i := range list.Items {
		newString := re.ReplaceAllString(list.Items[i], repl)
		newList.Items = append(newList.Items, newString)
	}

	return newList, nil
}
