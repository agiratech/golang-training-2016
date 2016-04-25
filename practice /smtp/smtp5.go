 package main

 import (
         "fmt"
         "net/smtp"
         "strconv"
         "errors"
         "github.com/mailgun/mailgun-go"
 )

 type EmailConfig struct {
         Username string 
         Password string  
         Host     string
         Port     int
 }

const (

         smtpUser = "arunkumar@agiratech.com"
         smtpPass = "agirah6b48"
         smtpHost = "smtp.gmail.com"
         smtpPort =  587
)

 
 func mail(Body error, receiver []string) {

         for i:=0;i<len(receiver);i++{
         
         errormessage := []byte("From: " + smtpUser + " \r\n"+
                                "To: " + receiver[i] + " \r\n" +
                                "Subject: Error Message!\r\n" +
                                "\r\n" + Body +"\r\n") 
         emailConf := &EmailConfig{smtpUser, smtpPass, smtpHost, smtpPort}
         emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)
         sender := smtpUser
         err := smtp.SendMail(smtpHost+":"+strconv.Itoa(emailConf.Port),emailauth,sender,receiver,errormessage,)
         if err != nil {
                 fmt.Println(err)
         }
 }
 }

 func main() {
       
       // var s string
       // err := errors.New("Some error in your coding")
       // fmt.Printf("%T",err)
      //  s = err.Error()
        receivers := []string{"arunkumar@agiratech.com","reddysai@agiratech.com"} 
       var a int = 3
       var b int = 4
       var c int
       c= a + b 
       if c!= 7 {
        mail(err, receivers)
       } else {
        mail(err, receivers)
       }

 }