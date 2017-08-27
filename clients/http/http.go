package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"

	"github.com/kryptn/modulario/data"
	"github.com/kryptn/modulario/proto"
	"time"
)

func Handle() {

	app := buildHandler()

	srv := &http.Server{
		Handler: app.router,
		Addr: ":5000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
	}

	log.Printf("Serving on localhost")
	log.Fatal(srv.ListenAndServe())
}

func buildHandler() *App {

	app := App{}
	app.engine = data.Engine{Dialect: "sqlite3", Args: "/tmp/gorm.db", LogMode: true}
	app.engine.InitDB()
	app.engine.InitSchema()

	app.router = mux.NewRouter()
	app.router.HandleFunc("/", index).Methods("GET")
	app.router.HandleFunc(`/api/v1/{key:[a-zA-Z0-9]{5,12}}`, app.GetPost(app.getPost)).Methods("GET").Name("post")
	app.router.HandleFunc(`/api/v1/view/{key:[a-zA-Z0-9]{5,12}}`, app.GetPost(app.viewPost)).Methods("GET").Name("view_post")
	app.router.HandleFunc(`/api/v1/delete/{key:[a-zA-Z0-9]{5,12}}`, app.DeletePost).Methods("GET").Name("delete_post")
	app.router.HandleFunc("/api/v1/create/", app.CreatePost).Methods("POST")

	return &app
}

type App struct {
	engine data.Engine
	router *mux.Router
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:3000/", http.StatusTemporaryRedirect)
}

func (app *App) GetPost(handler func(post *data.Post, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]
		post, err := app.engine.GetPostLinks(key)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Post not found: %s", err)
			return
		}
		handler(post, w, r)
	}

}

func (app *App) viewPost(post *data.Post, w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while marshalling: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (app *App) getPost(post *data.Post, w http.ResponseWriter, r *http.Request) {
	link, err := app.engine.VisitPost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while redirecting: %s", err)
		return
	}

	http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
}

func (app *App) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req proto.JsonCreateRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Shits broke up in CreatePost Decoding: ", err)
	}

	spew.Dump(req)

	// todo: this is where to inject users
	user := data.User{}
	user.ID = 0

	post, err := app.engine.CreatePost(user, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while creating post: %s", err)
	}

	path, _ := app.router.Get("view_post").URL("key", post.Key)

	http.Redirect(w, r, path.String(), http.StatusTemporaryRedirect)
}

func (app *App) DeletePost(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	key := args["key"]
	err := app.engine.DeletePost(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while deleting post: %s", err)
	}
}
