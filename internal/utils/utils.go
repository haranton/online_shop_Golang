package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetIDFromRequest(r *http.Request) (uint, error) {
	idStr := r.PathValue("id")
	if idStr == "" {
		return 0, fmt.Errorf("id is empty")
	}
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %w", err)
	}
	if idInt < 0 {
		return 0, fmt.Errorf("invalid id: negative")
	}
	return uint(idInt), nil
}
