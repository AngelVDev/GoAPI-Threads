package fetchbase

import (
	"encoding/json"
	"strings"
)

func constructBody(input FetchBaseInput) (*strings.Reader, error) {
	variablesJson, err := json.Marshal(input.Variables)
	if err != nil {
		return nil, err
	}
	body := strings.NewReader("lsd=jdFoLBsUcm9h-j90PeanuC&jazoest=21926&variables=" + string(variablesJson) + "&doc_id=" + input.DocumentId)
	return body, nil
}
