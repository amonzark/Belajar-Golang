package main

import (
	"fmt"
	//"html/template"
	"net/http"
	//"path"
)

//---------------------------------------------------------------------------------------------------------
/*func main() { //simple routing
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerPing)
	http.HandleFunc("/hello", handlerHello)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	var address = ":9000"
	fmt.Printf("server started at %s\n", address)

	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}

	//render html template link --> https://dasarpemrogramangolang.novalagung.com/B-template-render-html.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html") //menghubungkan path folder dengan file
		var tmpl, err = template.ParseFiles(filepath)   //parsing file template
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) //menandai respon error salah satunya 500 atau internal server error
			return
		}

		var data = map[string]interface{}{ //type data interface https://dasarpemrogramangolang.novalagung.com/A-interface.html
			"title": "Learning Golang Web",
			"name":  "Batman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}*/ //-------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------------------
/*func main() { //routing menggunakan parseglob
	var tmpl, errr = template.ParseGlob("views/*")
	if errr != nil {
		panic(errr.Error())
		return
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Arnothz"}
		errr = tmpl.ExecuteTemplate(w, "index", data)
		if errr != nil {
			http.Error(w, errr.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Arnothz"}
		errr = tmpl.ExecuteTemplate(w, "about", data)
		if errr != nil {
			http.Error(w, errr.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":8080", nil)
}*/ //---------------------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------------------
/*func main() { // routing menggunakan parsefile
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}*/ //---------------------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------------------
/*type M map[string]interface{}

func main() { //get post
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("ping"))
		case "GET":
			w.Write([]byte("ging"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}*/ //---------------------------------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------------------------------
func main() { //https://dasarpemrogramangolang.novalagung.com/B-form-value.html
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

//---------------------------------------------------------------------------------------------------------
