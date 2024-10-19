package heatmapDomain

import "encoding/json"

type ClientKey struct {
	ApiKey string `json:"api_key"`
}

func (c ClientKey) ToJSON() string {
	marshal, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type Device struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Manufacturer    string  `json:"manufacturer"`
	ScreenWidth     float64 `json:"screen_width"`
	ScreenHeight    float64 `json:"screen_height"`
	BuildID         string  `json:"build_id"`
	BatteryLevel    float64 `json:"battery_level"`
	BuildNumber     string  `json:"build_number"`
	BundleID        string  `json:"bundle_id"`
	FreeDiskStorage float64 `json:"free_disk_storage"`
	PowerState      string  `json:"power_state"`
	ReadableVersion string  `json:"readable_version"`
	SystemName      string  `json:"system_name"`
	SystemVersion   string  `json:"system_version"`
}

func (d Device) ToJSON() string {
	marshal, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type Capture struct {
	CoordinateX float64 `json:"coordinate_x"`
	CoordinateY float64 `json:"coordinate_y"`
	ScreenName  string  `json:"screen_name"`
	UUID        string  `json:"uuid"`
	Image       string  `json:"image"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func (c Capture) ToJSON() string {
	marshal, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(marshal)
}
