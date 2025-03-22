package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DSY4444/go-task-manager/database"
	"github.com/DSY4444/go-task-manager/models"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task

	if err := database.DB.First(&task, params["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&task)
	database.DB.Save(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// database.DB.Delete(&models.Task{}, params["id"])
	// w.WriteHeader(http.StatusNoContent)
	params := mux.Vars(r)
	var task models.Task

	if err := database.DB.Unscoped().First(&task, params["id"]).Error; err != nil {
		http.Error(w, "Task not found11", http.StatusNotFound)
		return
	}

	if err := database.DB.Unscoped().Delete(&task).Error; err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task permanently deleted from database",
	})

}
