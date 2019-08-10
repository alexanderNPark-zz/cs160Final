package main

import (
	"strings"
	"fmt"
	"encoding/json"
)

type Architect struct{
	Name string
	Stars int
	ImgPath string
	Projects map[string]*Project
	KeyWordTags []string
}

type Project struct{
   Client *Client
}


var Architects = make(map[string]*Architect)

func LoadArchitects(){
	Architects["Ted Mosby"] = &Architect{Name:"Ted Mosby", Stars:3, Projects:make(map[string]*Project), ImgPath:"/assets/imgs/ted_mosby.jpg"}
	Architects["Bob the Builder"] = &Architect{Name:"Bob the Builder", Stars:3, Projects:make(map[string]*Project),ImgPath:"/assets/imgs/bob_builder.png"}
	Architects["Bob Parr"] = &Architect{Name:"Bob Par", Stars:3, Projects:make(map[string]*Project),ImgPath:"/assets/imgs/Bob_Parr.jpg"}
	Architects["Jar-Jar"] = &Architect{Name:"Jar-Jar",Stars:1, Projects:make(map[string]*Project), ImgPath:"/assets/imgs/profile_jar_jar.jpg"}
}

func SearchQuery(querySearch string) string{
	querySearch = strings.ToLower(querySearch)
	resultQuery:=""
	for key := range Architects {
		lowerKey:= strings.ToLower(key)
		if(strings.HasPrefix(lowerKey,querySearch)){
			profile:=Architects[key]
			fmt.Println(profile)
			content,err :=json.Marshal(ArchQuery{Name:profile.Name, Stars:profile.Stars, Path:profile.ImgPath})
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
	return resultQuery
}