package routes

import (
	"API-RESTFUL/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/user", controllers.TodosUsuarios).Methods("Get")
	r.HandleFunc("/api/user/{username}", controllers.RetornaUsuario).Methods("Get")
	r.HandleFunc("/api/user", controllers.CriaNovoUsuario).Methods("Post")
	r.HandleFunc("/api/user/{username}", controllers.DeletarUsuario).Methods("Delete")
	r.HandleFunc("/api/user/{username}", controllers.EditaUsuario).Methods("Put")
	r.HandleFunc("/api/user/list", controllers.CriarUsuarioLista).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
