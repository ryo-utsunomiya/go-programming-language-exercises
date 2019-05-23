package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	prices := make(map[string]dollars)
	prices["shoes"] = 50
	prices["socks"] = 5
	db := database{prices: prices}
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	prices map[string]dollars
	sync.Mutex
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is not set")
		return
	}

	price, err := strconv.Atoi(req.FormValue("price"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price")
		return
	}

	if _, ok := db.prices[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "already exists")
		return
	}

	db.Lock()
	if db.prices == nil {
		db.prices = make(map[string]dollars)
	}
	db.prices[item] = dollars(price)
	db.Unlock()
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is not set")
		return
	}

	if _, ok := db.prices[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found")
		return
	}

	fmt.Fprintf(w, "%s: %d\n", item, db.prices[item])
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is not set")
		return
	}

	price, err := strconv.Atoi(req.FormValue("price"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price")
		return
	}

	if _, ok := db.prices[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found")
		return
	}

	db.Lock()
	db.prices[item] = dollars(price)
	db.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item is not set")
		return
	}

	if _, ok := db.prices[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found")
		return
	}

	db.Lock()
	delete(db.prices, item)
	db.Unlock()
}
