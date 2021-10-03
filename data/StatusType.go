package data

// StatusType represents various Status effects.
type StatusType uint8

// StringToStatusMap is as the name implies.
var StringToStatusMap = map[string]StatusType{
	"Falling":   Falling,
	"Squeezing": Squeezing,
	"Crouching": Crouching,
}

// StatusMapToString is as the name implies.
var StatusMapToString = map[StatusType]string{
	Falling:   "Falling",
	Squeezing: "Squeezing",
	Crouching: "Crouching",
}

// Our various status types.
const (
	Falling StatusType = 0 << iota
	Squeezing
	Crouching
)
