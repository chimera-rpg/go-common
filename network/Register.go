package network

import (
	"encoding/gob"
)

// RegisterCommands registers our various Command structures with their gob names.
func RegisterCommands() {
	gob.RegisterName("H", CommandHandshake{})
	gob.RegisterName("F", CommandFeatures{})
	gob.RegisterName("B", CommandBasic{})
	gob.RegisterName("M", CommandMap{})
	gob.RegisterName("L", CommandLogin{})
	gob.RegisterName("C", CommandCharacter{})
	gob.RegisterName("A", CommandAnimation{})
	gob.RegisterName("G", CommandGraphics{})
	gob.RegisterName("T", CommandTile{})
	gob.RegisterName("O", CommandObject{})
	gob.RegisterName("Oc", CommandObjectPayloadCreate{})
	gob.RegisterName("Od", CommandObjectPayloadDelete{})
	gob.RegisterName("Oa", CommandObjectPayloadAnimate{})
	gob.RegisterName("c", CommandCmd{})
	gob.RegisterName("e", CommandExtCmd{})
	gob.RegisterName("I", CommandInspect{})
	gob.RegisterName("Ii", CommandInspectPayloadInventory{})
	gob.RegisterName("Ic", CommandInspectPayloadCharacter{})
}
