package main

import (
	"log"
	"net/http"

	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/bindata"
	"gopkg.in/macaron.v1"
	"github.com/aqrun/gotest/web"
)

func main(){

	m := macaron.Classic()
	m.Use(macaron.Static("assets",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset: web.Asset,
				AssetDir: web.AssetDir,
				AssetNames: web.AssetNames,
				Prefix: "",
			}),
		}))

	m.Use(pongo2.Pongoer(pongo2.Options{
		TemplateFileSystem: bindata.Templates(bindata.Options{
			Asset:web.Asset,
			AssetDir: web.AssetDir,
			AssetNames: web.AssetNames,
			Prefix: "",
		}),
		Directory: "templates",
		Extensions: []string{".gohtml",".tmpl", ".html"},
	}))





	m.Get("/", myHandler)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe(":4000", m))

	m.Run()
}

func myHandler(ctx *macaron.Context) {
	ctx.Data["name"] = "Alex"
	ctx.HTML(200, "templates/index")
}