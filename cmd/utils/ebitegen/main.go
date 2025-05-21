package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*.gtpl
var TemplateFS embed.FS

type TemplateData struct {
	WindowTitle  string
	ScreenWidth  int
	ScreenHeight int
}

var path string
var windowTitle string
var screenWidth int
var screenHeight int

func ParseFlags() {
	flag.StringVar(&path, "f", "", "Filepath to generate the file.")
	flag.StringVar(&windowTitle, "t", "Example Title", "Title of the window.")
	flag.IntVar(&screenWidth, "w", 640, "Screen Width.")
	flag.IntVar(&screenHeight, "h", 240, "Screen Height.")
	flag.Parse()
}

func main() {
	ParseFlags()

	if path == "" {
		fmt.Println("Missing required -f flag")
		flag.Usage()
		os.Exit(1)
	}

	data := TemplateData{
		WindowTitle:  windowTitle,
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
	}

	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create directories:", err)
		os.Exit(1)
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	tpl, err := template.ParseFS(TemplateFS, "templates/ebitengine.tmpl")
	if err != nil {
		log.Fatal("parsing template:", err)
		os.Exit(1)
	}

	tpl.Execute(file, data)

	fmt.Println("File created at:", path)
}
