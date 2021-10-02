package main

import (
	"fmt"
	"github.com/Sh4yy/ImageRender/routes"
	"github.com/Sh4yy/ImageRender/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/render/{template}", routes.Render)
	r.HandleFunc("/{template}", routes.Capture)
	r.PathPrefix(utils.Config.Directory.Static).
		Handler(
			http.StripPrefix(utils.Config.Directory.Static,
			http.FileServer(http.Dir("."+utils.Config.Directory.Static))))

	addr := fmt.Sprintf("%s:%d", utils.Config.Server.Host, utils.Config.Server.Port)
	fmt.Printf("Running server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
