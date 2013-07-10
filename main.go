package main

import "encoding/json"
import "net/http"
import "log"
import "os"

import "github.com/gorilla/mux"

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/home", IndexHandler).Methods("GET")
	return router
}

func routerHandler(router *mux.Router) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		router.ServeHTTP(res, req)
	}
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'version':'1.0'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func main() {
	handler := routerHandler(router())
	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
