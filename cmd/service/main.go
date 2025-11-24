package main

import (
	"db/internal/database"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(1200, 600))

	loginLabel := widget.NewLabel("Login")
	login := widget.NewEntry()

	passwordLabel := widget.NewLabel("Пароль")
	password := widget.NewPasswordEntry()

	var connection string
	var db *database.Storage

	buttonLogIn := widget.NewButton("Войти", func() {
		connection = fmt.Sprintf("postgres://%s:%s@localhost:5432/postgres?sslmode=disable", login.Text, password.Text)

		db = database.NewStorage(connection)

	})

	defer db.Close()

	w.SetContent(container.NewVBox(loginLabel, login, passwordLabel, password, buttonLogIn))
	w.ShowAndRun()
}
