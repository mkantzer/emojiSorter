package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mkantzer/emojiSorter/internal/db"
)

// Structure:
// https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html

type server struct {
	db     *db.NotionDB
	router *mux.Router
}

func (s *server) routes() {
	// s.router.HandleFunc("/healthz", s.handleHealth())
	// s.router.HandleFunc("/elect", s.handleElect())
	s.router.HandleFunc("/", s.handleDummy())
}

// Return the handler:

func (s *server) handleDummy() http.HandlerFunc {
	// thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

////////////////////////////////////////////////////////

// // Arguments for handler-specific deps:

// func (s *server) handleGreeting(format string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, format, "World")
// 	}
// }

////////////////////////////////////////////////////////

// // Middleware functions take an http.HandlerFunc and return a new one that can run code before and/or after calling the original handler — or it can decide not to call the original handler at all.

// func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         if !currentUser(r).IsAdmin {
//             http.NotFound(w, r)
//             return
//         }
//         h(w, r)
//     }
// }

// // The logic inside the handler can optionally decide whether to call the original handler or not — in the example above, if IsAdmin is false, the handler will return an HTTP 404 Not Found and return (abort); notice that the h handler is not called.

// // If IsAdmin is true, execution is passed to the h handler that was passed in.

// // Usually I have middleware listed in the routes.go file:

// package app
// func (s *server) routes() {
//     s.router.HandleFunc("/api/", s.handleAPI())
//     s.router.HandleFunc("/about", s.handleAbout())
//     s.router.HandleFunc("/", s.handleIndex())
//     s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
// }

////////////////////////////////////////////////////////

// // Request and response types can go in there too
// // If an endpoint has its own request and response types, usually they’re only useful for that particular handler.

// // If that’s the case, you can define them inside the function.

// func (s *server) handleSomething() http.HandlerFunc {
//     type request struct {
//         Name string
//     }
//     type response struct {
//         Greeting string `json:"greeting"`
//     }
//     return func(w http.ResponseWriter, r *http.Request) {
//         ...
//     }
// }
// // This declutters your package space and allows you to name these kinds of types the same, instead of having to think up handler-specific versions.

// // In test code, you can just copy the type into your test function and do the same thing. Or… (continued in test file)

////////////////////////////////////////////////

// // sync.Once to setup dependencies
// // If I have to do anything expensive when preparing the handler, I defer it until when that handler is first called.

// // This improves application startup time.

// func (s *server) handleTemplate(files string...) http.HandlerFunc {
//     var (
//         init    sync.Once
//         tpl     *template.Template
//         tplerr  error
//     )
//     return func(w http.ResponseWriter, r *http.Request) {
//         init.Do(func(){
//             tpl, tplerr = template.ParseFiles(files...)
//         })
//         if tplerr != nil {
//             http.Error(w, tplerr.Error(), http.StatusInternalServerError)
//             return
//         }
//         // use tpl
//     }
// }
