//nolint:unused
package forms_test

import (
	"context"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/vporoshok/forms"
	"github.com/vporoshok/forms/renderer/bootstrap"
)

type Server struct {
	mux          *http.ServeMux
	pageTemplate *template.Template
}

func NewServer() (*Server, error) {
	renderer, err := bootstrap.New()
	if err != nil {
		return nil, err
	}

	pageTemplate, err := template.New("").Funcs(template.FuncMap{
		"renderForm": renderer.Render,
	}).Parse(`
	<!doctype html>
	<html lang="en">
	  <head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<link rel="stylesheet" href="/assets/bootstrap.min.css">
		<title>Hello, World!</title>
	  </head>
	  <body>
		<h1>Login</h1>

		{{ . | renderForm }}

		<script src="/assets/jquery-3.3.1.slim.min.js"></script>
		<script src="/assets/popper.min.js"></script>
		<script src="/assets/bootstrap.min.js"></script>
	  </body>
	</html>
	`)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		mux:          http.NewServeMux(),
		pageTemplate: pageTemplate,
	}

	srv.mux.HandleFunc("", srv.login)

	return srv, nil
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.mux.ServeHTTP(w, r)
}

type Region int8

const (
	Europe Region = iota
	Asia
	SouthAmerica
	NorthAmerica
)

type LoginData struct {
	Username   string `forms:"Email,placeholder(user@example.com),required"`
	Password   string `forms:",type(password),required"`
	RememberMe bool
	Region     Region
}

func (srv *Server) login(w http.ResponseWriter, r *http.Request) {
	data := LoginData{
		Region: Asia,
	}
	form, err := forms.From(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodPost {
		if form.Parse(r) {
			// Validate (just dummy example)
			if !strings.ContainsRune(data.Username, '@') {
				form.AddFieldError("Username", errors.New("should be a valid email"))
			}
			if data.Password != "password" {
				form.AddFormError(errors.New("invalid username or password"))
			}
			if form.IsValid() {
				// do login
				w.WriteHeader(http.StatusOK)
				return
			}
		}
	}
	if form.IsValid() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	_ = srv.pageTemplate.Execute(w, form)
}

func Example() {
	srv, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Addr:    ":8000",
		Handler: srv,
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		go func() {
			<-c
			cancel()
		}()
		_ = server.Shutdown(ctx)
	}()

	if err = server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
