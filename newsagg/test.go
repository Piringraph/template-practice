package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}
