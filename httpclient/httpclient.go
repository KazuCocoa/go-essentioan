package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	uri := "https://httpbin.org/html"
	resp, err := http.Get(uri)

	if err != nil {
		log.Fatalf("http.Get() failed with '%s'\n", err)
	}

	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() failed with '%s'\n", err)
	}

	contentType := resp.Header.Get("Content-Type")
	fmt.Printf("http.Get() returned content of type '%s' and size %d bytes.\nStatus code: %d\n", contentType, len(d), resp.StatusCode)

	// getting page that doesn't exist return 404
	uri = "https://httpbin.org/page-doesnt-exist"
	resp, err = http.Get(uri)
	if err != nil {
		log.Fatalf("http.Get() failed with '%s'\n", err)
	}

	contentType = resp.Header.Get("Content-Type")
	fmt.Printf("\nhttp.Get() returned content of type '%s' and size %d bytes.\nStatus code: %d\n", contentType, len(d), resp.StatusCode)

	uri = "http://website.not-exists.as/index.html"
	resp, err = http.Get(uri)
	if err != nil {
		fmt.Printf("\nhttp.Get() failed with: '%s'\nresp: %v\n", err, resp)
	}

	// client := &http.Client{}
	// client.Timeout = time.Millisecond * 100

	// uri_1 := "https://httpbin.org/delay/3"
	// _, err_1 := client.Get(uri_1)
	// if err_1 != nil {
	// 	log.Fatalf("http.Get() failed with '%s'\n", err_1)
	// }

	client_2 := &http.Client{}
	client_2.Timeout = time.Second * 15

	uri_2 := "https://httpbin.org/post"
	body_2 := bytes.NewBufferString("text we send")
	resp_2, err_2 := client_2.Post(uri_2, "text/plain", body_2)
	if err_2 != nil {
		log.Fatalf("client.Post() failed with '%s'\n", err_2)
	}
	defer resp_2.Body.Close()
	d, err_3 := ioutil.ReadAll(resp_2.Body)
	if err != nil {
		log.Fatalf("http.Get() failed with '%s'\n", err_3)
	}
	fmt.Printf("http.Post() returned statuts code %d, truncated text:\n%s...\n", resp_2.StatusCode, string(d)[:93])
}
