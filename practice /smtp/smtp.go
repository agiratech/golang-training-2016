 package main

 import (
         "fmt"
         "net/smtp"
         "strconv"
 )

 type EmailConfig struct {
         Username string
         Password string
         Host     string
         Port     int
 }

 func main() {
         //authentication configuration
        smtpHost := "smtp.gmail.com" // change to your SMTP provider address
        smtpPort :=  587 // change to your SMTP provider port number
        smtpPass := "agirah6b48" // change here
        smtpUser := "arunkumar@agiratech.com" // change here

         emailConf := &EmailConfig{smtpUser, smtpPass, smtpHost, smtpPort}

         emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)

         sender := "arunkumar@agiratech.com" // change here

         receivers := []string{
                 "reddysai@agiratech.com",
                 "arunkumar@agiratech.com", 
                 "arun@agiratech.com",
                 } // change here


         message := []byte("To: arunkumar@agiratech.com\r\n" +
                "Subject: Todays Summmary!\r\n" +
                "\r\n" +
                "Worked on SMTP.\r\n") // your message

         // send out the email
         err := smtp.SendMail(smtpHost+":"+strconv.Itoa(emailConf.Port),emailauth,sender,receivers,message,)
         //convert port number from int to string
         if err != nil {
                 fmt.Println(err)
         }
 }