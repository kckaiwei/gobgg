package gobgg

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"io/ioutil"
)

type User struct {
	XMLName 	xml.Name	`xml:"user"`
	Id 			int			`xml:"id,attr"`
	Username 	string 		`xml:"name,attr"`
	FirstName 	firstname	`xml:"firstname"`
	LastName	lastname	`xml:"lastname"`
	}

type firstname struct {
	firstname 			string `xml:"value,attr"`
}

type lastname struct {
	lastname 			string `xml:"value,attr"`
}

type Details struct {
	FirstName 			string	`xml:"value, attr"`
	LastName 			string	`xml:"lastname"`
	AvatarLink 			string	`xml:"avatarlink"`
	YearRegistered 		string	`xml:"yearregistered"`
	LastLogin 			string	`xml:"lastlogin"`
	StateOrProvince 	string	`xml:"stateorprovince"`
	Country 			string	`xml:"country"`
	Webaddress 			string	`xml:"webaddress"`
	Xboxaccount 		string	`xml:"xboxaccount"`
	Wiiaccount  		string	`xml:"wiiaccount"`
	Psnaccount 			string	`xml:"psnaccount"`
	Battlenetaccount 	string	`xml:"battlenetaccount"`
	Steamaccount 		string	`xml:"steamaccount"`
	Traderating 		string	`xml:"traderating"`
	Marketrating 		string	`xml:"marketrating"`
	Buddies				[]Buddy
}

type Buddy struct {
	ID 					int
	Name 				string
}

func GetUser(username, domain string, page int, buddies, guilds, hot, top bool){
	// Get User info

	user := User{}

	var Url *url.URL
	Url, err := url.Parse(BaseURL)
	if err != nil {
		log.Print("Error parsing url")
	}

	// Set defaults
	bd, gd, ht, tp := "0", "0", "0", "0"
	if domain != "" {
		domain = "boardgame"
	}
	if page <= 0 {
		page = 1
	}
	if buddies {
		bd = "1"
	}
	if guilds {
		gd = "1"
	}
	if hot {
		ht = "1"
	}
	if top {
		tp = "1"
	}

	Url.Path += UserSuffix
	parameters := url.Values{}
	parameters.Add("name", username)
	parameters.Add("buddies", bd)
	parameters.Add("guilds", gd)
	parameters.Add("hot", ht)
	parameters.Add("top", tp)
	parameters.Add("domain", domain)
	parameters.Add("page", strconv.Itoa(page))
	Url.RawQuery = parameters.Encode()

	log.Print(Url.String())
	// str name, bool buddies, guilds, hot, top, str domain, page
	resp, err := http.Get(Url.String())
	if err != nil{
		log.Print(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		log.Printf("Status error: %v", resp.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read body: %v", err)
	}
	log.Print(string(data))
	xml.Unmarshal(data, &user)

	log.Print(user)
	log.Print(user.Id)

	//log.Print(resp)
	//log.Print(resp.Body)
	time.Sleep(10)
	return
}
