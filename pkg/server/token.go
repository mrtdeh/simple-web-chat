package server

import (
	"fmt"
	"time"
)

func (tm *TokenManager) Create(userId uint32, username string) string {
	tm.l.Lock()
	defer tm.l.Unlock()
	token := tokenGenerator()
	tm.tokens[token] = &TokenData{
		UserID:     userId,
		Username:   username,
		ExpireTime: time.Now().Add(time.Minute * 5),
	}
	return token
}

func (tm *TokenManager) Get(key string) (*TokenData, error) {
	tm.l.RLock()
	defer tm.l.RUnlock()

	tokenData, ok := tm.tokens[key]
	if !ok {
		return nil, fmt.Errorf("token is not found")
	}
	return tokenData, nil
}

func (tm *TokenManager) Delete(key string) {
	tm.l.Lock()
	defer tm.l.Unlock()
	delete(tm.tokens, key)
}
