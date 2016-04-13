package relationDB

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func finishTest() {
	clear()
	Close()
}

type Relation struct {
	Text      string `json:"Text"`
	Relevance float32
}

func TestMultiInsert(t *testing.T) {
	assert.Nil(t, Open("http://localhost:7474/"))
	//defer finishTest()
	assert.Nil(t, clear())

	var _ = []string{
		"z", "1.0",
	}
	var relations = []Relation{
		{"x", 1.0},
		{"y", 1.0},
	}

	assert.Nil(t, Store("n"))
	assert.Nil(t, InsertRelations("n", "one", relations))
}

func TestInsert(t *testing.T) {
	// hits get by uuid and insert
	assert.Nil(t, Open("http://localhost:7474/"))
	defer finishTest()

	var tests = []struct {
		in  string
		err string
	}{
		{"hey", ""},
		{"hey", "uuid not unique"},
		{"hey hey", ""},
	}

	for i := range tests {
		err := Store(tests[i].in)
		errStr := tests[i].err
		if err != nil {
			assert.Equal(t, errStr, err.Error())
		}
	}
}
