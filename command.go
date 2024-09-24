package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

// Throws an error when a name is already in use
func (c *commands) register(name string, f func(*state, command) error) error {
	_, ok := c.cmds[name]
	if ok {
		return fmt.Errorf("there already exists function for the %s keyword.", name)
	}
	c.cmds[name] = f
	return nil
}

func (c *commands) unRegister(name string) {
	delete(c.cmds, name)
}
