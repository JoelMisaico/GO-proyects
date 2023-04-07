package main

import (
	"crud/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// esta función init carga las variables de entorno
func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error al cargar .env")
	}

	fmt.Println("Variables de entorno cargadas")
}

func main() {
	router := routes.Router()

	http.ListenAndServe(":3000", router)
}
