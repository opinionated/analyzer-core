package alchemy_test

import (
	"github.com/opinionated/analyzer-core/alchemy"
	"testing"
)

func ParseKimDavis() string {
	articleBody, err := alchemy.ParseArticle("test/kimDavisRelease.txt")
	if err != nil {
		panic(err) // shouldn't ever happen
	}
	return articleBody
}

func TestFetchKeywords(t *testing.T) {
	articleBody := ParseKimDavis()
	processed, err := alchemy.RequestKeywords(articleBody)
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
	processed, err := alchemy.RequestTaxonomy(articleBody)
	return
	if err != nil {
		t.Errorf("could not send request:", err)
	}

	// test that we got what we expected
	if processed.Status != "OK" {
		t.Errorf("expected status OK")
	}

	if processed.Taxonomys.Taxonomys[0].Label != "/society/social institution/marriage" {
		t.Errorf("expected top word to be marriage licenses but did not get")
	}
}
