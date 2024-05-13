package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Merchant struct {
	Nama          string `json:"nama"`
	NamaToko      string `json:"nama_toko"`
	JumlahProduct string `json:"jumlah_product"`
}
type Toko struct {
	NamaToko      string `json:"nama"`
	JumlahProduct string `json:"jumlah_product"`
}
func merchantMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		autor := r.Header.Get("Authorization")

		if autor != "merchant" {
			http.Error(w, "You are not authorized to access this resource", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func superMiddle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		autor := r.Header.Get("Authorization")

		if autor != "su-admin" {
			http.Error(w, "You are not authorized to access this resource", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func getMerchant(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:9000/get-merchant")
	if err != nil {
		http.Error(w, "Failed to fetch merchant data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read merchant data", http.StatusInternalServerError)
		return
	}

	var merch Merchant
	err = json.Unmarshal(data, &merch)
	if err != nil {
		http.Error(w, "Failed to unmarshal merchant data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(merch)
}

func getAllToko(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:9000/get-all-toko")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/get-merchant", merchantMiddle(getMerchant))
	mux.HandleFunc("/get-all-toko", superMiddle(getAllToko))

	fmt.Println("Server running on port 6000")
	http.ListenAndServe(":6000", mux)
}
