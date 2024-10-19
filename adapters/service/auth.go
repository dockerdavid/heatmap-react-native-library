package heatmapServiceAdapter

func (s Service) Auth(ApiKey string) error {
	if err := s.Repository.DeleteAuth(); err != nil {
		return err
	}

	if err := s.Repository.SaveAuth(ApiKey); err != nil {
		return err
	}

	return nil
}
