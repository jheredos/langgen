package main

import (
	"github.com/jheredos/langgen/phonotactics"
)

// Language is the top-level data structure wrapping a phonological inventory,
// the phonotactic tree, and everything else that comprises a language
type Language struct {
	ID string
	// phonology
	Consonants []string
	Vowels     []string
	// phonotactics
	PhonotacticTree *phonotactics.PhonotacticTreeNode
	WordGenerator   *phonotactics.WordGenerator
}
