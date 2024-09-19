package main

import (
	"fmt"
	"net/http"

	"github.com/kubewall/kubewall/backend/cmd"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type fwdHandler struct {
	format string
}

func (fw fwdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<meta http-equiv="refresh" content="0;url=http://localhost:7080/" />`)
}

func main() {
	go func() {
		cmd.Execute()
	}()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "kubewall",
		Width:  1200,
		Height: 1000,
		AssetServer: &assetserver.Options{
			Handler: fwdHandler{},
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
