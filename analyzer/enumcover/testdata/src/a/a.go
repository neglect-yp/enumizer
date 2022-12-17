package a

type EnumA string // want EnumA:"EnumAA, EnumAB, EnumAC"

// enumizer:generate
const (
	EnumAA EnumA = "a"
	EnumAB EnumA = "b"
	EnumAC EnumA = "c"
)

func Covered(a EnumA) {
	switch a {
	case EnumAA:
	case EnumAB:
	case EnumAC:
	}
}

func CoveredWithDefault(a EnumA) {
	switch a {
	case EnumAA:
	case EnumAB:
	case EnumAC:
	default:
	}
}

func CoveredWithMultipleValuesInCase(a EnumA) {
	switch a {
	case EnumAA:
	case EnumAB, EnumAC:
	default:
	}
}

func NotCovered(a EnumA) {
	switch a { // want "this switch statement doesn't cover enum variants. missing cases: EnumAC"
	case EnumAA:
	case EnumAB:
	}
}

func NotCovered2(a EnumA) {
	switch a { // want "this switch statement doesn't cover enum variants. missing cases: EnumAB, EnumAC"
	case EnumAA:
	}
}

func NotCoveredWithDefault(a EnumA) {
	switch a { // want "this switch statement doesn't cover enum variants. missing cases: EnumAC"
	case EnumAA:
	case EnumAB:
	default:
	}
}
