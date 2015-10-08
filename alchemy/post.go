package alchemy

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func GetKeywords(data string, t *Keywords) error {
	result := KeywordsResult{}
	err := Request(BuildRequest("Keywords", data), &result)
	if err != nil {
		return err
	}

	t.Keywords = result.Keywords.Keywords
	return nil
}

func GetTaxonomy(data string, t *Taxonomys) error {
	result := TaxonomyResult{}
	err := Request(BuildRequest("Taxonomy", data), &result)
	if err != nil {
		return err
	}

	t.Taxonomys = result.Taxonomys.Taxonomys
	return nil
}

// target is what we are after ie Keywords or Taxonomy
// data is article body
// you can append on params with "&param=value"
func BuildRequest(target, data string) string {
	// build default header values
	values := url.Values{}
	values.Set("apikey", "39995101e65858870797a627e548b1522f5c74a8")
	values.Add("text", data)

	// append target to the end of the default alchemy url
	ret := "http://gateway-a.watsonplatform.net/calls/text/TextGetRanked" + target
	ret = ret + "?" + values.Encode()

	return ret
}

// general request format
// create url using the BuildRequest method
// pass the target struct by reference in v
func Request(url string, v interface{}) error {
	// build request
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("error creating request:", err)
		return err

	}

	// send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error making request:", err)
		return err
	}
	defer response.Body.Close()

	// convert to the type we are after
	err = ToXML(response.Body, v)
	if err != nil {
		fmt.Println("wouldn't let it convert:", err)
		return err
	}

	return nil
}

// used largely for testing, parses an article from a file
func ParseArticle(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += " " + scanner.Text()
	}
	return lines, nil
}

// key 39995101e65858870797a627e548b1522f5c74a8
// curl --data "apikey=39995101e65858870797a627e548b1522f5c74a8&text=hello%20my%20name%20is%20test" http://access.alchemyapi.com/calls/text/TextGetRankedKeywords
// http://www.alchemyapi.com/api/keyword/textc.html
