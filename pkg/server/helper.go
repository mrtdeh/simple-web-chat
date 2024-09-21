package server

import (
	"crypto/rand"
	"fmt"
	"time"
)

func tokenGenerator() string {
	b := make([]byte, 10)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// sadasdasd
func checkToken(token string) (*TokenData, error) {
	tokenData, err := tm.Get(token)
	if err != nil {
		return nil, err
	}
	// Check token expiration
	if tokenData.ExpireTime.Before(time.Now()) {
		tm.Delete(token)
		return nil, fmt.Errorf("token is expired")
	}
	return tokenData, nil
}
