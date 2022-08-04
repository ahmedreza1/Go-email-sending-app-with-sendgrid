package main

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
    // Load variables from .env
    err := godotenv.Load()
    if err != nil {
            log.Fatal("Error loading .env file")
    }
    
    // Create a new SendGrid client
    from := mail.NewEmail(os.Getenv("SEND_FROM_NAME"), os.Getenv("SEND_FROM_ADDRESS"))
    subject := "You got a new message!"
    to := mail.NewEmail(os.Getenv("SEND_TO_NAME"), os.Getenv("SEND_TO_ADDRESS"))
    plainTextContent := "Hey Ahmed wassup?"
    htmlContent := "Hey Ahmed wassup?"
    message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

    // Start Sending the email
    client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := client.Send(message)
    if err != nil {
            fmt.Println("Unable to send your email")
            log.Fatal(err)
    }

    // Check if it was sent
    statusCode := response.StatusCode
    if statusCode == 200 || statusCode == 201 || statusCode == 202 {
            fmt.Println("Email sent!")
    }
}
