package phonotactics

import "github.com/jheredos/langgen/phonology"

// RuleFrequency is an enum to describe the frequency of phonotactic rules
// (always or never) and trends (very seldom to very often)
type RuleFrequency uint8

// RuleFrequency values. Frequencies are relative to sibling edge weights
const (
	UnspecifiedRF RuleFrequency = iota
	NeverRF                     // edge weight set to 0
	VerySeldomRF                // factor of 5 ^ -2
	SeldomRF                    // factor of 5 ^ -1
	SometimesRF                 // factor of 5 ^ 0
	OftenRF                     // factor of 5 ^ +1
	VeryOftenRF                 // factor of 5 ^ +2
	AlwaysRF                    // all sibling edge weights are set to 0
)

// Presets

// SetInitialNullOnset sets null onsets at the start of a word to the frequency provided.
// In some languages it is more common for a word to start with a vowel than a consonant,
// like Yoruba. In others it is illegal, like Hawaiian or Arabic (though they often give
// the appearance of null onsets with a simple glottal stop as the onset)
func (n *PhonotacticTreeNode) SetInitialNullOnset(frequency RuleFrequency) {
	n.SetFrequencyForPattern(frequency, phonology.WordBoundary{}, phonology.Vowel{}, WordStartPC)
}

// SetFinalNullCoda sets the likelihood of a word ending in a vowel
func (n *PhonotacticTreeNode) SetFinalNullCoda(frequency RuleFrequency) {
	n.SetFrequencyForPattern(frequency, phonology.Vowel{}, phonology.WordBoundary{}, WordEndPC)
}

// SetHiatus sets the frequency of a vowel across a syllable boundary. This tends to be
// relatively infrequent, except perhaps in moraic languages
func (n *PhonotacticTreeNode) SetHiatus(frequency RuleFrequency) {
	n.SetFrequencyForPattern(frequency, phonology.Vowel{}, phonology.Vowel{}, SyllableBoundaryPC)
}
