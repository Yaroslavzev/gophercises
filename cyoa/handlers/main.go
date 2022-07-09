package handlers

import (
	"cyoa"
	"cyoa/templates"
	"html/template"
	"log"
	"net/http"
)

func init() {
	tpl = template.Must(template.New("").Parse(templates.DefaultHandlerTmpl))
}

var tpl *template.Template

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s cyoa.Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn}

	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s      cyoa.Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

func defaultPathFn(r *http.Request) string {
	path := r.URL.Path
	//if path == "" || path == "/" {
	//	path = "/intro"
	//}

	return path[1:]
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	} else {
		http.Redirect(w, r, "/story/intro", http.StatusFound)

		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}
