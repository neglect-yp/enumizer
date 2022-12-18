package foo

type EnumA string // want EnumA:"EnumAA, EnumAB, EnumAC"

// enumizer:target
const (
	EnumAA EnumA = "a"
	EnumAB EnumA = "b"
	EnumAC EnumA = "c"
)
