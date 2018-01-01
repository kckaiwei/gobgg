# gobgg

Board game geek API in Go lang, support for XML API 2.

## Usage

Create a new Query, and pass it to the request type you would like to call.
```
type Query struct{
	Username 	string
	Domain 		string
	Id 			int
	Page 		int
	Buddies 	bool
	Guilds 		bool
	Hot 		bool
	Top 		bool
	Members		bool
	Sort 		string
	Type		string
}
```

### Example:

We create the query.

`u := gobgg.Query{Username: "whitesymphonia", Guilds: true}`

Followed by a calling the function GetUser, and passing the query.

`fmt.Println(gobgg.GetUser(u))`

## TODO

- Complete rest of XMLAPI2 commands
- Support logging of plays via cookies and passing username and password. Related link: (https://boardgamegeek.com/thread/1322486/api-feature-request-play-data)