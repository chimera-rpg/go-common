package data

// ArchetypeType is the numeric identifier for different archetype types.
type ArchetypeType uint8

const (
	ArchetypeUnknown ArchetypeType = iota
	ArchetypeGenus
	ArchetypeSpecies
	ArchetypePC
	ArchetypeNPC
	ArchetypeTile
	ArchetypeFloor
	ArchetypeWall
	ArchetypeItem
	ArchetypeBullet
	ArchetypeGeneric
)

// AsUint8 returns ArchetypeType as a uint8.
func (atype ArchetypeType) AsUint8() uint8 {
	return uint8(atype)
}
