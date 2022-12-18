package enumizer

import (
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/require"
)

func TestEnums_SortedEnums(t *testing.T) {
	enums := Enums{
		"Foo": {Name: "Foo", Variants: []string{"FooA", "FooB", "FooC"}},
		"Bar": {Name: "Bar", Variants: []string{"BarA", "BarB", "BarC"}},
	}

	sorted := enums.SortedEnums()
	require.EqualValues(t, []Enum{
		{Name: "Bar", Variants: []string{"BarA", "BarB", "BarC"}},
		{Name: "Foo", Variants: []string{"FooA", "FooB", "FooC"}},
	}, sorted)
}

func TestFindEnums(t *testing.T) {
	enumPackages, err := FindEnumPackages("./testdata/model1")
	require.NoError(t, err)
	require.Len(t, enumPackages, 1)
	require.EqualValues(t, EnumPackage{
		Path: "internal/enumizer/testdata/model1",
		Enums: Enums{
			"A":     {Name: "A", Variants: []string{"AA", "AB", "AC"}},
			"Iota":  {Name: "Iota", Variants: []string{"IotaZero", "IotaOne", "IotaTwo"}},
			"Split": {Name: "Split", Variants: []string{"SplitA", "SplitB", "SplitC"}},
		},
	}, enumPackages["model1"])
}

func TestFindEnums_NoMarker(t *testing.T) {
	enumPackages, err := FindEnumPackages("./testdata/nomarker")
	require.NoError(t, err)
	require.Len(t, enumPackages, 0)
}

func TestGenerateEnumHelpers(t *testing.T) {
	src, err := GenerateEnumHelpers("model", Enums{
		"Foo": {Name: "Foo", Variants: []string{"FooA", "FooB", "FooC"}},
		"Bar": {Name: "Bar", Variants: []string{"BarA", "BarB", "BarC"}},
	})
	require.NoError(t, err)

	g := goldie.New(
		t,
		goldie.WithNameSuffix(".golden.go"),
		goldie.WithFixtureDir("testdata/model"),
	)
	g.Assert(t, "model", src)
}
