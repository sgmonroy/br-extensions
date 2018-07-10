package shared

const Version string = "v1"

type Plugin_v1 interface {
	Init(BR_v1)
}

type BR_v1 interface {
	Name() string
	ID() (string, bool)
}
