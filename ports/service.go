package heatmapPorts

type Service interface {
	Auth(ApiKey string) error
	RegisterDevice(Device string) error
	CaptureHeatmap(Capture string, UUID string) error
	SendHeatmap() error
	DeleteOldCaptures() error
}
