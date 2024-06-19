package mail

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/gomail.v2"

	"github.com/thebravebyte/findr/internals"
)

func MailServer(mc internals.Mail) {
	d := gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("SOURCE_EMAIL_ADDRESS"), os.Getenv("APP_PASSWORD"))

	s, err := d.Dial()
	if err != nil {
		log.Panicf("Error connecting to the Mail Server: %v", err)
	}

	msg := gomail.NewMessage()

	msg.SetAddressHeader("From", mc.Source, os.Getenv("USERNAME"))
	msg.SetHeader("To", mc.Destination)
	msg.SetHeader("Subject", mc.Subject)
	if mc.Template == "" {
		msg.SetBody("text/html", mc.Message)
	} else {
		data, err := ioutil.ReadFile(fmt.Sprintf("./template/%s", mc.Template))
		if err != nil {
			log.Panic(err)
		}
		temp := string(data)
		mailToSend := strings.Replace(temp, "[%body%]", mc.Message, 1)
		msg.SetBody("text/html", mailToSend)
	}

	if err := gomail.Send(s, msg); err != nil {
		log.Printf("Mail Sever : %s %v\n", mc.Destination, err)
	}
	log.Printf("Mail Successfully Sent to %s", mc.Destination)
	msg.Reset()
}

// MailDelivery this function uses a goroutine to send mail to the users via email address
func MailDelivery(mc chan internals.Mail, worker int) {
	// Buffered channel for completion signals
	completionChan := make(chan bool, worker)

	for x := 0; x < worker; x += 1 {
		go func(x int) {
			// Signal completion
			defer func() {
				completionChan <- true
			}()

			for m := range mc {
				MailServer(m)
			}
		}(x)
	}

	// Wait for all goroutines to complete
	for x := 0; x < worker; x += 1 {
		<-completionChan
	}
}
