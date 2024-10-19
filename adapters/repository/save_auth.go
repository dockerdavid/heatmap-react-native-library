package heatmapRepositoryAdapter

func (r Repository) SaveAuth(ApiKey string) error {
	if _, err := r.DB.DB.Exec("INSERT INTO client (api_key) VALUES (?)", ApiKey); err != nil {
		return err
	}

	return nil
}
