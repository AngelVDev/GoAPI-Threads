package fetchbase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
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

func FetchUserIdByName(username string) (string, error) {
	if IS_DEBUG {
		fmt.Printf("https://www.threads.net/@%s\n", username)
	}
	resp, err := http.Get(fmt.Sprintf("https://www.threads.net/@%s", username))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`"user_id":"(\d+)"`)
	match := re.FindStringSubmatch(string(html))
	if len(match) > 1 {
		return match[1], nil
	}

	return "", nil
}
