package alchemy

import (
	"encoding/xml"
)

// data structures that the alchemy results get parsed into

// structures for keywords
// TODO: include sentiment here
type Keyword struct {
	Relevance float32 `xml:"relevance"`
	Text      string  `xml:"text"`
}

type Keywords struct {
	Keywords []Keyword `xml:"keyword"`
}

type KeywordsResult struct {
	XMLName  xml.Name `xml:"results"`
	Status   string   `xml:"status"`
	Keywords Keywords `xml:"keywords"`
}

// structures for taxonomy
type Taxonomy struct {
	Label string  `xml:"label"`
	Score float32 `xml:"score"`
}

type Taxonomys struct {
	Taxonomys []Taxonomy `xml:"element"`
}

type TaxonomyResult struct {
	XMLName   xml.Name  `xml:"results"`
	Status    string    `xml:"status"`
	Taxonomys Taxonomys `xml:"taxonomy"`
}

type EntitySentiment struct {
	// for whatever reason entity sentiment is not working.
	Score float32 `xml:"score"`
}

type EntityDisambiguated struct {
	Name     string   `xml:"name"`
	Subtypes []string `xml:"subType"`
}

// structures for entities
type Entity struct {
	Type          string                `xml:"type"`
	Relevance     float32               `xml:"relevance"`
	Count         float32               `xml:"count"`
	Text          string                `xml:"text"`
	Disambiguated []EntityDisambiguated `xml:"disambiguated"`

	// sentiment wouldn't work and it doesn't seem like it's that important
	//Sentiment     EntitySentiment       `xml:"sentiment"`
}

type Entities struct {
	Entities []Entity `xml:"entity"`
}

type EntityResult struct {
	XMLName  xml.Name `xml:"results"`
	Status   string   `xml:"status"`
	Entities Entities `xml:"entities"`
}

// structures for concepts
type Concept struct {
	Text      string  `xml:"text"`
	Relevance float32 `xml:"relevance"`
}

type Concepts struct {
	Concepts []Concept `xml:"concept"`
}

type ConceptResult struct {
	XMLName  xml.Name `xml:"results"`
	Status   string   `xml:"status"`
	Concepts Concepts `xml:"concepts"`
}
