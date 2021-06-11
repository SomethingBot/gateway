package api

import "sync"

type Command struct {
	//Name is the command Name
	Name string `json:"name"`

	//Prefix is true when Name is a prefix and not a command
	Prefix bool
}

type CommandHandler func() (err error)

type CommandParser struct {
	sync.RWMutex
	commands map[string]CommandHandler
}
