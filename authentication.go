package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginRequest LoginRequest
	err := decoder.Decode(&loginRequest)
	if err != nil {
		log.Println(err)
	}
	log.Println(loginRequest)
	if !loginValidation(loginRequest.User, loginRequest.Pass) {
		response := "Usuario o Contraseña Incorrecta"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}
	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Username: loginRequest.User,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(tokenString)
	response := LoginResponse{
		Message: "Inicio de sessión correctamente",
		Token:   tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)

}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	if user.Pass != "" {
		user.Pass = HashPassword(user.Pass)
	}

	if userExist(user) {
		response := "Usuario ya existe"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	fmt.Println(user)
	if insertUser(user) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}

}

func insertUser(user User) bool {
	db, err := sql.Open("mysql", "root:48821181Ap!@tcp(127.0.0.1:3306)/Practice")
	fmt.Println(user)
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO `users` (user, pass, email) values (?,?,?)"
	insert, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := insert.Exec(user.User, user.Pass, user.Email)
	fmt.Println(resp)
	insert.Close()

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func userExist(user User) bool {
	db, err := sql.Open("mysql", "root:48821181Ap!@tcp(127.0.0.1:3306)/Practice")
	if err != nil {
		fmt.Println(err)
	}
	query, err := db.Prepare(`SELECT count(*) FROM users where user like  ? `)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := query.Query(user.User)
	if err != nil {
		fmt.Println(err)
	}
	var count int
	rows.Scan(&count)

	fmt.Println(count)

	if count > 0 {
		return true
	}
	return false
}

func loginValidation(us string, password string) bool {
	db, err := sql.Open("mysql", "root:48821181Ap!@tcp(127.0.0.1:3306)/Practice")
	if err != nil {
		fmt.Println(err)
	}
	query, err := db.Prepare(`SELECT * FROM users where user like  ? `)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := query.Query(us)
	if err != nil {
		fmt.Println(err)
	}
	var user User
	for rows.Next() {

		err = rows.Scan(&user.Id, &user.User, &user.Pass, &user.Email)
		if err != nil {
			fmt.Println("ERROR en consulta", err)
			return false
		}
	}

	fmt.Println(user)

	if CheckPasswordHash(password, user.Pass) {
		return true
	}
	return false
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
