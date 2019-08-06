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


func indexHandler(w http.ResponseWriter, r *http.Request, ){
	fmt.Fprintln(w, "<!DOCTYPE html>")
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "<input type=\"password\" name =\"pass\" </input>")
	fmt.Fprintln(w, "<a href =\"initAuth\"> Authorize </a>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}

func author(w http.ResponseWriter, r *http.Request, ){
	if(active){
		fmt.Fprintln(w, "<!DOCTYPE html>")
		fmt.Fprintln(w, "<html>")
		fmt.Fprintln(w, "<body>")
		fmt.Fprintln(w, "Already Authorized")
		fmt.Fprintln(w, "<a href =\"invoke\"> Invoke </a>")
		fmt.Fprintln(w, "</body>")
		fmt.Fprintln(w, "</html>")
		return
	}

	r.ParseForm()

	combine:=""
	length:=16
	i:=0

	for i!=length{
		combine+=r.FormValue("h"+strconv.Itoa(i)+"h")
		i++
	}

	if(combine==data){
		fmt.Fprintf(w, "Authorized")
		active = true

	}
	browPass :=r.FormValue("pass")
	if(!active){
		if(browPass=="JibunWoo"){
			fmt.Fprintf(w, "Authorized")
			active = true

		} else{
			fmt.Fprintf(w, "Denied")
		}
	}


}





func true_index(w http.ResponseWriter, r *http.Request, ){
	//private_tmpl_files := []string{"templates/index.html"}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w,""); //change when template is generated
}


func Initialize(){
	port:= 2323
	ex,_:=os.Executable()
	location = filepath.Dir(ex)+"\\"



	handleMultiplex:=http.NewServeMux()

	
	handleMultiplex.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))




	handleMultiplex.HandleFunc("/", true_index)
	server:=http.Server{Addr:"0.0.0.0:"+strconv.Itoa(port), Handler:handleMultiplex,}
	server.ListenAndServe()
}