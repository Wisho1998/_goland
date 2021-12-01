package main

import "fmt"

// STRUCTS
// : SMS

type SMSNotification struct{}

type SMSNotificationSender struct{}

// : Email

type EmailNotification struct{}

type EmailNotificationSender struct{}

// INTERFACES

type INotification interface {
	SendNotification()
	GetSender() ISender
}
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// RECEIVERS FUNC
// : SMS

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification by SMS")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}
func (SMSNotificationSender) GetSender() ISender {
	return SMSNotificationSender{}
}

// : Email

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification by Email")
}
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// METHODS
// : SMS

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// : Email

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

// MAIN FLOW
func getNotification(notificationType string) (INotification, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf(" Unimplement notification type")
}

func sendNotification(f INotification) {
	f.SendNotification()
}

func getMethod(f INotification) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsNotification, _ := getNotification("SMS")
	emailNotification, _ := getNotification("Email")

	sendNotification(smsNotification)
	sendNotification(emailNotification)

	getMethod(smsNotification)
	getMethod(emailNotification)
}
