package belajardbgolang

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	queryScript := "INSERT INTO customer (id, name) VALUES('asep', 'ASEP')"

	_, err := db.ExecContext(ctx, queryScript)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

//kalau ExecContext itu biasanya digunakan untuk INSERT, UPDATE, DELETE, tapi kalau Query kyk SELECT itu lain lagi

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	queryScript := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, queryScript)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	queryScript := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"

	rows, err := db.QueryContext(ctx, queryScript)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString //kalau misalkan ada kolom yang Nullable
		var balance int32
		var rating float64
		var birth_date time.Time
		var created_at time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)

		if err != nil {
			panic(err)
		}

		fmt.Println("==============================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("birth_date:", birth_date)
		fmt.Println("rating:", rating)
		fmt.Println("married:", married)
		fmt.Println("created_at:", created_at)
	}

	fmt.Println("Success insert new customer")
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "asep'; #" //awas sql injection
	password := "kokok"

	script := "SELECT username FROM users WHERE username ='" + username + "' AND password = '" + password + "' LIMIT 1"

	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "asep'; #" //awas sql injection
	password := "kokok"

	script := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1" //ini sementara aman

	fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	username := "Konto"
	password := "Konto"

	queryScript := "INSERT INTO users (username, password) VALUES(?, ?)" // ini sementara aman

	_, err := db.ExecContext(ctx, queryScript, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	email := "Konto@gmail.com"
	comment := "pesekkk aowkaowkaowkawokawokawokawok"

	queryScript := "INSERT INTO comments (email, comment) VALUES(?, ?)" // ini sementara aman

	result, err := db.ExecContext(ctx, queryScript, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert comment with id", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments (email, comment) VALUES(?, ?)" // ini sementara aman
	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "asepwahyudi" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID ", id)
	}

} // kalau butuh melakukan query yang sama berkali dengan parameter yang berbeda lebih baik menggunakan prepare statement

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments (email, comment) VALUES(?, ?)" // ini sementara aman
	//do transaction
	for i := 0; i < 10; i++ {
		email := "kokochina" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment ID ", id)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
