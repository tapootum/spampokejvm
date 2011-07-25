package main

import (
	
	"smtp"
	"log"
	)

func main() {

	auth := smtp.PlainAuth(
			"",
			"gwean@grean.com",
			"password",
			"smtp.gmail.com",
		)
	//for i:=1; i<=5 ; i++ {
		err := smtp.SendMail(
				"smtp.gmail.com:587",
				auth,
				"tapootumfb@gmail.com",
				[]string{"tapootumfb@gmail.com"},
				[]byte("Gwean send spam mail"),
			)
		if err != nil {
			log.Fatal(err)
			}
	//	}

}
