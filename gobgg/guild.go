package gobgg

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"io/ioutil"
	"strings"
)

type Guild struct {
	Name 		string 		`xml:"name,attr"`
	Id 			int			`xml:"id,attr"`
	Created		string		`xml:"created,attr"`
	Category	string		`xml:"category"`
	Website		string		`xml:"website"`
	Manager		string		`xml:"manager"`
	Description	string		`xml:"description"`
	Addr1		string		`xml:"location>addr1"`
	Addr2		string		`xml:"location>addr2"`
	City		string		`xml:"location>city"`
	Stateorprovince	string	`xml:"location>stateorprovince"`
	Postalcode	string		`xml:"location>postalcode"`
	Country		string		`xml:"location>country"`
	Members 	[]Member	`xml:"members>member"`
}

type Member struct {
	Name		string		`xml:"name,attr"`
	Joindate		string		`xml:"date,attr"`
}

// Gets guild information via GET, from query struct, and fills in default values
// Returns struct of guild information from XML
// Uses the following parameters: id=NNN, members=1, sort=SORTTYPE(username,data), page=NNN
func GetGuild(q Query)(g Guild){

	guild := Guild{}

	var Url *url.URL
	Url, err := url.Parse(BaseURL)
	if err != nil {
		log.Print("Error parsing url")
	}
	// Not enough data to work with
	if q.Id <= 0 {
		return Guild{}
	}
	if q.Page <= 0 {
		q.Page = 1
	}
	sort := "username"
	if strings.ToLower(q.Sort) == "date" {
		sort = "date"
	}
	members := "0"
	if q.Members {
		members = "1"
	}

	Url.Path += GuildSuffix
	parameters := url.Values{}
	parameters.Add("id", strconv.Itoa(q.Id))
	parameters.Add("members", members)
	parameters.Add("sort", sort)
	parameters.Add("page", strconv.Itoa(q.Page))
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
	xml.Unmarshal(data, &guild)

	return guild
}

