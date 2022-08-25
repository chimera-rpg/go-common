package data

// ObjectInfo represents information about an object. Note that these fields are only sent if they are non-zero value, so this is efficient.
type ObjectInfo struct {
	Source   string
	Name     string
	Quality  string
	Weight   float64
	Worth    float64
	Count    int
	Matter   MatterType
	Material int
	Lore     string
	Value    float64
	Reach    int
	// Armor float64 // should this just be Value?
	//AttackTypes *AttackTypes
	//DamageTypes *DamageTypes
	// Spell ???
}

/*type DamageTypes struct {
	Value            int
	AttributeBonuses map[AttackType]map[uint8]float64
}*/
