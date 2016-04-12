package dbInterface

import (
	"github.com/opinionated/analyzer-core/analyzer"
	"testing"
)

func TestDBSimple() {

	dbInterface.Init()
	dbInterface.Store(&ArticleInfo{id.String(), "Multiple entries in database for this ID!!"})

}
