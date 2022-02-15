package phonology

import "errors"

// ToIPA returns the IPA representation of a Vowel struct
func (v Vowel) ToIPA() string {
	representation := "V"
	switch v.Frontness {
	case FrontVF:
		switch v.Height {
		case CloseVH:
			if v.Rounding == RoundedVR {
				representation = "y"
			} else {
				representation = "i"
			}
		case NearCloseVH:
			if v.Rounding == RoundedVR {
				representation = "ʏ"
			} else {
				representation = "ɪ"
			}
		case CloseMidVH, MidVH:
			if v.Rounding == RoundedVR {
				representation = "ø"
			} else {
				representation = "e"
			}
		case OpenMidVH:
			if v.Rounding == RoundedVR {
				representation = "œ"
			} else {
				representation = "ɛ"
			}
		case NearOpenVH:
			if v.Rounding == RoundedVR {
				representation = "œ"
			} else {
				representation = "æ"
			}
		case OpenVH:
			if v.Rounding == RoundedVR {
				representation = "œ"
			} else {
				representation = "a"
			}
		}
	case CentralVF:
		switch v.Height {
		case CloseVH, NearCloseVH:
			if v.Rounding == RoundedVR {
				representation = "ʉ"
			} else {
				representation = "ɨ"
			}
		case CloseMidVH, MidVH, OpenMidVH:
			if v.Rounding == RoundedVR {
				representation = "ɵ"
			} else {
				representation = "ə"
			}
		case NearOpenVH, OpenVH:
			if v.Rounding == RoundedVR {
				representation = "ɞ"
			} else {
				representation = "ɐ"
			}
		}
	case BackVF:
		switch v.Height {
		case CloseVH:
			if v.Rounding == RoundedVR {
				representation = "u"
			} else {
				representation = "ɯ"
			}
		case NearCloseVH:
			if v.Rounding == RoundedVR {
				representation = "ʊ"
			} else {
				representation = "ɯ"
			}
		case CloseMidVH, MidVH:
			if v.Rounding == RoundedVR {
				representation = "o"
			} else {
				representation = "ɤ"
			}
		case OpenMidVH:
			if v.Rounding == RoundedVR {
				representation = "ɔ"
			} else {
				representation = "ʌ"
			}
		case NearOpenVH, OpenVH:
			if v.Rounding == RoundedVR {
				representation = "ɒ"
			} else {
				representation = "ɑ"
			}
		}
	}
	switch v.Phonation {
	case DevoicedVP:
		representation += string(rune(0x0325)) // small circle below
	case CreakyVP:
		representation += string(rune(0x0330)) // tilde below
	case BreathyVP:
		representation += string(rune(0x0324)) // diaresis below
	}
	if v.Nasal == NasalVN {
		representation += string(rune(0x0303)) // tilde above
	}
	if v.Length == LongVL {
		representation += "ː"
	}
	return representation
}

