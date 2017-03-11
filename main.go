package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"flag"
)

var (
	text = "Soy Pepe"
	language = "en"
	baseUrl = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&dt=t"
)

func init() {
	flag.StringVar(&language, "l", "en", "Target language")
	flag.StringVar(&text, "t", "Hello", "The text you want to translate")
	flag.Parse()
}

func main() {
	msg := "Let's Go!"

	fmt.Println(msg)

	fmt.Println("Text:", text)

	baseUrl += "&tl=" + language
	baseUrl += fmt.Sprintf("&q=%s", url.QueryEscape(text))

	resp, err := http.Get(baseUrl)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Error: %s", resp.Status)
		return
	}

	// data := body

	fmt.Printf("Translation: %s", body)
}
