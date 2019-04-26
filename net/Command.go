package net

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

// GetType returns TYPE_BASIC
func (c CommandBasic) GetType() uint32 {
	return TYPE_BASIC
}

// CommandHandshake represents the handshake between the server and the client
// so as to ensure compatibility.
type CommandHandshake struct {
	Version int
	Program string
}

// GetType returns TYPE_HANDSHAKE
func (c CommandHandshake) GetType() uint32 {
	return TYPE_HANDSHAKE
}

// Versioning. This should probably be different.
const (
	VERSION = 1
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
	return TYPE_LOGIN
}

// These are the CommandLogin Types
const (
	QUERY    = 0
	LOGIN    = 1
	REGISTER = 2
	DELETE   = 3
)

type CommandCharacter struct {
	Type        uint8
	Race        string // Race used for query
	Class       string // Class used for query
	Image       []byte // Image for Race/Class/Character
	Character   string // Name to be created, loaded, or deleted.
	Description string // Description used for query
	Stats       []string
	Skills      []string
}

func (c CommandCharacter) GetType() uint32 {
	return TYPE_CHARACTER
}

const (
	QUERY_RACE          = iota // Queries Race information
	QUERY_CLASS                // Queries Class information
	QUERY_CHARACTER            // Query Character information
	CREATE_CHARACTER           // Creates a Character
	ADJUST_CHARACTER           // Adjusts an in-progress Character (race or class)
	LOAD_CHARACTER             // Loads a character (by name)
	DELETE_CHARACTER           // Deletes a character (by name)
	ROLL_ABILITY_SCORES        // Requests(client) or returns(server) rolls for ability scores
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
	NOK    = 0
	OK     = 1
	ONMAP  = 2
	SET    = 3
	GET    = 4
	REJECT = 5
	CYA    = 6
)

type CommandAnimation struct {
	Type        uint8         // ONMAP->, SET->, ->GET
	AnimationID int           // Animation ID in question
	Animations  map[int][]int // Animations[ANIM_GROUP]->GRAPHIC[...]
}

func (c CommandAnimation) GetType() uint32 {
	return TYPE_ANIMATION
}

const (
	GRAPHICS_PNG = 1
)

type CommandGraphics struct {
	Type       uint8 // SET->, ->GET
	GraphicsID int   //
	DataType   uint8 // GRAPHICS_PNG, ...
	Data       []byte
}

func (c CommandGraphics) GetType() uint32 {
	return TYPE_GRAPHICS
}

const (
	TRAVEL = 1
)

type CommandMap struct {
	Type   uint8  // TRAVEL
	Name   string // target map name
	Width  int
	Height int
}

func (c CommandMap) GetType() uint32 {
	return TYPE_MAP
}

type CommandObject struct {
	Type       uint8 // ADD->, REMOVE->, SET->
	GraphicsID int   //
	X          int   //
	Y          int   //
}

func (c CommandObject) GetType() uint32 {
	return TYPE_OBJECT_UPDATE
}

const (
	NORTH = 0
	SOUTH = 1
	EAST  = 2
	WEST  = 3
	DROP  = 4
	QUIT  = 5
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
	TYPE_BASIC = iota
	TYPE_HANDSHAKE
	TYPE_LOGIN
	TYPE_CHARACTER
	TYPE_DATA
	TYPE_ANIMATION
	TYPE_AUDIO
	TYPE_OBJECT_UPDATE
	TYPE_INVENTORY_UPDATE
	TYPE_STATUS_UPDATE
	TYPE_MAP
	TYPE_CMD
	TYPE_EXTCMD
	TYPE_GRAPHICS
)
