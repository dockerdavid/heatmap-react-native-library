package heatmapServiceAdapter

func (s Service) SendHeatmap() error {
	readClient, err := s.Repository.ReadClient()

	if err != nil {
		return err
	}

	readDevice, err := s.Repository.ReadDevice()

	if err != nil {
		return err
	}

	readCaptures, err := s.Repository.ReadCaptures()

	if err != nil {
		return err
	}

	if len(readCaptures) == 0 {
		return nil
	}

	if err = s.Repository.UpdateCaptures(); err != nil {
		return err
	}

	if err = s.HTTP.SendHeatmaps(readClient, readDevice, readCaptures); err != nil {
		return err
	}

	return nil
}
