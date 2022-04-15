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

func fileSystem() http.FileSystem {
	if os.Getenv("ENV") == "dev" {
		log.Info("Development mode")
		return http.Dir("www")
	}

	content, err := fs.Sub(fs.FS(www), "www")
	if err != nil {
		log.Fatal(err)
	}

	return http.FS(content)
}

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

	http.Handle("/", http.FileServer(fileSystem()))

	log.Infof("Listening on 0.0.0.0:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
