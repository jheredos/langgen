package phonology

// WordBoundary is used as a dummy Phoneme for phonotactic trees to
// mark word boundaries and boundaries between syllables
type WordBoundary struct {
	Initial bool `json:"initial"`
}

// ToIPA for SyllableBoundary, either a "." for word-medial syllable breaks,
// or empty string for word-boundaries
func (b WordBoundary) ToIPA() string {
	return ""
}

// AsVowel always returns an empty Vowel and false for a WordBoundary receiver
func (b WordBoundary) AsVowel() (Vowel, bool) {
	return Vowel{}, false
}

// AsConsonant always returns an empty Consonant and false for a WordBoundary receiver
func (b WordBoundary) AsConsonant() (Consonant, bool) {
	return Consonant{}, false
}

// Match always returns whether a target Phoneme is a WordBoundary
// by checking that the target is not a Consonant or Vowel
func (b WordBoundary) Match(target Phoneme) bool {
	_, isVowel := target.AsVowel()
	_, isConsonant := target.AsConsonant()
	return !isVowel && !isConsonant
}
