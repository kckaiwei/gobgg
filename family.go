package gobgg

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"io/ioutil"
	"errors"
)

type Family struct {
	FName 		name 		`xml:"item>name"`
	Thumbnail	string		`xml:"item>thumbnail"`
	Image 		string		`xml:"item>image"`
	Description	string		`xml:"item>description"`
	Members 	[]fMember	`xml:"item>link"`
}

type name struct {
	Name			string		`xml:"value,attr"`
	SortIndex		int			`xml:"sortindex,attr"`
	Type			string		`xml:"type,attr"`
}

type fMember struct {
	Name			string		`xml:"value,attr"`
	Type 			string		`xml:"type,attr"`
	Id 				int			`xml:"id,attr"`
	Inbound			bool		`xml:"inbound,attr"`
}

// Gets Family information via GET, from query struct, and fills in default values
// Returns struct of family information from XML
// Uses the following parameters: id=NNN, type=FAMILYTYPE(rpg,rpgperiodical,boardgamefamily)
func GetFamily(q Query)(f Family, e error){

	family := Family{}

	var Url *url.URL
	Url, err := url.Parse(BaseURL)
	if err != nil {
		log.Print("Error parsing url")
		return Family{}, errors.New("BaseURLInvalid")
	}
	// Not enough data to work with
	if q.Id <= 0 {
		return Family{}, errors.New("NoID")
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
		return Family{}, errors.New("GetRequestFailed")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		log.Printf("Status error: %v", resp.StatusCode)
		return Family{}, &StatusError{"StatusError", resp.StatusCode}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read body: %v", err)
		return Family{}, errors.New("ResponseReadError")
	}
	xml.Unmarshal(data, &family)

	return family, nil
}

