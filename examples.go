package main

import (
	"fmt"
	"gobgg/gobgg"
)

func main () {
	// Get User Example
	u := gobgg.Query{Username: "whitesymphonia", Guilds: true}
	fmt.Println(gobgg.GetUser(u))

	// Get Guild Example
	g := gobgg.Query{Id: 1290, Members: true}
	fmt.Println(gobgg.GetGuild(g))

	// Get Family Example
	f := gobgg.Query{Id: 2}
	fmt.Println(gobgg.GetFamily(f))
}
