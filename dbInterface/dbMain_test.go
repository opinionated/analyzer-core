package relationDB

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func finishTest() {
	clear()
	Close()
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
