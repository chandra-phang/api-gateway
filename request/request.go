package request

import (
	"bytes"
	"io"
	"net/http"
)

const jsonType = "application/json"

func Get(url string, authorization string) ([]byte, int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", jsonType)

	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}

	// Create an HTTP client
	client := http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return respBytes, resp.StatusCode, nil
}

func Post(url string, data []byte, authorization string) ([]byte, int, error) {
	body := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", jsonType)
	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}

	// Create an HTTP client
	client := http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return respBytes, resp.StatusCode, nil
}

func Put(url string, data []byte, authorization string) ([]byte, int, error) {
	body := bytes.NewBuffer(data)

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", jsonType)

	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}

	// Create an HTTP client
	client := http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return respBytes, resp.StatusCode, nil
}
