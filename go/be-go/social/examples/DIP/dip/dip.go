package main

import "fmt"

// Abstraction
type Notifier interface {
	Send(message string)
}

// Low-level modules
type EmailNotifier struct{}
func (e *EmailNotifier) Send(message string) {
	fmt.Println("Sending Email:", message)
}

type SMSNotifier struct{}
func (s *SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// High-level module
type Alert struct {
	notifier Notifier // depends on interface
}

func (a *Alert) Trigger(message string) {
	a.notifier.Send(message)
}

func main() {
	email := &EmailNotifier{}
	sms := &SMSNotifier{}

	alertEmail := &Alert{notifier: email}
	alertEmail.Trigger("Server down!")

	alertSMS := &Alert{notifier: sms}
	alertSMS.Trigger("High CPU usage!")
}

