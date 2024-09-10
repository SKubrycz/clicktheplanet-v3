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

	router.HandleFunc("/", s.handleGetHome)
	router.HandleFunc("/login", s.handleLogin)
	router.HandleFunc("/game", checkAuth(s.handleGetGame))
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

func (s *Server) handleGetHome(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{
		Title: "Home page",
		Desc:  "Welcome to the home page",
	}

	if r.Method == "GET" {
		jsonWelcome, err := json.Marshal(welcome)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonWelcome)
	} else {
		writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func (s *Server) handleGetGame(w http.ResponseWriter, r *http.Request) {
	// Here the user game state data is going to be fetched from the database

	welcome_game := Welcome{
		Title: "Hello to /game!",
		Desc:  "Here you can play since you're logged in!",
	}

	writeJSON(w, http.StatusOK, welcome_game)
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

type credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New from /login")
	if r.Method == "POST" {
		defer r.Body.Close()
		var credentials credentials
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

func writeJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}
