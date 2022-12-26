package main

import (
	"fmt"
	"github.com/SchmittyApps/GoArchitecture/dirOne"
	"github.com/SchmittyApps/GoArchitecture/dirOne/dirThree"
	"github.com/SchmittyApps/GoArchitecture/dirTwo"
	"html/template"
	"net/http"
	"path"
	"time"
)

func directory() {
	fmt.Println("Hello, World!")
	dirThree.SayHello()
	dirOne.SayGoodbye()
	dirTwo.SayShowdown()
}

func handle(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := ""
	if r.URL.Path == "/" {
		name = "index.html"
	} else {
		name = path.Base(r.URL.Path)
	}

	data := struct {
		Time time.Time
	}{
		Time: time.Now(),
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func main() {
	fmt.Println("http server up!")
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":0", nil)
}
