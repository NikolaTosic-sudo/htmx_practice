package main

import "html/template"

type Templates struct {
	templates *template.Template
}

type Count struct {
	Count int
}

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Id    int    `json:"id"`
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

type Page struct {
	Data Data
	Form FormData
}
