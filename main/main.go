package main

import (
	"net/http"
	"fmt"
	"strconv"
	"html/template"


)

func main(){
	Initialize()
}







func author(w http.ResponseWriter, r *http.Request, ){



}

func index(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/index.html")
	if(err!=nil){

		fmt.Println(err);
		return
	}
	t.Execute(w,""); //change when template is generated
}




func clientWebServer(w http.ResponseWriter, r *http.Request){
	//private_tmpl_files := []string{"templates/client.html"}
	t, err := template.ParseFiles("templates/client.html")
	if(err!=nil){

		fmt.Println(err);
		return
	}
	t.Execute(w,""); //change when template is generated
}



func multiplexers(handleMultiplex *http.ServeMux){
	handleMultiplex.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	handleMultiplex.HandleFunc("/author", clientWebServer)
	handleMultiplex.HandleFunc("/",index)
	handleMultiplex.HandleFunc("/aserver", StartWebserver)
}

var Multiplex *http.ServeMux


func Initialize(){
	port:= 2323

	handleMultiplex:=http.NewServeMux()
	Multiplex = handleMultiplex
	multiplexers(handleMultiplex)

	server:=http.Server{Addr:"0.0.0.0:"+strconv.Itoa(port), Handler:handleMultiplex,}
	server.ListenAndServe()
}