package heatmapPorts

type Repository interface {
	DeleteAuth() error
	ReadCaptures() (string, error)
	ReadDevice() (string, error)
	ReadClient() (string, error)
	SaveAuth(ApiKey string) error
	SaveCapture(Capture string, UUID string) error
	SaveDevice(Device string) error
	UpdateCaptures() error
	DeleteOldCaptures() error
}
