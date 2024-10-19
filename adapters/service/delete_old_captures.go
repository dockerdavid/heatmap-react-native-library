package heatmapServiceAdapter

func (s Service) DeleteOldCaptures() error {
	if err := s.Repository.DeleteOldCaptures(); err != nil {
		return err
	}

	return nil
}
