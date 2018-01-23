package main

import (
    "personal_website_backend/credentials"
    "encoding/json"
    "fmt"
    "net/http"
    "net/smtp"
    "log"
)

type ContactEmailResponse struct {
    Success bool
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
    // Get our email parameters.
    r.ParseForm()
    sender := r.Form["sender"][0]
    subject := r.Form["subject"][0]
    message := r.Form["message"][0]
    // Send out our email.
    myEmailAddress := credentials.Credentials.Email
    password := credentials.Credentials.Password
    msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\nSent by %s\n\n%s", sender, myEmailAddress, subject, sender, message)
    auth := smtp.PlainAuth("", myEmailAddress, password, "smtp.gmail.com")
    err := smtp.SendMail("smtp.gmail.com:587", auth, myEmailAddress, []string{myEmailAddress}, []byte(msg))
    contactEmailResponse := ContactEmailResponse { true }
    if err != nil {
        log.Printf("smtp error: %s", err)
        contactEmailResponse = ContactEmailResponse { false }
    }
    // Send response to client.
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(contactEmailResponse)
}

func main() {
    http.HandleFunc("/", sendEmail)           // set router
    err := http.ListenAndServe(":8000", nil)  // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
