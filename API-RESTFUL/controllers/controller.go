package controllers

import (
	"API-RESTFUL/database"
	"API-RESTFUL/models"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page")
}

func TodosUsuarios(w http.ResponseWriter, r *http.Request) {
	var arrayUsers []models.User
	database.DB.Find(&arrayUsers)
	json.NewEncoder(w).Encode(arrayUsers)
}

func RetornaUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	var userAtual models.User
	numeric := regexp.MustCompile(`\d`).MatchString(username)
	database.DB.Where("username = ?", username).First(&userAtual)
	if numeric {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username inválido")
	} else if userAtual.Username == "" {
		// 404
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuario não encontrado")
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userAtual)
	}

}

func CriaNovoUsuario(w http.ResponseWriter, r *http.Request) {

	var novo_user models.User
	json.NewDecoder(r.Body).Decode(&novo_user)
	var userTeste models.User
	database.DB.Where("username = ?", novo_user.Username).First(&userTeste)
	// is a bolean var
	numeric := regexp.MustCompile(`\d`).MatchString(novo_user.Username)

	if r.ContentLength == 0 {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Request não possui body")
	} else if numeric {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username inválido")
	} else if userTeste.Username == novo_user.Username {
		// 403
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Usuário já existe")
	} else {
		database.DB.Create(&novo_user)
		// 201
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Usuário criado com sucesso")
	}

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	var userTeste models.User
	database.DB.Where("username = ?", username).First(&userTeste)
	numeric := regexp.MustCompile(`\d`).MatchString(username)
	if numeric {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username inválido")
	} else if userTeste.Username == "" {
		// 404
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuário não encontrado")
	} else {
		// var userAtual models.User
		// json.NewDecoder(r.Body).Decode(&userAtual)
		database.DB.Delete(&userTeste, userTeste.Id)
		//200
		w.WriteHeader(http.StatusOK)
		// json.NewDecoder(r.Body).Decode(&userTeste)
		json.NewEncoder(w).Encode("Usuário deletado")
	}
}

func EditaUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	var userTeste models.User
	database.DB.Where("username = ?", username).First(&userTeste)
	numeric := regexp.MustCompile(`\d`).MatchString(username)

	if numeric {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Username inválido")
	} else if userTeste.Username == "" {
		// 404
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Usuário não encontrado")
	} else {
		// var userAtual models.User
		json.NewDecoder(r.Body).Decode(&userTeste)
		// database.DB.First(&userTeste, userTeste.Id)
		database.DB.Save(&userTeste)
		//200
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Usuário editado")
	}
}

func CriarUsuarioLista(w http.ResponseWriter, r *http.Request) {
	var arrayUsers []models.User
	json.NewDecoder(r.Body).Decode(&arrayUsers)
	// var ansRequest []models.User
	if len(arrayUsers) == 0 {
		// 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Request não possui body")
	} else {
		for i := 0; i < len(arrayUsers); i++ {
			var novo_user models.User = arrayUsers[i]
			database.DB.Create(&novo_user)
			// ansRequest = append(ansRequest, novo_user)
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Usuários criados com sucesso")
		// json.NewEncoder(w).Encode(ansRequest)
	}

	// json.NewEncoder(w).Encode(ansRequest)
}
