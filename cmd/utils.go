package main

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

var id int = 0

func newContact(name, email string) Contact {
	id++
	return Contact{
		Name:  name,
		Email: email,
		Id:    id,
	}
}

func (d *Data) indexOf(id int) (int, error) {
	for i, contact := range d.Contacts {
		if contact.Id == id {
			return i, nil
		}
	}
	return 0, errors.New("Contact not found")
}

func (d Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func (d *Data) updateContact(id int, name, email string) {
	for i := range d.Contacts {
		if d.Contacts[i].Id == id {
			d.Contacts[i].Name = name
			d.Contacts[i].Email = email
			return
		}
	}
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "jd@gmail.com"),
			newContact("Lane", "ld@gmail.com"),
			newContact("Dane", "dd@gmail.com"),
		},
	}
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}
