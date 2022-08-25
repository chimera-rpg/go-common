package network

import (
	"encoding/gob"

	"github.com/chimera-rpg/go-common/data"
)

// RegisterCommands registers our various Command structures with their gob names.
func RegisterCommands() {
	gob.RegisterName("H", CommandHandshake{})
	gob.RegisterName("F", CommandFeatures{})
	gob.RegisterName("B", CommandBasic{})
	gob.RegisterName("M", CommandMap{})
	gob.RegisterName("L", CommandLogin{})
	gob.RegisterName("R", CommandRejoin{})
	gob.RegisterName("C", CommandCharacter{})
	gob.RegisterName("A", CommandAnimation{})
	gob.RegisterName("G", CommandGraphics{})
	gob.RegisterName("T", CommandTile{})
	gob.RegisterName("Tl", CommandTileLight{})
	gob.RegisterName("O", CommandObject{})
	gob.RegisterName("Oc", CommandObjectPayloadCreate{})
	gob.RegisterName("Od", CommandObjectPayloadDelete{})
	gob.RegisterName("Oa", CommandObjectPayloadAnimate{})
	gob.RegisterName("Ov", CommandObjectPayloadViewTarget{})
	gob.RegisterName("c", CommandCmd{})
	gob.RegisterName("cl", CommandClearCmd{})
	gob.RegisterName("e", CommandExtCmd{})
	gob.RegisterName("r", CommandRepeatCmd{})
	gob.RegisterName("m", CommandMessage{})
	gob.RegisterName("s", CommandStatus{})
	gob.RegisterName("t", CommandStamina{})
	gob.RegisterName("I", CommandInspect{})
	gob.RegisterName("Ib", data.ObjectInfoBasic{})
	gob.RegisterName("Iq", data.ObjectInfoQuantity{})
	gob.RegisterName("Ii", data.ObjectInfoInventory{})
	gob.RegisterName("Iw", data.ObjectInfoWeapon{})
	gob.RegisterName("Ia", data.ObjectInfoArmor{})
	gob.RegisterName("It", data.ObjectInfoBlock{})
	gob.RegisterName("Vp", CommandViewport{})
	gob.RegisterName("S", CommandSound{})
	gob.RegisterName("a", CommandAudio{})
	gob.RegisterName("n", CommandNoise{})
	gob.RegisterName("Mu", CommandMusic{})
	gob.RegisterName("At", CommandAttack{})
	gob.RegisterName("D", CommandDamage{})
	gob.RegisterName("In", CommandInteract{})
}
