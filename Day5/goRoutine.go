package main
import (
	"fmt"
	"net/http"
	"time"
)
func checkLink(link string){
	_,err := http.Get(link)
	if err!=nil{
		fmt.Println(link,":-  Might Be DOWN")
		
	}else{
	   fmt.Println(link,":-   Is Up")
	}
}
func main(){
	links := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.golang.org",
		"https://www.gmail.com",
	}
	for _,link :=  range links{
	   go checkLink(link)
	}
    time.Sleep(10*time.Second)
}