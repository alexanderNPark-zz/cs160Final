package main

import (
	"github.com/gorilla/websocket"
	"fmt"
	"net/http"
	"encoding/json"
)

var interactionUpgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
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
			//write to server new text

			interactionWrite(typePerson,"server",packet.Content)
		} else{
			//write to client new links
			interactionWrite(typePerson,"client",packet.Content)
		}




	}

}

func interactionWrite(sender string, oppositeType string, content string) {


	packet:=UpdateMessage{Content:content}
	connection,ok:= oneTimePair[oppositeType]
	if(!ok){
		packet.StatusOfOther = "Dead"
		connectionSender,ok:= oneTimePair[sender]
		if(!ok){
			return
		}
		writer, err := connectionSender.NextWriter(websocket.TextMessage)
		if(err!=nil){
			fmt.Println("Failed 105")
			return
		}
		bytesTosend,err:=json.Marshal(packet)
		writer.Write(bytesTosend)


	}else{
		packet.StatusOfOther = "Alive"
		writer, err := connection.NextWriter(websocket.TextMessage)
		if(err!=nil){
			fmt.Println("Failure in lookup writing")
			return
		}
		bytesTosend,err:=json.Marshal(packet)
		writer.Write(bytesTosend)
	}




}








