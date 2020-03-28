package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func changePoint(id string, token string) {
	memberURL := "http://127.0.0.1:8000/pointcard/customers/" + id + "/"
	phURL := "http://127.0.0.1:8000/pointcard/pointHistory/"

	var bearer = "Bearer " + token

	_createupdatetime := randomDateTime(2000, 2020)
	client := &http.Client{}

	changeData := map[string]interface{}{
		"point":      rand.Intn(999),
		"created_at": _createupdatetime,
		"updated_at": _createupdatetime}
	requestBody, err := json.Marshal(changeData)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest("PATCH", memberURL, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("create err", err)
		return
	}

	if res.StatusCode != 202 {
		fmt.Println("http request fail", res.StatusCode)
		return
	}

	defer res.Body.Close()
	var changeResSt map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&changeResSt); err != nil {
		log.Fatalln(err)
	}
	ac, _ := strconv.Atoi(fmt.Sprintf("%v", changeResSt[`afterChange`]))
	bc, _ := strconv.Atoi(fmt.Sprintf("%v", changeResSt[`beforeChange`]))
	_createupdatetime = randomDateTime(2000, 2020)
	changeHistoryData := map[string]interface{}{
		"afterChange":  ac,
		"beforeChange": bc,
		"variation":    ac - bc,
		"shop":         changeResSt["shop"],
		"customer":     changeResSt["customer"],
		"created_at":   _createupdatetime,
		"updated_at":   _createupdatetime}
	requestBody, err = json.Marshal(changeHistoryData)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest("POST", phURL, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("create err", err)
		return
	}

	if res.StatusCode != 202 {
		fmt.Println("http request fail", res.StatusCode)
		defer res.Body.Close()
		var memberResSt map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&memberResSt); err != nil {
			log.Fatalln(err)
		}
		fmt.Println(memberResSt)
		return
	}

	defer res.Body.Close()
	var authResSt map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&authResSt); err != nil {
		log.Fatalln(err)
	}

}
