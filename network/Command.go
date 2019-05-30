package network

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

/*
Server -> Client
  ONMAP
    <ANIM_ID>
  SET
    <ANIM_ID>
    <DATA_TYPE>
    <PNG_DATA>
Client -> Server
  GET
    <ANIM_ID>
*/
const (
	Nokay = iota
	Okay
	OnMap
	Set
	Get
	Reject
	Cya
)

type CommandAnimation struct {
	Type        uint8         // ONMAP->, SET->, ->GET
	AnimationID int           // Animation ID in question
	Animations  map[int][]int // Animations[ANIM_GROUP]->GRAPHIC[...]
}

func (c CommandAnimation) GetType() uint32 {
	return TypeAnimation
}

const (
	GraphicsPng = iota
)

type CommandGraphics struct {
	Type       uint8 // SET->, ->GET
	GraphicsID int   //
	DataType   uint8 // GRAPHICS_PNG, ...
	Data       []byte
}

func (c CommandGraphics) GetType() uint32 {
	return TypeGraphics
}

const (
	Travel = iota
)

type CommandMap struct {
	Type   uint8  // TRAVEL
	Name   string // target map name
	Width  int
	Height int
}

func (c CommandMap) GetType() uint32 {
	return TypeMap
}

type CommandObject struct {
	Type       uint8 // ADD->, REMOVE->, SET->
	GraphicsID int   //
	X          int   //
	Y          int   //
}

func (c CommandObject) GetType() uint32 {
	return TypeObjectUpdate
}

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

type CommandCmd struct {
	Cmd  int
	Args []string
}
type CommandExtCmd struct {
	Cmd  string
	Args []string
}

const (
	TypeBasic = iota
	TypeHandshake
	TypeLogin
	TypeCharacter
	TypeData
	TypeAnimation
	TypeAudio
	TypeObjectUpdate
	TypeInventoryUpdate
	TypeStatusUpdate
	TypeMap
	TypeCmd
	TypeExtCmd
	TypeGraphics
)
