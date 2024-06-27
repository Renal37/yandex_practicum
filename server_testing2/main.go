package main

import (
    "encoding/json"
    "log"
    "net/http"
)

// User — основной объект для теста.
type User struct {
    ID        string
    FirstName string
    LastName  string
}

// UserViewHandler — хендлер, который нужно протестировать.
/* func UserViewHandler(users map[string]User) http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        userId := r.URL.Query().Get("user_id")
        if userId == "" {
            http.Error(rw, "user_id is empty", http.StatusBadRequest)
            return
        }

        user, ok := users[userId]
        if !ok {
            http.Error(rw, "user not found", http.StatusNotFound)
            return
        }

        jsonUser, err := json.Marshal(user)
        if err != nil {
            http.Error(rw, "can't provide a json. internal error", 
            http.StatusInternalServerError)
            return
        }

        rw.WriteHeader(http.StatusOK)
        rw.Write(jsonUser)
    }
}
 */
 func UserViewHandler(users map[string]User) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse the query parameters
        params := r.URL.Query()
        userID := params.Get("user_id")

        // Check if the user ID is provided
        if userID == "" {
            http.Error(w, "User ID is required", http.StatusBadRequest)
            return
        }

        // Check if the user exists
        user, ok := users[userID]
        if !ok {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }

        // Set the response headers
        w.Header().Set("Content-Type", "application/json")

        // Marshal the user data to JSON
        jsonResponse, err := json.Marshal(user)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Write the JSON response to the response writer
        w.WriteHeader(http.StatusOK)
        w.Write(jsonResponse)
    }
}
func main() {
    users := make(map[string]User)
    u1 := User{
        ID:        "u1",
        FirstName: "Misha",
        LastName:  "Popov",
    }
    u2 := User{
        ID:        "u2",
        FirstName: "Sasha",
        LastName:  "Popov",
    }
    users["u1"] = u1
    users["u2"] = u2

    http.HandleFunc("/users", UserViewHandler(users))
    log.Fatal(http.ListenAndServe(":8080", nil))
}