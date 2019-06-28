package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//t, _:= template.ParseFiles("login.html")
		t, _ := template.New("foo").Parse(`{{define "T"}}{{.}}!{{end}}`)
		t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
		//t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println(template.HTMLEscaper(r.Form.Get("username"), r.Form.Get("password")))
	}
}

func main()  {
	http.HandleFunc("/sayHello", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
