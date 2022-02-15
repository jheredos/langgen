package phonotactics

import (
	"github.com/jheredos/langgen/phonology"
)

// PhonotacticTreeNode is a node for a tree with weighted edges for
// generating sequences of phonemes. Consonants may have separate
// nodes for onset and coda position
type PhonotacticTreeNode struct {
	ID       string                 `json:"id"`
	Val      phonology.Phoneme      `json:"val"`
	Children []*PhonotacticTreeEdge `json:"-"`
	ChildIDs []string               `json:"childIds"`
	// Position?
}

// PhonotacticTreeEdge is the edge of a phonotactic tree with three
// weights depending on the syllable's position in the word
type PhonotacticTreeEdge struct {
	ID        string               `json:"id"`
	ChildNode *PhonotacticTreeNode `json:"-"`
	ChildID   string               `json:"childId"`
	Boundary  PhonotacticContext   `json:"boundary"`
	Weight    float32              `json:"weight"`
}

// PhonotacticContext as a field on a PhonotacticTreeEdge denotes
// an edge which crosses either from a WordBoundary to a Phoneme, or
// crosses over a syllable boundary
type PhonotacticContext uint8

// PhonotacticContext values
const (
	UnspecifiedPC PhonotacticContext = iota
	OnsetPC                          // Onset, Nucleus, and Coda refer to the start of the edge
	NucleusPC
	CodaPC
	WordStartPC
	WordEndPC
	SyllableBoundaryPC
)

func newPhonotacticTreeEdge(p *PhonotacticTreeNode, boundary PhonotacticContext) *PhonotacticTreeEdge {
	edge := &PhonotacticTreeEdge{
		ChildNode: p,
		Boundary:  boundary,
		Weight:    1,
	}
	return edge
}

func newConsonantNodeSlice(cs []phonology.Consonant) []*PhonotacticTreeNode {
	nodes := []*PhonotacticTreeNode{}
	for _, c := range cs {
		nodes = append(nodes, &PhonotacticTreeNode{Val: c})
	}
	return nodes
}

func newVowelNodeSlice(vs []phonology.Vowel) []*PhonotacticTreeNode {
	nodes := []*PhonotacticTreeNode{}
	for _, v := range vs {
		nodes = append(nodes, &PhonotacticTreeNode{Val: v})
	}
	return nodes
}

// NewPhonotacticTree creates a new phonotactic tree with uniformly
// weighted edges from three sonority hierarchies for the onset,
// nucleus, and coda, returning the root of the tree
func NewPhonotacticTree(onset ConsonantHierarchy, nucleus NucleusHierarchy, coda ConsonantHierarchy) (*PhonotacticTreeNode, error) {
	root := &PhonotacticTreeNode{ // Word start
		Val: phonology.WordBoundary{
			Initial: true,
		},
		Children: []*PhonotacticTreeEdge{},
	}

	end := &PhonotacticTreeNode{ // Word end
		Val: phonology.WordBoundary{
			Initial: false,
		},
		Children: []*PhonotacticTreeEdge{},
	}

	onsetRoots, onsetLeaves := createOnsets(onset)
	nucleusRoots, nucleusLeaves := createNuclei(nucleus)
	codaRoots, codaLeaves := createCodas(coda)

	// 1	# -> O
	attachNodes([]*PhonotacticTreeNode{root}, onsetRoots, WordStartPC)
	// 2	# -> N
	attachNodes([]*PhonotacticTreeNode{root}, nucleusRoots, WordStartPC)
	// 3 	O -> N
	attachNodes(onsetLeaves, nucleusRoots, OnsetPC)
	// 4	N -> C
	attachNodes(nucleusLeaves, codaRoots, NucleusPC)
	// 5	C -> #
	attachNodes(codaLeaves, []*PhonotacticTreeNode{end}, WordEndPC)
	// 6	N -> #
	attachNodes(nucleusLeaves, []*PhonotacticTreeNode{end}, WordEndPC)
	// 7	N -> O
	attachNodes(nucleusLeaves, onsetRoots, SyllableBoundaryPC)
	// 8	N -> N
	attachNodes(nucleusLeaves, nucleusRoots, SyllableBoundaryPC)
	// 9	C -> O
	attachNodes(codaLeaves, onsetRoots, SyllableBoundaryPC)
	// 10	C -> N
	attachNodes(codaLeaves, nucleusRoots, SyllableBoundaryPC)

	return root, nil
}

