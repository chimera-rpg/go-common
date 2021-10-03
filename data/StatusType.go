package data

// StatusType represents various Status effects.
type StatusType uint8

StringToStatusMap = map[string]StatusType{
	"Falling": Falling,
	"Squeezing": Squeezing,
	"Crouching": Crouching,
}

const (
	Falling StatusType = 0 << iota
	Squeezing
	Crouching
)
