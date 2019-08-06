package main

import (
	"github.com/gorilla/websocket"
	"fmt"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type ExJSON struct{

}


type Client struct{
	connection *websocket.Conn
	ID string

	channel chan string

}

var userIDs= make(map[string]*Client)
var unifyingRead = make(chan string)

func StartWebserver(w http.ResponseWriter, r *http.Request, ){
	connection,err:=upgrader.Upgrade(w,r,nil)
	if(err!=nil){
		fmt.Println("Failed to make connection to server")
		fmt.Println(err)
		return
	}
	address:=strings.Split(r.Header.Get("X-Forwarded-For"),",")
	newClient := &Client{connection:connection, ID:address[0]}
    go newClient.read()
    go newClient.write()

}

func (client *Client) read(){
	defer client.connection.Close()

	for{
		_, message, err := client.connection.ReadMessage()
		if(err!=nil){
			fmt.Println("Failure")
			return
		}
		client.channel<-string(message)
		fmt.Println(string(message))
	}

}

func (client *Client) write(){
	defer client.connection.Close()

	for{
		writer, err := client.connection.NextWriter(websocket.TextMessage)
		if(err!=nil){
			return
		}
		sendBack:="Received"+<-client.channel
		writer.Write([]byte(sendBack))
	}
}






