package fetchbase

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchBase(input FetchBaseInput) (map[string]interface{}, error) {
	body, err := constructBody(input)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", GRAPHQL_ENDPOINT, body)
	req.Header = initHeaders()

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var result map[string]interface{}

	json.Unmarshal(bodyBytes, &result)

	return result, nil
}
