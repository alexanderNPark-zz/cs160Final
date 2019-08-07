package main

import (
	"github.com/gorilla/websocket"
	"fmt"
	"net/http"
	"strings"

	"encoding/json"
)

var lookupUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ArchQuery struct{
     name string `id`

}





func LookupServer(w http.ResponseWriter, r *http.Request, ){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } //allow all hosts
	connection,err:=lookupUpgrader.Upgrade(w,r,nil)
	if(err!=nil){
		fmt.Println("Failed to make connection to server")
		fmt.Println(err)
		return
	}

	accept(connection)





}

func accept(connection *websocket.Conn) (resultQuery string) {
	defer connection.Close()



	_, message, err := connection.ReadMessage()

	if(err!=nil){
		fmt.Println("Failure in reading")
		return
	}
	resultQuery=""
	for key := range Architects {
		if(strings.Contains(key,string(message))){
			profile:=Architects[key]
			content,err :=json.Marshal(&ArchQuery{})
			if(err!=nil){
				break;
			}
			resultQuery+=string(content)
		}
	}
	return resultQuery
}

func write(queries string, connection *websocket.Conn){
	defer connection.Close()


	writer, err := connection.NextWriter(websocket.TextMessage)
	if(err!=nil){
		fmt.Println("Failure in writing")
		return
	}

	writer.Write([]byte(queries))

}






