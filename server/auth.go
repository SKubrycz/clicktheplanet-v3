package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func setSessionID(handlerFunc http.HandlerFunc, sessions []Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionClientCookie, err := r.Cookie("session")
		if err != nil {
			fmt.Println("Cookie doesn't exist")

			b := make([]byte, 32)
			if _, err := io.ReadFull(rand.Reader, b); err != nil {
				fmt.Println("Error session id reader")
			}

			value := base64.URLEncoding.EncodeToString(b)
			expires := time.Now().Add(time.Minute * 15)

			sessions = append(sessions, Session{
				Value:   value,
				Expires: expires,
			})
			fmt.Println(sessions)

			sessionCookie := http.Cookie{
				Name:     "session",
				Value:    value,
				SameSite: http.SameSiteStrictMode,
			}

			http.SetCookie(w, &sessionCookie)
			fmt.Println("session cookie set")
		}

		if err == nil {
			fmt.Println(sessionClientCookie.Value)

			exists := false
			index := -1
			for i, s := range sessions {
				if s.Value == sessionClientCookie.Value {
					exists = true
					index = i
					break
				}
			}

			if !exists {
				fmt.Println("SesssionId value from client doesn't exist on the server")
			} else {
				// Check for expiry
				currentTime := time.Now()
				verifyExpiry := currentTime.Compare(sessions[index].Expires)
				if verifyExpiry == -1 {
					fmt.Println("Not expired yet")
					fmt.Println(currentTime)
					fmt.Println(sessions[index].Expires)
				}
			}
		}

		handlerFunc(w, r)
	}
}

func assignJWT(w http.ResponseWriter, userId uint64) {
	fmt.Println("Assigning JWT...")
	REFRESH_MAX_AGE := 60 * 60 * 24 * 31
	ACCESS_MAX_AGE := 60 * 30
	refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")
	accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	refreshClaims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Duration(REFRESH_MAX_AGE) * time.Second),
		"userId":    userId,
	}
	accessClaims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Duration(ACCESS_MAX_AGE) * time.Second),
		"userId":    userId,
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

func verifyJWT(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("jwt verification unsuccessful: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

}

type Message struct {
	Message string `json:"message"`
}

func checkAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("checkAuth start...")

		refreshSecret := os.Getenv("REFRESH_TOKEN_SECRET")
		accessSecret := os.Getenv("ACCESS_TOKEN_SECRET")

		refreshToken, err := r.Cookie("refresh_token")
		if err != nil {
			fmt.Println("Couldn't find a cookie")
			error := Message{
				Message: "Not authorized",
			}
			writeJSON(w, http.StatusUnauthorized, error)
			return
		}
		refreshVerify, err := verifyJWT(refreshToken.Value, refreshSecret)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("refreshToken verified", refreshVerify.Claims)
		accessToken, err := r.Cookie("access_token")
		fmt.Println(accessToken)
		var accessTokenValue string
		if err != nil {
			if refreshVerify.Valid {
				fmt.Println("Refreshing access_token")
				ACCESS_MAX_AGE := 60 * 30
				accessClaims := &jwt.MapClaims{
					"expiresAt": time.Now().Add(time.Duration(ACCESS_MAX_AGE) * time.Second),
					"userId":    int(refreshVerify.Claims.(jwt.MapClaims)["userId"].(float64)),
				}
				accessTokenSigned, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(accessSecret))
				if err != nil {
					writeJSON(w, http.StatusInternalServerError, "Could not login")
					return
				}
				accessTokenValue = accessTokenSigned
				accessTokenCookie := http.Cookie{
					Name:     "access_token",
					Value:    accessTokenValue,
					MaxAge:   ACCESS_MAX_AGE,
					HttpOnly: true,
					SameSite: http.SameSiteStrictMode,
				}

				http.SetCookie(w, &accessTokenCookie)
				fmt.Println("accessTokenCookie set")
			} else {
				fmt.Println("Couldn't find a cookie")
				errMessage := Message{
					Message: "Not authorized",
				}
				writeJSON(w, http.StatusUnauthorized, errMessage)
				return
			}
		}
		if refreshVerify.Valid {
			var accessVerify *jwt.Token
			if accessToken == nil {
				accessVerify, err = verifyJWT(accessTokenValue, accessSecret)
			} else {
				accessVerify, err = verifyJWT(accessToken.Value, accessSecret)
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("accessToken verified", accessVerify.Claims)

			claims := accessVerify.Claims.(jwt.MapClaims)
			const userId UserId = "userid" // key of type UserId - it has to stay here

			ctx := context.WithValue(r.Context(), userId, int(claims["userId"].(float64)))
			fmt.Println("Context for userId has been created")
			handlerFunc(w, r.WithContext(ctx))
		} else {
			writeJSON(w, http.StatusUnauthorized, "Not authorized")
			return
		}
	}
}
