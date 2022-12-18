package model1

type A string

// enumizer:generate
const (
	AA A = "a"
	AB A = "b"
	AC A = "c"
)

type Split string

// enumizer:generate
const (
	SplitA Split = "A"
	SplitB Split = "B"
)

// enumizer:generate
const (
	SplitC Split = "C"
)

type ContaminationA string // ignored
type ContaminationB string

// enumizer:generate
const (
	ContaminationAA ContaminationA = "a"
	ContaminationBB ContaminationB = "b"
)

type Iota int

// enumizer:generate
const (
	IotaZero Iota = iota
	IotaOne
	IotaTwo
)

type WithImplicitType string // ignored

// enumizer:generate
const (
	WithImplicitTypeA WithImplicitType = "a"
	Foo                                = "foo"
)
