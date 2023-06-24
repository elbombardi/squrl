package main

import (
	"math/rand"
	"time"
)

var (
	alphanums = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func generateRandomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = alphanums[r.Intn(len(alphanums))]
	}

	return string(b)
}

func main() {
	hash := map[string]bool{}
	collisions := 0
	for {
		randomString := generateRandomString(6)
		_, ok := hash[randomString]
		if ok {
			collisions++
		}
		if len(hash) >= 1000000 {
			break
		}
		hash[randomString] = true

	}
	println(collisions, "Collisions!", len(hash))

}