// ToIPA returns the IPA representation of a Consonant struct
func (c Consonant) ToIPA() string {
	representation := "C"
	switch c.Manner {
	case NasalCM:
		switch c.Place {
		case BilabialCP, LabioDentalCP:
			representation = "m"
		case DentalCP:
			representation = "n" + string(rune(0x032A))
		case AlveolarCP, PostAlveolarCP:
			representation = "n"
		case RetroflexCP:
			representation = "ɳ"
		case PalatalCP:
			representation = "ɲ"
		case VelarCP:
			representation = "ŋ"
		case UvularCP, PharyngealCP, GlottalCP:
			representation = "ɴ"
		}
		if c.Voiced == UnvoicedCV {
			representation += string(rune(0x0325)) // devoiced nasal
		}
	case StopCM:
		switch c.Place {
		case BilabialCP, LabioDentalCP:
			if c.Voiced == VoicedCV {
				representation = "b"
			} else {
				representation = "p"
			}
		case DentalCP:
			if c.Voiced == VoicedCV {
				representation = "d"
			} else {
				representation = "t"
			}
			representation += string(rune(0x032A)) // dental
		case AlveolarCP, PostAlveolarCP:
			if c.Voiced == VoicedCV {
				representation = "d"
			} else {
				representation = "t"
			}
		case RetroflexCP:
			if c.Voiced == VoicedCV {
				representation = "ɖ"
			} else {
				representation = "ʈ"
			}
		case PalatalCP:
			if c.Voiced == VoicedCV {
				representation = "ɟ"
			} else {
				representation = "c"
			}
		case VelarCP:
			if c.Voiced == VoicedCV {
				representation = "g"
			} else {
				representation = "k"
			}
		case UvularCP:
			if c.Voiced == VoicedCV {
				representation = "ɢ"
			} else {
				representation = "q"
			}
		case PharyngealCP:
			representation = "ʡ"
		case GlottalCP:
			representation = "ʔ"
		}
	case AffricateCM: // U+0361 is the tie bar for affricates
		switch c.Place {
		case BilabialCP:
			if c.Voiced == VoicedCV {
				representation = "b" + string(rune(0x0361)) + "β"
			} else {
				representation = "p" + string(rune(0x0361)) + "ɸ"
			}
		case LabioDentalCP:
			if c.Voiced == VoicedCV {
				representation = "b" + string(rune(0x0361)) + "v"
			} else {
				representation = "p" + string(rune(0x0361)) + "f"
			}
		case DentalCP, AlveolarCP:
			if c.Sibilant == SibilantCS {
				if c.Voiced == VoicedCV {
					representation = "d" + string(rune(0x0361)) + "z"
				} else {
					representation = "t" + string(rune(0x0361)) + "s"
				}
			} else if c.Lateral == LateralCL {
				if c.Voiced == VoicedCV {
					representation = "d" + string(rune(0x0361)) + "ɮ"
				} else {
					representation = "t" + string(rune(0x0361)) + "ɬ"
				}
			} else {
				if c.Voiced == VoicedCV {
					representation = "d" + string(rune(0x0361)) + "ð"
				} else {
					representation = "t" + string(rune(0x0361)) + "θ"
				}
			}
		case PostAlveolarCP:
			if c.Sibilant == SibilantCS {
				if c.Voiced == VoicedCV {
					representation = "d" + string(rune(0x0361)) + "ʒ"
				} else {
					representation = "t" + string(rune(0x0361)) + "ʃ"
				}
			} else if c.Lateral == LateralCL {
				if c.Voiced == VoicedCV {
					representation = "d" + string(rune(0x0361)) + "ɮ"
				} else {
					representation = "t" + string(rune(0x0361)) + "ɬ"
				}
			} else {
				if c.Voiced == VoicedCV {
					representation = "d" + string(rune(0x0361)) + "ð"
				} else {
					representation = "t" + string(rune(0x0361)) + "θ"
				}
			}
		case RetroflexCP:
			if c.Sibilant == SibilantCS {
				if c.Voiced == VoicedCV {
					representation = "ɖ" + string(rune(0x0361)) + "ʐ"
				} else {
					representation = "ʈ" + string(rune(0x0361)) + "ʂ"
				}
			} else if c.Lateral == LateralCL {
				if c.Voiced == VoicedCV {
					representation = "ɖ" + string(rune(0x0361)) + "ɮ"
				} else {
					representation = "ʈ" + string(rune(0x0361)) + "ɬ"
				}
			} else {
				if c.Voiced == VoicedCV {
					representation = "d" + string(rune(0x0361)) + "ð"
				} else {
					representation = "t" + string(rune(0x0361)) + "θ"
				}
			}
		case PalatalCP:
			if c.Sibilant == SibilantCS {
				if c.Voiced == VoicedCV {
					representation = "d" + string(rune(0x0361)) + "ʑ"
				} else {
					representation = "t" + string(rune(0x0361)) + "ɕ"
				}
			} else if c.Lateral == LateralCL {
				if c.Voiced == VoicedCV {
					representation = "ɟ" + string(rune(0x0361)) + "ʎ̝"
				} else {
					representation = "c" + string(rune(0x0361)) + "ʎ̝̊"
				}
			} else {
				if c.Voiced == VoicedCV {
					representation = "ɟ" + string(rune(0x0361)) + "ʝ"
				} else {
					representation = "c" + string(rune(0x0361)) + "ç"
				}
			}
		case VelarCP:
			if c.Voiced == VoicedCV {
				representation = "ɡ" + string(rune(0x0361)) + "ɣ"
			} else {
				representation = "k" + string(rune(0x0361)) + "x"
			}
		case UvularCP:
			if c.Voiced == VoicedCV {
				representation = "ɢ" + string(rune(0x0361)) + "ʁ"
			} else {
				representation = "q" + string(rune(0x0361)) + "χ"
			}
		case PharyngealCP:
			representation = "ʡ" + string(rune(0x0361)) + "ʢ"
		case GlottalCP:
			representation = "ʔ" + string(rune(0x0361)) + "h"
		}
	case FricativeCM:
		switch c.Place {
		case BilabialCP:
			if c.Voiced == VoicedCV {
				representation = "β"
			} else {
				representation = "ɸ"
			}
		case LabioDentalCP:
			if c.Voiced == VoicedCV {
				representation = "v"
			} else {
				representation = "f"
			}
		case DentalCP, AlveolarCP:
			if c.Sibilant == SibilantCS {
				if c.Voiced == VoicedCV {
					representation = "z"
				} else {
					representation = "s"
				}
			} else if c.Lateral == LateralCL {
				if c.Voiced == VoicedCV {
					representation = "ɮ"
				} else {
					representation = "ɬ"
				}
			} else {
				if c.Voiced == VoicedCV {
					representation = "ð"
				} else {
					representation = "θ"
				}
			}
		case PostAlveolarCP:
			if c.Sibilant == SibilantCS {
				if c.Voiced == VoicedCV {
					representation = "ʒ"
				} else {
					representation = "ʃ"
				}
			} else if c.Lateral == LateralCL {
				if c.Voiced == VoicedCV {
					representation = "ɮ"
				} else {
					representation = "ɬ"
				}
			} else {
				if c.Voiced == VoicedCV {
					representation = "ð"
				} else {
					representation = "θ"
				}
			}
		case RetroflexCP:
			if c.Sibilant == SibilantCS {
				if c.Voiced == VoicedCV {
					representation = "ʐ"
				} else {
					representation = "ʂ"
				}
			} else if c.Lateral == LateralCL {
				if c.Voiced == VoicedCV {
					representation = "ɮ"
				} else {
					representation = "ɬ"
				}
			} else {
				if c.Voiced == VoicedCV {
					representation = "ð"
				} else {
					representation = "θ"
				}
			}
		case PalatalCP:
			if c.Sibilant == SibilantCS {
				if c.Voiced == VoicedCV {
					representation = "ʑ"
				} else {
					representation = "ɕ"
				}
			} else {
				if c.Voiced == VoicedCV {
					representation = "ʝ"
				} else {
					representation = "ç"
				}
			}
		case VelarCP:
			if c.Voiced == VoicedCV {
				representation = "ɣ"
			} else {
				representation = "x"
			}
		case UvularCP:
			if c.Voiced == VoicedCV {
				representation = "ʁ"
			} else {
				representation = "χ"
			}
		case PharyngealCP:
			if c.Voiced == VoicedCV {
				representation = "ħ"
			} else {
				representation = "ʕ"
			}
		case GlottalCP:
			if c.Voiced == VoicedCV {
				representation = "ɦ"
			} else {
				representation = "h"
			}
		}
	case ApproximantCM:
		switch c.Place {
		case BilabialCP: // technically labio-velar
			if c.Voiced == VoicedCV {
				representation = "w"
			} else {
				representation = "ʍ"
			}
		case LabioDentalCP:
			representation = "ʋ"
		case DentalCP, AlveolarCP, PostAlveolarCP:
			if c.Lateral == LateralCL {
				representation = "l"
			} else {
				representation = "ɹ"
			}
		case RetroflexCP:
			if c.Lateral == LateralCL {
				representation = "ɭ"
			} else {
				representation = "ɻ"
			}
		case PalatalCP:
			if c.Lateral == LateralCL {
				representation = "ʎ"
			} else {
				representation = "j"
			}
		case VelarCP, UvularCP, PharyngealCP, GlottalCP:
			if c.Lateral == LateralCL {
				representation = "ʟ"
			} else {
				representation = "ɰ"
			}
		}
		if c.Voiced == UnvoicedCV {
			representation += string(rune(0x0325)) // devoiced approximant
		}
	case TapCM:
		switch c.Place {
		case BilabialCP, LabioDentalCP:
			representation = "ⱱ"
		case DentalCP, AlveolarCP, PostAlveolarCP:
			if c.Lateral == LateralCL {
				representation = "ɺ"
			} else {
				representation = "ɾ"
			}
		case RetroflexCP:
			representation = "ɽ"
		case PalatalCP, VelarCP, UvularCP, PharyngealCP, GlottalCP:
			// these range from extremely rare to impossible
			representation = "ɽ"
		}
		if c.Voiced == UnvoicedCV {
			representation += string(rune(0x0325)) // devoiced tap
		}
	case TrillCM:
		switch c.Place {
		case BilabialCP, LabioDentalCP:
			representation = "ʙ"
		case DentalCP, AlveolarCP, PostAlveolarCP, PalatalCP:
			representation = "r"
		case RetroflexCP:
			representation = "ɽ" + string(rune(0x0361)) + "r"
		case VelarCP, UvularCP:
			// velar trill is impossible
			representation = "ʀ"
		case PharyngealCP, GlottalCP:
			// glottal trill is impossible
			representation = "ʢ"
		}
		if c.Voiced == UnvoicedCV {
			representation += string(rune(0x0325)) // devoiced trill
		}
	}

	switch c.Coarticulation {
	case LabialCC:
		representation += "ʷ"
	case PalatalCC:
		representation += "ʲ"
	case VelarCC:
		representation += "ˠ"
	case PharyngealCC:
		representation += "ˤ"
	}

	if c.Aspirated == AspiratedCA {
		representation += "ʰ"
	}

	if c.Geminate == GeminateCG {
		representation += "ː"
	}

	if c.NonPulmonic == EjectiveCNP {
		representation += "ʼ"
	}

	// implosives and clicks
	if c.NonPulmonic == ImplosiveCNP {
		switch c.Place {
		case BilabialCP, LabioDentalCP:
			representation = "ɓ"
		case DentalCP, AlveolarCP, PostAlveolarCP:
			representation = "ɗ"
		case RetroflexCP:
			representation = "ᶑ"
		case PalatalCP:
			representation = "ʄ"
		case VelarCP:
			representation = "ɠ"
		case UvularCP, PharyngealCP, GlottalCP:
			representation = "ʛ"
		}
		if c.Voiced == UnvoicedCV {
			representation += string(rune(0x0325))
		}
	} else if c.NonPulmonic == VelaricCNP {
		switch c.Place {
		case BilabialCP, LabioDentalCP:
			representation = "ʘ"
		case DentalCP:
			representation = "ǀ"
		case AlveolarCP, PostAlveolarCP:
			representation = "!"
		case RetroflexCP:
			representation = "‼"
		case PalatalCP:
			representation = "ǂ"
		}
		if c.Lateral == LateralCL {
			representation = "ǁ"
		}
	}

	return representation
}

