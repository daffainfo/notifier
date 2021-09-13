package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var apikey = "xxxxx"

func curl_line() {
	s := bufio.NewScanner(os.Stdin)
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	for s.Scan() {
		params := url.Values{}
		params.Add("message", s.Text())
		body := strings.NewReader(params.Encode())
		req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", body)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+apikey)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	}

}

func main() {
	curl_line()
}
