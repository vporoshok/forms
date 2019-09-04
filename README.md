# Forms

[![Travis Build](https://travis-ci.com/vporoshok/forms.svg?branch=master)](https://travis-ci.com/vporoshok/forms)
[![Go Report Card](https://goreportcard.com/badge/github.com/vporoshok/forms)](https://goreportcard.com/report/github.com/vporoshok/forms)
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/vporoshok/forms)
[![codecov](https://codecov.io/gh/vporoshok/forms/branch/master/graph/badge.svg)](https://codecov.io/gh/vporoshok/forms)
[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE)

> Generate HTML forms from struct

Just make struct of request form your service handler expect and generate, render, parse, validate and render errors quick and simple. This library provides form-abstraction that can be built manual, reflected from struct tags or pre-generated from struct tags. Forms may be rendered in an HTML presentation with optional renderers (you can create your renderer). Also, forms may be parsed from HTTP request form values or JSON-body. After parsing forms may be used to collect errors (by field) to render in response.

## Motivation

It is often difficult to understand what went wrong in a complex system built with microservices. Sometimes you need to just build request to microservice and look at the result. But microservices may have different and complex interfaces like [AMQP](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol), [gRPC](https://en.wikipedia.org/wiki/GRPC) or something else. In the following of my approach to project architecture ([in Russian](https://vporoshok.me/post/2018/04/clean-architect/)) business logic shouldn't depend on interface. So it should be easy to provide an alternative interface to service logic. Well, it may be some kind of rest API. But it would be quite good to give a simple HTML interface to the service. But build forms in HTML is a double work. So this library helps to render HTML forms from your Go code like a charm.

### Philosophy

The idea of this project is quite simple:

> Any form is a data expected by action

We have to describe this data in any case. So why don't we just use this description to build our interface?

## Usage

See full example in documentation.

```go
type LoginData struct {
    Username   string `forms:"Email,placeholder(user@example.com),required"`
    Password   string `forms:",type(password),required"`
    RememberMe bool
}

func (srv *Server) login(w http.ResponseWriter, r *http.Request) {
    data := LoginData{}
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
```

## How it works

Architect, conceptions, dependencies. This section should be written after implementation.

## Key features

* minimal additional code required;
* manual build forms (without reflection);
* build form as reflection by struct tags;
* generate form to optimize runtime speed;
* independent renderers (full customization);
* integration with [gorilla/csrf](https://github.com/gorilla/csrf);

## See also

This project inspired by next projects:
* [Symfony forms](https://symfony.com/doc/current/forms.html)
* [WTForms](https://wtforms.readthedocs.io/en/stable/)

## Roadmap

### v1.0
* examples and api stubs;
* base types;
* controls implementations;
* bootstrap renderer;
* parsing and validations;
* reflection builder;

### v1.1
* gorilla/csrf integration;

### v1.2
* code generation;

### Next
* guide: How to write renderer;
