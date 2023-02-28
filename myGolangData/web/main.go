package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev";
func main() {
	fmt.Println("LCO web request")
    resp,err := http.Get(url)
    
	if(err != nil){
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n",resp)

	defer resp.Body.Close()//caller resp to claose the connection

	// read response from web
	dataByte,err :=ioutil.ReadAll(resp.Body)
 if(err != nil){
	panic(err)
 }
     content := string(dataByte)
	 fmt.Println(content)
}