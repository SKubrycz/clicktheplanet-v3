package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
)

func assignJWT(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Assigning JWT...")
	REFRESH_MAX_AGE := 1000 * 60 * 60 * 24 * 31
	ACCESS_MAX_AGE := 1000 * 60 * 15
	refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	refreshClaims := &jwt.MapClaims{
		"expiresAt": REFRESH_MAX_AGE,
		//"id": user.id,
	}
	accessClaims := &jwt.MapClaims{
		"expiresAt": ACCESS_MAX_AGE,
		//"id": user.id,
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(refreshSecret))
	if err != nil {
		http.Error(w, "Could not login", http.StatusInternalServerError)
		return
	}
	refreshTokenCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		MaxAge:   REFRESH_MAX_AGE,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(accessSecret))
	if err != nil {
		http.Error(w, "Could not login", http.StatusInternalServerError)
		return
	}
	accessTokenCookie := http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		MaxAge:   ACCESS_MAX_AGE,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &refreshTokenCookie)
	http.SetCookie(w, &accessTokenCookie)

	fmt.Println("JWT has been assigned")

	//w.WriteHeader(200)
}
