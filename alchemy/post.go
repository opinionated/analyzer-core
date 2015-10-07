package main
import (
	"fmt"
    "bufio"
    "os"
    "os/exec"
	// "net/http"
	// "net/url"
	// "bytes"
	"io/ioutil"
	"strings"

	)
func parseArticle(filepath string) (string, error) {
    file, err := os.Open(filepath)
    if(err != nil){
        return "",err
    }
    defer file.Close()

    var lines string
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        lines += " " + scanner.Text()
    }
    lines = strings.Replace(lines, " ", "%20", -1)
    // fmt.Println(lines)
    return lines,nil
}
func keywordAPI(filepath string, sentiment bool) (string, error){
    cmd := "curl"
    args := []string{"--data", "apikey=39995101e65858870797a627e548b1522f5c74a8","http://access.alchemyapi.com/calls/text/TextGetRankedKeywords"}
    
    lines, err := parseArticle(filepath)
    if(err != nil){
        return "",err
    }
    args[1] += "&text=" + lines

    if sentiment {
        args[1] += "&sentiment=1"
    }

    // fmt.Println(args[1])
    out, err := exec.Command(cmd, args...).Output();
    if err != nil {
        os.Exit(1)
        return "",err
    }
    return string(out),nil
           
}
func taxonomyAPI(filepath string) (string, error){
    cmd := "curl"
    args := []string{"--data", "apikey=39995101e65858870797a627e548b1522f5c74a8","http://gateway-a.watsonplatform.net/calls/text/TextGetRankedTaxonomy"}
    
    lines, err := parseArticle(filepath)
    if(err != nil){
        return "",err
    }
    args[1] += "&text=" + lines

    fmt.Println(args[1])
    out, err := exec.Command(cmd, args...).Output();
    if err != nil {
        os.Exit(1)
        return "",err
    }
    return string(out),nil
}


func main() {
    keyxmlpath := "testkeyword.xml"
    xml1, err := keywordAPI("./samplebody.txt",false)
    if(err != nil){
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
}

// key 39995101e65858870797a627e548b1522f5c74a8
// curl --data "apikey=39995101e65858870797a627e548b1522f5c74a8&text=hello%20my%20name%20is%20test" http://access.alchemyapi.com/calls/text/TextGetRankedKeywords
// http://www.alchemyapi.com/api/keyword/textc.html
