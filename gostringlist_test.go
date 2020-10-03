package gostringlist

import (
	"strings"
	"testing"
)

func TestIsEqual(t *testing.T) {
	var list, list2 StringList

	list.Items = []string{"a", "b", "c"}
	list2.Items = []string{"a", "b", "c"}

	if !list.IsEqual(list2) {
		t.Fatal("IsEqual failed to compare two equal lists")
	}

	list2.Items = []string{"a", "b", "c", "d"}
	if list.IsEqual(list2) {
		t.Fatal("IsEqual failed to compare two different lists")
	}
}

func TestIndexOf(t *testing.T) {
	var list StringList

	list.Items = []string{"a", "b", "c", "d"}

	if list.IndexOf("c") != 2 {
		t.Fatal("IndexOf failed to return correct index")
	}

	if list.IndexOf("foo") >= 0 {
		t.Fatal("IndexOf failed to return correct index for nonexistent entry")
	}
}

func TestContains(t *testing.T) {
	var list StringList

	list.Items = []string{"a", "b", "c", "d"}

	if !list.Contains("c") {
		t.Fatal("Contains failed to find string")
	}

	if list.Contains("foo") {
		t.Fatal("Contains found a nonexistent entry")
	}
}

func TestInsert(t *testing.T) {
	var list, list2 StringList

	list.Items = []string{"a", "b", "d"}
	list2.Items = []string{"a", "b", "c", "d"}

	err := list.Insert("c", 2)
	if err != nil {
		t.Fatal("Bad index checking in Insert")
	}

	if !list.IsEqual(list2) {
		t.Fatalf("Insert failed to add item correctly\n%s\n", list.ToString())
	}
}

func TestRemove(t *testing.T) {
	var list, list2 StringList

	list.Items = []string{"a", "b", "c", "d"}
	list2.Items = []string{"a", "b", "d"}

	list.Remove("c")
	if !list.IsEqual(list2) {
		t.Fatalf("Remove failed to remove item correctly\n%s\n", list.ToString())
	}
}

func TestRemoveUnordered(t *testing.T) {
	var list, list2 StringList

	list.Items = []string{"a", "b", "c", "d", "e"}
	list2.Items = []string{"a", "b", "e", "d"}

	list.RemoveUnordered("c")
	if !list.IsEqual(list2) {
		t.Fatalf("RemoveUnordered failed to remove item correctly\n%s\n", list.ToString())
	}
}

func TestSort(t *testing.T) {
	var list, list2 StringList

	list.Items = []string{"b", "a", "d", "c"}
	list2.Items = []string{"a", "b", "c", "d"}

	list.Sort()
	if !list.IsEqual(list2) {
		t.Fatalf("Sort failed to sort list correctly\n%s\n", list.ToString())
	}
}

// Capitalization filter for TestFilter test
func toUpperOp(index int, list []string) (bool, string) {
	return true, strings.ToUpper(list[index])
}

func TestFilter(t *testing.T) {
	var list, list2 StringList

	list.Items = []string{"a", "b", "c", "d"}
	list2.Items = []string{"A", "B", "C", "D"}

	list = list.Filter(toUpperOp)
	if !list.IsEqual(list2) {
		t.Fatalf("Filter failed to process list correctly\n%s\n", list.ToString())
	}
}

func TestMatchFilter(t *testing.T) {
	var inlist, compareList StringList

	inlist.Items = []string{"apple", "Banana", "orange", "Pear"}
	compareList.Items = []string{"Banana", "Pear"}
	var outlist StringList

	var err error
	outlist, err = inlist.MatchFilter("[[:upper:]][[:lower:]]*")
	if err != nil || !outlist.IsEqual(compareList) {
		t.Fatalf("MatchFilter failed to process list correctly\n%s\n", outlist.ToString())
	}

	// Test case where there are no matches
	outlist, err = inlist.MatchFilter("[[:digit:]]")
	if err != nil || !outlist.IsEmpty() {
		t.Fatalf("MatchFilter failed to handle no matches correctly\n%s\n", outlist.ToString())
	}
}

func TestReplaceFilter(t *testing.T) {
	var inlist, compareList StringList

	inlist.Items = []string{"apple", "Banana", "orange", "Pear"}
	compareList.Items = []string{"Apple", "BAnAnA", "orAnge", "PeAr"}
	var outlist StringList

	var err error
	outlist, err = inlist.ReplaceAllFilter("a", "A")
	if err != nil || !outlist.IsEqual(compareList) {
		t.Fatalf("MatchFilter failed to process list correctly\n%s\n", outlist.ToString())
	}
}
