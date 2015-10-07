package alchemy

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// reads an alchemy file into its key words
func KeywordsFromFile(file *os.File) (KeywordsResult, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("oh nose, could not read file:", err)
		return KeywordsResult{}, err
	}

	// create struct to be read into
	keys := KeywordsResult{}

	err = xml.Unmarshal([]byte(data), &keys)
	if err != nil {
		fmt.Println("oh nose, err unmarshalling:", err)
		return KeywordsResult{}, err
	}
	return keys, nil
}
