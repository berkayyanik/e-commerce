package database

import (
	"database/sql"
	"fmt"
	"os"
)

func SetupPsql(host string, port int, user string, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigration(db *sql.DB) error {
	productData, err := os.ReadFile("internal/database/migrations/create_product_table.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(productData))
	if err != nil {
		return err
	}

	clothData, err := os.ReadFile("internal/database/migrations/create_cloth_table.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(clothData))
	if err != nil {
		return err
	}

	insertData, err := os.ReadFile("internal/database/migrations/insert_product_table.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(insertData))
	if err != nil {
		return err
	}

	clothesData, err := os.ReadFile("internal/database/migrations/insert_cloth_table.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(clothesData))
	if err != nil {
		return err
	}

	return nil
}
