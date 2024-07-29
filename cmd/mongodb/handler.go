package main

import (
	"net/http"

	"github.com/TobiasGleiter/go-databases/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	name := "Tobias Gleiter"
	email := "gleiter.tobias@gmail.com"
	password := "pa$$word"

	err := app.userExists(email)
	if err == models.ErrUserAlreadyExists {
		app.logger.Error(err.Error())
		return
	}

	err = app.users.Insert(name, email, password)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	app.logger.Info("inserted user", "name", "Tobias Gleiter")

	return
}

func (app *application) GetUser(w http.ResponseWriter, r *http.Request) {
	objId, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	user, err := app.users.GetByObjID(objId)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	app.logger.Info("found user", "user", user)

	return
}

func (app *application) UpdateUser(w http.ResponseWriter, r *http.Request) {
	name := "Gleiter, Tobias"
	email := "gleiter.tobias@gmail.com"
	password := "pa$$word"

	objId, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	err = app.users.Update(objId, name, email, password)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	return
}

func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {
	objId, err := primitive.ObjectIDFromHex(r.PathValue("id"))
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	err = app.users.DeleteByObjId(objId)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	return
}

func (app *application) userExists(email string) error {
	_, err := app.users.GetByEmail(email)
	if err == nil {
		return models.ErrUserAlreadyExists
	}
	return nil
}