func isIPAVowel(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch string([]rune(s)[0]) {
	case "y", "i", "ʏ", "ɪ", "ø", "e", "œ", "ɛ", "æ", "a", "ʉ", "ɨ", "ɵ", "ə", "ɞ", "ɐ", "u", "ɯ", "ʊ", "o", "ɤ", "ɔ", "ʌ", "ɒ", "ɑ":
		return true
	default:
		return false
	}
}

func isIPAConsonant(s string) bool {
	if len(s) == 0 {
		return false
	}
	switch string([]rune(s)[0]) {
	case "m", "n", "ɳ", "ɲ", "ŋ", "ɴ", "b", "p", "t", "d", "ɖ", "ʈ", "ɟ", "c", "g", "k", "ɢ", "q", "ʡ", "ʔ", "β", "ɸ", "v", "f", "z", "s", "ɮ", "ɬ", "ð", "θ", "ʒ", "ʃ", "ʐ", "ʂ", "ʑ", "ɕ", "ʝ", "ç", "ɣ", "x", "ʁ", "χ", "ħ", "ʕ", "ɦ", "h", "w", "ʍ", "ʋ", "l", "ɹ", "ɭ", "ɻ", "ʎ", "j", "ʟ", "ɰ", "ⱱ", "ɺ", "ɾ", "ɽ", "ʙ", "r", "ʀ", "ʢ", "ɓ", "ɗ", "ᶑ", "ʄ", "ɠ", "ʛ", "ʘ", "ǀ", "!", "‼", "ǂ", "ǁ":
		return true
	default:
		return false
	}
}

