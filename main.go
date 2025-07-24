package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
}

type Person struct {
	Name      string   `json:"nome"`
	Nickname  string   `json:"apelido"`
	BirthDate string   `json:"nascimento"`
	Stack     []string `json:"stack"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func handlerPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		createPerson(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(ErrorResponse{Error: "Método não permitido"})
	}
}

func parsePerson(person *Person) error {
	if person.Name == "" {
		return errors.New("nome não pode ser nulo")
	}

	if person.Nickname == "" {
		return errors.New("apelido não pode ser nulo")
	}
	return nil
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	check(err)
	err = parsePerson(&person)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(person)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/pessoas", handlerPeople)

	fmt.Println("Rodando em http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
