package analyzer

import (
	"github.com/opinionated/analyzer-core/alchemy"
)

// stores article data to be analyzed
type Analyzable struct {
	FileName  string
	Taxonomys alchemy.Taxonomys
	Keywords  alchemy.Keywords
}
