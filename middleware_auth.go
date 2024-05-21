package main

import (
	"fmt"
	"net/http"

	"github.com/sonclark/gorss/internal/auth"
	"github.com/sonclark/gorss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, error := cfg.DB.GetUserByAPKey(r.Context(), apiKey)
		if error != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", error))
			return
		}

		handler(w, r, user)
	}
}
