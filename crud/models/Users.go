package models

import (
	"crud/config"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Id       int
	Nombre   string
	Email    string
	Password string
}

// creamos una funcion para hashear nuestra constraseña
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 es la complejidad del algoritmo, 14 es lo recomendable

	return string(bytes), err
}

// Creamos una funcion que verifique las contranseñas
func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil //si el error retorna true siginifica que la contraseña coinside y si es false es que no coinciden las contraseñas
}

// estos modelos interactuan con la base de datos
func CreateUser(nombre string, email string, password string) {
	conexion := config.Conexion()

	query, err := conexion.Prepare("INSERT INTO usuario (nombre, email, password, activo) VALUE (?, ?, ?, ?)")

	if err != nil {
		panic("Error al crear usuario")
	}

	securePassword, err := HashPassword(password)
	if err != nil {
		panic("Error al cifrar contraseña")
	}

	//El exc me pide los argumentos que necesito
	query.Exec(nombre, email, securePassword, 1)

	conexion.Close()
}

func DeleteUser(id string) {
	conexion := config.Conexion()

	query, err := conexion.Prepare("UPDATE usuario SET activo = ? WHERE id = ?")

	if err != nil {
		panic("Error al eliminar ususario")
	}

	query.Exec(0, id)

	conexion.Close()
}

func UpdateUser(id, nombre, email, password string) {
	conexion := config.Conexion()

	query, err := conexion.Prepare("UPDATE usuario SET nombre = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		panic("Error al actualizar ususario")
	}

	securePassword, err := HashPassword(password)
	if err != nil {
		panic("Error al cifrar la contraseña")
	}

	query.Exec(nombre, email, securePassword, id)

	conexion.Close()
}

func ReadUser(id string) Usuario {
	conexion := config.Conexion()

	row := conexion.QueryRow("SELECT id, nombre, email, password FROM usuario WHERE activo = 1 AND id = ?", id)

	conexion.Close()

	var usuario Usuario

	row.Scan(&usuario.Id, &usuario.Nombre, &usuario.Email, &usuario.Password)

	return usuario
}

func ReadUsers() []Usuario {
	conexion := config.Conexion()

	rows, err := conexion.Query("SELECT id, nombre, email FROM usuario WHERE activo = 1")
	if err != nil {
		panic("Error al leer usuarios")
	}

	conexion.Close()

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario

		rows.Scan(&usuario.Id, &usuario.Nombre, &usuario.Email)
		usuarios = append(usuarios, usuario)
	}

	return usuarios
}
