package main
//this is a comment
import (
	"fmt"
	"net"
)

func main(){
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil{
		fmt.Println("Connection Successful")
	}
}