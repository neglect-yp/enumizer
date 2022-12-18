package model1

type A string

// enumizer:target
const (
	AA A = "a"
	AB A = "b"
	AC A = "c"
)

type Split string

// enumizer:target
const (
	SplitA Split = "A"
	SplitB Split = "B"
)

// enumizer:target
const (
	SplitC Split = "C"
)

type ContaminationA string // ignored
type ContaminationB string

// enumizer:target
const (
	ContaminationAA ContaminationA = "a"
	ContaminationBB ContaminationB = "b"
)

type Iota int

// enumizer:target
const (
	IotaZero Iota = iota
	IotaOne
	IotaTwo
)

type WithImplicitType string // ignored

// enumizer:target
const (
	WithImplicitTypeA WithImplicitType = "a"
	Foo                                = "foo"
)
