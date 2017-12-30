package gobgg

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"io/ioutil"
)

type Family struct {
	FName 		Name 		`xml:"item>name"`
	Thumbnail	string		`xml:"item>thumbnail"`
	Image 		string		`xml:"item>image"`
	Description	string		`xml:"item>description"`
	Members 	[]FMember	`xml:"item>link"`
}

type Name struct {
	Name			string		`xml:"value,attr"`
	SortIndex		int			`xml:"sortindex,attr"`
	Type			string		`xml:"type,attr"`
}

type FMember struct {
	Name			string		`xml:"value,attr"`
	Type 			string		`xml:"type,attr"`
	Id 				int			`xml:"id,attr"`
	Inbound			bool		`xml:"inbound,attr"`
}

// Gets Family information via GET, from query struct, and fills in default values
// Returns struct of family information from XML
func GetFamily(q Query)(f Family){

	family := Family{}

	var Url *url.URL
	Url, err := url.Parse(BaseURL)
	if err != nil {
		log.Print("Error parsing url")
	}
	// Not enough data to work with
	if q.Id <= 0 {
		return Family{}
	}

	Url.Path += FamilySuffix
	parameters := url.Values{}
	parameters.Add("id", strconv.Itoa(q.Id))
	parameters.Add("type", q.Type)
	Url.RawQuery = parameters.Encode()

	log.Print(Url.String())
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
	xml.Unmarshal(data, &family)

	return family
}

