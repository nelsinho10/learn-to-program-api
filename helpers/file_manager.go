package helpers

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

// CreateFile creates a new file with the given name and content
func CreateFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer file.Close()

	file.WriteString(content)
}

// RunPythonFile runs a python file
func RunPythonFile(fileName string) []byte {
	cmd := exec.Command("python3", fileName)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal("Error running python file:", err)
	}
	return out
}

// DeleteFile deletes a file
func DeleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Fatal("Error deleting file:", err)
	}
}

// Error writes the error to the response
func Error(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
