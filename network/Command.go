package network

import (
	"github.com/chimera-rpg/go-common/data"
)

// Command is our interface for all commands.
type Command interface {
	GetType() uint32
}

// CommandBasic represent very simple transmissions between the server and
// the client. This is used for disconnects among other things.
type CommandBasic struct {
	Type   uint8
	String string
}

// GetType returns TypeBasic
func (c CommandBasic) GetType() uint32 {
	return TypeBasic
}

// CommandHandshake represents the handshake between the server and the client
// so as to ensure compatibility.
type CommandHandshake struct {
	Version int
	Program string
}

// GetType returns TypeHandshake
func (c CommandHandshake) GetType() uint32 {
	return TypeHandshake
}

// Versioning. This should probably be different.
const (
	Version = iota
)

// CommandFeatures handles the communication of the features of the server, such as animations sizes, to the client.
type CommandFeatures struct {
	AnimationsConfig data.AnimationsConfig
}

// GetType returns TypeFeatures
func (c CommandFeatures) GetType() uint32 {
	return TypeFeatures
}

// CommandLogin handles the process of logging in, registering, recovering
// a password via email, and even deleting the account.
type CommandLogin struct {
	Type  uint8
	User  string
	Pass  string
	Email string
}

// GetType returns TYPE_LOGIN
func (c CommandLogin) GetType() uint32 {
	return TypeLogin
}

// These are the CommandLogin Types
const (
	Query = iota
	Login
	Register
	Delete
)

// CommandCharacter is the command involved in resolving:
//	  * Species, Culture, Training availability
//			* Image, AbilityScores, Skills, Description
//		* Querying, selection, creation, and deletion of Player Character(s)
//			* Species, Culture, Training, Character, Image, Description, AbilityScores, Skills
type CommandCharacter struct {
	Type          uint8
	Genera        []string // Genera used for query (humanoids, etc.)
	Species       []string // Species used for query (dwarf, elf, etc.)
	Cultures      []string // Cultures used for query (mountain dwarf, etc.)
	Trainings     []string // Trainings used for query
	Images        [][]byte // Images for Species/Culture/Training
	Characters    []string // Name(s) to be created, loaded, or deleted.
	Levels        []uint16 // Level of character.
	Descriptions  []string // Description used for query
	AbilityScores [][]string
	Skills        [][]string
}

// GetType returns TypeCharacter
func (c CommandCharacter) GetType() uint32 {
	return TypeCharacter
}

// These const values provide the sub-types for CommandCharacter
const (
	QueryGenera       = iota // Query of Genera, Image(?), Description. Sent when CreateCharacter.
	QuerySpecies             // Query of Genera+Species, Image, Description, AbilityScores, Skills
	QueryCultures            // Query of Genera+Species+Cultures, Image(?), Description, AbilityScores, Skills
	QueryTrainings           // Query of Genera+Species+Cultures+Trainings, Image(?), Description, Skills
	QueryCharacters          // Query of available characters. Sent when client connects.
	CreateCharacter          // Creates a Character->(name).
	AdjustCharacter          // Adjusts an in-progress Character Character->(Species, Culture, Training, AbilityScores)
	ChooseCharacter          // Chooses a character Character->() and logs in.
	DeleteCharacter          // Deletes a character Character->()
	RollAbilityScores        // Requests(client) or returns(server) rolls for ability scores Character->(AbilityScores)
)

// Our basic return types
const (
	Nokay = iota
	Okay
	OnMap
	Set
	Get
	Reject
	Cya
)

// CommandAnimation is for setting and/or getting animation ID->FaceIDs->Frames
type CommandAnimation struct {
	Type        uint8                       // ONMAP->, SET->, ->GET
	AnimationID uint32                      // Animation ID in question
	Faces       map[uint32][]AnimationFrame // FaceID to Frames
}

// AnimationFrame represents an imageID and how long it should play.
type AnimationFrame struct {
	ImageID uint32
	Time    int
}

// GetType returns TypeAnimation.
func (c CommandAnimation) GetType() uint32 {
	return TypeAnimation
}

// Our Graphics data types.
const (
	GraphicsPng = iota
)

