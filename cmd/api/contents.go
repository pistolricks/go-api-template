package main

import (
	"fmt"
	"github.com/pistolricks/go-api-template/internal/data"
	"github.com/pistolricks/validation"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (app *application) uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	err := r.ParseMultipartForm(10)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	var input struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"-"`
		Name      string    `json:"name"`
		Src       string    `json:"src"`
		Type      string    `json:"type"`
		Size      int32     `json:"size"`
		Width     float32   `json:"width"`
		Height    float32   `json:"height"`
		SortOrder int16     `json:"sort_order"`
		UserID    string    `json:"user_id"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	file, handler, err := r.FormFile(input.Src)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			app.badRequestResponse(w, r, err)
		}
	}(file)

	var filename = handler.Filename
	var option = handler.Header["Option"][0]

	fmt.Println("Past readJSON")

	dst, err := app.createFile(w, r, handler.Filename)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}
	}(dst)

	content := &data.Content{
		ID:        input.ID,
		CreatedAt: input.CreatedAt,
		Name:      filename,
		Src:       input.Src,
		Type:      option,
		Size:      input.Size,
		Width:     input.Width,
		Height:    input.Height,
		SortOrder: input.SortOrder,
		UserID:    input.UserID,
	}

	v := validation.New()

	if data.ValidateContent(v, content); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Contents.EncodeWebP(content)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/contents/%s", content.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"content": content}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createFile(w http.ResponseWriter, r *http.Request, filename string) (*os.File, error) {
	// Create an uploads directory if it doesn’t exist
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err := os.Mkdir("uploads", 0755)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	}

	// Build the file path and create it
	dst, err := os.Create(filepath.Join("uploads", filename))
	if err != nil {
		return nil, err
	}

	return dst, nil
}
