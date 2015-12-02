package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

func viewDayTotals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(getDayTotals()); err != nil {
		panic(err)
	}
}

func viewCategoryTotals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(getCategoryTotals()); err != nil {
		panic(err)
	}
}
func addExpense(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var expense Expense
	err = json.Unmarshal(body, &expense)

	if err != nil {
		panic(err)
	}

	if expense.Amount <= 0 {
		http.Error(w, "Invalid Amount", http.StatusBadRequest)
		return
	}
	if err != nil {
		panic(err)
	}

	if len(expense.Date.Format(time.RFC3339)) == 0 {
		http.Error(w, "Invalid Date", http.StatusBadRequest)
		return
	}

	if expense.Category == "" {
		http.Error(w, "Invalid Category", http.StatusBadRequest)
		return
	}

	err = expense.Write()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type EMServer struct {
	router *mux.Router
}

func (server *EMServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", `POST, GET, OPTIONS,
        	PUT, DELETE`)
		w.Header().Set("Access-Control-Allow-Headers",
			`Accept, Content-Type, Content-Length, Accept-Encoding,
            X-CSRF-Token, Authorization`)
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	server.router.ServeHTTP(w, r)
}

func AddHandlers(router *mux.Router) {
	router.HandleFunc("/totals/day", viewDayTotals).Methods("GET")
	router.HandleFunc("/totals/category", viewCategoryTotals).Methods("GET")
	router.HandleFunc("/expense", addExpense).Methods("POST")
}
