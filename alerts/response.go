package alerts

import (
	"time"

	"github.com/jmhobbs/nws/alerts/certainty"
	"github.com/jmhobbs/nws/alerts/messagetype"
	"github.com/jmhobbs/nws/alerts/severity"
	"github.com/jmhobbs/nws/alerts/status"
	"github.com/jmhobbs/nws/alerts/urgency"
)

// Response to /alerts
type AlertsResponse struct {
	Context  []interface{} `json:"@context"`
	Type     string        `json:"type"`
	Features []struct {
		ID         string      `json:"id"`
		Type       string      `json:"type"`
		Geometry   interface{} `json:"geometry"`
		Properties struct {
			IDURL    string `json:"@id"`
			Type     string `json:"@type"`
			ID       string `json:"id"`
			AreaDesc string `json:"areaDesc"`
			Geocode  struct {
				UGC  []string `json:"UGC"`
				SAME []string `json:"SAME"`
			} `json:"geocode"`
			AffectedZones []string `json:"affectedZones"`
			References    []struct {
				IDURL      string `json:"@id"`
				Identifier string `json:"identifier"`
				Sender     string `json:"sender"`
				Sent       string `json:"sent"`
			} `json:"references"`
			Sent        string                  `json:"sent"`
			Effective   string                  `json:"effective"`
			Onset       string                  `json:"onset"`
			Expires     string                  `json:"expires"`
			Ends        string                  `json:"ends"`
			Status      status.Status           `json:"status"`
			MessageType messagetype.MessageType `json:"messageType"`
			Category    string                  `json:"category"`
			Severity    severity.Severity       `json:"severity"`
			Certainty   certainty.Certainty     `json:"certainty"`
			Urgency     urgency.Urgency         `json:"urgency"`
			Event       string                  `json:"event"`
			Sender      string                  `json:"sender"`
			SenderName  string                  `json:"senderName"`
			Headline    string                  `json:"headline"`
			Description string                  `json:"description"`
			Instruction string                  `json:"instruction"`
			Response    string                  `json:"response"`
			Parameters  struct {
				NWSheadline     []string `json:"NWSheadline"`
				HazardType      []string `json:"HazardType"`
				VTEC            []string `json:"VTEC"`
				EASORG          []string `json:"EAS-ORG"`
				PIL             []string `json:"PIL"`
				BLOCKCHANNEL    []string `json:"BLOCKCHANNEL"`
				EventEndingTime []string `json:"eventEndingTime"`
			} `json:"parameters"`
			ReplacedBy string `json:"replacedBy"`
			ReplacedAt string `json:"replacedAt"`
		} `json:"properties"`
	} `json:"features"`
	Title      string    `json:"title"`
	Updated    time.Time `json:"updated"`
	Pagination struct {
		Next string `json:"next"`
	} `json:"pagination"`
}

// Response to /alerts/types
type alertTypesResponse struct {
	Context    []interface{} `json:"@context"`
	EventTypes []string      `json:"eventTypes"`
}
