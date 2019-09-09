package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const SendgridMailSendAPI = "https://api.sendgrid.com/v3/mail/send"

type MailContent struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type MailTo struct {
	Email string `json:"email"`
}

type MailFrom struct {
	Email string `json:"email"`
}

type Personalization struct {
	To []MailTo `json:"to"`
}

type Mail struct {
	Personalizations []Personalization `json:"personalizations"`
	From             MailFrom          `json:"from"`
	Subject          string            `json:"subject"`
	Content          []MailContent     `json:"content"`
}

func SendNotification(apiKey string, email string, subject string, content string) error {
	mail := Mail{
		Personalizations: []Personalization{{
			To: []MailTo{{Email: email}},
		}},
		From:    MailFrom{Email: email},
		Subject: subject,
		Content: []MailContent{{
			Type:  "text/plain",
			Value: content,
		}},
	}
	data, err := json.Marshal(&mail)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", SendgridMailSendAPI, bytes.NewReader(data))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != 202 {
		return fmt.Errorf("failed to send notification email: API responded with status code %d", response.StatusCode)
	}

	return nil
}
