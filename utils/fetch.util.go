package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// POST utiliti; return response body and error
func Post(rpcUrl string, request interface{}) ([]byte, error) {
	requestBody, marshalErr := json.Marshal(request)

	if marshalErr != nil {
		log.Printf("[POST] Marshal Request Error: %v", marshalErr)
		return nil, marshalErr
	}

	response, fetchErr := http.Post(rpcUrl, "application/json", bytes.NewBuffer(requestBody))

	if fetchErr != nil {
		log.Printf("[POST] Send Request Error: %v", fetchErr)
		return nil, fetchErr
	}

	defer response.Body.Close()

	body, readErr := io.ReadAll(response.Body)

	if readErr != nil {
		log.Printf("[POST] Read Response Body Error: %v", readErr)
		return nil, readErr
	}

	return body, nil
}
