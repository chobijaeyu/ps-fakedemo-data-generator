package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func create(token string, member member, idCh chan string, wg *sync.WaitGroup) {
	memberURL := "http://127.0.0.1:8000/pointcard/customers/"
	var bearer = "Bearer " + token
	requestBody, err := json.Marshal(member)
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", memberURL, bytes.NewBuffer(requestBody))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("create err", err)
		return
	}

	if res.StatusCode != 201 {
		fmt.Println("http request fail", res.StatusCode)
		wg.Done()
		return
	}

	defer res.Body.Close()
	var memberResSt map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&memberResSt); err != nil {
		log.Fatalln(err)
	}
	idCh <- fmt.Sprintf("%v", memberResSt[`id`])
}
