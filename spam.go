package main

import (
	"fmt"
	"smtp"
	"log"
	"os"
	"strconv"
	"rand"
	"io/ioutil"
	"runtime"
	)
var wait = make (chan int)
func sendmail (file string, dir string){
		da := "subject/" + dir
		ms , _ := ioutil.ReadFile(file)
		sub , _ := ioutil.ReadFile(da)
		subj := string(sub)
		mss := string(ms)
		msg := "Subject:" + subj + "\n" + mss
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
	runtime.GOMAXPROCS(4)
	if len(os.Args) == 1 {
		fmt.Println("Don't send mail")
	}else {
	
	s := os.Args[1]
	n , _ := strconv.Atoi64(s)

	for i:=1 ; int64(i)<=n ; i++{	
		num := rand.Int63n(20)
		ff := strconv.Itoa64(num)
		file := ff + ".txt"

		num2 := rand.Int63n(20)
		aa := strconv.Itoa64(num2)
		dir := aa + ".txt"
		go sendmail(file,dir)
		}
	for i:=1 ; int64(i)<=n ; i++{
		<- wait
		}

	}
}
