package trie_test

import (
	"testing"
	"github.com/lucjohnson/challenges-lucjohnson/suggestion-challenge/trie"
)

func TestBuildTrie(t *testing.T) {
	trie := trie.NewTrie()
	if trie == nil {
		t.Error("trie should have been made")
	}
}

func TestAddEntry(t *testing.T) {
	trie := trie.NewTrie()
	trie.AddEntry("hello")
	trie.AddEntry("h")
	trie.AddEntry("HOLa")
	trie.AddEntry("!hoLLa!")
}

func TestFindEntries(t *testing.T) {
	trie := trie.NewTrie()
	trie.AddEntry("h")
	trie.AddEntry("hello")
	trie.AddEntry("HOLa")
	trie.AddEntry("!hoLLa!")
	
	test1 := trie.FindEntries("h", 10)
	if len(test1) != 3 {
		t.Error("There was an issue finding entries")
	}
	
	test2 := trie.FindEntries("p", 10)
	if len(test2) > 0 {
		t.Error("There should be no 'p' words")
	}
	
	test3 := trie.FindEntries("h", 1)
	if len(test3) > 1 {
		t.Error("Should have only received 1 result")
	}
	
	test4 := trie.FindEntries("H", 1)
	if test4[0] != "h" {
		t.Error("Check to make sure no casing issues")
	}
	
	test5 := trie.FindEntries("h", 10)
	if test5[0] != "h" && test5[1] != "hello" {
		t.Error("Entries were not returned in alphabetical order")
	}
}

func TestGetSize(t *testing.T) {
	trie := trie.NewTrie()
	trie.AddEntry("a")
	trie.AddEntry("b")
	trie.AddEntry("c")
	
	if trie.GetSize() != 3 {
		t.Error("Size is off")
	}
}