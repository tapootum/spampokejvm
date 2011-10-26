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
	"time"
	)

var wait = make (chan int)
var count int

// AUTHEN 
var auth = smtp.PlainAuth(
                        "",
                        "anchalee@bangkokranch.co.th",
                        "11111",
                        "mail2.bangkokranch.co.th",  // smtp protocal //No port
                )

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
		err := smtp.SendMail(
			"mail2.bangkokranch.co.th:25",  
			auth,
			"anchalee@mail2bangkokranch.co.th",   // from 
			[]string{"anchalee@bangkokranch.co.th"},  //to
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

func send(n int64){
		for i:=1 ; int64(i)<=n ; i++ {
			go sendmail()
		}

			fmt.Printf("|||||||||||||||||||||| %d\n",n)

		for j:=1 ; int64(j)<=n ; j++{
			<- wait
		}

}

func main() {
	times1 , _ , _ := os.Time()
	runtime.GOMAXPROCS(4)
	if len(os.Args) == 1 {
		fmt.Println("Don't send mail")
	}else {

	s := os.Args[1]
	n , _ := strconv.Atoi64(s)

	var tt1 int64
	tt1 = 0
	v := 0

		for {
			tt2:=time.Seconds()
				if tt1 < tt2 {
					go send(n)
					v++
					wait <- 1
				}
				if v == 5{
					break
					fmt.Println("tttttttttttttttttttttttttttttttttttttttttt")
				}
			tt1=tt2
		}
	}


		for j:=1 ; int64(j)<=5 ; j++{
			<- wait
		}

	times2 , _ , _ := os.Time()
	tt := times2-times1
		fmt.Printf("Send spam mail = ")
		fmt.Println(count)
		fmt.Printf(".............%d s...........\n",tt)

}
