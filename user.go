package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//variaveis

var DB *gorm.DB
var err error

//conexão com o banco de dados
const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

//criando a estrutura usuario, definindo suas propriedades e usando json para receber e enviar as informações
type User struct {
	//usando gorm para criar um modelo que ajuda a lidar com o BD, adinicando id etc
	gorm.Model

	PrimeiroNome string `json:"primeironome"`
	SegundoNome  string `json:"segundonome"`
	Email        string `json:"email"`
}

//criando funcções

//conectando com o banco de dados usando o gorm e o DNS para conectar com o MYSQL e verificando erro de conecção
func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Falha ao conectar com o banco de dados")
	}
	DB.AutoMigrate(&User{})
}

//pegar todos usuarios usando GET
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//salvando em um slice
	var users []User
	//Buscando dados no BD
	DB.Find(&users)
	//pegando/passando dados do json
	json.NewEncoder(w).Encode(users)
}

//pegar usuarios por id usando GET
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	//Buscando dados no BD
	DB.First(&user, params["id"])
	//pegando/passando dados do json
	json.NewEncoder(w).Encode(user)
}

//criar usuario usando POST
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	//pegando/passando dados do json
	json.NewDecoder(r.Body).Decode(&user)
	//salvando dados no BD
	DB.Create(&user)
	//pegando/passando dados do json
	json.NewEncoder(w).Encode(user)
}

//atualizar usuarios por id usando PUT
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	//buscando dados no BD
	DB.First(&user, params["id"])
	//pegando/passando dados do json
	json.NewDecoder(r.Body).Decode(&user)
	//salvando dados no BD
	DB.Save(&user)
	//pegando/passando dados do json
	json.NewEncoder(w).Encode(user)
}

//deletar usuario por id usando DELETE
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	//deletando dados no BD
	DB.Delete(&user, params["id"])
	//pegando/passando dados do json
	json.NewEncoder(w).Encode("Usuario deletato!")
}
