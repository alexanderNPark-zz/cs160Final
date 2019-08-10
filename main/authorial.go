package main

type Architect struct{
	Name string
	Stars int
	ImgPath string
	Projects map[string]*Project
}

type Project struct{
   Client *Client
}


var Architects = make(map[string]*Architect)

func LoadArchitects(){
	Architects["Bob"] = &Architect{Name:"Bob", Stars:3, Projects:make(map[string]*Project)}
	Architects["Bob the Builder"] = &Architect{Name:"Bob the Builder", Stars:3, Projects:make(map[string]*Project),ImgPath:"/assets/imgs/bob_builder.png"}
	Architects["Jar-Jar"] = &Architect{Name:"Jar-Jar",Stars:1, Projects:make(map[string]*Project), ImgPath:"/assets/imgs/profile_jar_jar.jpg"}
}

