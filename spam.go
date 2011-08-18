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
var wait2 = make (chan int)
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
			"kawaoisoki",
			"tapootum.com",  // smtp protocal
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
		fmt.Println("Send mail")	
	// channel
 	 wait <- 1

}
func gomail(){
	 for i := 1 ; i <= 10 ; i++{
		go sendmail()
		<- wait
	}
wait2 <- 1
}
func main() {
	//times1,timens1,_ := os.Time()
	runtime.GOMAXPROCS(4)
	if len(os.Args) == 1 {
		fmt.Println("Don't send mail")
	}else {
	
	s := os.Args[1]
	n , _ := strconv.Atoi64(s)

	for i:=1 ; int64(i)<=n ; i++{
		go gomail()
	fmt.Printf("||||||||||||||||||||||||||||||||||||||||||| %d\n",i)
		<- wait
		}
	for i:=1 ; int64(i)<=n*10 ; i++{
		<-wait
		<- wait2
		}
	}	

	//times2,timens2,_ := os.Time()
 
		fmt.Printf("Send spam mail = ")
		fmt.Println(count)
	//	fmt.Printf("%.2f s =  %.2f ns/n",times2,timens2)
}
