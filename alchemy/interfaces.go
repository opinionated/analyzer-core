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
