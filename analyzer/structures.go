package analyzer

import (
	"github.com/opinionated/analyzer-core/alchemy"
)

// stores article data to be analyzed
type Analyzable struct {
	Name      string
	FileName  string
	Taxonomys alchemy.Taxonomys
	Keywords  alchemy.Keywords
	score     float64
}

func BuildAnalyzable() Analyzable {
	return Analyzable{"", "", alchemy.Taxonomys{}, alchemy.Keywords{}, 0.0}
}
