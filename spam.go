package main

import (
	"fmt"
	"smtp"
	"log"
	"os"
	"strconv"
	"rand"
	"io/ioutil"
	)
var wait = make (chan int)
func sendmail (file string){

		st , _ := ioutil.ReadFile(file)
		msg := string(st)
		
		auth := smtp.PlainAuth(
			"",
			"spam@tapootum.com",
			"kawaoisoki",
			"tapootum.com",
		)
		err := smtp.SendMail(
			"tapootum.com:25",
			auth,
			"spam@tapootum.com",
			[]string{"spamtest@tapootum.com"},
			[]byte(msg),
			)
		if err != nil {
			log.Fatal(err)
			}
		fmt.Println("Send mail")
	
 wait <- 1
}
func main() { 

	if len(os.Args) == 1 {
		fmt.Println("Don't send mail")
	}else {
	
	s := os.Args[1]
	n , _ := strconv.Atoi64(s)

	for i:=1 ; int64(i)<=n ; i++{	
		num := rand.Int63n(20)
		ff := strconv.Itoa64(num)
		file := ff + ".txt"
		go sendmail(file)
		}
	for i:=1 ; int64(i)<=n ; i++{
		<- wait
		}

	}
}
