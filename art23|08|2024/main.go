package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)


func main() {
	
	server := gin.Default()

	server.POST("/sendEmail", func(ctx *gin.Context) {
		if err := SendEmail(); err != nil {
			fmt.Println("error ", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error sending email"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Email send successfully"})
	})

	server.Run()

}


func SendEmail() error {
	m := gomail.NewMessage()

    m.SetHeader("From", "Gomail<hey@example.com>")
    m.SetHeader("To", "example@gmail.com")
    m.SetHeader("Subject", "Hello there")
    m.SetBody("text/plain", "Are you getting this")

    d := gomail.NewDialer("mail.privateemail.com", 465, "hey@example.com", "*********")

    if err := d.DialAndSend(m); err != nil {
        fmt.Println("Failed to send email", err)
		return err
    }

    fmt.Println("Email Sent Successfully!")
	return nil
}