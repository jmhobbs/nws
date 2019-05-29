package alerts

import (
	"testing"

	"github.com/jmhobbs/nws/geo"
)

func TestAlertOptions(t *testing.T) {
	// Area, Point, Region, RegionType and Zone are all incompatible
	expectedErrors := []AlertsOptions{
		// RegionType
		AlertsOptions{AlertListOptions: AlertListOptions{RegionType: "land", Area: "NE"}},
		AlertsOptions{AlertListOptions: AlertListOptions{RegionType: "land", Point: &geo.Point{0, 0}}},
		AlertsOptions{AlertListOptions: AlertListOptions{RegionType: "land", Region: "NW"}},
		AlertsOptions{AlertListOptions: AlertListOptions{RegionType: "land", Zone: "NE"}},
		// Area
		AlertsOptions{AlertListOptions: AlertListOptions{Area: "NE", Point: &geo.Point{0, 0}}},
		AlertsOptions{AlertListOptions: AlertListOptions{Area: "NE", Region: "NW"}},
		AlertsOptions{AlertListOptions: AlertListOptions{Area: "NE", Zone: "NE"}},
		// Point
		AlertsOptions{AlertListOptions: AlertListOptions{Point: &geo.Point{0, 0}, Region: "NW"}},
		AlertsOptions{AlertListOptions: AlertListOptions{Point: &geo.Point{0, 0}, Zone: "NE"}},
		// Region
		AlertsOptions{AlertListOptions: AlertListOptions{Region: "NW", Zone: "NE"}},
	}

	for _, ao := range expectedErrors {
		_, err := ao.Values()
		if err == nil {
			t.Errorf("expected incompatible parameters but did not receive one: %v", ao)
		}
	}
}
