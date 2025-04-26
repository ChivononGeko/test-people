package enrichment

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"test-people/internal/ports/out"
)

type externalDataAdapter struct{}

func NewExternalDataAdapter() out.EnrichmentClient {
	return &externalDataAdapter{}
}

func (e *externalDataAdapter) GetAge(name string) (*int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	slog.Info("Sending request to Agify API", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Failed to get age", "error", err, "url", url)
		return nil, fmt.Errorf("failed to get age: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		Age int `json:"age"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		slog.Error("Failed to decode age response", "error", err)
		return nil, fmt.Errorf("failed to decode age response: %w", err)
	}

	slog.Info("Successfully retrieved age", "age", data.Age)

	return &data.Age, nil
}

func (e *externalDataAdapter) GetGender(name string) (*string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	slog.Info("Sending request to Genderize API", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Failed to get gender", "error", err, "url", url)
		return nil, fmt.Errorf("failed to get gender: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		Gender string `json:"gender"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		slog.Error("Failed to decode gender response", "error", err)
		return nil, fmt.Errorf("failed to decode gender response: %w", err)
	}

	slog.Info("Successfully retrieved gender", "gender", data.Gender)
	return &data.Gender, nil
}

func (e *externalDataAdapter) GetNationality(name string) (*string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	slog.Info("Sending request to Nationalize API", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Failed to get nationality", "error", err, "url", url)
		return nil, fmt.Errorf("failed to get nationality: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		Country string `json:"country"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		slog.Error("Failed to decode nationality response", "error", err)
		return nil, fmt.Errorf("failed to decode nationality response: %w", err)
	}

	slog.Info("Successfully retrieved nationality", "country", data.Country)
	return &data.Country, nil
}
