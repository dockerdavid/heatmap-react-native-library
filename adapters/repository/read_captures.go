package heatmapRepositoryAdapter

import (
	"encoding/json"
	heatmapDomain "heatmap_mobile/domain"
)

func (r Repository) ReadCaptures() (string, error) {
	const query = `SELECT coordinate_x, coordinate_y, screen_name, uuid, image, status, created_at, updated_at FROM capture WHERE status = '0'`

	var captures []heatmapDomain.Capture

	prepare, err := r.DB.DB.Prepare(query)

	if err != nil {
		return "", err
	}

	defer prepare.Close()

	rows, err := prepare.Query()

	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		var capture heatmapDomain.Capture

		if err = rows.Scan(&capture.CoordinateX, &capture.CoordinateY, &capture.ScreenName, &capture.UUID, &capture.Image, &capture.Status, &capture.CreatedAt, &capture.UpdatedAt); err != nil {
			return "", err
		}

		captures = append(captures, capture)
	}

	marshal, err := json.Marshal(captures)

	if err != nil {
		return "", err
	}

	return string(marshal), nil
}
