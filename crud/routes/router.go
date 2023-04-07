package routes

import (
	"crud/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/users/", http.StatusFound)
	})

	//todas las rutas tienen este prefijo /users entonces creamos esta varibale para decir que voy a usar en todas el prefijo /users
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("/", handlers.ReadUsers).Methods("GET")
	userRouter.HandleFunc("/create", handlers.CreateUser).Methods("GET", "POST")
	userRouter.HandleFunc("/delete", handlers.DeleteUser).Methods("GET")
	userRouter.HandleFunc("/update", handlers.UpdateUser).Methods("GET", "POST")

	return router
}
