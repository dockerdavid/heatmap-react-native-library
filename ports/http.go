package heatmapPorts

type HTTP interface {
	SendHeatmaps(client string, device string, heatmaps string) error
}
