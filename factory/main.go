package main

import "fmt"

// Interfaces

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// Implentacion ISender para SMS

type SMSNotificationSender struct{}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Implementacion de INotificacion
type SMSNotification struct{}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification of SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// Implementacion  Isender Email
type EmailNotificationSender struct{}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

// Implementacion Inotificacion Email
type EmailNotification struct{}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification of Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No Notification type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	sendNotification(smsFactory)
	getMethod(smsFactory)
}