func createOnsets(hierarchy ConsonantHierarchy) ([]*PhonotacticTreeNode, []*PhonotacticTreeNode) {
	roots, leaves := []*PhonotacticTreeNode{}, []*PhonotacticTreeNode{}

	// convert hierarchy to PhonotacticTreeNode
	tiers := [][]*PhonotacticTreeNode{}
	for _, tier := range hierarchy.Tiers {
		tiers = append(tiers, newConsonantNodeSlice(tier))
	}
	noCluster := newConsonantNodeSlice(hierarchy.NoCluster)

	// attach nodes
	for _, tier := range tiers {
		attachNodes(roots, tier, OnsetPC)
		roots = append(roots, tier...)
		leaves = append(leaves, tier...)
	}
	roots = append(roots, noCluster...)
	leaves = append(leaves, noCluster...)
	return roots, leaves
}

func createNuclei(hierarchy NucleusHierarchy) ([]*PhonotacticTreeNode, []*PhonotacticTreeNode) {
	roots, leaves := []*PhonotacticTreeNode{}, []*PhonotacticTreeNode{}

	onglides := newVowelNodeSlice(hierarchy.Onglides)
	nuclei := newVowelNodeSlice(hierarchy.Nuclei)
	offglides := newVowelNodeSlice(hierarchy.Offglides)
	monophthongs := newVowelNodeSlice(hierarchy.Monophthongs)
	consonants := newConsonantNodeSlice(hierarchy.Consonants)

	// Onglides
	roots = append(roots, onglides...)
	attachNodes(onglides, nuclei, NucleusPC)
	// Nuclei
	roots = append(roots, nuclei...)
	leaves = append(leaves, nuclei...)
	// Offglides
	attachNodes(nuclei, offglides, NucleusPC)
	leaves = append(leaves, offglides...)
	// Monophthongs
	roots = append(roots, monophthongs...)
	leaves = append(leaves, monophthongs...)
	// Syllabic Consonants
	roots = append(roots, consonants...)
	leaves = append(leaves, consonants...)

	return roots, leaves
}

func createCodas(hierarchy ConsonantHierarchy) ([]*PhonotacticTreeNode, []*PhonotacticTreeNode) {
	roots, leaves := []*PhonotacticTreeNode{}, []*PhonotacticTreeNode{}

	// convert hierarchy to PhonotacticTreeNode
	tiers := [][]*PhonotacticTreeNode{}
	for _, tier := range hierarchy.Tiers {
		tiers = append(tiers, newConsonantNodeSlice(tier))
	}
	noCluster := newConsonantNodeSlice(hierarchy.NoCluster)

	for _, tier := range tiers {
		attachNodes(roots, tier, CodaPC)
		roots = append(roots, tier...)
		leaves = append(leaves, tier...)
	}
	roots = append(roots, noCluster...)
	leaves = append(leaves, noCluster...)
	return roots, leaves
}

// attachNodes attaches all node(s) of the second param as children to all the nodes in the first,
// with the phonotactic boundaries that the new edges may cross described by the variadic final param
func attachNodes(parents []*PhonotacticTreeNode, nodes []*PhonotacticTreeNode, contexts ...PhonotacticContext) {
	for _, parent := range parents {
		for _, node := range nodes {
			for _, context := range contexts {
				parent.Children = append(parent.Children, newPhonotacticTreeEdge(node, context))
			}
			if len(contexts) == 0 {
				parent.Children = append(parent.Children, newPhonotacticTreeEdge(node, UnspecifiedPC))
			}
		}
	}
}
