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
var sumTime float64
// AUTHEN 
var auth = smtp.PlainAuth(
                        "",
                        "",
                        "",
                        "mail2.bangkokranch.co.th",  // smtp protocal //No port
                )

const ( 
  dateStamp="2006-01-02-15:04:05" 
  dateLayout="20060102150405" 
) 

func getTime() { 
 // var ds string 
  //t:=time.Seconds() 
  lt:=time.LocalTime() 
  ts:=lt.Format(dateStamp) 
  //ds="2011-01-28-16:15:00" 
  //tp,_:=time.Parse(dateStamp, ds) 
  //println("lt: ",lt) 
  //print("Unix Epoch: ",t,"\n") 
  print("Start Time: ",ts,"\n") 
}



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
			"tapootumchannel@gmail.com",   // from 
			[]string{"anchalee@bangkokranch.co.th"},  //to
			[]byte(msg),
			)
		if err != nil {
			log.Fatal(err)
			}
		count++
//		fmt.Print(count)
//		fmt.Println("   Send mail")
	// channel
//	wait <- 1

}

func send(n int64,x int){
	timee1 , time1 , _ := os.Time()
		for i:=1 ; int64(i)<=n ; i++ {
			 sendmail()
		}

//			fmt.Printf("|||||||||||||||||||||| %d\n",n)

//		for j:=1 ; j<=10 ; j++{
//			<- wait
//		}

	timee2 , time2 , _ := os.Time()
	time := time2 - time1
	timee := timee2 - timee1
	times := float64(1e9*timee+time)*0.000000001
	fmt.Printf("Thread %d time = %.2f s\n",x,times)

	sumTime = sumTime + float64(times)

	wait <- 1
}

func main() {
	print("\n\n")
	getTime()
	print("\n")
	times1 , timen1 , _ := os.Time()
	runtime.GOMAXPROCS(4)
	if len(os.Args) == 1 {
		fmt.Println("Don't send mail")
	}else {

	s := os.Args[1]
	t , _ := strconv.Atoi64(s)
	y := os.Args[2]
	m , _ := strconv.Atoi64(y)


		for i:=1 ; int64(i)<=t ; i++{
			go send(m,i)
		}

		for i:=1 ; int64(i)<=t ; i++{
			<- wait
		}




	times2 , timen2 , _ := os.Time()
	tt := times2-times1
	nn := timen2-timen1
	time := float64(1e9*tt+nn)*0.000000001
//	mm := 60
		avgTime := sumTime / float64(t)
		fmt.Println("\n==========================================================")
		fmt.Printf("%d Thread\n%d Mail\n",t,m)
		fmt.Printf("Send spam %d mail\n",count)
                fmt.Printf("All time send %.2f s\n",time)
		fmt.Printf("Average thread time = %.2f s\n",avgTime)
		fmt.Printf("Average = %.2f / mail\n",avgTime/float64(t*m))
		fmt.Println("===========================================================")

	}

}
