package alchemy_test

import (
	"encoding/xml"
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
	url := alchemy.BuildRequest("Keywords", articleBody)

	processed := alchemy.KeywordsResult{}
	err := alchemy.Request(url, &processed)

	if err != nil {
		t.Errorf("could not send request:", err)
	}

	// test that we got what we expected
	if processed.Status != "OK" {
		t.Errorf("expected status OK")
	}

	if processed.Keywords.Keywords[0].Text != "marriage licenses" {
		t.Errorf("expected top word to be marriage licenses but did not get")
	}
}

func TestFetchTaxonomy(t *testing.T) {
	articleBody := ParseKimDavis()

	processed, err := alchemy.GetTaxonomy(articleBody)
	if err != nil {
		t.Errorf("could not send request:", err)
	}

	// test that we got what we expected

	if processed.Taxonomys.Taxonomys[0].Label != "/society/social institution/marriage" {
		t.Errorf("expected top word to be marriage licenses but did not get")
	}
}

func TestFetchEntities(t *testing.T) {
	articleBody := ParseKimDavis()

	resp, err := alchemy.GetEntities(articleBody)
	if err != nil {
		t.Errorf("unexpected error parsing: %s\n", err)
	}

	for _, entity := range resp.Entities.Entities {
		fmt.Println(entity.Text, ":", entity.Sentiment)
		//fmt.Println(entity.Text, ":", entity.Disambiguated)
	}

	_, err = xml.MarshalIndent(resp, " ", " ")
	if err != nil {
		panic(err)
	}

}

func TestFetchConcepts(t *testing.T) {
	articleBody := ParseKimDavis()

	resp, err := alchemy.GetConcepts(articleBody)
	if err != nil {
		t.Errorf("unexpected error parsing: %s\n", err)
	}

	if len(resp.Concepts.Concepts) != 8 {
		t.Errorf("expected len 8, but got len %d\n",
			len(resp.Concepts.Concepts))
	}
	if resp.Concepts.Concepts[0].Text != "President of the United States" {
		t.Errorf("expected \"prez of us\", but got \"%s\"\n",
			resp.Concepts.Concepts[0].Text)
	}
}
