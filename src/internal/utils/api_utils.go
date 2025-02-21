package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"knull/internal/dtos"
	"net/http"
	"os"
)

func JsonResponse(w http.ResponseWriter, payload dtos.ResponseDto) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)

}

// ParseRequest parses the headers and body of an HTTP request into a struct.
// The `headers` parameter should be a pointer to a struct that matches the header structure.
// The `body` parameter should be a pointer to a struct that matches the expected JSON body.
func ParseRequest(r *http.Request, headers interface{}, body interface{}) error {
	// Parse headers
	if headers != nil {
		err := parseHeaders(r, headers)
		if err != nil {
			return err
		}
	}

	// Parse body
	if body != nil {
		err := parseBody(r, body)
		if err != nil {
			return err
		}
	}

	return nil
}

// parseHeaders extracts headers from the request and maps them to the provided struct.
func parseHeaders(r *http.Request, headers interface{}) error {
	headerMap := make(map[string]string)
	for key, values := range r.Header {
		headerMap[key] = values[0] // Take the first value for simplicity
	}

	// Convert the header map to JSON and then unmarshal into the struct
	headerJSON, err := json.Marshal(headerMap)
	if err != nil {
		return err
	}

	return json.Unmarshal(headerJSON, headers)
}

// parseBody reads the request body and unmarshals it into the provided struct.
func parseBody(r *http.Request, body interface{}) error {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.Unmarshal(bodyBytes, body)
}

func IsDirEmpty(dir string) (bool, error) {
	// Open the directory
	f, err := os.Open(dir)
	if err != nil {
		return false, fmt.Errorf("failed to open directory: %v", err)
	}
	defer f.Close()

	// Read the directory contents (read up to 1 entry)
	_, err = f.Readdir(1)

	// If EOF is returned, the directory is empty
	if err == io.EOF {
		return true, nil
	}

	// If another error occurs, return it
	if err != nil {
		return false, fmt.Errorf("failed to read directory: %v", err)
	}

	// If no error and no EOF, the directory is not empty
	return false, nil
}
