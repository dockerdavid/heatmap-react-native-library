package heatmapRepositoryAdapter

func (r Repository) DeleteAuth() error {
	if _, err := r.DB.DB.Exec("DELETE FROM client"); err != nil {
		return err
	}

	return nil
}
