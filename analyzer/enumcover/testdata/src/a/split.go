package a

type SplitA int // want SplitA:"SplitAA, SplitAB, SplitAC"

// enumizer:generate
const (
	SplitAA SplitA = 0
	SplitAB SplitA = 1
)

// enumizer:generate
const (
	SplitAC SplitA = 2
)

func SplitACovered(a SplitA) {
	switch a {
	case SplitAA:
	case SplitAB:
	case SplitAC:
	}
}

func SplitANotCovered(a SplitA) {
	switch a { // want "this switch statement doesn't cover enum variants. missing cases: SplitAC"
	case SplitAA:
	case SplitAB:
	}
}
