package main

import (
	"html/template"
	"net/smtp"
)

func main() {
	html := template.New("")

	smtp.SendMail()
}
