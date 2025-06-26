package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFooList(t *testing.T) {
	require.ElementsMatch(t, []Foo{FooA, FooB, FooC}, FooList())
}

func TestFoo_IsValid(t *testing.T) {
	require.True(t, FooA.IsValid())
	require.False(t, Foo("invalid").IsValid())
}

func TestFoo_Validate(t *testing.T) {
	require.NoError(t, FooA.Validate())
	require.Error(t, Foo("invalid").Validate())
}

func TestBarList(t *testing.T) {
	require.ElementsMatch(t, []Bar{BarA, BarB, BarC}, BarList())
}

func TestBar_IsValid(t *testing.T) {
	require.True(t, BarA.IsValid())
	require.False(t, Bar(0).IsValid())
}

func TestBar_Validate(t *testing.T) {
	require.NoError(t, BarA.Validate())
	require.Error(t, Bar(0).Validate())
}