// NewVowelFromIPA creates a Vowel from an IPA string
func NewVowelFromIPA(s string) (Vowel, error) {
	runes := []rune(s)
	vowel := Vowel{}

	// Frontness
	switch string(runes[0]) {
	case "y", "i", "ʏ", "ɪ", "ø", "e", "œ", "ɛ", "æ", "a":
		vowel.Frontness = FrontVF
	case "ʉ", "ɨ", "ɵ", "ə", "ɞ", "ɐ":
		vowel.Frontness = CentralVF
	case "u", "ɯ", "ʊ", "o", "ɤ", "ɔ", "ʌ", "ɒ", "ɑ":
		vowel.Frontness = BackVF
	default:
		return vowel, errors.New("Failed to parse vowel string: \"" + s + "\"")
	}

	// Height
	switch string(runes[0]) {
	case "y", "i", "ʉ", "ɨ", "u", "ɯ":
		vowel.Height = CloseVH
	case "ʏ", "ɪ", "ʊ":
		vowel.Height = NearCloseVH
	case "ø", "e", "ɵ", "o", "ɤ":
		vowel.Height = CloseMidVH
	case "ə":
		vowel.Height = MidVH
	case "œ", "ɛ", "ɞ", "ɔ", "ʌ":
		vowel.Height = OpenMidVH
	case "æ", "ɐ":
		vowel.Height = NearOpenVH
	case "a", "ɒ", "ɑ":
		vowel.Height = OpenVH
	default:
		return vowel, errors.New("Failed to parse vowel string: \"" + s + "\"")
	}

	// Rounded
	switch string(runes[0]) {
	case "y", "ʏ", "ø", "œ", "ʉ", "ɵ", "ɞ", "u", "ʊ", "o", "ɔ", "ɒ":
		vowel.Rounding = RoundedVR
	default:
		vowel.Rounding = UnroundedVR
	}

	vowel.Phonation = ModalVP
	vowel.Nasal = OralVN
	vowel.Length = ShortVL

	// Nasal, Long, Phonation
	for i := 1; i < len(runes); i++ {
		switch string(runes[i]) {
		case string(rune(0x0325)): // devoiced
			vowel.Phonation = DevoicedVP
		case string(rune(0x0330)): // creaky
			vowel.Phonation = CreakyVP
		case string(rune(0x0324)): // breathy
			vowel.Phonation = BreathyVP
		case string(rune(0x0303)): // nasal
			vowel.Nasal = NasalVN
		case "ː": // long
			vowel.Length = LongVL
		}
	}

	return vowel, nil
}

