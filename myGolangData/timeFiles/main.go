package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("itd time to learn about time in go lang")
    
	// time now
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
    
	 createdDate := time.Date(2021,time.August,11,10,10,10,10,time.UTC)
     fmt.Println(createdDate)
	 // formating
	 fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))




}