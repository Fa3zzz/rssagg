package main

import (
	"fmt"
	"net/http"

	"github.com/Fa3zzz/rssagg/internal/database"
	"github.com/Fa3zzz/rssagg/internal/database/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middleWareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Coudln't get user: %v", err))
			return
		}

		handler(w, r, user)

	}
}
