package go_percent_encoding_test

import (
	"fmt"
	"net/url"

	go_percent_encoding "github.com/leoleaf/go-percent-encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func ExampleUtf8ToOther() {
	u, err := url.Parse("http://www.example.com/register")
	if err != nil {
		fmt.Println(err)
		return
	}

	// add query string
	query := u.Query()
	query.Add("age", "19")
	query.Add("name", "张三")

	u.RawQuery = query.Encode()

	fmt.Println(u.String())

	// if the web server encoding is not utf-8, you need to encode the query string.
	// assume the web server encoding is GBK
	enc := simplifiedchinese.GBK.NewEncoder()
	targetUrl, err := go_percent_encoding.Utf8ToOther(u.String(), enc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(targetUrl)
	// Output:
	// http://www.example.com/register?age=19&name=%E5%BC%A0%E4%B8%89
	// http://www.example.com/register?age=19&name=%D5%C5%C8%FD
}

func ExampleOther2Utf8() {
	// assume the web server encoding is GBK
	// get the url from the web server
	urlFromServer := "http://www.example.com/register?age=19&name=%D5%C5%C8%FD"
	u, err := url.Parse(urlFromServer)
	if err != nil {
		fmt.Println(err)
		return
	}

	dec := simplifiedchinese.GBK.NewDecoder()
	u.RawQuery, err = go_percent_encoding.Other2Utf8(u.RawQuery, dec)
	if err != nil {
		fmt.Println(err)
		return
	}

	query := u.Query()

	fmt.Println(query.Get("age"))
	fmt.Println(query.Get("name"))
	// Output:
	// 19
	// 张三
}
