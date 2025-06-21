package main

import "html/template"

type Templates struct {
	templates *template.Template
}

type Count struct {
	Count int
}

type Contact struct {
	Name  string
	Email string
	Id    int
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
