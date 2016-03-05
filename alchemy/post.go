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
	response, err := Request(BuildRequest("Keywords", data))
	if err != nil {
		return err
	}
	err = ConvertResponseXML(response, &result)
	if err != nil {
		return err
	}

	t.Keywords = result.Keywords.Keywords
	return nil
}

func GetTaxonomy(data string, t *Taxonomys) error {
	result := TaxonomyResult{}
	response, err := Request(BuildRequest("Taxonomy", data))
	if err != nil {
		return err
	}
	err = ConvertResponseXML(response, &result)
	if err != nil {
		return err
	}
	fmt.Println("Result:", result)

	t.Taxonomys = result.Taxonomys.Taxonomys
	return nil
}

/*
func GetEntities(data string, e *Entities) error {
	result := EntityResult{}
	_, err := Request(BuildRequest("Entity", data), &result)
	if err != nil {
		return err
	}

	e.Entities = result.Entities.Entities
	return nil
}
*/

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
// Return raw response
func Request(url string) (*http.Response, error) {
	// build request
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("error creating request:", err)
		return nil, err
	}

	// send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error making request:", err)
		return nil, err
	}

	return response, nil
}

func ConvertResponseXML(response *http.Response, v interface{}) error {
	// convert to the type we are after
	err := ToXML(response.Body, v)
	if err != nil {
		fmt.Println("Conversion Failed: ", err)
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
// gateway http://gateway-a.watsonplatform.net/calls/text/TextGetRanked
