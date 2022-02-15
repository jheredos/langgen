package phonology

// Inventory represents the phonological inventory of a language,
// essentially just lists of consonant and vowel phonemes
type Inventory struct {
	LanguageID string      `json:"lang_id" bson:"lang_id"`
	Vowels     []Vowel     `json:"vowels" bson:"vowels"`
	Consonants []Consonant `json:"consonants" bson:"consonants"`
}

// InventoryIPA ...
type InventoryIPA struct {
	LanguageID string   `json:"lang_id" bson:"lang_id"`
	Vowels     []string `json:"vowels" bson:"vowels"`
	Consonants []string `json:"consonants" bson:"consonants"`
}

// ToIPA converts a Inventory struct to an InventoryIPA struct to send to the frontend
func (i *Inventory) ToIPA() InventoryIPA {
	res := InventoryIPA{LanguageID: i.LanguageID}
	for _, v := range i.Vowels {
		res.Vowels = append(res.Vowels, v.ToIPA())
	}
	for _, c := range i.Consonants {
		res.Consonants = append(res.Consonants, c.ToIPA())
	}
	return res
}

func (i *Inventory) addPhoneme(s string) {
	if isIPAVowel(s) {
		v, _ := NewVowelFromIPA(s)
		for _, vw := range i.Vowels {
			if v == vw {
				return
			}
		}
		i.Vowels = append(i.Vowels, v)
	} else if isIPAConsonant(s) {
		c, _ := NewConsonantFromIPA(s)
		for _, ct := range i.Consonants {
			if c == ct {
				return
			}
		}
		i.Consonants = append(i.Consonants, c)
	}
}

// addPhonemes takes a slice of string representing IPA phonemes and
// adds them to the receiver Inventory
func (i *Inventory) addPhonemes(ps []string) {
	for _, p := range ps {
		i.addPhoneme(p)
	}
}

// NewInventory creates a pointer to a phonological Inventory out of
// a slice of IPA strings
func NewInventory(phonemes []string) *Inventory {
	inv := &Inventory{
		Consonants: []Consonant{},
		Vowels:     []Vowel{},
	}

	inv.addPhonemes(phonemes)

	return inv
}
