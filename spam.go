package main

import (
	"fmt"
	"net/smtp"
	"log"
//	"os"
	"strconv"
	"math/rand"
	"io/ioutil"
	"runtime"
	)




var wait = make (chan int)
var count int


func sendmail (){
	//	random file	
                num := rand.Int63n(20)
                ff := strconv.FormatInt(num,10)
                file := ff + ".txt"
	//	random subject
                num2 := rand.Int63n(20)
                aa := strconv.FormatInt(num2,10)
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
			"tapootumchannel@gmail.com",
			"p@ssw0rd",
			"smtp.gmail.com",  // smtp protocal ,, No port
		)
		err := smtp.SendMail(
			"smtp.gmail.com:587",  
			auth,
			"tapootumchannel@gmail.com",   // from 
			[]string{"tapootumchannel@gmail.com"},  //to
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
//	times1 , _ , _ := os.Time()
	runtime.GOMAXPROCS(2)
	//if len(os.Args) == 1 {
	//	fmt.Println("Don't send mail")
	//}else {
	
	//s := os.Args[1]
	//n , _ := strconv.FormatInt(s,10)
	n := 1

	//for i:=1 ; int64(i)<=n ; i++{
	for i:=1 ; i<=n ; i++{
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
	//}	

//	times2 , _ , _ := os.Time()
 //	tt := times2-times1
		fmt.Printf("Send spam mail = ")
		fmt.Println(count)
//		fmt.Printf(".............%d s...........\n",tt)
}
