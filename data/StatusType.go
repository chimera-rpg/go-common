package data

// StatusType represents various Status effects.
type StatusType uint8

// Our various status types.
const (
	NoStatus                 = 0
	FallingStatus StatusType = 1 << iota
	SqueezingStatus
	CrouchingStatus
)

// StringToStatusMap is as the name implies.
var StringToStatusMap = map[string]StatusType{
	"Falling":   FallingStatus,
	"Squeezing": SqueezingStatus,
	"Crouching": CrouchingStatus,
}

// StatusMapToString is as the name implies.
var StatusMapToString = map[StatusType]string{
	FallingStatus:   "Falling",
	SqueezingStatus: "Squeezing",
	CrouchingStatus: "Crouching",
}
