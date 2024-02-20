package models

import (
	"database/sql"
	"fmt"
	"go-postgres-crud/config"
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

// Buku schema dari tabel Buku
// kita coba dengan jika datanya null
// jika return datanya ada yang null, silahkan pake NullString, contoh nya dibawah
// Penulis - config.NullString `json:"penulis"`
type Buku struct {
	ID            int64  `json:"id"`
	Judul_buku    string `json:"judul_buku"`
	Penulis       string `json:"penulis"`
	Tgl_publikasi string `json:"tgl_publikasi"`
}

// tambahkan data buku baru
func TambahBuku(buku Buku) int64 {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya diakhir proses
	defer db.Close()

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari buku yang dimasukan ke db
	sqlStatement := `INSERT INTO buku (judul_buku, penulis, tgl_publikasi) VALUES ($1, $2, $3) RETURNING id`

	// id yang dimasukan akan disimpan di id ini
	var id int64

	// scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, buku.Judul_buku, buku.Penulis, buku.Tgl_publikasi).Scan(&id)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query insert. %v", err)
	}

	fmt.Printf("Insert data single successfully %v", id)

	return id
}

// ambil semua buku
func AmbilSemuaBuku() ([]Buku, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya diakhir proses
	defer db.Close()

	var bukus []Buku

	// kita buat select query
	sqlStatement := `SELECT * FROM buku`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query %v", err)
	}

	// kita tutup eksekusi proses sql querynya
	defer rows.Close()

	// kita literasi mengambil datanya
	for rows.Next() {
		var buku Buku

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&buku.ID, &buku.Judul_buku, &buku.Penulis, &buku.Tgl_publikasi)

		if err != nil {
			log.Fatalf("Tidak bisa mengambil data. %v", err)
		}

		// masukan kedalam slice bukus
		bukus = append(bukus, buku)
	}

	// return empty buku atau jika error
	return bukus, err
}

// mengambil satu data buku by id
func AmbilSatuBuku(id int64) (Buku, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var buku Buku

	// buat sql query
	sqlStatement := `SELECT * FROM buku WHERE id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&buku.ID, &buku.Judul_buku, &buku.Penulis, &buku.Tgl_publikasi)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return buku, nil
	case nil:
		return buku, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return buku, err
}

// update buku by id
func UpdateBuku(id int64, buku Buku) int64 {
	// mengkoneksikan ke DB
	db := config.CreateConnection()

	// kitu tutup koneksinya di akhir proses
	defer db.Close()

	// query update
	sqlStatement := `UPDATE buku SET judul_buku=$2, penulis=$3, tgl_publikasi=$4 WHERE id=$1`

	// eksekusi sql query update
	res, err := db.Exec(sqlStatement, id, buku.Judul_buku, buku.Penulis, buku.Tgl_publikasi)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa banyak row/data yang di update
	rowsAffected, err := res.RowsAffected()

	// kita cek
	if err != nil {
		log.Fatalf("Error ketika mengechek rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang di update %v\n", rowsAffected)

	return rowsAffected
}

// hapus data buku by id
func HapusBuku(id int64) int64 {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksi db di akhir proses
	defer db.Close()

	// buat query sql
	sqlStatement := `DELETE FROM buku WHERE id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa jumlah row/data yang dihapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang dihapus %v", rowsAffected)

	return rowsAffected
}
