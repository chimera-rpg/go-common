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

// ArchetypeToStringMap maps ArchetypeTypes to string representations
var ArchetypeToStringMap = map[ArchetypeType]string{
	ArchetypeUnknown: "Unknown",
	ArchetypeGenus:   "Genus",
	ArchetypeSpecies: "Species",
	ArchetypePC:      "PC",
	ArchetypeNPC:     "NPC",
	ArchetypeTile:    "Tile",
	ArchetypeBlock:   "Block",
	ArchetypeItem:    "Item",
	ArchetypeBullet:  "Bullet",
	ArchetypeGeneric: "Generic",
}

// StringToArchetypeMap maps string representations to ArchetypeTypes.
var StringToArchetypeMap = map[string]ArchetypeType{
	"Unknown": ArchetypeUnknown,
	"Genus":   ArchetypeGenus,
	"Species": ArchetypeSpecies,
	"PC":      ArchetypePC,
	"NPC":     ArchetypeNPC,
	"Tile":    ArchetypeTile,
	"Block":   ArchetypeBlock,
	"Item":    ArchetypeItem,
	"Bullet":  ArchetypeBullet,
	"Generic": ArchetypeGeneric,
}

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
	if v, ok := StringToArchetypeMap[value]; ok {
		*atype = v
		return nil
	}
	*atype = ArchetypeUnknown
	return fmt.Errorf("Unknown Type '%s'", value)
}

// MarshalYAML marshals an ArchetypeType into a string.
func (atype ArchetypeType) MarshalYAML() (interface{}, error) {
	if v, ok := ArchetypeToStringMap[atype]; ok {
		return v, nil
	}
	return "Unknown", nil
}
