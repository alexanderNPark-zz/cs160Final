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
     Name string `json:"name"`
     Stars int   `json:"stars"`
     Path string  `json:"path"`
}


var delimter string = "</Delim>"


func LookupServer(w http.ResponseWriter, r *http.Request, ){
	lookupUpgrader.CheckOrigin = func(r *http.Request) bool { return true } //allow all hosts
	connection,err:=lookupUpgrader.Upgrade(w,r,nil)
	if(err!=nil){
		fmt.Println("Failed to make connection to server")
		fmt.Println(err)
		return
	}

	response:=make(chan string)
	go readUntilClose(connection, response)
	go writeUntilClose(connection,response)



}

func readUntilClose(connection *websocket.Conn, response chan string){
	defer connection.Close()

	for{

		_, message, err := connection.ReadMessage()
		querySearch:=string(message)
		if(querySearch=="<Finished>"){
			break;
		}

		if(err!=nil){
			fmt.Println("Failure in lookup reading")
			return
		}
		querySearch = strings.ToLower(querySearch)
		resultQuery:=""
		for key := range Architects {
			lowerKey:= strings.ToLower(key)
			if(strings.Contains(lowerKey,querySearch)){
				profile:=Architects[key]
				fmt.Println(profile)
				content,err :=json.Marshal(ArchQuery{Name:profile.Name, Stars:profile.Stars, Path:"/fake"})
				if(err!=nil){
					break;
				}
				fmt.Println("Found Query:",string(content))
				resultQuery+=string(content)+delimter
			}
		}
		if(resultQuery==""){
			resultQuery = "null"
		}
		response<-resultQuery
	}

}

func writeUntilClose(connection *websocket.Conn, response chan string){
	defer connection.Close()

	for{
		writer, err := connection.NextWriter(websocket.TextMessage)
		if(err!=nil){
			fmt.Println("Failure in lookup writing")
			return
		}

		writer.Write([]byte(<-response))
	}
}









