package alerts

import (
	"encoding/json"
	"fmt"

	"github.com/jmhobbs/nws"
)

type Service struct {
	client *nws.Client
}

func New(client *nws.Client) *Service {
	return &Service{client}
}

//func (s *Service) Alerts() {}

//func (s *Service) ActiveAlerts() {}

// Get all of the possible Alert types.
func (s *Service) AlertTypes() ([]string, error) {
	var types alertTypesResponse

	resp, err := s.client.LDGet("https://api.weather.gov/alerts/types")
	if err != nil {
		return []string{}, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	err = dec.Decode(&types)
	if err != nil {
		return []string{}, fmt.Errorf("error decoding API response: %v", err)
	}

	return types.EventTypes, nil
}
