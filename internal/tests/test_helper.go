package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getBodyFromResponse(response *http.Response, data any) error {

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if err := json.Unmarshal(bodyBytes, data); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return nil
}
