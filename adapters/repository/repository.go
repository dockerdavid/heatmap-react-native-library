package heatmapRepositoryAdapter

import (
	"heatmap_mobile/pkg/database/sqlite"
)

type Repository struct {
	DB *sqlite.Client
}
