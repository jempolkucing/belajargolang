package module

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Kategori struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

var kategori = []Kategori{
	{ID: 1, Nama: "Makanan"},
	{ID: 2, Nama: "Minuman"},
	{ID: 3, Nama: "Sembako"},
}

// GET http://localhost:8080/tugas1/api/categories
func GetSemuaKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}

// POST localhost:8080/tugas1/api/categories
func KategoriBaru(w http.ResponseWriter, r *http.Request) {
	var kategoriBaru Kategori
	err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// masukkin data ke dalam variable categories
	kategoriBaru.ID = len(kategori) + 1
	kategori = append(kategori, kategoriBaru)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kategoriBaru)
}

// POST localhost:8080/tugas1/api/categories/{id}
func KategoribyID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tugas1/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for _, k := range kategori {
		if k.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(k)
			return
		}
	}

	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}

// PUT localhost:8080/tugas1/api/categories/{id}
func UpdateKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tugas1/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	var updatedCategory Kategori
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, k := range kategori {
		if k.ID == id {
			kategori[i].Nama = updatedCategory.Nama

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(kategori[i])
			return
		}
	}

	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}

// DELETE localhost:8080/tugas1/api/categories/{id}
func HapusKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tugas1/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for i, k := range kategori {
		if k.ID == id {
			kategori = append(kategori[:i], kategori[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "sukses delete kategori"})
			return
		}
	}

	http.Error(w, "Kategori tidak ditemukan", http.StatusNotFound)
}
