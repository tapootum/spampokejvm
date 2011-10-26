package main 

import (
	"smtp"
	"log"
	"fmt"
)

func main() {
	/*
	auth := smtp.PlainAuth(
			"",
			"example@mail.com",
			"password",
			"smtp.gmail.com",
		)
	for i := 1; i <= 10; i++ {
		err := smtp.SendMail(
				"smtp.gmail.com:587",
				auth,
				"wingyminus@gmail.com",
				[]string{"wingyminus@gmail.com"},
				[]byte("test greans send mail."),
			)
		if err != nil {
			log.Fatal(err)
		}
	}
	*/
	
	client, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dial PASS!!")
	
	err = client.StartTLS(nil)
	if err != nil {
		log.Fatal(err)
	}
	auth := smtp.PlainAuth(
			"",
			"example@mail.com",
			"password",
			"smtp.gmail.com",
		)
	if err = client.Auth(auth); err != nil {
		log.Fatal(err)
	}
	fmt.Println("AUTH PASS!!")
	
	// Send Command MAIL FROM:
	if err = client.Mail("wingyminus@gmail.com"); err != nil {
		log.Fatal(err)
	}
	
	// Send Command RCPT TO:
	if err = client.Rcpt("wingyminus@gmail.com"); err != nil {
		log.Fatal(err)
	}
	
	// Create body
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	
	defer client.Quit()
	defer wc.Close()
	
	// Write body
	if _, err = wc.Write([]uint8("This is greans send body of mail!!")); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Send Mail Success!!")
	
	
	
	
}
