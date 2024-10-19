package heatmapServiceAdapter

func (s Service) CaptureHeatmap(capture string, uuid string) error {
	if err := s.Repository.SaveCapture(capture, uuid); err != nil {
		return err
	}

	return nil
}
