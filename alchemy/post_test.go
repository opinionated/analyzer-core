package alchemy_test

import (
	"fmt"
	"github.com/opinionated/analyzer-core/alchemy"
	"testing"
)

func TestFetchKeywords(t *testing.T) {
	articleBody, err := alchemy.ParseArticle("test/kimDavisRelease.txt")
	if err != nil {
		t.Errorf("could not open body:", err)
	}

	processed, err := alchemy.RequestKeywords(articleBody)
	if err != nil {
		t.Errorf("could not send request:", err)
	}
	fmt.Println("result is:", processed)
}
