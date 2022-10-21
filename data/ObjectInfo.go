package data

// ObjectInfo represents information about an object. Note that these fields are only sent if they are non-zero value, so this is efficient.
type ObjectInfo struct {
	Source    string
	Near      bool
	Name      string
	Quality   string
	Weight    float64
	Worth     float64
	Count     int
	Matter    MatterType
	Material  int
	Lore      string
	Value     float64
	Reach     int
	Slots     ObjectInfoSlots
	TypeHints []uint32
	// Armor float64 // should this just be Value?
	//AttackTypes *AttackTypes
	//DamageTypes *DamageTypes
	// Spell ???
}

type ObjectInfoSlots struct {
	Has   []uint32
	Uses  []uint32
	Needs struct {
		Min []uint32
		Max []uint32
	}
	Gives []uint32
}

/*type DamageTypes struct {
	Value            int
	AttributeBonuses map[AttackType]map[uint8]float64
}*/
