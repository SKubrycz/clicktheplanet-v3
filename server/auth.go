package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

func assignJWT(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Assigning JWT...")
	REFRESH_MAX_AGE := 60 * 60 * 24 * 31
	ACCESS_MAX_AGE := 60 * 15
	refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	refreshClaims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Duration(REFRESH_MAX_AGE) * time.Second),
		//"id": user.id,
	}
	accessClaims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Duration(ACCESS_MAX_AGE) * time.Second),
		//"id": user.id,
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(refreshSecret))
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, "Could not login")
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
		writeJSON(w, http.StatusInternalServerError, "Could not login")
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

type Message struct {
	Message string `json:"message"`
}

func checkAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("checkAuth start...")

		cookie, err := r.Cookie("access_token")
		if err != nil {
			fmt.Println("Couldn't find a cookie")
			error := Message{
				Message: "Not authorized",
			}
			writeJSON(w, http.StatusUnauthorized, error)
			return
		}

		fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)

		handlerFunc(w, r)
	}
}
