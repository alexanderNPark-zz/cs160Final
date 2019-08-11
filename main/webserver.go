package main

import (
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}




var unifyingRead = make(chan string)


//test variable-delete later
var activated bool = false;

/*
func StartWebserver(w http.ResponseWriter, r *http.Request, ){
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } //allow all hosts
	connection,err:=upgrader.Upgrade(w,r,nil)
	if(err!=nil){
		fmt.Println("Failed to make connection to server")
		fmt.Println(err)
		return
	}
	address:=strings.Split(r.Header.Get("X-Forwarded-For"),",")

	fmt.Println("Contact")
    go newClient.read()
    go newClient.write()

}

func (client *Client) read(){
	defer client.connection.Close()

	for{
		fmt.Println("Reading Message")
		_, message, err := client.connection.ReadMessage()
		fmt.Println("read the message")
		if(err!=nil){
			fmt.Println("Failure in reading")
			return
		}

		client.channel<-string(message)
		if(!activated){
			activated = !activated
			testHandleGenerator()
		}
		fmt.Println(string(message))
	}

}

func (client *Client) write(){
	defer client.connection.Close()

	for{
		writer, err := client.connection.NextWriter(websocket.TextMessage)
		if(err!=nil){
			fmt.Println("Failure in writing")
			return
		}
		sendBack:="Received"+<-client.channel
		writer.Write([]byte(sendBack))
	}
}

func testHandleGenerator(){
	beta:=func(w http.ResponseWriter, r *http.Request){
		t, err := template.ParseFiles("templates/workspace.html")
		if(err!=nil){

			fmt.Println(err);
			return
		}
		t.Execute(w,""); //change when template is generated
	}
	Multiplex.HandleFunc("/workspace",beta)
}
*/





