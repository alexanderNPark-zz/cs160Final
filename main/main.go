package main

import (
	"net/http"
	"fmt"
	"strconv"
	"os"
	"path/filepath"
	"html/template"
)

func main(){
	Initialize()
}


var active bool = false
var data string = "255712391091861932312181243211672021836999"
var location string




func author(w http.ResponseWriter, r *http.Request, ){


}





func index(w http.ResponseWriter, r *http.Request, ){
	//private_tmpl_files := []string{"templates/index.html"}
	t, err := template.ParseFiles("templates/index.html")
	if(err!=nil){
		fmt.Println(err);
		return
	}
	t.Execute(w,""); //change when template is generated
}


func Initialize(){
	port:= 2323
	ex,_:=os.Executable()
	location = filepath.Dir(ex)+"\\"



	handleMultiplex:=http.NewServeMux()

	
	handleMultiplex.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))




	handleMultiplex.HandleFunc("/", index)
	server:=http.Server{Addr:"0.0.0.0:"+strconv.Itoa(port), Handler:handleMultiplex,}
	server.ListenAndServe()
}