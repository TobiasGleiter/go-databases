package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{id}", app.GetUser)
	mux.HandleFunc("PATCH /users/{id}", app.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", app.DeleteUser)
	mux.HandleFunc("POST /users/create", app.CreateUser)

	return mux
}
