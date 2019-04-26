package net

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type Connection struct {
	IsConnected bool
	Conn        net.Conn
	Encoder     *gob.Encoder
	Decoder     *gob.Decoder
	CmdChan     chan Command  // Becomes valid for reading after ConnectTo(...). See LoopCmd
	ClosedChan  chan struct{} // Has close(...) called upon it in Close()
}

func (c *Connection) SetConn(conn net.Conn) {
	if c.IsConnected == true {
		c.Close()
	}
	c.Conn = conn
	c.Encoder = gob.NewEncoder(conn)
	c.Decoder = gob.NewDecoder(conn)
	c.CmdChan = make(chan Command)
	c.ClosedChan = make(chan struct{})
	c.IsConnected = true
}

func (c *Connection) ConnectTo(address string) (err error) {
	c.Conn, err = net.Dial("tcp", address)
	if err != nil {
		return
	}
	c.Encoder = gob.NewEncoder(c.Conn)
	c.Decoder = gob.NewDecoder(c.Conn)
	c.CmdChan = make(chan Command)
	c.ClosedChan = make(chan struct{})
	c.IsConnected = true
	// I'm unsure if we should start our Command Loop channel coroutine here as it prevents and use of Send/Receive to the owner of the Connection. However, I suspect it is fine, as we should probably just use the LoopCmd 100% of the time when a client is connected to a server.
	go c.LoopCmd()
	return
}

func (c *Connection) Send(cmd Command) (err error) {
	err = c.Encoder.Encode(&cmd)
	return
}

func (c *Connection) Receive(cmd *Command) (err error) {
	err = c.Decoder.Decode(&cmd)
	return
}

func (c *Connection) ReceiveCommandBasic() (b CommandBasic) {
	var command Command
	c.Receive(&command)
	switch t := command.(type) {
	case CommandBasic:
		b = t
	default:
		panic(fmt.Errorf("Expected Net.CommandBasic(%d), got: %d\n", TYPE_BASIC, t.GetType()))
	}
	return
}

func (c *Connection) ReceiveCommandHandshake() (hs CommandHandshake) {
	var command Command
	c.Receive(&command)
	switch t := command.(type) {
	case CommandHandshake:
		hs = t
	default:
		panic(fmt.Errorf("Expected Net.CommandHandshake(%d), got: %d\n", TYPE_BASIC, t.GetType()))
	}
	return
}

func (c *Connection) Close() {
	if c.IsConnected == false {
		return
	}
	c.IsConnected = false
	if r := recover(); r != nil {
		log.Print("Closing due to problematic connection.")
	} else {
		c.Send(CommandBasic{
			Type: CYA,
		})
	}
	c.Conn.Close()
	var blank struct{}
	c.ClosedChan <- blank
}

func (c *Connection) LoopCmd() {
	var cmd Command
	var err error
	for c.IsConnected {
		err = c.Receive(&cmd)
		if err != nil {
			c.Close()
			break
		}
		c.CmdChan <- cmd
	}
}
