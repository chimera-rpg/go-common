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
	Type        uint8                       // ONMAP->, SET->, ->GET
	AnimationID uint32                      // Animation ID in question
	Faces       map[uint32][]AnimationFrame // FaceID to Frames
}

type AnimationFrame struct {
	ImageID uint32
	Time    int
}

func (c CommandAnimation) GetType() uint32 {
	return TypeAnimation
}

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

// CommandTile is a list of tiles at a given Tile. This might be expanded to also have a brightness/visibility value.
type CommandTile struct {
	X, Y, Z   uint32
	ObjectIDs []uint32
}

// GetType returns TypeTileUpdate
func (c CommandTile) GetType() uint32 {
	return TypeTileUpdate
}

// ...is it appropriate to use interfaces within gobs as we are below...?
// CommandObject is the command type used to create, delete, and update objects.
type CommandObject struct {
	Type     uint8  //
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
	AnimationID uint32
	FaceID      uint32
	X, Y, Z     uint32
}

// CommandObjectPayloadDelete is the type indicating that an object should be deleted.
type CommandObjectPayloadDelete struct {
}

// CommandObjectPayloadTravel is the type used for doing interpolated travel from one position to another.
type CommandObjectPayloadTravel struct {
	X, Y, Z uint32
}

// CommandObjectPayloadMove is the type used for doing a strict move from one position to another.
type CommandObjectPayloadMove struct {
	X, Y, Z uint32
}

// CommandObjectPayloadAnimate is the type used for updating an objects animation and face.
type CommandObjectPayloadAnimate struct {
	AnimationID uint32 //
	FaceID      uint32 //
}

const (
	ObjectCreate  = iota // used to create an object with given id.
	ObjectDelete         // used to completely delete given object.
	ObjectTravel         // used for client-side interpolated movement.
	ObjectMove           // teleport given object.
	ObjectAnimate        // whether used to set AnimationID and FaceID.
)

// CommandCmd is used for player commands to interact with the game world.
type CommandCmd struct {
	Cmd  int
	Args []string
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

// CommandExtCmd is used for extended player commands with variadic inputs.
type CommandExtCmd struct {
	Cmd  string
	Args []string
}

// A list of all our command types.
const (
	TypeBasic = iota
	TypeHandshake
	TypeLogin
	TypeCharacter
	TypeData
	TypeAnimation
	TypeAudio
	TypeTileUpdate
	TypeObjectUpdate
	TypeInventoryUpdate
	TypeStatusUpdate
	TypeMap
	TypeCmd
	TypeExtCmd
	TypeGraphics
)