// NewConsonantFromIPA creates a Consonant from an IPA string
func NewConsonantFromIPA(s string) (Consonant, error) {
	runes := []rune(s)
	cons := Consonant{}

	// Manner
	switch string(runes[0]) {
	case "m", "n", "ɳ", "ɲ", "ŋ", "ɴ":
		cons.Manner = NasalCM
	case "b", "p", "t", "d", "ɖ", "ʈ", "ɟ", "c", "g", "k", "ɢ", "q", "ʡ", "ʔ", "ɓ", "ɗ", "ᶑ", "ʄ", "ɠ", "ʛ", "ʘ", "ǀ", "!", "‼", "ǂ", "ǁ":
		cons.Manner = StopCM
	case "β", "ɸ", "v", "f", "z", "s", "ɮ", "ɬ", "ð", "θ", "ʒ", "ʃ", "ʐ", "ʂ", "ʑ", "ɕ", "ʝ", "ç", "ɣ", "x", "ʁ", "χ", "ħ", "ʕ", "ɦ", "h":
		cons.Manner = FricativeCM
	case "w", "ʋ", "l", "ɹ", "ɭ", "ɻ", "ʎ", "j", "ʟ", "ɰ", "ʍ":
		cons.Manner = ApproximantCM
	case "ⱱ", "ɺ", "ɾ", "ɽ":
		cons.Manner = TapCM
	case "ʙ", "r", "ʀ", "ʢ":
		cons.Manner = TrillCM
	default:
		return cons, errors.New("Failed to parse consonant string: \"" + s + "\"")
	}

	// Place
	switch string(runes[0]) {
	case "m", "b", "p", "β", "ɸ", "w", "ʍ", "ʙ", "ɓ", "ʘ":
		cons.Place = BilabialCP
	case "v", "f", "ʋ", "ⱱ":
		cons.Place = LabioDentalCP
	case "ð", "θ", "ǀ":
		cons.Place = DentalCP
	case "n", "t", "d", "z", "s", "ɮ", "ɬ", "l", "ɹ", "ɺ", "ɾ", "r", "ɗ", "!", "ǁ":
		cons.Place = AlveolarCP
	case "ʒ", "ʃ":
		cons.Place = PostAlveolarCP
	case "ɳ", "ɖ", "ʈ", "ʐ", "ʂ", "ɭ", "ɻ", "ɽ", "ᶑ", "‼":
		cons.Place = RetroflexCP
	case "ɲ", "ɟ", "c", "ʑ", "ɕ", "ʝ", "ç", "ʎ", "j", "ʄ", "ǂ":
		cons.Place = PalatalCP
	case "ŋ", "g", "k", "ɣ", "x", "ʟ", "ɰ", "ɠ":
		cons.Place = VelarCP
	case "ɴ", "ɢ", "q", "ʁ", "χ", "ʀ", "ʛ":
		cons.Place = UvularCP
	case "ʡ", "ħ", "ʕ", "ʢ":
		cons.Place = PharyngealCP
	case "ʔ", "ɦ", "h":
		cons.Place = GlottalCP
	default:
		return cons, errors.New("Failed to parse consonant string: \"" + s + "\"")
	}

	// Sibilant
	switch string(runes[0]) {
	case "z", "s", "ʒ", "ʃ", "ʐ", "ʂ", "ʑ", "ɕ":
		cons.Sibilant = SibilantCS
	default:
		cons.Sibilant = NonsibilantCS
	}

	// Voiced
	switch string(runes[0]) {
	case "m", "n", "ɳ", "ɲ", "ŋ", "ɴ", "b", "d", "ɖ", "ɟ", "g", "ɢ", "β", "v", "z", "ɮ", "ð", "ʒ", "ʐ", "ʑ", "ʝ", "ɣ", "ʁ", "ʕ", "ɦ", "w", "ʋ", "l", "ɹ", "ɭ", "ɻ", "ʎ", "j", "ʟ", "ɰ", "ⱱ", "ɺ", "ɾ", "ɽ", "ʙ", "r", "ʀ", "ʢ", "ɓ", "ɗ", "ᶑ", "ʄ", "ɠ", "ʛ":
		cons.Voiced = VoicedCV
	default:
		cons.Voiced = UnvoicedCV
	}

	// Lateral
	switch string(runes[0]) {
	case "ɮ", "ɬ", "l", "ɺ", "ǁ", "ɭ", "ʎ", "ʟ":
		cons.Lateral = LateralCL
	default:
		cons.Lateral = CentralCL
	}

	// NonPulmonic
	switch string(runes[0]) {
	case "ɓ", "ɗ", "ᶑ", "ʄ", "ɠ", "ʛ":
		cons.NonPulmonic = ImplosiveCNP
	case "ʘ", "ǀ", "!", "‼", "ǂ", "ǁ":
		cons.NonPulmonic = VelaricCNP
	default:
		cons.NonPulmonic = PulmonicCNP
	}

	cons.Aspirated = UnaspiratedCA
	cons.Geminate = SingletonCG
	cons.Coarticulation = NoneCC

	// Aspiration, Affricates, Ejectives, Coarticulation, Gemination
	for i := 1; i < len(runes); i++ {
		switch string(runes[i]) {
		case "ː": // Geminated
			cons.Geminate = GeminateCG
		case "ʰ": // Aspiration
			cons.Aspirated = AspiratedCA
		case "ˤ": // Coarticulation
			cons.Coarticulation = PharyngealCC
		case "ˠ":
			cons.Coarticulation = VelarCC
		case "ʲ":
			cons.Coarticulation = PalatalCC
		case "ʷ":
			cons.Coarticulation = LabialCC
		case "ʼ":
			cons.NonPulmonic = EjectiveCNP
		case string(rune(0x032A)): // Dental
			cons.Place = DentalCP
		case string(rune(0x0325)): // Devoiced
			cons.Voiced = UnvoicedCV
		case "ɮ", "ɬ", "ɭ", "ʎ", "β", "ɸ", "v", "f", "z", "s", "ð", "θ", "ʒ", "ʃ", "ʐ", "ʂ", "ʑ", "ɕ", "ʝ", "ç", "ɣ", "x", "ʁ", "χ", "ħ", "ʕ", "ɦ", "h": // Affricates
			cons.Manner = AffricateCM
			// fricative part adds sibilance or lateral release
			switch string(runes[i]) {
			case "ɮ", "ɬ", "ɭ", "ʎ":
				cons.Lateral = LateralCL
			case "z", "s", "ʒ", "ʃ", "ʐ", "ʂ", "ʑ", "ɕ":
				cons.Sibilant = SibilantCS
			}
			// fricative part changes place of articulation
			switch string(runes[i]) {
			case "ɮ", "ɬ", "z", "s":
				cons.Place = AlveolarCP
			case "ʒ", "ʃ":
				cons.Place = PostAlveolarCP
			case "ʐ", "ʂ", "ɭ":
				cons.Place = RetroflexCP
			case "ʑ", "ɕ", "ʎ":
				cons.Place = PalatalCP
			}
		}
	}

	return cons, nil
}
