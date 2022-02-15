package phonotactics

import (
	"math/rand"
	"time"
)

// WordGenerator wraps a Phonotactic tree
type WordGenerator struct {
	Root *PhonotacticTreeNode
}

// NewWordGenerator creates a new WordGenerator from the root
// of a phonotactic tree
func NewWordGenerator(root *PhonotacticTreeNode) *WordGenerator {
	wg := &WordGenerator{
		Root: root,
	}

	rand.Seed(time.Now().UnixNano())

	return wg
}

// NewWord generates a new word of the specified number of syllables as a string
func (g *WordGenerator) NewWord(syllables int) string {
	syllable := []*PhonotacticTreeNode{}
	prev := g.Root
	word := ""

	for i := 0; i < syllables; i++ {
		syllable, prev = prev.newSyllable(i == syllables-1)
		word += phonotacticPathToIPA(syllable)
		if i < syllables-1 {
			word += "."
		}
	}

	return word
}

func phonotacticPathToIPA(path []*PhonotacticTreeNode) string {
	s := ""
	for _, n := range path {
		s += n.Val.ToIPA()
	}
	return s
}

// newSyllable generates random nodes from the receiver node until hitting a syllable or word boundary.
// It returns a slice of nodes and the final node that crossed the boundary, either WordBoundary or the
// first node of the next syllable. The final param allows the caller to determine when to end the word
func (n *PhonotacticTreeNode) newSyllable(final bool) ([]*PhonotacticTreeNode, *PhonotacticTreeNode) {
	syll := []*PhonotacticTreeNode{n}
	node, boundary := n.randomNode(WordStartPC, OnsetPC, NucleusPC, CodaPC)

	for boundary != WordEndPC && boundary != SyllableBoundaryPC {
		syll = append(syll, node)
		if final {
			node, boundary = node.randomNode(OnsetPC, NucleusPC, CodaPC, WordEndPC)
		} else {
			node, boundary = node.randomNode(OnsetPC, NucleusPC, CodaPC, SyllableBoundaryPC)
		}
	}

	return syll, node
}

// randomNode returns a random child of the receiver node over any PhonotacticContext specified
// in the params, according to the weights of those edges
func (n *PhonotacticTreeNode) randomNode(boundaries ...PhonotacticContext) (*PhonotacticTreeNode, PhonotacticContext) {
	var wsum float32 = 0
	edges := []*PhonotacticTreeEdge{}

	for _, edge := range n.Children {
		for _, b := range boundaries {
			if edge.Boundary == b {
				edges = append(edges, edge)
				wsum += edge.Weight
			}
		}
	}

	k := rand.Float32() * wsum
	for _, edge := range edges {
		k -= edge.Weight
		if k <= 0 {
			return edge.ChildNode, edge.Boundary
		}
	}

	return edges[len(edges)-1].ChildNode, edges[len(edges)-1].Boundary
}

func round(n float32) int {
	return int(n*10) / 10
}

//
func skewedRand(max int) int {
	a := rand.Float32() * float32(max)
	b := rand.Float32() * float32(max)
	if a < b {
		return round(a)
	}
	return round(b)
}

// GetWordLength ... this seems to skew really high...
func GetWordLength(min, median, max WordLength) int {
	lens := []int{1, 1, 2, 3, 6, 10, 15}
	n := lens[median] + skewedRand(lens[median]/2+1) - skewedRand(lens[median]/2+1)
	return n - skewedRand(lens[min]) + skewedRand(lens[max])
}
