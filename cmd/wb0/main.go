package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zykunov/courseGoFirst/WB0/internal/app/back"
	"github.com/zykunov/courseGoFirst/WB0/storage"
)

func main() {

	DBconnection := storage.GetDB()
	defer DBconnection.Close()

	cache := back.GetCache(DBconnection)

	c := back.CacheStruct{Cache: cache}

	r := mux.NewRouter()
	r.HandleFunc("/items/{key}", c.ItemPage)
	r.HandleFunc("/", c.AllItemsPage)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)

}
