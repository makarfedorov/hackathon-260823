package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type user struct {
	name     string
	password string
}

// test users
var roleByUser = map[user]string{
	user{
		name:     "John",
		password: "Mnemonic",
	}: "administrator",

	user{
		name:     "Kate",
		password: "qwerty",
	}: "consultant",
}

func (u user) authenticate() (string, bool) {
	role, exist := roleByUser[u]
	return role, exist
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		fmt.Fprintf(w, "No chatID in request\n")
		return
	}
	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Can't parse chatID\n")
		return
	}

	if r.Method == http.MethodPost {
		role, exist := user{
			name:     r.FormValue("username"),
			password: r.FormValue("password"),
		}.authenticate()

		if exist {
			userByChatID[id] = NewUser(id, role)
			fmt.Fprintf(w, "Welcome! You can continue using bot as %q\n", role)
		} else {
			fmt.Fprintf(w, "Login failed\n")
		}
		return
	}

	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Login</title>
		</head>
		<body>
			<h1>Login</h1>
			<form method="post">
				<label for="username">Username:</label>
				<input type="text" id="username" name="username"><br><br>
				<label for="password">Password:</label>
				<input type="password" id="password" name="password"><br><br>
				<input type="submit" value="Login">
			</form>
		</body>
		</html>
	`

	fmt.Fprintf(w, html)
}

func startLoginService() {
	http.HandleFunc("/login", loginHandler)
	log.Println("Authentication service started on :8080")
	http.ListenAndServe(":8080", nil)
}
