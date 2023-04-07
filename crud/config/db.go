package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //el _ es que voy a importar este paquete
)

// Es la conexion a la base de datos
func Conexion() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_CONECTION"))

	if err != nil {
		log.Fatal("DB error...")
	}

	fmt.Println("DB conection...")

	return db
}
