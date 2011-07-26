package main

import (
	"time"
	"fmt"
	"runtime"
	"smtp"
	"log"
	)
var c = make (chan int)
func sendmail (){

		c <- 1
	
		auth := smtp.PlainAuth(
			"",
			"tapootumfb@gmail.com",
			"password",
			"smtp.gmail.com",
		)
		err := smtp.SendMail(
			"smtp.gmail.com:587",
			auth,
			"tapootumfb@gmail.com",
			[]string{"tapootumfb@gmail.com"},
			[]byte("Gwean send spam mailgo "),
			)
		if err != nil {
			log.Fatal(err)
			}

	

}

func main() {
	n := 5
	runtime.GOMAXPROCS(n)
	
	for i:=1; i<=5 ; i++ {
	go  sendmail()		
        time.Sleep(600000)	
	<- c
	fmt.Printf("Email %d\n",i)	
	
	}


}
