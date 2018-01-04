# gobgg 

[![GoDoc](https://godoc.org/github.com/kckaiwei/gobgg/gobgg?status.svg)](https://godoc.org/github.com/kckaiwei/gobgg/gobgg)

Board game geek API in Go lang, support for XML API 2.

## Usage

Create a new Query, and pass it to the request type you would like to call.
```
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
```

Acceptable query parameters for each command can be found here:
https://boardgamegeek.com/wiki/page/BGG_XML_API2

### Example:

We create the query.

`u := gobgg.Query{Username: "whitesymphonia", Guilds: true}`

Followed by a calling the function GetUser, and passing the query.

`fmt.Println(gobgg.GetUser(u))`

Full example:

```golang
package main

import (
	"fmt"
	"gobgg/gobgg"
)

func main () {
	// Get User Example
	u := gobgg.Query{Username: "whitesymphonia", Guilds: true}
	user, err := gobgg.GetUser(u)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}

	// Get User Example with Error
	us := gobgg.Query{Username: "", Guilds: true}
	username, err := gobgg.GetUser(us)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(username)
	}

	// Get Guild Example
	g := gobgg.Query{Id: 1290, Members: true}
	guild, err := gobgg.GetGuild(g)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(guild)
	}

	// Get Family Example
	f := gobgg.Query{Id: 2}
	family, err := gobgg.GetFamily(f)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(family)
	}

	// Get ForumList Example
	fl := gobgg.Query{Id: 174430, Type: "thing"}
	forumlist, err := gobgg.GetForumlist(fl)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(forumlist)
	}
}

```

## TODO

- Complete rest of XMLAPI2 commands
- Datetime XML unmarshal conversion. Related link: (https://stackoverflow.com/questions/17301149/golang-xml-unmarshal-and-time-time-fields)
- Support logging of plays via cookies and passing username and password. Related link: (https://boardgamegeek.com/thread/1322486/api-feature-request-play-data)