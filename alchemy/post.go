package alchemy

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

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

func RequestKeywords(data string) (KeywordsResult, error) {
	// build params
	values := url.Values{}
	values.Set("apikey", "39995101e65858870797a627e548b1522f5c74a8")
	values.Add("text", data)

	// build params into url
	target := "http://gateway-a.watsonplatform.net/calls/text/TextGetRankedKeywords"
	target = target + "?" + values.Encode()

	// build request
	request, err := http.NewRequest("POST", target, nil)
	if err != nil {
		fmt.Println("error creating request:", err)
		return KeywordsResult{}, err

	}

	// send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error making request:", err)
		return KeywordsResult{}, err
	}
	defer response.Body.Close()

	// read request results into KeywordsResults struct
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err reading body:", err)
		return KeywordsResult{}, err
	}

	keywords := KeywordsResult{}
	err = xml.Unmarshal(body, &keywords)

	return keywords, nil
}

func taxonomyAPI(filepath string) (string, error) {
	cmd := "curl"
	args := []string{"--data", "apikey=39995101e65858870797a627e548b1522f5c74a8", "http://gateway-a.watsonplatform.net/calls/text/TextGetRankedTaxonomy"}

	lines, err := ParseArticle(filepath)
	if err != nil {
		return "", err
	}
	args[1] += "&text=" + lines

	fmt.Println(args[1])
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		os.Exit(1)
		return "", err
	}
	return string(out), nil
}

func main() {
	/*	keyxmlpath := "testkeyword.xml"
		xml1, err := keywordAPI("./samplebody.txt", false)
		if err != nil {
			panic(err)
		}
		// fmt.Println(xml1)
		err = ioutil.WriteFile(keyxmlpath, []byte(xml1), 0644)
		if err != nil {
			panic(err)
		}
		fmt.Println("wrote keyword xml to " + keyxmlpath)

		// xml2, err := taxonomyAPI("./samplebody.txt")
		// if(err != nil){
		//     fmt.Println(err)
		// }
		// fmt.Println(xml2)
	*/
}

// key 39995101e65858870797a627e548b1522f5c74a8
// curl --data "apikey=39995101e65858870797a627e548b1522f5c74a8&text=hello%20my%20name%20is%20test" http://access.alchemyapi.com/calls/text/TextGetRankedKeywords
// http://www.alchemyapi.com/api/keyword/textc.html
