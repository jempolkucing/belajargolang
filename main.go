package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tugas1/module"
)

func main() {
	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// ROute Profuk
	http.HandleFunc("/tugas1/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			module.GetSemuaProduk(w, r)
		} else if r.Method == "POST" {
			module.ProdukBaru(w, r)
		}
	})

	http.HandleFunc("/tugas1/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			module.ProdukbyID(w, r)
		} else if r.Method == "PUT" {
			module.UpdateProduk(w, r)
		} else if r.Method == "DELETE" {
			module.HapusProduk(w, r)
		}
	})

	// Route Kategori
	http.HandleFunc("/tugas1/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			module.GetSemuaKategori(w, r)
		} else if r.Method == "POST" {
			module.KategoriBaru(w, r)
		}
	})

	http.HandleFunc("/tugas1/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			module.KategoribyID(w, r)
		} else if r.Method == "PUT" {
			module.UpdateKategori(w, r)
		} else if r.Method == "DELETE" {
			module.HapusKategori(w, r)
		}
	})

	fmt.Println("Server running di localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
