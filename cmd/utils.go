package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"os"

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

func loadContactsFromJSON(path string) ([]Contact, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var contacts []Contact
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

func newData() (Data, error) {
	contacts, err := loadContactsFromJSON("data/contacts.json")

	if err != nil {
		return Data{}, errors.New("could not read from the file")
	}

	var newContacts []Contact

	newContacts = append(newContacts, contacts...)

	return Data{
		Contacts: newContacts,
	}, nil
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func newPage() Page {

	data, _ := newData()

	return Page{
		Data: data,
		Form: newFormData(),
	}
}
