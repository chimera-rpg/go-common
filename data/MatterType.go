package data

// MatterType represents the given matter state of an object.
type MatterType uint8

// StringToMatterMap maps string representations to MatterTypes.
var StringToMatterMap = map[string]MatterType{
	"None":     NoMatter,
	"Solid":    SolidMatter,
	"Liquid":   LiquidMatter,
	"Gas":      GasMatter,
	"Plasma":   PlasmaMatter,
	"Physical": PhysicalMatter,
	"Spirit":   SpiritMatter,
	"Arcane":   ArcaneMatter,
}

const (
	// NoMatter represents something that takes up no space.
	NoMatter = 0
	// SolidMatter represents solid objects, such as walls.
	SolidMatter MatterType = 1 << iota
	// LiquidMatter represents liquid objects, such as water.
	LiquidMatter
	// GasMatter represents gas objects.
	GasMatter
	// PlasmaMatter represents a state which we'll probably never use.
	PlasmaMatter
	// PhysicalMatter represents all matter states in the physical world.
	PhysicalMatter
	// SpiritMatter represents all matter states in the spirit world.
	SpiritMatter
	// ArcaneMatter represents all matter states in the arcane world.
	ArcaneMatter
)

// UnmarshalYAML unmarshals an MatterType from a string.
func (m *MatterType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var value []string
	if err := unmarshal(&value); err != nil {
		return err
	}
	for _, v := range value {
		if v, ok := StringToMatterMap[v]; ok {
			*m |= v
		}
	}
	return nil
}

// MarshalYAML marshals an MatterType into a string.
func (m MatterType) MarshalYAML() (interface{}, error) {
	var out []string
	if m.Is(SolidMatter) {
		out = append(out, "Solid")
	}
	if m.Is(LiquidMatter) {
		out = append(out, "Liquid")
	}
	if m.Is(GasMatter) {
		out = append(out, "Gas")
	}
	if m.Is(PlasmaMatter) {
		out = append(out, "Plasma")
	}
	if m.Is(PhysicalMatter) {
		out = append(out, "Physical")
	}
	if m.Is(SpiritMatter) {
		out = append(out, "Spirit")
	}
	if m.Is(ArcaneMatter) {
		out = append(out, "Arcane")
	}
	return out, nil
}

// Is returns whether one MatterType contains the types of another MatterType.
func (m *MatterType) Is(o MatterType) bool {
	return *m&o != 0
}
