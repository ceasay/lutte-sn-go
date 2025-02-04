package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const videoDir = "video"
const port = 8080
const videExtenstion = "m3u8"

func main() {
	// configure the songs directory name and port

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h1>Video Streaming Server</h1>")
	})

	http.HandleFunc("/videos", func(w http.ResponseWriter, r *http.Request) {
		files, err := os.ReadDir("output")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h1>Liste des vidéos disponibles </h1> <br> "+getListOfVideos())
	})

	// add a handler for the song files
	http.Handle("/", addHeaders(http.FileServer(http.Dir(videoDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", videoDir, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func getListOfVideos() string {
	files, err := os.ReadDir(videoDir)
	if err != nil {
		log.Fatal(err)
	}
	var html string

	var host = "http://localhost:" + fmt.Sprintf("%d", port) + "/"

	for _, file := range files {
		var name = file.Name()

		if file.IsDir() {
			videoFiles, err := os.ReadDir(videoDir + "/" + name)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range videoFiles {
				if filepath.Ext(file.Name()) != "."+videExtenstion {
					continue
				}
				html += "<video width='320' height='240' controls>"
				html += "<source src='" + host + name + "/" + file.Name() + "' type='video/' + " + videExtenstion + ">"
				html += "</video> <br>"
			}
		}
	}

	return html
}
