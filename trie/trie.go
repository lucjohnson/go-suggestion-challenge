package trie 

import (
    "fmt"
    "strings"    
    "sort"
)

type Trie struct {
    root *node
    size int
}

type node struct {
    validWord bool
    nodes map[rune]*node
}

// AddEntry adds a string to the given trie
func (trie *Trie) AddEntry(entry string) {
    n := trie.root
    entry = strings.ToLower(entry)
    for idx, r := range entry {
        // check to see if this is the last character
        fullWord := idx == len(entry) - 1
        
        // if there is nothing there in this node's map, make it
        if (n.nodes[r] == nil) {
            n.nodes[r] = &node {nodes: make(map[rune]*node), validWord: fullWord}
        } else if (n.nodes[r] != nil && fullWord) {
            n.nodes[r].validWord = true
        }
        n = n.nodes[r]
    }
    trie.size++
}

// GetSize returns the size of the given trie as an int
func (trie *Trie) GetSize() int {
    return trie.size
}

// FindEntriesHelper is a function that enables FindEntries to do so recursively
func (trie *Trie) FindEntriesHelper(node *node, entries *[]string, prefix string) []string { 
    var keys []rune
    for k := range node.nodes {
        keys = append(keys, k)
    }
    sort.Sort(RuneSlice(keys))
    
    for _, value := range keys {
        prefix := prefix + string(value)
        if (node.nodes[value].validWord) {
            *entries = append(*entries, prefix)
        }
        if (len(node.nodes[value].nodes) > 0) {
            *entries = trie.FindEntriesHelper(node.nodes[value], entries, prefix)
        }
    }
    return *entries    
}

// FindEntries takes a given string and returns an array of strings that contain the given prefix string
func (trie *Trie) FindEntries(prefix string, max int) []string  {
    var entries []string
    n := trie.root
    
    prefix = strings.ToLower(prefix)

    // this gets you to where the possible entries are
    for _, r := range prefix {
        if (n.nodes[r] == nil) {
            fmt.Println("'" + prefix + "'", "prefix does not exist")
            return nil
        } 
        n = n.nodes[r]
    }
    if (n.validWord) {
        entries = append(entries, prefix)
    }
    entries = trie.FindEntriesHelper(n, &entries, prefix)
    
    if len(entries) < max {
        return entries
    } 
    return entries[:max] 
}

// RuneSlice is defined in order to properly sort a slice of runes
type RuneSlice []rune

func (p RuneSlice) Len() int { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap (i, j int) { p[i], p[j] = p[j], p[i] }

// NewTrie initializes a new, empty trie
func NewTrie() *Trie {
    return &Trie {root: &node{nodes: make(map[rune]*node)}, size: 0}
}
