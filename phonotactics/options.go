package phonotactics

// PhonotacticOptions holds a languages suprasegmental features like
// stress and tone, along with word length
type PhonotacticOptions struct {
	MinWordLength    WordLength              `json:"minWordLength"`
	MedianWordLength WordLength              `json:"medianWordLength"`
	MaxWordLength    WordLength              `json:"maxWordLength"`
	StressType       `json:"stressType"`     // Is there contrastive stress, is it predictable, and where does it fall?
	StressPosition   `json:"stressPosition"` // Is stress position determined by the stem or the whole word?
	ToneType         `json:"toneType"`       // Is there contrastive tone? Register or contour?
	TonePosition     `json:"tonePosition"`   // Are all syllables marked for tone, or just stressed (pitch accent)?
	ToneCategories   []ToneCategory          `json:"toneCategories"`
}

// WordLength is a categorical clasification of word length, from
// monosyllabic to XXL, 12+ syllables.
type WordLength uint8

// WordLength values
const (
	UnspecifiedWL  WordLength = iota
	MonosyllabicWL            // 1 syllable
	ShortWL                   // 2
	MediumWL                  // 3
	LongWL                    // 6
	XLongWL                   // 10
	XXLongWL                  // 15
)

// StressType encompasses whether a language has stress, if it is fixed
// or variable, and where it falls
type StressType uint8

// StressType values
const (
	UnspecifiedST StressType = iota
	NoneST
	InitialST
	SecondST
	ThirdST
	FinalST
	PenultimateST
	AntepenultimateST
	VariableST
)

// StressPosition is whether stress is determined by the syllables of the
// word or of the stem
type StressPosition uint8

// StressPosition values
const (
	UnspecifiedSP StressPosition = iota
	WordSP
	StemSP
)

// ToneType is whether a language has contrastive tone, and whether it
// has a register or contour system
type ToneType uint8

// ToneType values
const (
	UnspecifiedTT ToneType = iota
	NoneTT
	RegisterTT
	ContourTT
)

// TonePosition is whether tone is marked for all syllables or just for
// stressed syllables (i.e. pitch accent)
type TonePosition uint8

// TonePosition values
const (
	UnspecifiedTP TonePosition = iota
	SyllableTP
	AccentTP
)

// ToneCategory is an integer where each place represents a pitch level
// between 1-5, i.e. values 1-5 represent register tones, 11-55
// represent rising/falling contour tones, 111-555 rising+falling tones, etc.
// up to 5 places for complex tones
type ToneCategory uint16
