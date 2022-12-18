package a

type IotaA int // want IotaA:"IotaAZero, IotaAOne, IotaATwo"

// enumizer:target
const (
	IotaAZero IotaA = iota
	IotaAOne
	IotaATwo
)

func IotaACovered(a IotaA) {
	switch a {
	case IotaAZero:
	case IotaAOne:
	case IotaATwo:
	}
}

func IotaNotCovered(a IotaA) {
	switch a { // want "this switch statement doesn't cover enum variants. missing cases: IotaATwo"
	case IotaAZero:
	case IotaAOne:
	}
}
