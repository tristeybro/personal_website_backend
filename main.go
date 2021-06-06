package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type ContactEmailResponse struct {
	Success bool
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	// Get our email parameters.
	r.ParseMultipartForm(32 << 20)
	sender := r.FormValue("sender")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	// Construct our email
	fmt.Println(r.FormValue("sender"))
	from := mail.NewEmail("Example User", "tristers.b@gmail.com")
	body := "Received a message from: " + sender + ".\n" + message
	htmlContent := "<strong>and easy to do anywhere</strong>"
	to := mail.NewEmail("Tristan Benavides", "tristers.b@gmail.com")
	email := mail.NewSingleEmail(from, subject, to, body, htmlContent)

	// Send out our email.
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(email)
	contactEmailResponse := ContactEmailResponse{true}
	if err != nil {
		log.Println(err)
		contactEmailResponse = ContactEmailResponse{false}
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	// Send response to client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contactEmailResponse)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/sendEmail", sendEmail)
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
