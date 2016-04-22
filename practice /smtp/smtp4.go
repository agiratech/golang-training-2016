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


func sendmemail(m []byte) {
         
         var smtpUser, smtpPass string 
         receiver := []string{}
         fmt.Println("Enter the Username and Password")
         fmt.Scanf("%s",&smtpUser)
         fmt.Scanf("%s",&smtpPass)
         smtpHost := "smtp.gmail.com" 
         smtpPort :=  587 
         //smtpPass = "agirah6b48"
        // smtpUser = "arunkumar@agiratech.com"
         emailConf := &EmailConfig{smtpUser, smtpPass, smtpHost, smtpPort}
         emailauth := smtp.PlainAuth("", emailConf.Username, emailConf.Password, emailConf.Host)
         //if err!=nil{
           //   fmt.Println(err)
         //}else {
         fmt.Println("Successfully Logged")
         //}
         fmt.Println("Please Enter the Receiver Email-id to SendMail")
         var c="y"
         var receivers string
         for c=="y"{
               fmt.Scanf("%s",&receivers)
               receiver = append(receiver,receivers)
                fmt.Println("if u want to add another receiver press 'y'")
                fmt.Scanf("%s",&c)
         }
        // fmt.Scanf("%s",&receiver)
         sender := smtpUser 
        // receiver = //[]string{
                // "reddysai@agiratech.com",
                // "arunkumar@agiratech.com", 
                // "arun@agiratech.com",
                 //} 
        // message := []byte("To: arunkumar@agiratech.com\r\n" +
                   //        "Subject: Todays Summmary!\r\n" +
                   //        "\r\n" +"Worked on SMTP.\r\n")   
         err := smtp.SendMail(smtpHost+":"+strconv.Itoa(emailConf.Port),emailauth,sender,receiver,m,)
         if err != nil {
                 fmt.Println(err)
         }else {
                fmt.Println("Your mail to  is Successfully Delivered")  
         }
  

}
func main() {
         
        message := []byte("To: arunkumar@agiratech.com\r\n" +
                           "Subject: Todays Summmary!\r\n" +
                           "\r\n" +"Worked on SMTP.\r\n")  
        sendmemail(message)
         
 }