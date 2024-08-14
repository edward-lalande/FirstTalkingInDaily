package main

import (
	"math/rand"
	"time"
)

func GetRandomName(name []string) string {
	rand.Seed(time.Now().UnixNano())
	nameChoose := string(name[rand.Int()%len(name)])

	if nameChoose == "\n" || nameChoose == "" {
		GetRandomName(name)
	}

	return nameChoose
}
