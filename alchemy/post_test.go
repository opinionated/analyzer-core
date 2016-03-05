package alchemy_test

import (
	"fmt"
	"github.com/opinionated/analyzer-core/alchemy"
	"testing"
)

// quickly test each of the alchemy calls
func ParseKimDavis() string {
	articleBody, err := alchemy.ParseArticle("test/kimDavisRelease.txt")
	if err != nil {
		panic(err) // shouldn't ever happen
	}
	return articleBody
}

func TestFetchKeywords(t *testing.T) {
	articleBody := ParseKimDavis()

	processed := alchemy.Keywords{}
	err := alchemy.GetKeywords(articleBody, &processed)

	if err != nil {
		t.Errorf("could not send request:", err)
	}

	if processed.Keywords[0].Text != "marriage licenses" {
		t.Errorf("expected top word to be marriage licenses but did not get")
	}
}

func TestFetchTaxonomy(t *testing.T) {
	articleBody := ParseKimDavis()
	url := alchemy.BuildRequest("Taxonomy", articleBody)

	processed := alchemy.Taxonomys{}
	response, err := alchemy.Request(url)

	if err != nil {
		t.Errorf("invalid response:", err)
	}
	err = alchemy.ConvertResponseXML(response, &processed)
	if err != nil {
		t.Errorf("could not send request:", err)
	}

	fmt.Println(processed.Taxonomys)
	if processed.Taxonomys[0].Label != "/society/social institution/marriage" {
		t.Errorf("expected top word to be marriage licenses but did not get")
	}
}
