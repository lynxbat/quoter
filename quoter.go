package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	URL = "https://andruxnet-random-famous-quotes.p.mashape.com/cat=all"
)

type Quote struct {
	Quote    string `json:quote`
	Author   string `json:author`
	Category string `json:category`
}

func RandomQuote(w http.ResponseWriter, req *http.Request) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	req.Header.Add("X-Mashape-Key", "5uKIAGiQVrmshQxvJDrANVT2Lw9bp17qnlCjsnAj2KGJQ9Ht3a")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	q := &Quote{}
	json.Unmarshal(b, q)
	output := fmt.Sprintf("_`”%s”`_ - *%s*", q.Quote, q.Author)
	io.WriteString(w, output)
}

func main() {
	fmt.Println("The Intel SDI-X Quoter API is started")
	http.HandleFunc("/randomquote", RandomQuote)
	http.ListenAndServe(":8080", nil)
}
