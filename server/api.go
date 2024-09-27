package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	addr string
	db   Database
}

func NewServer(addr string, db Database) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Start() {
	router := mux.NewRouter()
	router.Use(setCors)

	router.HandleFunc("/", s.handleHome)
	router.HandleFunc("/login", s.handleLogin)
	router.HandleFunc("/register", s.handleRegister)
	router.HandleFunc("/game", checkAuth(s.handleGame))
	router.HandleFunc("/ws_game", checkAuth(s.handleGetWsGame))

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
		writeJSON(w, http.StatusOK, "Welcome to /game")
		return
	} else {
		writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
}

func (s *Server) handleGetWsGame(w http.ResponseWriter, r *http.Request) {
	const userId UserId = "userid" // key of type UserId - it has to stay here
	id, ok := r.Context().Value(userId).(int)
	if !ok {
		fmt.Println(id)
		fmt.Println(ok)
		writeJSON(w, http.StatusForbidden, "Not authorized")
		return
	}

	gameData, err := s.db.GetGameByUserId(id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, "Internal server error")
	}
	game := NewGame(gameData)

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "http://localhost:3000"
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			fmt.Printf("%v: Saving game...\n", time.Now())
			s.db.SaveGameProgress(id, game)
		}
	}()

	defer close(game.Ch)
	go func() {
		for message := range game.Ch {
			switch message {
			case "click":
				fmt.Printf("%v: Planet destroyed - Saving game...\n", time.Now())
				err := s.db.SaveGameProgress(id, game)
				if err != nil {
					fmt.Println("!Error: SaveGameProgress: ", err)
				}
			case "upgrade":
				fmt.Printf("%v: Element upgraded - Saving game...\n", time.Now())
				err := s.db.SaveGameProgress(id, game)
				if err != nil {
					fmt.Println("!Error: SaveGameProgress: ", err)
				}
			}
		}
	}()

	dpsSent := false
	var dps *time.Ticker
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		var response []byte
		if len(string(p)) > 0 && string(p) != " " {
			response = ActionHandler(game, string(p))
			if err := conn.WriteMessage(messageType, response); err != nil {
				log.Println(err)
				return
			}
		}
		if string(p) == "dps" {
			if dpsSent {
				dps.Stop()
			}
			dpsSent = true
			dps = time.NewTicker(100 * time.Millisecond)
			defer dps.Stop()
			go func() {
				for range dps.C {
					response = DealDps(game)
					if err := conn.WriteMessage(messageType, response); err != nil {
						log.Println(err)
						return
					}
				}
			}()
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

		user, err := s.db.GetAccountByLogin(req.Login)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, "Log in error")
			return
		}
		if user == nil {
			writeJSON(w, http.StatusNotFound, "User not found")
			return
		}
		fmt.Println(user.Login)

		compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if compare != nil {
			writeJSON(w, http.StatusBadRequest, "Invalid credentials")
			return
		}

		assignJWT(w, user.Id)
		return
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
		if len([]rune(req.Password)) < 8 {
			writeJSON(w, http.StatusBadRequest, "Password has to be at least 8 characters long")
			return
		}
		isValidPassword := [3]bool{false, false, false}
		for _, ch := range req.Password {
			switch {
			case unicode.IsNumber(ch):
				isValidPassword[0] = true
			case unicode.IsUpper(ch):
				isValidPassword[1] = true
			case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
				isValidPassword[2] = true
			default:

			}
		}
		if !isValidPassword[0] || !isValidPassword[1] || !isValidPassword[2] {
			writeJSON(w, http.StatusBadRequest, "Invalid password")
			return
		}

		user, err := NewUser(req.Login, req.Email, req.Password)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, "Something went wrong")
			return
		}

		fmt.Println("...Validation process finished")

		u, err := NewUser(req.Login, req.Email, req.Password)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		fmt.Println(u)

		if err := s.db.CreateAccount(u); err != nil {
			fmt.Println("Here error2")
			writeJSON(w, http.StatusInternalServerError, "Internal server error")
			return
		}

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

// type WithDbFunc func(http.ResponseWriter, *http.Request)

// func returnHandlerFunc(f WithDbFunc, p *Postgres) http.HandlerFunc {
// 	return func (w http.ResponseWriter, r *http.Request) {
// 	}
// }
