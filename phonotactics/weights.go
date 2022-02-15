package phonotactics

import "github.com/jheredos/langgen/phonology"

// SetFrequencyForPattern finds all occurrences in a phonotactic tree of phoneme pattern A followed
// by pattern B and sets the weights of those edges according to the desired frequency. The patterns
// can be Vowels or Consonants, fully or partially described, and the caller can optionally specify
// the context(s), i.e. syllable internal or cross-syllable
func (n *PhonotacticTreeNode) SetFrequencyForPattern(frequency RuleFrequency, patternA phonology.Phoneme, patternB phonology.Phoneme, contexts ...PhonotacticContext) {
	nodes, edges := n.findPattern(patternA, patternB, contexts)
	setWeights(nodes, edges, frequency)
}

// findPattern finds a pattern of two phonemes, a and b, across given contexts (syllable internal,
// word-final, etc). The phonemes can be consonants, vowels, or word boundaries, including partially
// specified phonemes, like unvoiced consonants or rounded vowels
func (n *PhonotacticTreeNode) findPattern(patternA phonology.Phoneme, patternB phonology.Phoneme, contexts []PhonotacticContext) ([]*PhonotacticTreeNode, []*PhonotacticTreeEdge) {
	aNodes := n.findPhoneme(patternA)
	edges := []*PhonotacticTreeEdge{}

	for _, node := range aNodes {
		for _, edge := range node.Children {
			if len(contexts) == 0 && edge.ChildNode.Val.Match(patternB) {
				edges = append(edges, edge)
				continue
			}
			for _, context := range contexts {
				if edge.Boundary == context && edge.ChildNode.Val.Match(patternB) {
					edges = append(edges, edge)
					break
				}
			}
		}
	}

	return aNodes, edges
}

// findPhoneme finds all occurrences of a given phoneme in a phonotactic tree. The phoneme
// can be partially specified, like all unvoiced consonants (i.e. Consonant{Voiced: VoicedCV})
// or all rounded vowels (Vowel{Rounding: RoundedVR})
func (n *PhonotacticTreeNode) findPhoneme(pattern phonology.Phoneme) []*PhonotacticTreeNode {
	nodes := []*PhonotacticTreeNode{}
	seen := map[*PhonotacticTreeNode]bool{}

	row, nextRow := []*PhonotacticTreeNode{n}, []*PhonotacticTreeNode{}
	for len(row) > 0 {
		for _, node := range row {
			if node.Val.Match(pattern) {
				nodes = append(nodes, node)
			}
			for _, edge := range node.Children {
				if _, alreadySeen := seen[edge.ChildNode]; alreadySeen {
					continue
				}
				nextRow = append(nextRow, edge.ChildNode)
				seen[edge.ChildNode] = true
			}
		}
		row, nextRow = nextRow, []*PhonotacticTreeNode{}
	}

	return nodes
}

// setWeights adjusts the weights of a slice of phonotactic tree edges by a float32 weight
func setWeights(nodes []*PhonotacticTreeNode, edges []*PhonotacticTreeEdge, frequency RuleFrequency) {
	var weight float32 = 1
	var factor float32 = 2
	switch frequency {
	case NeverRF:
		weight = 0
	case VerySeldomRF:
		weight = 1 / (factor * factor) // factor ^ -2
	case SeldomRF:
		weight = 1 / factor // factor ^ -1
	case SometimesRF:
		weight = 1 // factor ^ 0
	case OftenRF:
		weight = factor // factor ^ 1
	case VeryOftenRF:
		weight = factor * factor // factor ^ 2
	}
	if frequency < AlwaysRF {
		for _, edge := range edges {
			edge.Weight = weight
		}
	} else {
		alwaysEdges := map[*PhonotacticTreeEdge]bool{}
		for _, edge := range edges {
			alwaysEdges[edge] = true
		}
		for _, node := range nodes {
			for _, edge := range node.Children {
				if _, always := alwaysEdges[edge]; !always {
					edge.Weight = 0
				}
			}
		}
	}
}
