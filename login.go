package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func login() (token, shopid string) {
	loginURL := "http://127.0.0.1:8000/api-token-auth/"
	requestBody, err := json.Marshal(map[string]string{"username": "demoshop", "password": "demo0000"})
	res, err := http.Post(loginURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var authResSt map[string]interface{}
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(string(body))
	if err := json.NewDecoder(res.Body).Decode(&authResSt); err != nil {
		log.Fatalln(err)
	}
	return fmt.Sprintf("%v", authResSt[`token`]), fmt.Sprintf("%v", authResSt[`shop`])
}
