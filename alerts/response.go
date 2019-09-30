package alerts

import (
	"time"

	"github.com/jmhobbs/nws/alerts/certainty"
	"github.com/jmhobbs/nws/alerts/messagetype"
	"github.com/jmhobbs/nws/alerts/severity"
	"github.com/jmhobbs/nws/alerts/urgency"
)

// Response to /alerts
type AlertsResponse struct {
	Context struct {
		Wx    string `json:"wx"`
		Vocab string `json:"@vocab"`
	} `json:"@context"`
	Graph []struct {
		URI_ID   string `json:"@id"`
		Type     string `json:"@type"`
		ID       string `json:"id"`
		AreaDesc string `json:"areaDesc"`
		Geometry string `json:"geometry"`
		Geocode  struct {
			UGC  []string `json:"UGC"`
			SAME []string `json:"SAME"`
		} `json:"geocode"`
		AffectedZones []string `json:"affectedZones"`
		References    []struct {
			ID         string `json:"@id"`
			Identifier string `json:"identifier"`
			Sender     string `json:"sender"`
			Sent       string `json:"sent"`
		} `json:"references"`
		Sent        string                  `json:"sent"`
		Effective   string                  `json:"effective"`
		Onset       string                  `json:"onset"`
		Expires     string                  `json:"expires"`
		Ends        string                  `json:"ends"`
		Status      string                  `json:"status"`
		MessageType messagetype.MessageType `json:"messageType"`
		Category    string                  `json:"category"`
		Severity    severity.Severity       `json:"severity"`
		Certainty   certainty.Certainty     `json:"certainty"`
		Urgency     urgency.Urgency         `json:"urgency"`
		Event       string                  `json:"event"` // TODO: Type/Enum?
		Sender      string                  `json:"sender"`
		SenderName  string                  `json:"senderName"`
		Headline    string                  `json:"headline"`
		Description string                  `json:"description"`
		Instruction string                  `json:"instruction"`
		Response    string                  `json:"response"`
		Parameters  struct {
			VTEC            []string `json:"VTEC"`
			EASORG          []string `json:"EAS-ORG"`
			PIL             []string `json:"PIL"`
			BLOCKCHANNEL    []string `json:"BLOCKCHANNEL"`
			EventEndingTime []string `json:"eventEndingTime"`
		} `json:"parameters"`
	} `json:"@graph"`
	Title   string    `json:"title"`
	Updated time.Time `json:"updated"`
}

// Response to /alerts/types
type alertTypesResponse struct {
	Context    []interface{} `json:"@context"`
	EventTypes []string      `json:"eventTypes"`
}
