package alchemy_test

import (
	"fmt"
	"github.com/opinionated/analyzer-core/alchemy"
	"os"
	"testing"
)

// test by visual inspection on this one
func TestToXML(t *testing.T) {
	file, err := os.Open("test/kimDavisRelease_AlcKeywords.xml")
	if err != nil {
		t.Errorf("oh nose, err opening file:", err)
	}

	defer file.Close()

	keywords := alchemy.KeywordsResult{}
	err = alchemy.ToXML(file, &keywords)
	if err != nil {
		t.Errorf("oh nose, err reading keywords:", err)
	}
	fmt.Println("keywords are:", keywords)
}

// read out of one file then try to write to other
// then read the data back out of the file to check if it is OK
func TestXmlToFile(t *testing.T) {
	// open actual file
	file, err := os.Open("test/kimDavisRelease_AlcKeywords.xml")
	if err != nil {
		t.Errorf("oh nose, err opening file:", err)
	}
	defer file.Close()

	// read from file
	keywords := alchemy.KeywordsResult{}
	err = alchemy.ToXML(file, &keywords)
	if err != nil {
		t.Errorf("oh nose, err reading keywords:", err)
	}

	// write actual results to a test file
	alchemy.MarshalToFile("test/tmp.txt", keywords)
	if err != nil {
		t.Errorf("oh nose, err writing to file:", err)
	}

	// open the test file
	tmp, err := os.Open("test/tmp.txt")
	if err != nil {
		t.Errorf("oh nose, err opening file:", err)
	}
	defer tmp.Close()

	// read data back out of test file
	test_keys := alchemy.KeywordsResult{}
	err = alchemy.ToXML(tmp, &test_keys)
	if err != nil {
		t.Errorf("oh nose, err reading keywords:", err)
	}

	// check if data got changed
	if len(keywords.Keywords.Keywords) != len(test_keys.Keywords.Keywords) {
		t.Errorf("expected size:", len(keywords.Keywords.Keywords), "got:", len(test_keys.Keywords.Keywords))
	}

	for i := range keywords.Keywords.Keywords {
		actual_v := test_keys.Keywords.Keywords[i]
		expected_v := keywords.Keywords.Keywords[i]
		if actual_v.Text != expected_v.Text || actual_v.Relevance != expected_v.Relevance {
			t.Errorf("error, files are not the same!")
		}
	}
}
