package main

import (
	"fmt"
//	"runtime"
	"smtp"
	"log"
	)
//var gm = make (chan int)
//var m = make (chan int)
var wait = make (chan int)
func sendmail (){

		
	
		auth := smtp.PlainAuth(
			"",
			"spamgrean@gmail.com",
			"greangrean",
			"smtp.gmail.com",
		)
		err := smtp.SendMail(
			"smtp.gmail.com:587",
			auth,
			"spamgrean@gmail.com",
			[]string{"spamgrean@gmail.com"},
			[]byte("Gwean send spam mailgo "),
			)
		if err != nil {
			log.Fatal(err)
			}
		fmt.Println("Send mail")
	
 wait <- 1
}
func main() {
	for i:=1 ; i<=10 ; i++{
		go sendmail()
	}
	for i:=1 ; i<=10 ; i++{
		<- wait
	}


}
