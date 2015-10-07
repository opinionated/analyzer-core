package alchemy

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// reads an xml file containing alchemy keyword data
// reader == where to read from
// v == type to read as, pass it as a reference
func ToXML(in io.Reader, v interface{}) error {
	data, err := ioutil.ReadAll(in)
	if err != nil {
		fmt.Println("oh nose, could not read file:", err)
		return err
	}

	err = xml.Unmarshal([]byte(data), v)
	if err != nil {
		fmt.Println("data is:", data)
		fmt.Println("oh nose, err unmarshalling:", err)
		return err
	}

	return nil
}

// create a new file and write the data structure into the file
// TODO: think about making this be able to write more generically
func MarshalToFile(name string, v interface{}) error {

	body, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Println("oh nose, could not marshal:", err)
		return err
	}

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