// CommandGraphics are for setting and requesting images.
type CommandGraphics struct {
	Type       uint8  // SET->, ->GET
	GraphicsID uint32 //
	DataType   uint8  // GRAPHICS_PNG, ...
	Data       []byte
}

// GetType returns TypeGraphics.
func (c CommandGraphics) GetType() uint32 {
	return TypeGraphics
}

// Our CommandMap.Type constants.
const (
	Travel = iota
)

// CommandMap is a basic command for creating a map of a given name and ID at provided dimensions.
type CommandMap struct {
	Type   uint8 // TRAVEL
	MapID  uint32
	Name   string // target map name
	Height int
	Width  int
	Depth  int
}

// GetType returns TypeMap
func (c CommandMap) GetType() uint32 {
	return TypeMap
}

// CommandTile is a list of tiles at a given Tile. This might be expanded to also have a brightness/visibility value.
type CommandTile struct {
	X, Y, Z   uint32
	ObjectIDs []uint32
}

// GetType returns TypeTileUpdate
func (c CommandTile) GetType() uint32 {
	return TypeTileUpdate
}

// CommandObject is the command type used to create, delete, and update objects.
type CommandObject struct {
	ObjectID uint32 // id of target object
	Payload  CommandObjectPayload
}

// GetType returns TypeObjectUpdate
func (c CommandObject) GetType() uint32 {
	return TypeObjectUpdate
}

// CommandObjectPayload is a generic interface for actual payloads.
type CommandObjectPayload interface {
}

// CommandObjectPayloadCreate is the type for creating a new object.
type CommandObjectPayloadCreate struct {
	TypeID               uint8
	AnimationID          uint32
	FaceID               uint32
	Height, Width, Depth uint8
}

// CommandObjectPayloadDelete is the type indicating that an object should be deleted.
type CommandObjectPayloadDelete struct {
}

// CommandObjectPayloadAnimate is the type used for updating an object's animation and face.
type CommandObjectPayloadAnimate struct {
	AnimationID uint32 //
	FaceID      uint32 //
}

// CommandObjectPayloadViewTarget is the type used for marking a given object as the client's view target.
type CommandObjectPayloadViewTarget struct {
}

// Our Object types (unused)
const (
	ObjectCreate     = iota // used to create an object with given id.
	ObjectDelete            // used to completely delete given object.
	ObjectAnimate           // whether used to set AnimationID and FaceID.
	ObjectViewTarget        // used to target the object as the client's view.
)

// CommandInspect is used to inspect objects, characters, inventory, and similar.
type CommandInspect struct {
	updates []CommandInspectPayload
}

// GetType returns TypeInspect
func (c CommandInspect) GetType() uint32 {
	return TypeInspect
}

// CommandInspectPayload is our interface for inspection types.
type CommandInspectPayload interface {
}

// CommandInspectPayloadInventory is used to inspect the player's own inventory.
type CommandInspectPayloadInventory struct {
	ObjectID uint32
	Name     string
	Equipped bool
}

// CommandInspectPayloadCharacter is used to inspect a character.
type CommandInspectPayloadCharacter struct {
	ObjectID uint32
	Name     string
}

// CommandCmd is used for player commands to interact with the game world.
type CommandCmd struct {
	Cmd  int
	Args []string
}

// GetType returns TypeCmd
func (c CommandCmd) GetType() uint32 {
	return TypeCmd
}

// Our various CommandCmd.Cmd values
const (
	North = iota
	South
	East
	West
	Northeast
	Northwest
	Southeast
	Southwest
	Brace
	Drop
	Quit
)

// CommandExtCmd is used for extended player commands with variadic inputs.
type CommandExtCmd struct {
	Cmd  string
	Args []string
}

// GetType returns TypeExtCmd
func (c CommandExtCmd) GetType() uint32 {
	return TypeExtCmd
}

// A list of all our command types.
const (
	TypeBasic = iota
	TypeHandshake
	TypeFeatures
	TypeLogin
	TypeCharacter
	TypeData
	TypeAnimation
	TypeAudio
	TypeTileUpdate
	TypeObjectUpdate
	TypeInventoryUpdate
	TypeInspect
	TypeStatusUpdate
	TypeMap
	TypeCmd
	TypeExtCmd
	TypeGraphics
)
