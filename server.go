// Go Time Server Lab
// This is a simple Go server that returns the current time.
// commit to master
package main

//імпорти
import (
    "encoding/json"
    "net/http"
    "time"
)

// Структура для JSON-відповіді
type TimeResponse struct {
    Time string `json:"time"`
}

// Обробник HTTP-запиту для отримання поточного часу
func timeHandler(w http.ResponseWriter, r *http.Request) {
    response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}


func main() {
    // Реєструємо обробник для маршруту "/time", який буде викликати функцію timeHandler
    http.HandleFunc("/time", timeHandler)
    // Запускаємо HTTP-сервер на порту 8795, який оброблятиме запити
    http.ListenAndServe(":8795", nil)
}
