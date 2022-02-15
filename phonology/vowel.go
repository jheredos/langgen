package phonology

import (
	"encoding/json"
)

// Vowel represents a vowel phoneme
type Vowel struct {
	Height    VowelHeight    `json:"height" bson:"height"`       // Close, NearClose, CloseMid, etc.
	Frontness VowelFrontness `json:"frontness" bson:"frontness"` // Front, Mid, Back
	Phonation VowelPhonation `json:"phonation" bson:"phonation"` // Modal, Breathy, Creaky, Devoiced
	Rounding  VowelRounding  `json:"rounding" bson:"rounding"`   // Rounded / Unrounded
	Nasal     VowelNasality  `json:"nasal" bson:"nasal"`         // Nasal / Oral
	Length    VowelLength    `json:"length" bson:"length"`       // Long / Short
}

// VowelHeight is the height of the vowel
type VowelHeight uint8

// VowelHeight values
const (
	UnspecifiedVH VowelHeight = iota
	CloseVH
	NearCloseVH
	CloseMidVH
	MidVH
	OpenMidVH
	NearOpenVH
	OpenVH
)

// VowelFrontness is the frontness of the vowel
type VowelFrontness uint8

// VowelFrontness values
const (
	UnspecifiedVF VowelFrontness = iota
	FrontVF
	CentralVF
	BackVF
)

// VowelPhonation includes features such as creaky voice and devoiced vowels
type VowelPhonation uint8

// VowelPhonation values
const (
	UnspecifiedVP VowelPhonation = iota
	ModalVP
	DevoicedVP
	CreakyVP
	BreathyVP
)

// VowelRounding is whether a vowel is rounded or not
type VowelRounding uint8

// VowelRounding values
const (
	UnspecifiedVR VowelRounding = iota
	RoundedVR
	UnroundedVR
)

// VowelNasality is whether a vowel is nasal or oral
type VowelNasality uint8

// VowelNasality values
const (
	UnspecifiedVN VowelNasality = iota
	OralVN
	NasalVN
)

// VowelLength is whether a vowel is long or short
type VowelLength uint8

// VowelLength values
const (
	UnspecifiedVL VowelLength = iota
	ShortVL
	LongVL
	ExtraShortVL
	ExtraLongVL
)

// VowelJSON is an intermediate form of the Vowel type, mirroring
// the simplified version used on the frontend, used for
// marshalling and unmarshalling
type VowelJSON struct {
	IPA       string `json:"ipa"`
	Height    string `json:"height"`
	Frontness string `json:"frontness"`
	Rounding  bool   `json:"rounding"`
	Nasal     bool   `json:"nasal"`
	Length    bool   `json:"long"`
}

// MarshalJSON to implement Marshaler interface for type Vowel
// converts Vowel structs into the JS format used on the frontend
func (v Vowel) MarshalJSON() ([]byte, error) {
	return json.Marshal(&VowelJSON{
		IPA:       v.ToIPA(),
		Height:    heightToString(v.Height),
		Frontness: frontnessToString(v.Frontness),
		Rounding:  v.Rounding == RoundedVR,
		Nasal:     v.Nasal == NasalVN,
		Length:    v.Length == LongVL,
	})
}

// UnmarshalJSON implements the Unmarshaler interface for type Vowel
func (v *Vowel) UnmarshalJSON(data []byte) error {
	var vj VowelJSON
	json.Unmarshal(data, &vj)

	v.Height = heightFromString(vj.Height)
	v.Frontness = frontnessFromString(vj.Frontness)
	v.Phonation = ModalVP

	if vj.Rounding {
		v.Rounding = RoundedVR
	} else {
		v.Rounding = UnroundedVR
	}

	if vj.Nasal {
		v.Nasal = NasalVN
	} else {
		v.Nasal = OralVN
	}

	if vj.Length {
		v.Length = LongVL
	} else {
		v.Length = ShortVL
	}

	return nil
}

func heightToString(h VowelHeight) string {
	var s string
	switch h {
	case CloseVH:
		s = "close"
	case NearCloseVH:
		s = "near-close"
	case CloseMidVH:
		s = "close-mid"
	case MidVH:
		s = "mid"
	case OpenMidVH:
		s = "open-mid"
	case NearOpenVH:
		s = "near-open"
	case OpenVH:
		s = "open"
	default:
		s = ""
	}
	return s
}

func heightFromString(s string) VowelHeight {
	var h VowelHeight
	switch s {
	case "close":
		h = CloseVH
	case "near-close":
		h = NearCloseVH
	case "close-mid":
		h = CloseMidVH
	case "mid":
		h = MidVH
	case "open-mid":
		h = OpenMidVH
	case "near-open":
		h = NearOpenVH
	case "open":
		h = OpenVH
	}
	return h
}

func frontnessToString(f VowelFrontness) string {
	var s string
	switch f {
	case FrontVF:
		s = "front"
	case CentralVF:
		s = "central"
	case BackVF:
		s = "back"
	default:
		s = ""
	}
	return s
}

func frontnessFromString(s string) VowelFrontness {
	var f VowelFrontness
	switch s {
	case "front":
		f = FrontVF
	case "central":
		f = CentralVF
	case "back":
		f = BackVF
	}
	return f
}
