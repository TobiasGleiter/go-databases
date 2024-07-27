package main

import "net/http"

func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	app.users.Insert("Tobias Gleiter", "gleiter.tobias@gmail.com", "pa$$word")
	return
}
