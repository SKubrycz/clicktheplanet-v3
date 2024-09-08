package main

type User struct {
	Id       uint64 `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func returnUser() User {
	user1 := User{
		Id:       1,
		Login:    "Someone",
		Email:    "someone@mail.com",
		Password: "itwillbehashedlater",
	}

	return user1
}
