package main

type RegisterCredentials struct {
	Login         string `json:"login"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	AgainPassword string `json:"againPassword"`
}

type LoginCredentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
