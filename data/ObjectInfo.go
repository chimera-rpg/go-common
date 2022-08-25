package data

// ObjectInfo is the result of an inspection action.
type ObjectInfo struct {
	Remote []RemoteInfo
	Near   []NearInfo
}

// RemoteInfo represents information about an object when the character is far from it.
type RemoteInfo struct {
	Source  string // Source is where the knowledge came from. Empty means mundane.
	Name    string
	Quality string // Knowledgeable field
}

// NearInfo represents information about an object when the character is near it (within reach). Note that it's okay to send this much information, as the gob encoder does not sent zero/default values for types across the network.
type NearInfo struct {
	Source   string
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
