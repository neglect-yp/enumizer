package a

type NotTarget string

const (
	NotTargetA = "a"
	NotTargetB = "b"
	NotTargetC = "c"
)

func NotTargetNotCovered(a NotTarget) {
	switch a {
	case NotTargetA:
	case NotTargetB:
	}
}
