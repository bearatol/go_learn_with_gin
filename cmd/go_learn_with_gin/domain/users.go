package domain

type user struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

var usersList = []user{
	user{ID: 1, Login: "admin", Password: "admin"},
	user{ID: 2, Login: "qwerty", Password: "qwerty1"},
}

func GetAllUsers() []user {
	return usersList
}
