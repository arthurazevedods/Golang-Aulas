package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Finish      bool   `json:"finish"`
}

var tarefas []Task

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/tarefas", handleTarefas)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem-vindo à API de tarefas!"))
}
func handleTarefas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listarTarefas(w, r)
	case http.MethodPost:
		criarTarefa(w, r)
	case http.MethodPut:
		finalizarTarefa(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}
func listarTarefas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tarefas)
}
func criarTarefa(w http.ResponseWriter, r *http.Request) {
	var tarefa Task
	err := json.NewDecoder(r.Body).Decode(&tarefa)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tarefa.ID = len(tarefas) + 1
	tarefa.Finish = false
	tarefas = append(tarefas, tarefa)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tarefa)
}

func finalizarTarefa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var aDeletar Task

	json.NewEncoder(w).Encode(tarefas)

	json.NewDecoder(r.Body).Decode(&aDeletar)

	fmt.Println(aDeletar.ID)
	id := aDeletar.ID
	for i := 0; i < len(tarefas); i++ {
		if tarefas[i].ID == id {
			tarefas = append(tarefas[:i], tarefas[i+1:]...)
			break
		}
	}
	fmt.Println("Depois:")
	fmt.Println(tarefas)
}
