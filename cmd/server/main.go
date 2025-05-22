package main

import (
	"log"
	"net/http"

	"github.com/Kalebhawkins/natureofgo/web"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := web.Must(web.ParseTemplate(web.TemplateFS, "templates/index.gtpl"))

		tmpl.Execute(w, data)
	})

	r.Get("/run", func(w http.ResponseWriter, r *http.Request) {
		tmpl := web.Must(web.ParseTemplate(web.TemplateFS, "templates/wasm_runner.gtpl"))

		wasmFile := r.URL.Query().Get("wasm")
		if wasmFile == "" {
			http.Error(w, "Missing wasm file", http.StatusBadRequest)
			return
		}

		tmpl.Execute(w, struct{ WasmFile string }{WasmFile: wasmFile})
	})

	log.Println("Started listening on http://127.0.0.1:3000")
	http.ListenAndServe("127.0.0.1:3000", r)
}
