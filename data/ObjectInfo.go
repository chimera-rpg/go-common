package data

type ObjectInfoBasic struct {
	Name        string
	Description string
}

type ObjectInfoQuantity struct {
	Count int
}

type ObjectInfoInventory struct {
	Equipped bool
}

type ObjectInfoMatter struct {
	Matter MatterType
}

type ObjectInfoWeapon struct {
	Reach       int
	AttackTypes AttackTypes
	Damage      struct {
		Value            int
		AttributeBonuses map[AttackType]map[uint8]float64
	}
}

type ObjectInfoArmor struct {
	Resistances AttackTypes
	Armor       float64
}

type ObjectInfoBlock struct {
	// ??
}
