package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"regexp"
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
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		var req LoginCredentials
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println(err)
			return
		}

		//... something something
		//check verify the input
		//if ok then give jwt:

		fmt.Println(req)
		if req.Login == "frank" {
			assignJWT(w, r)
		}
	}
}

func (s *Server) handleRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == "POST" {
		var req RegisterCredentials
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println(err)
			return
		}

		// verify users input
		// check if the user already exists in the database
		// if ok:
		// hash the password
		// save to the database

		fmt.Println("Beginning the validation process...")

		if req.Login == "" || len([]rune(req.Login)) < 3 {
			writeJSON(w, http.StatusBadRequest, "Login is incorrect")
			return
		}
		isValidEmail, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, req.Email)
		if err != nil {
			fmt.Println(err)
			return
		}
		if req.Email == "" || !isValidEmail {
			writeJSON(w, http.StatusBadRequest, "Email is incorrect")
			return
		}
		if req.Password != req.AgainPassword {
			writeJSON(w, http.StatusBadRequest, "Passwords are not the same")
			return
		}

		user, err := NewUser(req.Login, req.Email, req.Password)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, "Something went wrong")
			return
		}

		fmt.Println("...Validation process finished")

		u, err := NewUser(req.Login, req.Email, req.Password)
		fmt.Println(u)

		writeJSON(w, http.StatusOK, user)
		return

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
