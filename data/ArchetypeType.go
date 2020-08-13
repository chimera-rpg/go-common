package data

import "fmt"

// ArchetypeType is the numeric identifier for different archetype types.
type ArchetypeType uint8

const (
	// ArchetypeUnknown represents an unknown archetype.
	ArchetypeUnknown ArchetypeType = iota
	// ArchetypeGenus represents a genus archetype.
	ArchetypeGenus
	// ArchetypeSpecies represents a species archetype.
	ArchetypeSpecies
	// ArchetypePC represents a PC archetype.
	ArchetypePC
	// ArchetypeNPC represents a NPC archetype.
	ArchetypeNPC
	// ArchetypeTile represents a Tile archetype.
	ArchetypeTile
	// ArchetypeBlock represents a Block archetype.
	ArchetypeBlock
	// ArchetypeItem represents an Item Archetype.
	ArchetypeItem
	// ArchetypeBullet represents a Bullet Archetype.
	ArchetypeBullet
	// ArchetypeGeneric represents a Generic Archetype.
	ArchetypeGeneric
	// ArchetypeSkill represents a Skill Archetype.
	ArchetypeSkill
	// ArchetypeArmor represents a Weapon Archetype.
	ArchetypeArmor
	// ArchetypeShield represents a Shield Archetype.
	ArchetypeShield
	// ArchetypeWeapon represents an Armor Archetype.
	ArchetypeWeapon
	// ArchetypeFood represents a tasty morsel.
	ArchetypeFood
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
	ArchetypeSkill:   "Skill",
	ArchetypeArmor:   "Armor",
	ArchetypeShield:  "Shield",
	ArchetypeWeapon:  "Weapon",
	ArchetypeFood:    "Food",
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
	"Skill":   ArchetypeSkill,
	"Armor":   ArchetypeArmor,
	"Shield":  ArchetypeShield,
	"Weapon":  ArchetypeWeapon,
	"Food":    ArchetypeFood,
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
