package heatmapServiceAdapter

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func (s Service) SendHeatmap(TimeToBeSent string) error {
	c := cron.New()

	_, err := c.AddFunc(fmt.Sprintf("@every %s", TimeToBeSent), func() {
		fmt.Println("Sending heatmaps")
		readClient, err := s.Repository.ReadClient()

		if err != nil {
			return
		}

		readDevice, err := s.Repository.ReadDevice()

		if err != nil {
			return
		}

		readCaptures, err := s.Repository.ReadCaptures()

		if err != nil {
			return
		}

		if len(readCaptures) == 0 {
			return
		}

		if err = s.Repository.UpdateCaptures(); err != nil {
			return
		}

		if err = s.HTTP.SendHeatmaps(readClient, readDevice, readCaptures); err != nil {
			return
		}

	})

	if err != nil {
		return err
	}

	c.Start()

	select {}
}
