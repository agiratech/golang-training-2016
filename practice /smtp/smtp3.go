package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "arunkumar@agiratech.com", "agirah6b48", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"arunkumar@agiratech.com"}
	msg := []byte("To: arunkumar@agiratech.com\r\n" +
		"Subject: Todays Summmary!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("smtp.gmail.com:587", auth, "arunkumar@agiratech.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}