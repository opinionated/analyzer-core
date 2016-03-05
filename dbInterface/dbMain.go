import (
	"analyzer/structures"
)

// keyQuery and taxQuery store the same information, differentiation is so search-field can be implicit
type keyQuery struct {
	name      string
	sentiment string
	threshold float64
}

type taxQuery struct {
	name      string
	sentiment string
	threshold float64
}

struct articleInfo struct {
	id         int
	author     string
	body       string //filename
	keywords   []Keyword
	taxonomies []Taxonomy
}
