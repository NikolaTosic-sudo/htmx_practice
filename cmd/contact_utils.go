package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func (page *Page) AddNewContact() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			return c.Render(http.StatusUnprocessableEntity, "form", formData)
		}

		contact := newContact(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)

		c.Render(http.StatusOK, "form", newFormData())
		return c.Render(http.StatusOK, "oob-contact", contact)
	}
}

func (page *Page) DeleteContact() echo.HandlerFunc {
	return func(c echo.Context) error {
		time.Sleep(3 * time.Second)
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		index := page.Data.indexOf(id)
		if index == -1 {
			return c.String(http.StatusNotFound, "Contact not found")
		}

		page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)

		return c.NoContent(http.StatusOK)
	}
}

func (page *Page) UpdateContact() echo.HandlerFunc {
	return func(c echo.Context) error {
		strId := c.Param("id")
		email := c.QueryParam("email")
		id, err := strconv.Atoi(strId)
		name := c.FormValue("update-name")

		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid Id")
		}

		index := page.Data.indexOf(id)

		if index == -1 {
			return c.String(http.StatusNotFound, "Contact not found")
		}

		page.Data.updateContact(id, name, email)
		contact := page.Data.Contacts[index]

		c.Render(http.StatusOK, "button-edit", contact)
		return c.Render(http.StatusOK, "oob-update-contact", contact)
	}
}

func (page *Page) GetEditContact() echo.HandlerFunc {
	return func(c echo.Context) error {

		strId := c.Param("id")
		id, err := strconv.Atoi(strId)

		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid id")
		}

		index := page.Data.indexOf(id)

		if index == -1 {
			return c.String(http.StatusNotFound, "Contact not found")
		}

		contact := page.Data.Contacts[index]

		return c.Render(http.StatusOK, "edit-contact", contact)
	}
}
