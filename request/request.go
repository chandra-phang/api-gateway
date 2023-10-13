package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const jsonType = "application/json"

func Get(url string) ([]byte, int, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil, response.StatusCode, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}
	return responseBytes, response.StatusCode, nil
}

func Post(url string, data []byte) ([]byte, int, error) {
	body := bytes.NewBuffer(data)

	response, err := http.Post(url, jsonType, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, response.StatusCode, err
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}
	return responseBytes, response.StatusCode, nil
}

func Put(url string, data []byte) ([]byte, int, error) {
	body := bytes.NewBuffer(data)

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", jsonType)

	// Create an HTTP client
	client := http.Client{}

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return nil, http.StatusInternalServerError, err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}
	return responseBytes, response.StatusCode, nil
}
