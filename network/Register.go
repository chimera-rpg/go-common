package network

import (
	"encoding/gob"
)

func RegisterCommands() {
	gob.RegisterName("H", CommandHandshake{})
	gob.RegisterName("B", CommandBasic{})
	gob.RegisterName("M", CommandMap{})
	gob.RegisterName("L", CommandLogin{})
	gob.RegisterName("C", CommandCharacter{})
}
