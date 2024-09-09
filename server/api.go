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
	router.HandleFunc("/ws_game", s.handleWsGame)

	http.ListenAndServe(s.addr, router)
}

func setCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authentication, Upgrade, Connection, Host, Accept")
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

	fmt.Println("HERE")

	if r.Method == "GET" {
		jsonWelcome, err := json.Marshal(welcome)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		assignJWT(w, r) //later to be moved to /login POST
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonWelcome)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (s *Server) handleWsGame(w http.ResponseWriter, r *http.Request) {
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
