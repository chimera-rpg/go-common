package data

import "fmt"

// ArchetypeType is the numeric identifier for different archetype types.
type ArchetypeType uint8

const (
	ArchetypeUnknown ArchetypeType = iota
	ArchetypeGenus
	ArchetypeSpecies
	ArchetypePC
	ArchetypeNPC
	ArchetypeTile
	ArchetypeBlock
	ArchetypeItem
	ArchetypeBullet
	ArchetypeGeneric
)

// AsUint8 returns ArchetypeType as a uint8.
func (atype ArchetypeType) AsUint8() uint8 {
	return uint8(atype)
}

// UnmarshalYAML unmarshals an ArchetypeType from a string.
func (atype *ArchetypeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value string
	if err := unmarshal(&value); err != nil {
		return err
	}
	switch value {
	case "Genus":
		*atype = ArchetypeGenus
	case "Species":
		*atype = ArchetypeSpecies
	case "PC":
		*atype = ArchetypePC
	case "NPC":
		*atype = ArchetypeNPC
	case "Tile":
		*atype = ArchetypeTile
	case "Block":
		*atype = ArchetypeBlock
	case "Item":
		*atype = ArchetypeItem
	case "Bullet":
		*atype = ArchetypeBullet
	case "Generic":
		*atype = ArchetypeGeneric
	default:
		*atype = ArchetypeUnknown
		return fmt.Errorf("Unknown Type '%s'", value)
	}
	return nil
}
