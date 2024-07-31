package main

import (
	"fmt"
	"net/http"

	"github.com/RuxinZ/rssagg/internal/auth"
	"github.com/RuxinZ/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden,fmt.Sprintf("Auth error: %v", err) )
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound,fmt.Sprintf("Couldn't get user: %v", err) )
		}
		handler(w,r,user)
	}


}