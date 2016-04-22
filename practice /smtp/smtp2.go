package main

import (
    "gopkg.in/gomail.v2"
)

func main() {
    m := gomail.NewMessage()
    m.SetAddressHeader("From", "sender@example.com", "Sandy Sender")
    m.SetAddressHeader("To", "recipient@example.com")
    m.SetHeader("Subject", "Hello!")
    m.SetBody("text/plain", "This is the body of the message.")

    d := gomail.NewPlainDialer("smtp.example.com", 587, "user", "123456")

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}