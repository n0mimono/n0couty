package application

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"n0couty/app/config"
	"net/http"
	"strconv"
)

func (app *App) Listen() {
	// static
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("tsx/dist/assets"))))

	// index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filename := "index.html"

		tmpl, err := template.
			New(filename).
			Delims("[[", "]]").
			ParseFiles("tsx/dist/" + filename)
		if err != nil {
			log.Fatal(err)
		}

		tmpl.Execute(w, nil)
	})

	// api
	http.Handle("/api/crawl", &CrawlHandler{app: app})
	http.Handle("/api/users/list", &UsersHandler{app: app})
	http.Handle("/api/users", &DetailHandler{app: app})
	http.Handle("/api/users/items", &UserItemsHandler{app: app})
	http.Handle("/api/users/scout", &UserScoutHandler{app: app})

	// api for ml
	to := "http://" + config.ML_HOST + ":" + config.ML_PORT
	http.HandleFunc("/api/ml/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			resp := proxyApiGet(to+r.RequestURI, r.Method)
			fmt.Fprintf(w, resp)
			return
		case "POST":
			resp := proxyApiPost(to+r.RequestURI, r.Method, r.Body)
			fmt.Fprintf(w, resp)
			return
		}
	})

	// websocket
	notificator := NewNotifyHandler()
	http.Handle("/socket/", notificator)
	go notificator.Run()

	// register
	app.Channel.Crawl.Forward = notificator.Forward

	// http
	port := config.PORT
	http.ListenAndServe(":"+port, nil)
}

type CrawlHandler struct {
	app *App
}

func (hndl *CrawlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data, _ := json.Marshal(hndl.app.State.Crawl)
		fmt.Fprintf(w, string(data))
		return
	case "POST":
		hndl.app.UcStartCrawl(InStartCrawl{Start: true})

		data, _ := json.Marshal(hndl.app.State.Crawl)
		fmt.Fprintf(w, string(data))
		return
	case "PUT":
		hndl.app.UcStartCrawl(InStartCrawl{Start: false})

		data, _ := json.Marshal(hndl.app.State.Crawl)
		fmt.Fprintf(w, string(data))
		return
	}
}

type UsersHandler struct {
	app *App
}

func (hndl *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qToN := func(q string, max int, def int) int {
		n, err := strconv.Atoi(q)
		if n < 1 || err != nil {
			return def
		}
		if max > 0 && n >= max {
			return max
		}
		return n
	}
	qToB := func(q string, def bool) bool {
		b, err := strconv.ParseBool(q)
		if err != nil {
			return def
		}
		return b
	}

	switch r.Method {
	case "GET":
		page := r.URL.Query().Get("page")
		perPage := r.URL.Query().Get("perpage")
		onlyStarred := r.URL.Query().Get("onlyStarred")
		out := hndl.app.UcGetUsers(InGetUsers{
			Page:        qToN(page, -1, 1),
			PerPage:     qToN(perPage, 20, 20),
			OnlyStarred: qToB(onlyStarred, false),
		})

		data, _ := json.Marshal(out)
		fmt.Fprintf(w, string(data))
		return
	}
}

type DetailHandler struct {
	app *App
}

func (hndl *DetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			fmt.Fprintf(w, `{ error: "Error" }`)
			return
		}

		out := hndl.app.UcGetDetail(InGetDetail{Id: uint(id)})
		data, _ := json.Marshal(out)
		fmt.Fprintf(w, string(data))

		return
	}
}

type UserItemsHandler struct {
	app *App
}

func (hndl *UserItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			fmt.Fprintf(w, `{ error: "Error" }`)
			return
		}

		out := hndl.app.UcGetUserItems(InGetUserItems{Id: uint(id)})
		data, _ := json.Marshal(out)
		fmt.Fprintf(w, string(data))

		return
	}
}

type UserScoutHandler struct {
	app *App
}

func (hndl *UserScoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			fmt.Fprintf(w, `{ error: "Error" }`)
			return
		}
		star, err := strconv.ParseBool(r.URL.Query().Get("star"))
		if err != nil {
			fmt.Fprintf(w, `{ error: "Error" }`)
			return
		}

		out := hndl.app.UcUpdateUserScout(InUpdateUserScout{
			Id:   uint(id),
			Star: star,
		})
		data, _ := json.Marshal(out)
		fmt.Fprintf(w, string(data))

		return
	}
}
