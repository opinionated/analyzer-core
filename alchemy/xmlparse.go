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

// create a new file and write the data structure into the file
func MarshalToFile(name string, v interface{}) error {

	body, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Println("oh nose, could not marshal:", err)
		return err
	}

	// TODO: find out why this number is set for file write mode
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("oh nose, could not create file:", err)
		return err
	}
	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		fmt.Println("oh nose, could not write file:", err)
		return err
	}

	return nil
}
