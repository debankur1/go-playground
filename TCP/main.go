package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)
func main(){
	srv,err := net.Listen("tcp",":8081")
	if err!=nil{
		log.Panic("")
	}
	for {
		conn,err :=srv.Accept()
		if err!=nil{
			log.Panic("error")
		}
		io.WriteString(conn,"Hello from Golang")
		go HandleData(conn)
	}
}
func HandleData(con net.Conn){
	data:= bufio.NewScanner(con)
	for data.Scan(){
		fmt.Println(data.Text())
	}
	defer  con.Close()
}
