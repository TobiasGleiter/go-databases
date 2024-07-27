package main

import "net/http"

func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	// This follows the UserModelInterface.
	// If you swtch the database, you need to implement this function defined in the interface.
	app.users.Insert("Tobias Gleiter", "gleiter.tobias@gmail.com", "pa$$word")
	return
}
