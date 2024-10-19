package heatmapRepositoryAdapter

import (
	"encoding/json"
	heatmapDomain "heatmap_mobile/domain"
)

func (r Repository) ReadClient() (string, error) {
	const query = `SELECT api_key FROM client LIMIT 1`

	var clientKey heatmapDomain.ClientKey

	prepare, err := r.DB.DB.Prepare(query)

	if err != nil {
		return "", err
	}

	defer prepare.Close()

	if err = prepare.QueryRow().Scan(&clientKey.ApiKey); err != nil {
		return "", err
	}

	marshal, err := json.Marshal(clientKey)

	return string(marshal), nil
}
