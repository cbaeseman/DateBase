package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/google/uuid"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&User{})

	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/upload", uploadFile).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20) // 10 MB limit

    file, handler, err := r.FormFile("file")
    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }
    defer file.Close()

    tempFile, err := ioutil.TempFile("uploads", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    tempFile.Write(fileBytes)

    filePath := tempFile.Name()
    response := map[string]string{"file_path": filePath}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

