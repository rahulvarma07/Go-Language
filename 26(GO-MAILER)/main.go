// Import this package...
// go get gopkg.in/gomail.v2

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

}

func main() {
	// initialize...
	sendEmail := gomail.NewMessage()
	sendEmail.SetHeader("From", os.Getenv("Mail"))
	sendEmail.SetHeader("To", os.Getenv("Sendmail"))
	sendEmail.SetHeader("Subject", "Let's GO with go!")
	sendEmail.SetBody("text/plain", "This is a test email sent from Gmail to Outlook using Go!")

	mail := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("Mail"), os.Getenv("Pass"))

	err := mail.DialAndSend(sendEmail)

	if err != nil {
		panic(err)
	}
	log.Println("Sen't mail Successs")
}
