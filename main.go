package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	source = rand.NewSource(time.Now().UnixNano())
	rng    = rand.New(source)
)

func main() {
	urls := make(map[string]string)

	for {
		fmt.Print("Enter your url : ")
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')

		urls[generateShortKey()] = line

		for k, v := range urls {
			fmt.Printf("%s =====> %s", k, v)
		}
	}

}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 9
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rng.Intn(len(charset))]
	}
	return string(shortKey)
}
