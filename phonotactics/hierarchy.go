package phonotactics

import "github.com/jheredos/langgen/phonology"

// ConsonantHierarchy is a ranked list of subsets of the consonant inventory,
// with lower indices indicating lower sonority and vice versa. If Onset is
// true, the hierarchy applies to syllable onset, otherwise to coda
type ConsonantHierarchy struct {
	Onset     bool                    `json:"onset"`
	NoCluster []phonology.Consonant   `json:"noCluster"`
	Tiers     [][]phonology.Consonant `json:"tiers"`
}

// NucleusHierarchy contains the phonemes that can form a syllabic nucleus.
// Onglides and Offglides pair with Nuclei to form diphthongs and triphthongs.
// Monopthongs includes all of a language's vowels. Consonants are any phonemes
// that act as syllabic consonants.
type NucleusHierarchy struct {
	Onglides     []phonology.Vowel     `json:"onglides"`
	Nuclei       []phonology.Vowel     `json:"nuclei"`
	Offglides    []phonology.Vowel     `json:"offglides"`
	Monophthongs []phonology.Vowel     `json:"monophthongs"`
	Consonants   []phonology.Consonant `json:"consonants"`
}
