package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

func fetech(url string, resCh chan string) {
	fmt.Println("Fetch Url", url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("http get err", err)
		return
	}

	if res.StatusCode != 200 {
		fmt.Println("http request fail", res.StatusCode)
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read error", err)
		return
	}

	resCh <- string(body)
}

func parseName(body string, nameChan chan string, wg *sync.WaitGroup) {
	var re = regexp.MustCompile(
		`(?m)<td class="name">
      <a href=".*">(.*)<\/a>
      <a href=".*">(.*)<\/a>
    <\/td>`)

	rename := re.FindAllStringSubmatch(body, -1)
	i := 0
	for _, match := range rename {
		// fmt.Println(match[1], match[2], "found at index", i)
		nameChan <- match[1] + " " + match[2]
		i++
		fmt.Println(i)
		wg.Add(1)
	}
}
