package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/manucorporat/try"
	"github.com/valyala/fastjson"
)

func main() {
	apikey := os.Args[1] // api key (get it here https://saucenao.com/user.php?page=search-api)
	url := os.Args[2]    // url to find the sauce of

	client := &http.Client{}
	var data = strings.NewReader(``)
	var url1 = "https://saucenao.com/search.php?output_type=2&testmode=1&numres=5&url=" + url + "&api_key=" + apikey

	req, _ := http.NewRequest("GET", url1, data)
	res, _ := client.Do(req)

	var p fastjson.Parser

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()

	parsed, _ := p.Parse(newStr)

	for _, element := range parsed.GetArray("results") {
		ext_urls := element.GetObject("data").Get("ext_urls")

		try.This(func() {
			str := ext_urls.GetArray()[0]
			ass := str.String()
			fmt.Println(cleanStr(ass))
		}).Catch(func(e try.E) {
			// ignore the error LOL
		})
	}
}

func cleanStr(str string) string {
	var s string
	s = strings.Replace(str, "\\", "SLASH", -1)
	s = strings.Replace(s, "u0026", "AND", -1)

	s = strings.Replace(s, "SLASH", "", -1)
	s = strings.Replace(s, "AND", "&", -1)
	return s
}
