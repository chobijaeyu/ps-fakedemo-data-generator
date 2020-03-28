package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func randomDate(start, end int) string {
	min := time.Date(start, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(end, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02")
}

func randomDateTime(start, end int) string {
	min := time.Date(start, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(end, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format(time.RFC3339)
}

func randomPhone() string {
	return "0" + strconv.Itoa(7099999999-rand.Intn(99999999))
}

func randomCardid() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

func randmoGener() string {
	if randBool() {
		return "M"
	}
	return "F"
}

func randBool() bool {
	return rand.Int()%2 == 0
}
