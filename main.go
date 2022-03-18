package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

//go:embed www
var www embed.FS

func main() {
	port := "8080"
	if wwwPort := os.Getenv("WWW_PORT"); len(wwwPort) > 0 {
		port = wwwPort
	}

	wsUrl := fmt.Sprintf("window.wsUrl = '%s';", os.Getenv("WS_URL"))
	http.Handle("/js/config.js",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(wsUrl))
		}))

	content, err := fs.Sub(fs.FS(www), "www")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(content)))

	log.Infof("Listening on 0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
