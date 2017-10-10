package apl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// OfferingIndex is
type OfferingIndex struct {
	Disclaimer      string `json:"disclaimer"`
	FormatVersion   string `json:"formatVersion"`
	Offers          Offers `json:"offers"`
	PublicationDate string `json:"publicationDate"`
}

// Offers is map of Offer
type Offers map[string]Offer

// Offer is
type Offer struct {
	CurrentRegionIndexURL string `json:"currentRegionIndexUrl"`
	CurrentVersionURL     string `json:"currentVersionUrl"`
	OfferCode             string `json:"offerCode"`
	VersionIndexURL       string `json:"versionIndexUrl"`
}

func (o OfferingIndex) fetch() {

	raw, err := ioutil.ReadFile("./index.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(raw, &o)

}

// ListOfferCode returns list of OfferCode
func (o OfferingIndex) ListOfferCode() []string {
	var ret []string

	o.fetch()

	for _, v := range o.Offers {
		ret = append(ret, v.OfferCode)
	}
	return ret
}
