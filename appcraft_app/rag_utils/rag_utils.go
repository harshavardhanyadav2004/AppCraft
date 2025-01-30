package rag_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ragAPIEndpoint = "https://api-inference.huggingface.co/models/your-username/your-model-name"
	apiToken = "your-api-token"
)

type RAGResponse struct {
	Answer string `json:"answer"`
}

func ExtractAnswer(query string) (string, error) {
	req, err := http.NewRequest("POST", ragAPIEndpoint, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ragResponse RAGResponse
	err = json.Unmarshal(body, &ragResponse)
	if err != nil {
		return "", err
	}

	return ragResponse.Answer, nil
}