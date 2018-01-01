/*
Package gobgg is an api library for Boardgamegeek's XML API 2

The package provides command functions that only require a Query to be passed.

type Query struct{
	Username    string
	Domain      string
	Id          int
	Page        int
	Buddies     bool
	Guilds      bool
	Hot         bool
	Top         bool
	Member      bool
	Sort        string
	Type        string
}

An Example:

u := gobgg.Query{Username: "whitesymphonia", Guilds: true}

Acceptable query parameters for each command can be found here:
https://boardgamegeek.com/wiki/page/BGG_XML_API2

After creating a Query, we pass it as an argument to any of the commands in the package, and the XML data will be returned in a struct.

fmt.Println(gobgg.GetUser(u))

 */
package gobgg
