package store

import (
	"encoding/json"
	"fmt"
	"github.com/treenq/treenq-cli/src/dto"
	"net/http"
)

func GetInfo(url string) (dto.InfoResponse, error) {
	resp, err := http.Get(url + "/info")
	var info dto.InfoResponse
	if err != nil || resp.StatusCode != http.StatusOK {
		return info, fmt.Errorf("failed to get info: %w", err)
	}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return info, fmt.Errorf("failed to decode info: %w", err)
	}
	if info.Version == "" {
		return info, fmt.Errorf("response does not contain version: %w", err)
	}
	return info, nil
}
