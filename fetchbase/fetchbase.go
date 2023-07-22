package fetchbase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

	match := extractUserID(string(html))
	return match, nil
}
func extractUserID(html string) string {
	startIndex := strings.Index(html, `"user_id":"`)
	if startIndex == -1 {
		return ""
	}

	endIndex := strings.Index(html[startIndex+len(`"user_id":"`):], `"`)
	if endIndex == -1 {
		return ""
	}

	return html[startIndex+len(`"user_id":"`) : startIndex+len(`"user_id":"`)+endIndex]
}

// func fetchUserProfile(userId, userName string) (string, error) {
// 	if userName != "" && userId == "" {
// 		fetchedUserId, err := FetchUserIdByName(userName)
// 		if err != nil {
// 			return "", err
// 		}
// 		userId = fetchedUserId
// 	}

// 	variables := map[string]string{"userID": userId}
// 	documentId := USER_PROFILE
// 	data, err := FetchBase(variables)
// 	if err != nil {
// 		return "", err
// 	}

// 	return mapUserProfile(data), nil
// }
