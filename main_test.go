package main

import (
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/require"
)

func TestFindEnums(t *testing.T) {
	enums, err := FindEnums("./testdata/model1")
	require.NoError(t, err)
	require.EqualValues(t, []Enum{
		{Name: "A", Variants: []string{"AA", "AB", "AC"}},
		{Name: "Iota", Variants: []string{"IotaZero", "IotaOne", "IotaTwo"}},
		{Name: "Split", Variants: []string{"SplitA", "SplitB", "SplitC"}},
	}, enums)
}

func TestGenerateEnumHelpers(t *testing.T) {
	src, err := GenerateEnumHelpers([]Enum{
		{Name: "Foo", Variants: []string{"FooA", "FooB", "FooC"}},
		{Name: "Bar", Variants: []string{"BarA", "BarB", "BarC"}},
	})
	require.NoError(t, err)

	g := goldie.New(
		t,
		goldie.WithNameSuffix(".golden.go"),
		goldie.WithFixtureDir("testdata/model"),
	)
	g.Assert(t, "model", src)
}
