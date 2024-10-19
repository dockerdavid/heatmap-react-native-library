package heatmapRepositoryAdapter

import (
	"encoding/json"
	heatmapDomain "heatmap_mobile/domain"
	"time"
)

func (r Repository) SaveCapture(capture string, uuid string) error {
	prepare, err := r.DB.DB.Prepare(`INSERT INTO capture (coordinate_x, coordinate_y, screen_name, uuid, image, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	var c heatmapDomain.Capture

	err = json.Unmarshal([]byte(capture), &c)

	if err != nil {
		return err
	}

	_, err = prepare.Exec(c.CoordinateX, c.CoordinateY, c.ScreenName, uuid, c.Image, "0", time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))

	if err != nil {
		return err
	}

	return nil
}
