package analyzer

import (
//	"github.com/opinionated/analyzer-core/alchemy"
)

// stores article data to be analyzed
type Analyzable struct {
	Name     string
	FileName string
	//	Taxonomys alchemy.Taxonomys
	//	Keywords  alchemy.Keywords
	Score float64
	index int
}

func BuildAnalyzable() Analyzable {
	return Analyzable{"", "", 0.0, 0}
}
