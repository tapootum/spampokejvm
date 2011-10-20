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
var count int


func sendmail (){
	//	random file	
                num := rand.Int63n(20)
                ff := strconv.Itoa64(num)
                file := ff + ".txt"
	//	random subject
                num2 := rand.Int63n(20)
                aa := strconv.Itoa64(num2)
                dir := aa + ".txt"


		da := "subject/" + dir
		ms , _ := ioutil.ReadFile(file)
		sub , _ := ioutil.ReadFile(da)
		subj := string(sub)
		mss := string(ms)

	//     send mail of smtp
		msg := "Subject:" + subj + "\n" + mss
		auth := smtp.PlainAuth(
			"",
			"spam@tapootum.com",
			"XXXXXX",
			"tapootum.com",  // smtp protocal ,, No port
		)
		err := smtp.SendMail(
			"tapootum.com:25",  
			auth,
			"spam@tapootum.com",   // from 
			[]string{"spamtest@tapootum.com"},  //to
			[]byte(msg),
			)
		if err != nil {
			log.Fatal(err)
			}
		count++
		fmt.Print(count)
		fmt.Println("   Send mail")	
	// channel
 	 wait <- 1

}

func main() {
	times1 , _ , _ := os.Time()
	runtime.GOMAXPROCS(2)
	if len(os.Args) == 1 {
		fmt.Println("Don't send mail")
	}else {
	
	s := os.Args[1]
	n , _ := strconv.Atoi64(s)

	for i:=1 ; int64(i)<=n ; i++{
		go sendmail()
		go sendmail()
		go sendmail()
		go sendmail()
		go sendmail()
		go sendmail()
		go sendmail()
		go sendmail()
		go sendmail()
		go sendmail()

	fmt.Printf("|||||||||||||||||||||| %d\n",i)
		<- wait
		<- wait
		<- wait
		<- wait
		<- wait
		<- wait
		<- wait
		<- wait
		<- wait
		<- wait
		}   
	}	

	times2 , _ , _ := os.Time()
 	tt := times2-times1
		fmt.Printf("Send spam mail = ")
		fmt.Println(count)
		fmt.Printf(".............%d s...........\n",tt)
}
