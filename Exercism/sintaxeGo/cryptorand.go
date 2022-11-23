package main

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
)

func hash(b []byte) string {
	h := sha1.New()
	h.Write(b)
	sum := h.Sum(nil)
	armored := fmt.Sprintf("%x", sum)
	return armored
}

func randomString() (string, error) {
	b := make([]byte, 10)
	_, err := rand.Read(b)

	if err != nil {
		fmt.Printf("erro: %v", err)
		return "", err
	}

	armored := hash(b)
	return armored, err
}

func main() {
	armored, _ := randomString()
	fmt.Printf("string randomica %s\r\n", armored)
}
