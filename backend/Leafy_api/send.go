package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var base_url = "127.0.0.1:8080/"

func SendRequest(method string, url string, body string, authtoken string) {
	client := http.Client{}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		fmt.Println(err, "=")
		return
	}
	req, err := http.NewRequest(method, base_url+url, &buf)
	req.Header.Add("Authorization", "Bearer "+authtoken)
	if err != nil {
		fmt.Println(err, "==")
		return
	}
	_, err = client.Do(req)

	if err != nil {
		fmt.Println(err, "===")
		return
	}
}
