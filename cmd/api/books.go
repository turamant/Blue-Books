package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/turamant/bluebooks/internal/data"

)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
	app.notFoundResponse(w, r)
		return
	}

	book := data.Book {
		ID: id,
		CreatedAt: time.Now(),
		Title: "The simple Golang",
		Pages: 402,
		Genres: []string{"language", "it", "development"},
		Language: "english",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book":book}, nil)
	if err != nil {
		app.logger.Println(err)
		app.serverErrorResponse(w, r, err)
	}
}

