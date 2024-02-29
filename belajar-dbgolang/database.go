package belajardbgolang

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golangdb?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  //Pengaturan berapa jumlah koneksi minimal yang dibuat
	db.SetMaxOpenConns(100)                 //Pengaturan berapa jumlah koneksi maksimal yang dibuat
	db.SetConnMaxIdleTime(5 * time.Minute)  // pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) //pengaturan berapa lama koneksi yang boleh digunakan

	return db
}
