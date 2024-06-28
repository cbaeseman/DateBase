package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
)

// UploadFileHandler handles the file upload and storage
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Parse the multipart form, limit the size to 10MB
    err := r.ParseMultipartForm(10 << 20) // 10MB
    if err != nil {
        http.Error(w, "Error parsing form data", http.StatusBadRequest)
        return
    }

    // Get the file from the form data
    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving the file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create the file path
    filePath := filepath.Join("filestore", handler.Filename)
    os.MkdirAll(filepath.Dir(filePath), os.ModePerm) // Ensure the directory exists

    // Create the file on the server
    dst, err := os.Create(filePath)
    if err != nil {
        http.Error(w, "Error saving the file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    // Copy the uploaded file to the server file
    _, err = io.Copy(dst, file)
    if err != nil {
        http.Error(w, "Error copying the file", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

func main() {
    http.HandleFunc("/upload", UploadFileHandler)

    // Start the server on port 8080
    fmt.Println("Server started at :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}

