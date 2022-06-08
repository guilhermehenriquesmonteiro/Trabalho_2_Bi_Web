package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	//criando nova rota e passando valor para ela
	r := mux.NewRouter()

	//definindo as handlers functions

	//pegar todos usuarios usando GET
	r.HandleFunc("/users", GetUsers).Methods("GET")
	//pegar usuarios por id usando GET
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	//criar usuario usando POST
	r.HandleFunc("/users", CreateUser).Methods("POST")
	//atualizar usuarios por id usando PUT
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	//deletar usuario por id usando DELETE
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}
