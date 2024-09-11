package main

type User struct {
	Id       uint64 `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) createUser() {

}
