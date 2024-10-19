package heatmapHttpAdapter

import (
	"bytes"
	"encoding/json"
	"errors"
	heatmapDomain "heatmap_mobile/domain"
	"net/http"
)

const (
	URL = "https://take-it.amovil.com.co:42281/api/stats/heatmap"
)

var (
	ErrorApiKeyRequired   = errors.New("api key is required")
	ErrorMarshallingData  = errors.New("error marshalling data")
	ErrorCreatingRequest  = errors.New("error creating request")
	ErrorSendingRequest   = errors.New("error sending request")
	ErrorUnexpectedStatus = errors.New("unexpected status code")
)

func (h HTTP) SendHeatmaps(client string, device string, heatmaps string) error {
	var valueClient heatmapDomain.ClientKey
	var valueDevice heatmapDomain.Device
	var valueHeatmaps []heatmapDomain.Capture

	_ = json.Unmarshal([]byte(client), &valueClient)
	_ = json.Unmarshal([]byte(device), &valueDevice)
	_ = json.Unmarshal([]byte(heatmaps), &valueHeatmaps)

	if valueClient.ApiKey == "" {
		return ErrorApiKeyRequired
	}

	var i = map[string]interface{}{
		"device":   valueDevice,
		"heatmaps": valueHeatmaps,
	}

	data, err := json.Marshal(i)
	if err != nil {
		return ErrorMarshallingData
	}

	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewBuffer(data))
	if err != nil {
		return ErrorCreatingRequest
	}

	req.Header.Set("X-API-KEY", valueClient.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return ErrorSendingRequest
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return ErrorUnexpectedStatus
	}

	return nil
}
