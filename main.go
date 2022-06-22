package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Url struct {
	XMLName xml.Name `xml:"url"`
	Url     []News   `xml:"news:news"`
	Name    string
}

type News struct {
	XMLName xml.Name `xml:"news:news"`
	Title   string   `xml:"news:title"`
	Date    string   `xml:"news:publication_date"`
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s Url
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]Url)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx := range n.Publications {
			news_map[n.Titles[idx]] = Url{n.Publications[idx], n.Locations[idx]}
		}
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var url Url
	xml.Unmarshal(byteValue, &url)

	p := Url{Name: "Amazing News Aggregator", Url: Date}
	t, _ := template.ParseFiles("newsaggtemplate.html")
	t.Execute(w, p)

	fmt.Println(news_map)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
