package phonology

// Phoneme is a consonant or a vowel that can be represented with
// IPA, and perhaps with XSAMPA and various orthographies in the future
// The Match method returns whether a Phoneme matches the specified features,
// and the asVowel, asConsonant methods are slightly hacky ways to make
// Match work across Phoneme types
type Phoneme interface {
	ToIPA() string
	Match(Phoneme) bool
	asVowel() (Vowel, bool)
	asConsonant() (Consonant, bool)
}

// asVowel for a Vowel receiver simply forwards the Vowel struct and returns true
// for the second parameter to denote that the receiver is indeed a Vowel
func (v Vowel) asVowel() (Vowel, bool) {
	return v, true
}

// asConsonant for a Vowel receiver returns an empty Consonant struct and false
// to denote that the receiver is not actually a Consonant
func (v Vowel) asConsonant() (Consonant, bool) {
	return Consonant{}, false
}

// asVowel for a Consonant receiver returns an empty Vowel struct and false
// to denote that the receiver is not actually a Vowel
func (c Consonant) asVowel() (Vowel, bool) {
	return Vowel{}, false
}

// asConsonant for a Consonant receiver simply forwards the Consonant struct and returns
// true for the second parameter to denote that the receiver is indeed a Consonant
func (c Consonant) asConsonant() (Consonant, bool) {
	return c, true
}

// Match returns true if the Vowel matches the specified features in the pattern
// Each field's zero value denotes an unspecified feature, such that Vowel{} as
// a pattern param will match any vowel. If the pattern is not a Vowel, however,
// it will be caught with the asVowel() Phoneme method and return false
// Note that the pattern param may be partially described, but the receiver should not
func (v Vowel) Match(pattern Phoneme) bool {
	patternv, isVowel := pattern.asVowel()
	if !isVowel {
		return false
	}
	if patternv.Height != 0 && v.Height != patternv.Height {
		return false
	}
	if patternv.Frontness != 0 && v.Frontness != patternv.Frontness {
		return false
	}
	if patternv.Phonation != 0 && v.Phonation != patternv.Phonation {
		return false
	}
	if patternv.Rounding != 0 && v.Rounding != patternv.Rounding {
		return false
	}
	if patternv.Nasal != 0 && v.Nasal != patternv.Nasal {
		return false
	}
	if patternv.Length != 0 && v.Length != patternv.Length {
		return false
	}
	return true
}

// Match returns true if the Consonant matches the specified features in the pattern
// Each field's zero value denotes an unspecified feature, such that Consonant{} as
// a pattern param will match any consonant. If the pattern is not a Consonant, however,
// it will be caught with the asConsonant() Phoneme method and return false
// Note that the pattern param may be partially described, but the receiver should not
func (c Consonant) Match(pattern Phoneme) bool {
	patternc, isConsonant := pattern.asConsonant()
	if !isConsonant {
		return false
	}
	if patternc.Place != 0 && c.Place != patternc.Place {
		return false
	}
	if patternc.Manner != 0 && c.Manner != patternc.Manner {
		return false
	}
	if patternc.Coarticulation != 0 && c.Coarticulation != patternc.Coarticulation {
		return false
	}
	if patternc.NonPulmonic != 0 && c.NonPulmonic != patternc.NonPulmonic {
		return false
	}
	if patternc.Voiced != 0 && c.Voiced != patternc.Voiced {
		return false
	}
	if patternc.Aspirated != 0 && c.Aspirated != patternc.Aspirated {
		return false
	}
	if patternc.Lateral != 0 && c.Lateral != patternc.Lateral {
		return false
	}
	if patternc.Sibilant != 0 && c.Sibilant != patternc.Sibilant {
		return false
	}
	if patternc.Geminate != 0 && c.Geminate != patternc.Geminate {
		return false
	}
	return true
}
