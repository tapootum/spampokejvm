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
			"tapootum@localhost",
			"1234",
			"smtp.localhost",
		)
		err := smtp.SendMail(
			"smtp.localhost:25",
			auth,
			"tapootum@localhost",
			[]string{"tapootum@localhost"},
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
	n , _ := strconv.Atoi(s)

	for i:=1 ; i<=n ; i++{	
		num := rand.Int63n(n)
		ff := strconv.Itoa64(num)
		file := ff + ".txt"
		go sendmail(file)
		}
	for i:=1 ; i<=n ; i++{
		<- wait
		}

	}
}
