package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func readapikey() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	data, err := ioutil.ReadFile(dirname + "/.config/notifier/key.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return string(data)
}

func showBanner() {
	fmt.Println(`
____    ___   ______  ____  _____  ____    ___  ____  
|    \  /   \ |      Tl    j|     |l    j  /  _]|    \ 
|  _  YY     Y|      | |  T |   __j |  T  /  [_ |  D  )
|  |  ||  O  |l_j  l_j |  | |  l_   |  | Y    _]|    / 
|  |  ||     |  |  |   |  | |   _]  |  | |   [_ |    \ 
|  |  |l     !  |  |   j  l |  T    j  l |     T|  .  Y
l__j__j \___/   l__j  |____jl__j   |____jl_____jl__j\_j

Name: Muhammad Daffa
Version: 0.0.1`)
}

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
		req.Header.Set("Authorization", "Bearer "+readapikey())
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	}
}

func main() {
	showBanner()
	curl_line()
}
