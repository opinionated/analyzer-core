package alchemy_test

import (
	"fmt"
	"github.com/opinionated/analyzer-core/alchemy"
	"os"
	"testing"
)

func TestKeywordsFromFile(t *testing.T) {
	file, err := os.Open("test/kimDavisRelease_AlcKeywords.xml")
	if err != nil {
		t.Errorf("oh nose, err opening file:", err)
	}

	defer file.Close()

	keywords, err := alchemy.KeywordsFromFile(file)
	if err != nil {
		t.Errorf("oh nose, err reading keywords:", err)
	}
	fmt.Println("keywords are:", keywords)
}

func TestXmlToFile(t *testing.T) {
	file, err := os.Open("test/kimDavisRelease_AlcKeywords.xml")
	if err != nil {
		t.Errorf("oh nose, err opening file:", err)
	}

	defer file.Close()

	keywords, err := alchemy.KeywordsFromFile(file)
	if err != nil {
		t.Errorf("oh nose, err reading keywords:", err)
	}

	alchemy.MarshalToFile("test/tmp.txt", keywords)
	if err != nil {
		t.Errorf("oh nose, err writing to file:", err)
	}

	tmp, err := os.Open("test/tmp.txt")

	if err != nil {
		t.Errorf("oh nose, err opening file:", err)
	}

	defer tmp.Close()
	test_keys, err := alchemy.KeywordsFromFile(tmp)
	if err != nil {
		t.Errorf("oh nose, err reading keywords:", err)
	}
	for i := range keywords.Keywords.Keywords {
		actual_v := test_keys.Keywords.Keywords[i]
		expected_v := keywords.Keywords.Keywords[i]
		if actual_v.Text != expected_v.Text || actual_v.Relevance != expected_v.Relevance {
			t.Errorf("error, files are not the same!")
		}
	}
}
