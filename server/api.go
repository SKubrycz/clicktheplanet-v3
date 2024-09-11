package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.Use(setCors)

	router.HandleFunc("/", s.handleHome)
	router.HandleFunc("/login", s.handleLogin)
	router.HandleFunc("/register", s.handleRegister)
	router.HandleFunc("/game", checkAuth(s.handleGame))
	router.HandleFunc("/ws_game", s.handleGetWsGame)

	http.ListenAndServe(s.addr, router)
}

func setCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authentication, Upgrade, Connection, Host, Accept, Cookie, Set-Cookie")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
}

type Welcome struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{
		Title: "Home page",
		Desc:  "Welcome to the home page",
	}

	if r.Method == "GET" {
		writeJSON(w, http.StatusOK, welcome)
		return
	} else {
		writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func (s *Server) handleGame(w http.ResponseWriter, r *http.Request) {
	// Here the user game state data is going to be fetched from the database

	if r.Method == "GET" {
		welcome_game := Welcome{
			Title: "Hello to /game!",
			Desc:  "Here you can play since you're logged in!",
		}

		writeJSON(w, http.StatusOK, welcome_game)
		return
	} else {
		writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func (s *Server) handleGetWsGame(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// origin := r.Header.Get("Origin")
			// return origin == "http://localhost:3000"
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		fmt.Println(string(p))
		if err != nil {
			log.Println(err)
			return
		}
		if bytes.Equal(p, []byte("Message to server")) {
			p = []byte("Message to client")
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		var credentials Credentials
		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			fmt.Println(err)
			return
		}

		//... something something
		//check verify the input
		//if ok then give jwt:

		fmt.Println(credentials)
		if credentials.Login == "frank" {
			assignJWT(w, r)
		}
	}
}

type RegisterCredentials struct {
	Login         string `json:"login"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	AgainPassword string `json:"againPassword"`
}

func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		var registerCredentials RegisterCredentials
		if err := json.NewDecoder(r.Body).Decode(&registerCredentials); err != nil {
			fmt.Println(err)
			return
		}

		// verify users input
		// check if the user already exists in the database
		// if ok:
		// hash the password
		// save to the database

	} else {
		writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}
