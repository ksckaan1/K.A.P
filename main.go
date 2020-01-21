package main

import (
	"net/http"

	"github.com/rakyll/statik/fs"
	"github.com/zserge/webview"

	_ "./statik" // TODO: Oluşturulmuş statik.go dosyasının konumu
)

//Versiyon ...
var Versiyon = "0.7"

func main() {
	statikFS, _ := fs.New()

	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	go serverbaslat() //server ayrı olarak kurulsun
	uygulama := webview.New(webview.Settings{
		Title:                  "K.A.P",
		URL:                    "http://localhost:5555",
		Width:                  800,
		Height:                 600,
		Resizable:              false,
		ExternalInvokeCallback: Yakala,
		Debug:                  true,
	})
	defer uygulama.Exit()
	uygulama.SetColor(uint8(23), uint8(23), uint8(35), uint8(255))
	uygulama.Run()
	uygulama.Terminate()

}

func serverbaslat() {
	http.ListenAndServe(":5555", nil)
}
