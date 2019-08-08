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



func search(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("templates/search.html")
	if(err!=nil){

		fmt.Println(err);
		return
	}
	t.Execute(w,""); //change when template is generated
}



func makeClientRequest(w http.ResponseWriter, r *http.Request ){
	/*
	architectNames, provided := r.URL.Query()["architect"]
	//var architectName string;
	if(!provided || len(architectNames) < 1){
		architectName = "bob";
	}else{
		architectName=architectNames[0]
	}

	//profile,_:=Architects[architectName]
*/

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
	handleMultiplex.HandleFunc("/aserver", LookupServer)
	handleMultiplex.HandleFunc("/search", search)
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