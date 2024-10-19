package heatmapRepositoryAdapter

func (r Repository) DeleteOldCaptures() error {
	prepare, err := r.DB.DB.Prepare("DELETE FROM capture WHERE created_at < datetime('now', '-7 day')")

	if err != nil {
		return err
	}

	_, err = prepare.Exec()

	if err != nil {
		return err
	}

	return nil
}
