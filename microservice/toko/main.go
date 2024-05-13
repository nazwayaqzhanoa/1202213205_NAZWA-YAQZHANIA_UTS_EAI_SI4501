package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getDetailToko(w http.ResponseWriter, r *http.Request) {
	// Kita sudah mendapatkan data dari database untuk detail toko
	data := map[string]interface{}{
		"nama":           "NANAS Store",
		"jumlah_product": 10,
	}
	json.NewEncoder(w).Encode(data)
}

func getAllToko(w http.ResponseWriter, r *http.Request) {
	// Kita sudah mendapatkan data dari database untuk semua toko
	data := []map[string]string{
		{
			"nama":           "NANAZ Store",
			"jumlah_product": "10",
		},
		{
			"nama":           "LALA Store",
			"jumlah_product": "10",
		},
		{
			"nama":           "HONEY Store",
			"jumlah_product": "100",
		},
	}
	json.NewEncoder(w).Encode(data)
}

func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/get-detail-toko", getDetailToko)
	mux.HandleFunc("/get-all-toko", getAllToko)

	fmt.Println("Server started at localhost:9000")

	log.Fatal(http.ListenAndServe(":9000", mux))
}
