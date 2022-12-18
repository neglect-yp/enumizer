// Code generated by enumizer; DO NOT EDIT.
package myenum

import "fmt"

var myEnumSet = map[MyEnum]struct{}{
	A: {},
	B: {},
	C: {},
}

func MyEnumList() []MyEnum {
	ret := make([]MyEnum, 0, len(myEnumSet))
	for v := range myEnumSet {
		ret = append(ret, v)
	}
	return ret
}

func (m MyEnum) IsValid() bool {
	_, ok := myEnumSet[m]
	return ok
}

func (m MyEnum) Validate() error {
	if !m.IsValid() {
		return fmt.Errorf("MyEnum(%v) is invalid", m)
	}
	return nil
}
