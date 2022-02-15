package phonology

import (
	"encoding/json"
)

// Consonant represents a consonant phoneme
type Consonant struct {
	Place          ConsonantPlace          `json:"place"`          // Dental, Velar, etc.
	Manner         ConsonantManner         `json:"manner"`         // Plosive, Nasal, Approximant, etc.
	Coarticulation ConsonantCoarticulation `json:"coarticulation"` // Labialized, Palatalized, Velarized, etc.
	NonPulmonic    ConsonantNonPulmonic    `json:"nonpulmonic"`    // Ejective, Implosive, Velaric
	Voiced         ConsonantVoice          `json:"voiced"`         // Voiced / Voiceless
	Aspirated      ConsonantAspiration     `json:"aspirated"`      // Aspirated / Unaspirated
	Lateral        ConsonantLateral        `json:"lateral"`        // Lateral / Central
	Sibilant       ConsononantSibilance    `json:"sibilant"`       // Sibilant / Nonsibilant
	Geminate       ConsonantGeminate       `json:"geminate"`       // Geminated / Singleton
}

// ConsonantPlace is the place of articulation for consonants
type ConsonantPlace uint8

// ConsonantPlace values
const (
	UnspecifiedCP ConsonantPlace = iota
	BilabialCP
	LabioDentalCP
	DentalCP
	AlveolarCP
	PostAlveolarCP
	RetroflexCP
	PalatalCP
	VelarCP
	UvularCP
	PharyngealCP
	GlottalCP
)

// ConsonantManner is the manner of articulation for consonants
type ConsonantManner uint8

// ConsonantManner values
const (
	UnspecifiedCM ConsonantManner = iota
	NasalCM
	StopCM
	AffricateCM
	FricativeCM
	ApproximantCM
	TapCM
	TrillCM
	ClickCM
)

// ConsonantCoarticulation is the presence of coarticulations with a consonant,
// like palatization or pharyngealization
type ConsonantCoarticulation uint8

// ConsonantCoarticulation values
const (
	UnspecifiedCC ConsonantCoarticulation = iota
	NoneCC
	LabialCC
	PalatalCC
	VelarCC
	PharyngealCC
	PrenasalCC
)

// ConsonantNonPulmonic denotes non-pulmonic articulations, like ejectives and clicks
type ConsonantNonPulmonic uint8

// ConsonantNonPulmonic values
const (
	UnspecifiedCNP ConsonantNonPulmonic = iota
	PulmonicCNP
	EjectiveCNP
	ImplosiveCNP
	VelaricCNP
)

// ConsonantVoice is whether a consonant is voiced or not
type ConsonantVoice uint8

// ConsonantVoice values
const (
	UnspecifiedCV ConsonantVoice = iota
	UnvoicedCV
	VoicedCV
	PrevoicedCV // this could alternatively be included in Aspiration
)

// ConsonantAspiration is whether a consonant is aspirated or not
type ConsonantAspiration uint8

// ConsonantAspiration values
const (
	UnspecifiedCA ConsonantAspiration = iota
	UnaspiratedCA
	AspiratedCA
)

// ConsonantLateral denotes whether a consonant is articulated centrally or laterally
type ConsonantLateral uint8

// ConsonantLateral values
const (
	UnspecifiedCL ConsonantLateral = iota
	CentralCL
	LateralCL
)

// ConsononantSibilance is whether a consonant is sibilant or not
type ConsononantSibilance uint8

// ConsononantSibilance values
const (
	UnspecifiedCS ConsononantSibilance = iota
	NonsibilantCS
	SibilantCS
)

// ConsonantGeminate is whether a consonant is a singleton or geminate
type ConsonantGeminate uint8

// ConsonantGeminate values
const (
	UnspecifiedCG ConsonantGeminate = iota
	SingletonCG
	GeminateCG
)

// ConsonantJSON is an intermediate form of the Consonant type,
// mirroring the simplified version used on the frontend, used
// for marshalling and unmarshalling
type ConsonantJSON struct {
	IPA         string `json:"ipa"`
	Manner      string `json:"manner"`
	Place       string `json:"place"`
	Voiced      bool   `json:"voiced"`
	NonPulmonic string `json:"nonpulmonic"`
	Aspirated   bool   `json:"aspirated"`
	Lateral     bool   `json:"lateral"`
	Sibilant    bool   `json:"sibilant"`
}

// MarshalJSON to implement Marshaler interface for type Consonant
// converts Consonant structs into the JS format used on the frontend
func (c Consonant) MarshalJSON() ([]byte, error) {
	return json.Marshal(&ConsonantJSON{
		IPA:         c.ToIPA(),
		Manner:      mannerToString(c.Manner),
		Place:       placeToString(c.Place),
		Voiced:      c.Voiced == VoicedCV,
		NonPulmonic: nonPulmonicToString(c.NonPulmonic),
		Aspirated:   c.Aspirated == AspiratedCA,
		Lateral:     c.Lateral == LateralCL,
		Sibilant:    c.Sibilant == SibilantCS,
	})
}

// UnmarshalJSON implements the Unmarshaler interface for type Consonant
func (c *Consonant) UnmarshalJSON(data []byte) error {
	var cj ConsonantJSON
	json.Unmarshal(data, &cj)

	c.Place = placeFromString(cj.Place)
	c.Manner = mannerFromString(cj.Manner)
	c.NonPulmonic = nonPulmonicFromString(cj.NonPulmonic)

	if cj.Voiced {
		c.Voiced = VoicedCV
	} else {
		c.Voiced = UnvoicedCV
	}

	if cj.Aspirated {
		c.Aspirated = AspiratedCA
	} else {
		c.Aspirated = UnaspiratedCA
	}

	if cj.Lateral {
		c.Lateral = LateralCL
	} else {
		c.Lateral = CentralCL
	}

	if cj.Sibilant {
		c.Sibilant = SibilantCS
	} else {
		c.Sibilant = NonsibilantCS
	}

	c.Coarticulation = NoneCC
	c.Geminate = SingletonCG

	return nil
}

