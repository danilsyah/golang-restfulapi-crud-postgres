package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	"strconv"

	// package yang digunakan untuk mengubah string menjadi tipe int
	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"go-postgres-crud/models" //models package dimana Buku didefinisikan

	// digunakan untuk mendapatkan parameter dari router
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // postgres golang driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Buku `json:"data"`
}

// Tambah Buku
func TambahBuku(w http.ResponseWriter, r *http.Request) {
	// create an empty user of type models.User
	// kita buat empty buku dengan tipe models.Buku
	var buku models.Buku

	// decode data json request ke buku
	err := json.NewDecoder(r.Body).Decode(&buku)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body. %v", err)
	}

	// panggil modelsnya lalu insert buku
	insertID := models.TambahBuku(buku)

	// format response objectnya
	res := response{
		ID:      insertID,
		Message: "Data buku telah ditambahkan",
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

// Ambil Semua data buku
func AmbilSemuaBuku(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// memanggil models AmbilSemuaBuku
	bukus, err := models.AmbilSemuaBuku()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = 200
	response.Message = "Success"
	response.Data = bukus

	// kirim semua response
	json.NewEncoder(w).Encode(response)
}

// AmbilBuku mengambil single data dengan paremeter id
func AmbilBuku(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// dapatkan idbuku dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari string ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	// memanggil models AmbilSatuBuku dengan parameter id yang nantinya akan mengambil single data
	buku, err := models.AmbilSatuBuku(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data buku. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(buku)
}

func UpdateBuku(w http.ResponseWriter, r *http.Request) {
	// kita ambil request parameter id nya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	// buat variabel buku dengan type models.Buku
	var buku models.Buku

	// decode json request ke variabel buku
	err = json.NewDecoder(r.Body).Decode(&buku)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body. %v", err)
	}

	// panggil UpdateBuku untuk mengupdate data
	updateRows := models.UpdateBuku(int64(id), buku)

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Buku telah berhasil diupdate, jumlah yang di update %v rows/record", updateRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

func HapusBuku(w http.ResponseWriter, r *http.Request) {
	// kita ambil request parameter id nya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int. %v", err)
	}

	// panggil fungsi HapusBuku, dan convert int ke int64
	deletedRows := models.HapusBuku(int64(id))

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Buku sukses dihapus. Total data yang dihapus %v", deletedRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send response
	json.NewEncoder(w).Encode(res)
}
