package heatmapServiceAdapter

func (s Service) RegisterDevice(device string) error {
	if err := s.Repository.SaveDevice(device); err != nil {
		return err
	}

	return nil
}
