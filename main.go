package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type member struct {
	id          string
	Name        string `json:"name"`
	Memo        string `json:"memo"`
	Gender      string `json:"gender"`
	Birthday    string `json:"birthday"`
	Phone       string `json:"phone"`
	Cardid      string `json:"cardid"`
	Point       int    `json:"point"`
	CheckinShop string `json:"checkin_shop"`
	ShopsList   []string
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

func main() {
	start := time.Now()
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	nameURL := "https://namegen.jp/?sex=female&country=japan&middlename_cond=fukumu&middlename_rarity_cond=ika&lastname_cond=fukumu&lastname_rarity_cond=ika&lastname_type=name&firstname=&firstname_cond=fukumu&firstname_rarity_cond=ika&firstname_type=name"
	token, shopid := login()

	nameResChan := make(chan string, 10)
	nameChan := make(chan string, 60)
	memberChan := make(chan member, 60)
	memberIDChan := make(chan string, 100)
	wg.Add(1)
	for i := 0; i < 2; i++ {
		go fetech(nameURL, nameResChan)
		go parseName(<-nameResChan, nameChan, &wg)
	}
	// for {
	// 	// create(token,memberChan)
	// 	fmt.Println(<-memberChan)
	// }
	go func() {
		for {
			select {
			case name := <-nameChan:
				go fakeMemberGenerator(name, shopid, memberChan)
			case member := <-memberChan:
				go create(token, member, memberIDChan, &wg)
			case id := <-memberIDChan:
				go func() {
					for i := 0; i < rand.Intn(10); i++ {
						changePoint(id, token)
					}
				}()
				wg.Done()
			}
		}
	}()

	wg.Done()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}
