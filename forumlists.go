package gobgg

import (
	"encoding/xml"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"io/ioutil"
	"strings"
	"errors"
)

type ForumList struct {
	Type 		string		`xml:"type,attr"`
	Subforum	[]sforum	`xml:"forum"`
}

type sforum struct {
	Title        string `xml:"title,attr"`
	Groupid      int    `xml:"groupid,attr"`
	Forumid      int    `xml:"id,attr"`
	Description  string `xml:"description,attr"`
	NoPosting    bool   `xml:"noposting,attr"`
	Numthreads   int    `xml:"numthreads,attr"`
	Numposts     int    `xml:"numposts,attr"`
	LastPostDate string `xml:"lastpostdate,attr"`
}

// Gets Forumlist information via GET, from query struct, and fills in default values
// Returns struct of forumlist information from XML
// Uses the following parameters: id=NNN, type=[thing,family]
func GetForumlist(q Query)(f ForumList, e error){

	fl := ForumList{}

	var Url *url.URL
	Url, err := url.Parse(BaseURL)
	if err != nil {
		log.Print("Error parsing url")
		return ForumList{}, errors.New("BaseURLInvalid")
	}
	// Not enough data to work with
	if q.Id <= 0 {
		return ForumList{}, errors.New("NoIDGiven")
	}

	// Must have type to work
	if q.Type == "" {
		return ForumList{}, errors.New("NoTypeGiven")
	}

	if strings.ToLower(q.Type) != "thing" && strings.ToLower(q.Type) != "family"{
		return ForumList{}, errors.New("InvalidTypeGiven")
	}

	Url.Path += ForumListSuffix
	parameters := url.Values{}
	parameters.Add("id", strconv.Itoa(q.Id))
	parameters.Add("type", q.Type)
	Url.RawQuery = parameters.Encode()

	log.Print(Url.String())
	resp, err := http.Get(Url.String())
	if err != nil{
		log.Print(err)
		return ForumList{}, errors.New("GetRequestFailed")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		log.Printf("Status error: %v", resp.StatusCode)
		return ForumList{}, &StatusError{"StatusError", resp.StatusCode}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read body: %v", err)
		return ForumList{}, errors.New("ResponseReadError")
	}
	xml.Unmarshal(data, &fl)

	return fl, nil
}

