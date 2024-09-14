package server

import (
	"crypto/rand"
	"fmt"
)

func tokenGenerator() string {
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
