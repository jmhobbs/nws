package alerts

import (
	"fmt"
	"net/url"
	"time"

	"github.com/jmhobbs/nws/alerts/certainty"
	"github.com/jmhobbs/nws/alerts/messagetype"
	"github.com/jmhobbs/nws/alerts/severity"
	"github.com/jmhobbs/nws/alerts/status"
	"github.com/jmhobbs/nws/alerts/urgency"
	"github.com/jmhobbs/nws/geo"
)

type APIOptions interface {
	Values() (*url.Values, error)
}

// Query options shared among all alerts endpoints
type AlertListOptions struct {
	Status      status.Status
	MessageType messagetype.MessageType
	EventName   string
	EventCode   string
	// Region type (land or marine). This parameter is incompatible with the following parameters: area, point, region, zone
	RegionType string // TODO: Enum
	// Latitude and Longitude point. This parameter is incompatible with the following parameters: area, region, region_type, zone
	Point *geo.Point
	// Marine region code. This parameter is incompatible with the following parameters: area, point, region_type, zone
	Region string
	// State/marine area code. This parameter is incompatible with the following parameters: point, region, region_type, zone
	Area string
	// Zone ID (forecast or county). This parameter is incompatible with the following parameters: area, point, region, region_type
	Zone      string
	Urgency   urgency.Urgency
	Severity  severity.Severity
	Certainty certainty.Certainty
}

// Options for the alerts endpoint.
type AlertsOptions struct {
	AlertListOptions
	// Only active alerts
	Active bool
	// Start time
	Start time.Time
	// End time
	End time.Time
	// Limit number of results
	Limit int
	// Pagination cursor
	Cursor string
}

func (a AlertsOptions) Values() (*url.Values, error) {
	if a.RegionType != "" && (a.Area != "" || a.Point != nil || a.Region != "" || a.Zone != "") {
		return nil, fmt.Errorf("RegionType is incompatible with: Area, Point, Region and Zone")
	}
	if a.Area != "" && (a.Point != nil || a.Region != "" || a.Zone != "") {
		return nil, fmt.Errorf("Area is incompatible with: RegionType, Point, Region and Zone")
	}
	if a.Point != nil && (a.Region != "" || a.Zone != "") {
		return nil, fmt.Errorf("Point is incompatible with: RegionType, Area, Region and Zone")
	}
	if a.Region != "" && a.Zone != "" {
		return nil, fmt.Errorf("Region is incompatible with: RegionType, Point, Area and Zone")
	}
	// TODO
	return &url.Values{}, nil
}

type ActiveAlertsOptions struct {
	AlertListOptions
	Limit int
}
