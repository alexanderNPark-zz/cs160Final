package main

import (
	"github.com/gorilla/websocket"
	"fmt"
	"net/http"
	"encoding/json"
)

var interactionUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}




var aliveStatus string = "</Alive>"

type PacketMessage struct{
	Content string `json:"content"`
}

type UpdateMessage struct{
	StatusOfOther string `json:"statusOfOther"`
	Content string `json:"content"`
}

var oneTimePair map[string]*websocket.Conn

func InteractionServer(w http.ResponseWriter, r *http.Request, ){
	interactionUpgrader.CheckOrigin = func(r *http.Request) bool { return true } //allow all hosts
	connection,err:=interactionUpgrader.Upgrade(w,r,nil)
	if(err!=nil){
		fmt.Println("Failed to make connection to server")
		fmt.Println(err)
		return
	}

	interactionRead(connection,statusActivate(connection))



}

func statusActivate(connection *websocket.Conn) string{
	_, message, err := connection.ReadMessage()
	if(err!=nil){
		fmt.Println("Failure in lookup reading")
		return ""
	}
	name:=string(message)
	oneTimePair[name] = connection
    return name

}

func interactionRead(connection *websocket.Conn, typePerson string){
	close:=func(){
		delete(oneTimePair,typePerson)
		connection.Close()
	}
	defer close()

	packet:=&PacketMessage{}
	for{

		_, message, err := connection.ReadMessage()
		if(err!=nil){
			fmt.Println("Failure in lookup reading")
			return
		}
		json.Unmarshal(message,packet)
		if(packet.Content==aliveStatus){
			continue;
		}
		if(typePerson=="client"){
			//write to server
		} else{
			//write to client new links
		}




	}

}

func interactionWrite(connection *websocket.Conn, message []byte){

	writer, err := connection.NextWriter(websocket.TextMessage)
	if(err!=nil){
		fmt.Println("Failure in lookup writing")
		return
	}
	writer.Write([]byte(message))

}









