package heatmapServiceAdapter

import heatmapPorts "heatmap_mobile/ports"

type Service struct {
	Repository heatmapPorts.Repository
	HTTP       heatmapPorts.HTTP
}
