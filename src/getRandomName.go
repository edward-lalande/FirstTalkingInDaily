package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandomName(name []string) string {
	rand.Seed(time.Now().UnixNano())
	nameChoose := string(name[rand.Int()%len(name)])
	fmt.Printf("Random name: (%s)\n", nameChoose)
	if nameChoose == "\n" || nameChoose == "" {
		GetRandomName(name)
	}

	return nameChoose
}
