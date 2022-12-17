package across_package

import (
	"github.com/neglect-yp/foo"
)

func AcrossPackageCovered(a foo.EnumA) {
	switch a {
	case foo.EnumAA:
	case foo.EnumAB:
	case foo.EnumAC:
	}
}

func AcrossPackageNotCovered(a foo.EnumA) {
	switch a { // want "this switch statement doesn't cover enum variants. missing cases: EnumAC"
	case foo.EnumAA:
	case foo.EnumAB:
	}
}
