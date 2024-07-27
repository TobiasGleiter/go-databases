package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
	mux.HandleFunc("POST /users/create", app.CreateUser)

	return mux
}
