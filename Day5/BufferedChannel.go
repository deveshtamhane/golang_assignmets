package main
import (
	"fmt"
	"net/http"
)
func checkLink(ch chan string, link string){
	_,err := http.Get(link)
	if err!=nil{
		//fmt.Println(link,":-  Might Be DOWN")
		ch <- link+"  Might Be DOWN"		
	}else{
	   //fmt.Println(link,":-   Is Up")
	   ch <- link+"   Is Up"
	}
}
func main(){
	links := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.golang.org",
		"https://www.gmail.com",
	}
	//buffered channel
	ch := make(chan string,2)
	for _,link :=  range links{
	   go checkLink(ch,link)
	   fmt.Println(<-ch)
	}
}