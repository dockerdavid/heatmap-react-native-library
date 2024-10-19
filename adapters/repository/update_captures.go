package heatmapRepositoryAdapter

func (r Repository) UpdateCaptures() error {
	prepare, err := r.DB.DB.Prepare(`UPDATE capture SET status = '1' WHERE status = '0'`)

	if err != nil {
		return err
	}

	_, err = prepare.Exec()

	if err != nil {
		return err
	}

	return nil
}
