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

	query:=accept(connection)


	if(query=="") {
		query="null"
	}
	fmt.Println(query)
	writeIt(query, connection)





}

func accept(connection *websocket.Conn) (resultQuery string) {

	_, message, err := connection.ReadMessage()
	fmt.Println(string(message))
	if(err!=nil){
		fmt.Println(err)
		return
	}
	resultQuery=""
	for key := range Architects {
		if(strings.Contains(key,string(message))){
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
	return resultQuery
}

func writeIt(queries string, connection *websocket.Conn){



	writer, err := connection.NextWriter(websocket.TextMessage)
	if(err!=nil){
		fmt.Println(err)
		return
	}

	writer.Write([]byte(queries))

}







