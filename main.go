package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type PageData struct {
	PageTitle string
	Images    []Image
	Ratio     float64
	Columns   int
}

type Image struct {
	Src string
}

func main() {
	// Get Gallery Name from Flags
	title := flag.String("t", "default", "Name of the Gallery")
	ratio := flag.Float64("r", 1.0, "Ratio of the images")
	port := flag.Int("p", 8888, "Port to use for web server.")
	columns := flag.Int("c", 6, "Number of columns")
	flag.Parse()

	tmpl := template.Must(template.ParseFiles("layout.tmpl"))
	files, err := ioutil.ReadDir("./images")
	var images []Image
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			images = append(images, Image{Src: "/static/" + f.Name()})
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(images), func(i, j int) { images[i], images[j] = images[j], images[i] })
	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: *title,
			Images:    images,
			Ratio:     *ratio,
			Columns:   *columns,
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
