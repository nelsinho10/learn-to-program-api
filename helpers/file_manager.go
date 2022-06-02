package helpers

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// CreateFile creates a new file with the given name and content
func CreateFile(fileName string, content string) error {
	path := filepath.Join("temp", fileName)
	file, err := os.Create(path)
	if err != nil {
		log.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	file.WriteString(content)
	return nil
}

// RunPythonFile runs a python file
func RunPythonFile(fileName string) ([]byte, error) {
	path := filepath.Join("temp", fileName)
	cmd := exec.Command("python3", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error running python file:", err)
		return nil, err
	}
	return out, nil
}

// DeleteFile deletes a file
func DeleteFile(fileName string) error {
	path := filepath.Join("temp", fileName)
	err := os.Remove(path)
	if err != nil {
		log.Println("Error deleting file:", err)
		return err
	}
	return nil
}
