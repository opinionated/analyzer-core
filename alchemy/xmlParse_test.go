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
