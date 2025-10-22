package main

import "fmt"

// Low-level module
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// High-level module
type Alert struct {
	emailNotifier *EmailNotifier // depends directly on low-level
}

func (a *Alert) Trigger(message string) {
	a.emailNotifier.Send(message)
}

func main() {
	email := &EmailNotifier{}
	alert := &Alert{emailNotifier: email}

	alert.Trigger("Server down!")
}