// FromJSON converts ConsonantJSON structs into Consonant
func (cj ConsonantJSON) FromJSON() Consonant {
	var c Consonant

	switch cj.Place {
	case "bilabial":
		c.Place = BilabialCP
	case "labio-dental":
		c.Place = LabioDentalCP
	case "dental":
		c.Place = DentalCP
	case "alveolar":
		c.Place = AlveolarCP
	case "post-alveolar":
		c.Place = PostAlveolarCP
	case "retroflex":
		c.Place = RetroflexCP
	case "palatal":
		c.Place = PalatalCP
	case "velar":
		c.Place = VelarCP
	case "uvular":
		c.Place = UvularCP
	case "pharyngeal":
		c.Place = PharyngealCP
	case "glottal":
		c.Place = GlottalCP
	}

	switch cj.Manner {
	case "nasal":
		c.Manner = NasalCM
	case "stop":
		c.Manner = StopCM
	case "affricate":
		c.Manner = AffricateCM
	case "fricative":
		c.Manner = FricativeCM
	case "approximant":
		c.Manner = ApproximantCM
	case "tap":
		c.Manner = TapCM
	case "trill":
		c.Manner = TrillCM
	case "click":
		c.Manner = ClickCM
	}

	switch cj.NonPulmonic {
	case "ejective":
		c.NonPulmonic = EjectiveCNP
	case "implosive":
		c.NonPulmonic = ImplosiveCNP
	case "velaric":
		c.NonPulmonic = VelaricCNP
	default:
		c.NonPulmonic = PulmonicCNP
	}

	if cj.Voiced {
		c.Voiced = VoicedCV
	} else {
		c.Voiced = UnvoicedCV
	}

	if cj.Aspirated {
		c.Aspirated = AspiratedCA
	} else {
		c.Aspirated = UnaspiratedCA
	}

	if cj.Lateral {
		c.Lateral = LateralCL
	} else {
		c.Lateral = CentralCL
	}

	if cj.Sibilant {
		c.Sibilant = SibilantCS
	} else {
		c.Sibilant = NonsibilantCS
	}

	c.Coarticulation = NoneCC
	c.Geminate = SingletonCG

	return c
}

func mannerToString(m ConsonantManner) string {
	var s string
	switch m {
	case NasalCM:
		s = "nasal"
	case StopCM:
		s = "stop"
	case AffricateCM:
		s = "affricate"
	case FricativeCM:
		s = "fricative"
	case ApproximantCM:
		s = "approximant"
	case TapCM:
		s = "tap"
	case TrillCM:
		s = "trill"
	default:
		s = ""
	}
	return s
}

func mannerFromString(s string) ConsonantManner {
	var m ConsonantManner
	switch s {
	case "nasal":
		m = NasalCM
	case "stop":
		m = StopCM
	case "affricate":
		m = AffricateCM
	case "fricative":
		m = FricativeCM
	case "approximant":
		m = ApproximantCM
	case "tap":
		m = TapCM
	case "trill":
		m = TrillCM
	}
	return m
}

func placeToString(p ConsonantPlace) string {
	var s string
	switch p {
	case BilabialCP:
		s = "bilabial"
	case LabioDentalCP:
		s = "labio-dental"
	case DentalCP:
		s = "dental"
	case AlveolarCP:
		s = "alveolar"
	case PostAlveolarCP:
		s = "post-alveolar"
	case RetroflexCP:
		s = "retroflex"
	case PalatalCP:
		s = "palatal"
	case VelarCP:
		s = "velar"
	case UvularCP:
		s = "uvular"
	case PharyngealCP:
		s = "pharyngeal"
	case GlottalCP:
		s = "glottal"
	default:
		s = ""
	}
	return s
}

func placeFromString(s string) ConsonantPlace {
	var p ConsonantPlace
	switch s {
	case "bilabial":
		p = BilabialCP
	case "labio-dental":
		p = LabioDentalCP
	case "dental":
		p = DentalCP
	case "alveolar":
		p = AlveolarCP
	case "post-alveolar":
		p = PostAlveolarCP
	case "retroflex":
		p = RetroflexCP
	case "palatal":
		p = PalatalCP
	case "velar":
		p = VelarCP
	case "uvular":
		p = UvularCP
	case "pharyngeal":
		p = PharyngealCP
	case "glottal":
		p = GlottalCP
	}
	return p
}

func nonPulmonicToString(np ConsonantNonPulmonic) string {
	var s string
	switch np {
	case PulmonicCNP:
		s = "pulmonic"
	case EjectiveCNP:
		s = "ejective"
	case ImplosiveCNP:
		s = "implosive"
	case VelaricCNP:
		s = "velaric"
	default:
		s = ""
	}
	return s
}

func nonPulmonicFromString(s string) ConsonantNonPulmonic {
	var np ConsonantNonPulmonic
	switch s {
	case "pulmonic":
		np = PulmonicCNP
	case "ejective":
		np = EjectiveCNP
	case "implosive":
		np = ImplosiveCNP
	case "velaric":
		np = VelaricCNP
	}
	return np
}
