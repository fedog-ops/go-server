package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"acme/db"
	"strconv"
)


func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}

func handleUsers(writer http.ResponseWriter, request *http.Request) {
    
    fmt.Printf("got /api/users request\n")

	users := db.GetUsers()
	json.NewEncoder(writer).Encode(users)
}

func createUser(writer http.ResponseWriter, request *http.Request) {

        
    var user db.User
    err := json.NewDecoder(request.Body).Decode(&user)
    if err != nil {
        fmt.Println("Error decoding request body:", err)
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

    id := db.AddUser(user)
    writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "User created successfully: %d", id)

}
func getSingleUser(writer http.ResponseWriter, request *http.Request) {

    idStr := request.PathValue("id")
    
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

    user := db.GetUser(id)

    json.NewEncoder(writer).Encode(user)

}
func deleteSingleUser(writer http.ResponseWriter, request *http.Request) {

    idStr := request.PathValue("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

	db.DeleteUser(id)
	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "User deleted successfully: %d", id)

}

func putSingleUser(writer http.ResponseWriter, request *http.Request) {

	idStr := request.PathValue("id")

    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

    var updatedUser db.User
    err2 := json.NewDecoder(request.Body).Decode(&updatedUser)
    if err2 != nil {
        fmt.Println("Error decoding request body:", err2)
        http.Error(writer, "Bad Request", http.StatusBadRequest)
        return
    }

    success := db.PutUser(updatedUser, id)
	if success {
    	writer.WriteHeader(http.StatusCreated)
		fmt.Fprintf(writer, "User updated successfully: %d", updatedUser.ID)
	} else{
		http.Error(writer, "User not found", http.StatusNotFound)
	}
}


func main() {

	//db connection  code
	router := http.NewServeMux()

    router.HandleFunc("GET /", rootHandler)
    router.HandleFunc("GET /api/users", handleUsers)
    router.HandleFunc("POST /api/users", createUser)
	router.HandleFunc("GET /api/users/{id}", getSingleUser)
	router.HandleFunc("DELETE /api/users/{id}", deleteSingleUser)
	router.HandleFunc("PUT /api/users/{id}", putSingleUser)


	// Starting the HTTP server on port 8080
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", CorsMiddleware(router))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        writer.Header().Set("Access-Control-Allow-Origin", "*")
        // Continue with the next handler
        next.ServeHTTP(writer, request)
    })
}
